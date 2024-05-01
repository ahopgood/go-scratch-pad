package puml

import "fmt"

type Uml struct {
	start   string
	diagram UmlDiagram
	end     string
}

func NewUml(umlDiagram UmlDiagram) Uml {
	return Uml{
		start:   "@startuml\n",
		diagram: umlDiagram,
		end:     "@enduml",
	}
}

func (uml Uml) Contents() string {
	return uml.start + uml.diagram.Contents() + uml.end
}

type UmlDiagram interface {
	Contents() string
}

type Digraph struct {
	start        string
	dependencies []Dependency
	end          string
}

func NewDigraph(dependencies []Dependency) Digraph {
	return Digraph{
		start:        "digraph test {\n",
		dependencies: dependencies,
		end:          "}\n",
	}
}

// Represents a dependency From one thing To another
// E.g. "salt-master" -> "salt-common"
type Dependency struct {
	From string
	To   string
}

// Uml Diagram implementation
func (d Digraph) Contents() string {
	fmt.Println("Building diagraph contents")
	output := d.start + "\n"
	for _, value := range d.dependencies {
		output = output + "\t" + "\"" + value.From + "\"" + " -> " + "\"" + value.To + "\"" + "\n"

	}
	output = output + "\n" + d.end + "\n"
	return output
}
