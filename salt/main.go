package main

import (
	"com/alexander/scratch/salt/debendency"
	"com/alexander/scratch/salt/puml"
	"flag"
	"fmt"
)

func main() {

	var conf Config
	flag.StringVar(&conf.packageName, "p", "", ".deb package name to calculate dependencies for")
	flag.BoolVar(&conf.generateSalt, "s", false, "output dependencies as salt code")
	flag.BoolVar(&conf.generateDiagram, "d", false, "output dependencies as a diagram")
	flag.StringVar(&conf.installerLocation, "o", "", "output directory to save installer files to")
	flag.Parse()

	fmt.Printf("%#v", conf)

	if true == conf.generateDiagram {
		packageModelMap := make(map[string]*debendency.PackageModel)
		//rootPackageModel := debendency.NewPackager().BuildPackage(conf.packageName, packageModelMap)
		debendency.NewPackager().BuildPackage(conf.packageName, packageModelMap)
		fmt.Println(puml.GenerateDiagram(packageModelMap).Contents())
	}
}

type Config struct {
	packageName       string
	generateSalt      bool
	generateDiagram   bool
	installerLocation string
}
