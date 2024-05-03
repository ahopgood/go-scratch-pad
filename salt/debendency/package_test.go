package debendency_test

import (
	"com/alexander/scratch/salt/debendency"
	"com/alexander/scratch/salt/internal"
	"com/alexander/scratch/salt/puml"
	"errors"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// ginkgo --focus Package
var _ = Describe("Package", func() {

	// packageModel := salt.PackageModel{}
	// Possibly create a table test here using: command, standard out, standard error, exit status?
	// When("Package not recognised", func() {
	// 	It("fails with error message", func() {
	// 		// vagrant@vagrant:~$ apt download test
	// 		// E: Unable to locate debendency test
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

				model := debendency.PackageModel{}
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

	jq := "Get:1 http://gb.archive.ubuntu.com/ubuntu focal-updates/universe amd64 jq amd64 1.6-1ubuntu0.20.04.1 [50.2 kB]\nFetched 50.2 kB in 0s (1,666 kB/s)\n"
	jqFile := "jq_1.6-1ubuntu0.20.04.1_amd64.deb"
	libjq1 := "Get:1 http://gb.archive.ubuntu.com/ubuntu focal-updates/universe amd64 libjq1 amd64 1.6-1ubuntu0.20.04.1 [121 kB]\nFetched 121 kB in 0s (2,898 kB/s)\n"
	libjq1File := "libjq1_1.6-1ubuntu0.20.04.1_amd64.deb"
	libonig5 := "Get:1 http://gb.archive.ubuntu.com/ubuntu focal/universe amd64 libonig5 amd64 6.9.4-1 [142 kB]\nFetched 142 kB in 0s (3,228 kB/s)\n"
	libonig5File := "libonig5_6.9.4-1_amd64.deb"
	libc6 := "Get:1 http://gb.archive.ubuntu.com/ubuntu focal-updates/main amd64 libc6 amd64 2.31-0ubuntu9.15 [2,723 kB]\nFetched 2,723 kB in 0s (20.9 MB/s)\n"
	libc6File := "libc6_2.31-0ubuntu9.15_amd64.deb"

	FWhen("BuildPackage", func() {
		When("Package does not exist", func() {
			It("Produces nothing", func() {
				// Fail to download the debendency
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturns("", 0, errors.New("not found"))

				dpkger := &internal.FakeDpkg{}
				packager := debendency.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]*debendency.PackageModel)
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
				// Successfully download the debendency
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturns(successMessage, 0, nil)

				// No dependencies found via dpkg -I
				dpkger := &internal.FakeDpkg{}
				dpkger.IdentifyDependenciesReturns([]string{})

				packager := debendency.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]*debendency.PackageModel)
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

		When("Package has one dependency", func() {
			It("Produces two models", func() {
				// Successfully download the debendency
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturnsOnCall(0, successMessage, 0, nil)
				apter.DownloadPackageReturnsOnCall(1, libc6SuccessMessage, 0, nil)

				// No dependencies found via dpkg -I
				dpkger := &internal.FakeDpkg{}
				dpkger.IdentifyDependenciesReturnsOnCall(0, []string{"libc6"})
				dpkger.IdentifyDependenciesReturnsOnCall(1, []string{})

				packager := debendency.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]*debendency.PackageModel)
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

		FWhen("Package has shared dependencies", func() {
			It("Produces four models", func() {
				// Successfully download the debendency
				apter := &internal.FakeApt{}
				apter.DownloadPackageReturnsOnCall(0, jq, 0, nil)
				apter.DownloadPackageReturnsOnCall(1, libjq1, 0, nil)
				apter.DownloadPackageReturnsOnCall(2, libonig5, 0, nil)
				apter.DownloadPackageReturnsOnCall(3, libc6, 0, nil)

				// No dependencies found via dpkg -I
				dpkger := &internal.FakeDpkg{}
				dpkger.IdentifyDependenciesReturnsOnCall(0, []string{"libjq1", "libc6"})
				dpkger.IdentifyDependenciesReturnsOnCall(1, []string{"libonig5", "libc6"})
				dpkger.IdentifyDependenciesReturnsOnCall(2, []string{"libc6"})
				dpkger.IdentifyDependenciesReturnsOnCall(3, []string{})

				packager := debendency.Packager{
					Apt:  apter,
					Dpkg: dpkger,
				}

				modelMap := make(map[string]*debendency.PackageModel)
				_ = packager.BuildPackage("", modelMap)

				By("Invoking Apt and Dpkg twice", func() {
					Expect(apter.DownloadPackageCallCount()).To(Equal(4))
					Expect(dpkger.IdentifyDependenciesCallCount()).To(Equal(4))
				})
				By("Adding four models to the map", func() {
					Expect(len(modelMap)).To(Equal(4))
				})
				By("Returning jq model", func() {
					model := modelMap["jq"]
					Expect(model.Name).To(Equal("jq"))
					//Expect(model.Version).To(Equal("7.4.0-2"))
					Expect(model.Filepath).To(Equal(jqFile))
				})
				By("Returning libjq1 model", func() {
					model := modelMap["libjq1"]
					Expect(model.Name).To(Equal("libjq1"))
					//Expect(model.Version).To(Equal("2.31-0ubuntu9.15"))
					Expect(model.Filepath).To(Equal(libjq1File))
				})
				By("Returning libonig5 model", func() {
					model := modelMap["libonig5"]
					Expect(model.Name).To(Equal("libonig5"))
					//Expect(model.Version).To(Equal("2.31-0ubuntu9.15"))
					Expect(model.Filepath).To(Equal(libonig5File))
				})

				By("Returning libc6 model", func() {
					model := modelMap["libc6"]
					Expect(model.Name).To(Equal("libc6"))
					//Expect(model.Version).To(Equal("2.31-0ubuntu9.15"))
					Expect(model.Filepath).To(Equal(libc6File))
				})
				By("Adding libc6 to the jq dependencies map", func() {
					model := modelMap["jq"]
					Expect(model.Dependencies["libc6"]).To(Not(BeNil()))
				})
				By("Adding libc6 to the jq dependencies map", func() {
					model := modelMap["libjq1"]
					Expect(model.Dependencies["libc6"]).To(Not(BeNil()))
				})
				By("Adding libonig5 to the  dependencies map", func() {
					model := modelMap["libonig5"]
					Expect(model.Dependencies["libc6"]).To(Not(BeNil()))
				})
				By("Producing a puml diagram", func() {
					dependencies := make([]puml.Dependency, 0)
					// Loop through the model dependencies?
					for key, fromModel := range modelMap {
						fmt.Printf("%s %#v\n", key, fromModel)
						for _, toModel := range fromModel.Dependencies {
							fmt.Printf("From %s to %s\n", fromModel.Name, toModel.Name)
							dependencies = append(dependencies, puml.Dependency{
								From: fromModel.Name,
								To:   toModel.Name,
							})
						}
					}
					puml := puml.NewUml(
						puml.NewDigraph(dependencies),
					)
					fmt.Println(puml.Contents())
				})
			})
		})
	})
})
