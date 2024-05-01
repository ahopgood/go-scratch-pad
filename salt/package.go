package salt

import (
	"fmt"
	"strings"

	"com/alexander/scratch/salt/commands"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o internal/fake_dpkger.go com/alexander/scratch/salt/commands.Dpkg
//counterfeiter:generate -o internal/fake_apter.go com/alexander/scratch/salt/commands.Apt

type PackageModel struct {
	Filepath     string
	Name         string
	Version      string
	Dependencies map[string]*PackageModel
}

type Packager struct {
	Apt  commands.Apt
	Dpkg commands.Dpkg
}

func NewPackager() Packager {
	command := commands.LinuxCommand{}
	apter := commands.Apter{
		Cmd: command,
	}
	dpkger := commands.Dpkger{
		Cmd: command,
	}
	return Packager{
		Apt:  apter,
		Dpkg: dpkger,
	}
}

func (packager Packager) Start(name string) *PackageModel {
	modelMap := make(map[string]*PackageModel)
	return packager.BuildPackage(name, modelMap)
}

func (packager Packager) BuildPackage(name string, modelMap map[string]*PackageModel) *PackageModel {
	model, ok := modelMap[name]
	if ok {
		return model
	}
	// start with 1st package
	standardOut, _, err := packager.Apt.DownloadPackage(name)

	if err != nil {
		return &PackageModel{}
	}
	packageModel := PackageModel{}
	packageModel.GetPackageFilename(standardOut)

	// Add packageModel to map under package name
	// packageModel.Name
	modelMap[packageModel.Name] = &packageModel
	// Iterate through dependencies
	dependencyNames := packager.Dpkg.IdentifyDependencies(packageModel.Filepath)

	fmt.Printf("Dependencies %#v\n", dependencyNames)
	fmt.Printf("Dependencies length %d\n", len(dependencyNames))
	packageModel.Dependencies = make(map[string]*PackageModel, len(dependencyNames))

	for _, name := range dependencyNames {
		fmt.Printf("Building dependency [%s] for %s \n", name, packageModel.Name)
		dep := packager.BuildPackage(name, modelMap)
		packageModel.Dependencies[dep.Name] = dep
	}
	return &packageModel
}

func (packageModel *PackageModel) GetPackageFilename(name string) {
	fmt.Printf("Package download output: %#v\n", name)
	outputArray := strings.Split(name, "\n")
	fmt.Printf("Number of lines: %d\n", len(outputArray))
	for index, _ := range outputArray {
		fmt.Println(outputArray[index])
	}

	downloadOutputLine := strings.Split(outputArray[0], " ")
	// Length should be 8
	// Get:1 http://gb.archive.ubuntu.com/ubuntu focal/universe amd64 dos2unix amd64 7.4.0-2 [374 kB]
	// Get:1 https://repo.saltproject.io/py3/ubuntu/20.04/amd64/3004 focal/main amd64 salt-master all 3004.2+ds-1 [40.9 kB]
	packageName := downloadOutputLine[4]
	fmt.Printf("PackageName: %s\n", packageName)
	arch := downloadOutputLine[5]
	fmt.Printf("Arch: %s\n", arch)
	version := downloadOutputLine[6]
	fmt.Printf("Version: %s\n", version)

	fileName := packageName + "_" + version + "_" + arch + ".deb"
	fmt.Printf("Filename: %s\n", fileName)
	//Check file exists
	packageModel.Filepath = fileName
	packageModel.Name = packageName
	packageModel.Version = version
}

func (packageModel PackageModel) Generate(filepath string) PackageModel {
	// native call to dpkg -I filepath
	// extract package name
	// extract package version
	// extract dependencies
	// create packageModel
	// Loop through dependencies
	// Download package
	// apt-get download packagename
	// Recursively call Generate(filepath string)
	// Add new PackageModel to dependencies array
	return PackageModel{}
}

func CreateSaltState(packageModel PackageModel) {
	// Create an .sls file from a package model
	// packageModel.name:
	// source: salt://{somepath}/packageModel.filepath
	// source filepath may vary depending on whether we're dealing with a dependency or not
	// version: packageModel.version
	// requires:
	// 		- pkg: packageModel.dependencies[0]
	// recursively call CreateSaltSate(packageModel PackageModel) for each dep
}

func SetupVariables() {
	// use os.getenv()
	// to retrieve download location where files are downloaded to
	// salt state installer file destination
	// output filename (if one already exists create one with a timestamp)
	// then call CreateSaltState?
}
