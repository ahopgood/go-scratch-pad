package salt_test

import (
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

	When("getInstallerPath", func() {

		When("Downloaded Successfully", func() {
			It("Should construct file name correctly", func() {
				successMessage := `
				Get:1 http://gb.archive.ubuntu.com/ubuntu focal/universe amd64 dos2unix amd64 7.4.0-2 [374 kB]
				Fetched 374 kB in 0s (4,447 kB/s)
				`

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

	When("BuildPackage", func() {
		When("Package has no dependencies", func() {
		})
		When("Package has one dependency", func() {
		})
		When("Package has shared dependency", func() {
		})
	})
})
