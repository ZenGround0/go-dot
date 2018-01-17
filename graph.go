package dot

import (
	"fmt"
	"io"
)

type DotGraph struct {
	Name   string
	Body []Line
	w      io.Writer
}

type Line interface {
	Write(io.Writer) error
}

type VertexDescription struct {
	ID    string
	Label string
	Color string
}

type EdgeDescription struct {
	From     VertexDescription
	To       VertexDescription
	Directed bool
}

type Literal struct {
	Line string
}

// Write the vertex description to file
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

// Write the edge description to file
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

func (lit *Literal) Write(w io.Writer) error {
	_, err := io.WriteString(w, lit.Line)
	return err
}

// Get a new DotGraph that will write to w
func NewGraph(name string, w io.Writer) DotGraph {
	return DotGraph {
		w: w,
		Name: name,
	}
}

func (graph *DotGraph) AddComment(text string) {
	commentStr := fmt.Sprintf("/* %s */", text)
	line := &Literal {
		Line: commentStr,
	}
	graph.Body = append(graph.Body, line)
}

func (graph *DotGraph) AddNewLine() {
	line := &Literal {
		Line: "", // newline already printed for every line
	}
	graph.Body = append(graph.Body, line)
}

// Add the vertex 
func (graph *DotGraph) AddVertex(v *VertexDescription) {
	graph.Body = append(graph.Body, v)
}

// Add an edge, directed or undirected, between the two vertices.  
func (graph *DotGraph) AddEdge(v1 *VertexDescription, v2 *VertexDescription, directed bool) {
	edge := &EdgeDescription{
		From: *v1,
		To:   *v2,
		Directed: directed,
	}
	graph.Body = append(graph.Body, edge)
}

func (graph *DotGraph) WriteDot() error {
	title := fmt.Sprintf("digraph %s {\n\n", graph.Name)
	_, err := io.WriteString(graph.w, title)
	if err != nil {
		return err
	}

	for _, line := range graph.Body {
		err = line.Write(graph.w)
		_, err2 := io.WriteString(graph.w, "\n")
		if err != nil || err2 != nil {
			return err
		}
		
	}

	_, err = io.WriteString(graph.w, "\n }")
	return err
}
