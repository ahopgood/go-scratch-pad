package salt_test

import (
	"com/alexander/scratch/salt/internal"
	"errors"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"com/alexander/scratch/salt"
)

// ginkgo --focus Package
var _ = Describe("Package", func() {

	// packageModel := salt.PackageModel{}
	// Possibly create a table test here using: command, standard out, standard error, exit status?
	// When("Package not recognised", func() {
	// 	It("fails with error message", func() {
	// 		// vagrant@vagrant:~$ apt download test
	// 		// E: Unable to locate package test
	// 		// echo $?
	// 		// 130
	// 		packageModel.DownloadPackage("")

	// 	})
	// })

	// When("Package not recognised", func() {
	// 	It("fails with error message", func() {
	// 		// vagrant@vagrant:~$ sudo apt download test
	// 		// W: Download is performed unsandboxed as root as file '/home/vagrant/dos2unix_7.4.0-2_amd64.deb' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
	// 		// echo $?
	// 		// 0
	// 		packageModel.DownloadPackage("")
	// 	})
	// })
	// When("Package recognised", func() {
	// 	It("downloads to current working directory", func() {
	// 		// vagrant@vagrant:~$ apt download dos2unix
	// 		// Get:1 http://gb.archive.ubuntu.com/ubuntu focal/universe amd64 dos2unix amd64 7.4.0-2 [374 kB]
	// 		// Fetched 374 kB in 0s (5,732 kB/s)
	// 		// echo $?
	// 		// 0
	// 		packageModel.DownloadPackage("")
	// 	})
	// })

	When("GetPackageFilename", func() {

		When("Downloaded Successfully", func() {
			It("Should construct file name correctly", func() {
				successMessage :=
					`Get:1 http://gb.archive.ubuntu.com/ubuntu focal/universe amd64 dos2unix amd64 7.4.0-2 [374 kB]
Fetched 374 kB in 0s (4,447 kB/s)`

				model := salt.PackageModel{}
				model.GetPackageFilename(successMessage)
				Expect(model.Name).To(Equal("dos2unix"))
				Expect(model.Version).To(Equal("7.4.0-2"))
				Expect(model.Filepath).To(Equal("dos2unix_7.4.0-2_amd64.deb"))
			})
		})

	})
	AfterEach(func() {
		os.Remove("dos2unix_7.4.0-2_amd64.deb")
	})

	successMessage :=
		`Get:1 http://gb.archive.ubuntu.com/ubuntu focal/universe amd64 dos2unix amd64 7.4.0-2 [374 kB]
Fetched 374 kB in 0s (4,447 kB/s)`

	libc6SuccessMessage :=
		`Get:1 http://gb.archive.ubuntu.com/ubuntu focal-updates/main amd64 libc6 amd64 2.31-0ubuntu9.15 [2,723 kB]
Fetched 2,723 kB in 0s (12.4 MB/s)`

	FWhen("BuildPackage", func() {
		When("Package does not exist", func() {
			It("Produces nothing", func() {
				// Fail to download the package
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturns("", 0, errors.New("not found"))

				dpkger := &internal.FakeDpkg{}
				packager := salt.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]salt.PackageModel)
				model := packager.BuildPackage("", modelMap)
				By("Not adding a model to the map", func() {
					Expect(len(modelMap)).To(Equal(0))
				})
				By("Returning an empty model", func() {
					Expect(model.Name).To(BeEmpty())
					Expect(model.Version).To(BeEmpty())
					Expect(model.Filepath).To(BeEmpty())
				})
				By("Apt being only invoked once", func() {
					Expect(apter.DownloadPackageCallCount()).To(Equal(1))
				})
				By("Dpkg not being invoked", func() {
					Expect(dpkger.IdentifyDependenciesCallCount()).To(Equal(0))
				})
			})
		})

		When("Package has no dependencies", func() {
			It("Produces a single model", func() {
				// Successfully download the package
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturns(successMessage, 0, nil)

				// No dependencies found via dpkg -I
				dpkger := &internal.FakeDpkg{}
				dpkger.IdentifyDependenciesReturns([]string{})

				packager := salt.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]salt.PackageModel)
				model := packager.BuildPackage("", modelMap)

				By("Adding a model to the map", func() {
					Expect(len(modelMap)).To(Equal(1))
				})
				By("Returning a populated model", func() {
					Expect(model.Name).To(Equal("dos2unix"))
					Expect(model.Version).To(Equal("7.4.0-2"))
					Expect(model.Filepath).To(Equal("dos2unix_7.4.0-2_amd64.deb"))
				})
				By("Invoking Apt and Dpkg only once", func() {
					Expect(apter.DownloadPackageCallCount()).To(Equal(1))
					Expect(dpkger.IdentifyDependenciesCallCount()).To(Equal(1))
				})
			})
		})

		FWhen("Package has one dependency", func() {
			It("Produces two models", func() {
				// Successfully download the package
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturnsOnCall(0, successMessage, 0, nil)
				apter.DownloadPackageReturnsOnCall(1, libc6SuccessMessage, 0, nil)

				// No dependencies found via dpkg -I
				dpkger := &internal.FakeDpkg{}
				dpkger.IdentifyDependenciesReturnsOnCall(0, []string{"libc6"})
				dpkger.IdentifyDependenciesReturnsOnCall(1, []string{})

				packager := salt.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]salt.PackageModel)
				_ = packager.BuildPackage("", modelMap)

				By("Invoking Apt and Dpkg twice", func() {
					Expect(apter.DownloadPackageCallCount()).To(Equal(2))
					Expect(dpkger.IdentifyDependenciesCallCount()).To(Equal(2))
				})
				By("Adding two models to the map", func() {
					Expect(len(modelMap)).To(Equal(2))
				})
				By("Returning dos2unix model", func() {
					model := modelMap["dos2unix"]
					Expect(model.Name).To(Equal("dos2unix"))
					Expect(model.Version).To(Equal("7.4.0-2"))
					Expect(model.Filepath).To(Equal("dos2unix_7.4.0-2_amd64.deb"))
				})
				By("Returning libc6 model", func() {
					model := modelMap["libc6"]
					Expect(model.Name).To(Equal("libc6"))
					Expect(model.Version).To(Equal("2.31-0ubuntu9.15"))
					Expect(model.Filepath).To(Equal("libc6_2.31-0ubuntu9.15_amd64.deb"))
				})
				By("Adding libc6 to the dos2unix dependencies map", func() {
					model := modelMap["dos2unix"]
					Expect(model.Dependencies["libc6"]).To(Not(BeNil()))
				})

			})
		})

		When("Package has shared dependency", func() {
		})
	})
})
