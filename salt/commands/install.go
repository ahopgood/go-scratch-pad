package commands

import "strings"

// Interfaces may allow for faking when testing our native commands
type Dpkg interface {
	IdentifyDependencies(filename string) []string
	ParseDependencies(output string) []string
}

type Dpkger struct {
	Cmd Command
}

func (dpkg Dpkger) IdentifyDependencies(filename string) []string {
	output, exitCode, err := dpkg.Cmd.Command("dpkg", "-I", filename)

	if err != nil {
		//shit the bed
		//log stuff
		//propagate error
	}

	if exitCode != 0 {
		// shit the bed
		// log output
		// propagate an error
	}

	//parse output into an array of package names
	return dpkg.ParseDependencies(output)

}

func (dpkg Dpkger) ParseDependencies(output string) []string {

	//Find the line Depends:
	_, after, found := strings.Cut(output, "Depends:")
	if found {
		depends := strings.Split(after, "\n")[0]
		dependencies := strings.Split(depends, ",")
		for i := range dependencies {
			dependencies[i] = strings.TrimSpace(dependencies[i])
		}
		return dependencies
	}
	return make([]string, 0)
}
