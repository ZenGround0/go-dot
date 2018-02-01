// Package dot provides an abstraction for painlessly building and printing
// graphviz dot-files
package dot

import (
	"fmt"
	"io"
)

// Graph is the graphviz dot-file graph representation.
type Graph struct {
	Name   string
	Body []Element
}

// Element captures the information of a dot-file element,
// typically corresoponding to one line of the file
type Element interface {
	Write(io.Writer) error
}

// VertexDescription is an element containing all the information needed to
// fully describe a dot-file vertex
type VertexDescription struct {
	ID    string
	Label string
	Color string
}

// EdgeDescription is an element containing all the information needed to
// fully describe a dot-file edge
type EdgeDescription struct {
	From     VertexDescription
	To       VertexDescription
	Directed bool
}

// Literal is an element consisting of the corresponding literal string
// printed in the dot-file
type Literal struct {
	Line string
}

// Write writes the vertex description to a writer
func (v *VertexDescription) Write(w io.Writer) error {
	nodeStr := fmt.Sprintf("%s ", v.ID)
	// TODO: This is clunky already and certainly won't scale to more
	// attributes.  Should have a map of attribute labels (reflectors?)
	if v.Label != "" && v.Color == "" {
		nodeStr += fmt.Sprintf( "[label=\"%s\"]", v.Label)
	}
	if v.Label == "" && v.Color != "" {
		nodeStr += fmt.Sprintf( "[color=\"%s\"]", v.Color)
	}
	if v.Label != "" && v.Color != "" {
		nodeStr += fmt.Sprintf( "[label=\"%s\" color=\"%s\"]", v.Label, v.Color)
	}
	_, err := io.WriteString(w, nodeStr)
	return err
}

// Write writes the edge description to a writer
func (e *EdgeDescription) Write(w io.Writer) error {
	var arrow string
	if e.Directed {
		arrow = "->"
	} else {
		arrow = "--"
	}
	edgeStr := fmt.Sprintf("%s %s %s", e.From.ID, arrow, e.To.ID)
	_, err := io.WriteString(w, edgeStr)
	return err
}

// Write writes the literal to a writer
func (lit *Literal) Write(w io.Writer) error {
	_, err := io.WriteString(w, lit.Line)
	return err
}

// NewGraph returns a new dot-file graph object given the provided name
func NewGraph(name string) Graph {
	return Graph {
		Name: name,
	}
}

// AddComment interprets the given argument as the text of a comment and
// schedules the comment to be written in the output dotfile
func (graph *Graph) AddComment(text string) {
	commentStr := fmt.Sprintf("/* %s */", text)
	line := &Literal {
		Line: commentStr,
	}
	graph.Body = append(graph.Body, line)
}

// AddNewLine schedules a newline to be written in the output dotfile
func (graph *Graph) AddNewLine() {
	line := &Literal {
		Line: "", // newline already printed for every line
	}
	graph.Body = append(graph.Body, line)
}

// AddVertex schedules the vertexdescription to be written in the output
// dotfile
func (graph *Graph) AddVertex(v *VertexDescription) {
	graph.Body = append(graph.Body, v)
}

// AddEdge constructs an edgedescription connecting the two vertices given
// as parameters and schedules this element to be written in the output dotfile
func (graph *Graph) AddEdge(v1 *VertexDescription, v2 *VertexDescription, directed bool) {
	edge := &EdgeDescription{
		From: *v1,
		To:   *v2,
		Directed: directed,
	}
	graph.Body = append(graph.Body, edge)
}

// WriteDot writes the elements scheduled on this Graph to the provided
// writer to construct a valid dot-file
func (graph *Graph) WriteDot(w io.Writer) error {
	title := fmt.Sprintf("digraph %s {\n\n", graph.Name)
	_, err := io.WriteString(w, title)
	if err != nil {
		return err
	}

	for _, line := range graph.Body {
		err = line.Write(w)
		_, err2 := io.WriteString(w, "\n")
		if err != nil || err2 != nil {
			return err
		}
		
	}

	_, err = io.WriteString(w, "\n }")
	return err
}
