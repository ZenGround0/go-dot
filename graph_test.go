package dot

import (
	"bytes"
	"testing"
)

var emptyGraph = `digraph testGraph {


 }`

func TestMakeWriteEmptyGraph(t *testing.T) {
	buf := new(bytes.Buffer)
	g := NewGraph("testGraph", buf)
	err := g.WriteDot()
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if s != emptyGraph {
		t.Errorf("unexpected output: \n%s", s)
	}
}

var edgeGraph = `digraph testGraph {

vTo -> vFrom

 }`

func TestAddEdge(t *testing.T) {
	buf := new(bytes.Buffer)
	g := NewGraph("testGraph", buf)
	vTo := &VertexDescription {
		ID: "vTo",
	}
	vFrom := &VertexDescription {
		ID: "vFrom",
	}
	g.AddEdge(vTo, vFrom, true)
	err := g.WriteDot()
	if err != nil{
		t.Fatal(err)
	}
	s := buf.String()
	if s != edgeGraph {
		t.Errorf("unexpected output: \n%s",s)
	}

}

var vertexGraph = `digraph testGraph {

v [label="vertex"]

 }`

func TestAddVertex(t *testing.T) {
	buf := new(bytes.Buffer)
	g := NewGraph("testGraph", buf)
	v := &VertexDescription {
		ID: "v",
		Label: "vertex",
	}
	g.AddVertex(v)
	err := g.WriteDot()
	if err != nil{
		t.Fatal(err)
	}
	s := buf.String()
	if s != vertexGraph {
		t.Errorf("unexpected output: \n%s\n",s)
		t.Errorf("expected outpuf: \n%s", vertexGraph)
	}
}

var commentGraph = `digraph testGraph {

/* This is a comment */

 }`

func TestAddComment(t *testing.T) {
	buf := new(bytes.Buffer)
	g := NewGraph("testGraph", buf)
	g.AddComment("This is a comment")
	err := g.WriteDot()
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if s!= commentGraph {
		t.Errorf("unexpeted output: \n%s\n",s)
	}
}

var newlineGraph = `digraph testGraph {



 }`

func TestAddNewLine(t *testing.T) {
	buf := new(bytes.Buffer)
	g := NewGraph("testGraph", buf)
	g.AddNewLine()
	err := g.WriteDot()
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if s!= newlineGraph {
		t.Errorf("unexpeted output: \n%s\n",s)
		t.Errorf("expected ouput: \n%s\n", newlineGraph)
	}
}

var basicGraph = `digraph cluster {

/* The nodes of the connectivity graph */
/* The cluster-service peers */
C0 [label="EhD" color="blue2"]
C1 [label="DQJ" color="blue2"]
C2 [label="mJu" color="blue2"]

/* The ipfs peers */
I0 [label="Ssq" color="goldenrod"]
I1 [label="ZDV" color="goldenrod"]
I2 [label="suL" color="goldenrod"]

/* Edges representing active connections in the cluster */
/* The connections among cluster-service peers */
C0 -> C1
C0 -> C2
C1 -> C0
C1 -> C2
C2 -> C0
C2 -> C1

/* The connections between cluster peers and their ipfs daemons */
C0 -> I1
C1 -> I0
C2 -> I2

/* The swarm peer connections among ipfs daemons in the cluster */
I0 -> I1
I0 -> I2
I1 -> I0
I1 -> I2
I2 -> I0
I2 -> I1

 }`


func TestPrintBasicGraph(t *testing.T) {
	buf := new(bytes.Buffer)
	g := NewGraph("cluster", buf)
	g.AddComment("The nodes of the connectivity graph")
	g.AddComment("The cluster-service peers")
	c0 := &VertexDescription{
		ID: "C0",
		Label: "EhD",
		Color: "blue2",
	}
	c1 := &VertexDescription{
		ID: "C1",
		Label: "DQJ",
		Color: "blue2",
	}
	c2 := &VertexDescription{
		ID: "C2",
		Label: "mJu",
		Color: "blue2",
	}	
	g.AddVertex(c0)
	g.AddVertex(c1)
	g.AddVertex(c2)
	g.AddNewLine()

	g.AddComment("The ipfs peers")
	i0 := &VertexDescription{
		ID: "I0",
		Label: "Ssq",
		Color: "goldenrod",
	}
	i1 := &VertexDescription{
		ID: "I1",
		Label: "ZDV",
		Color: "goldenrod",
	}
	i2 := &VertexDescription{
		ID: "I2",
		Label: "suL",
		Color: "goldenrod",
	}
	g.AddVertex(i0)
	g.AddVertex(i1)
	g.AddVertex(i2)
	g.AddNewLine()

	g.AddComment("Edges representing active connections in the cluster")
	g.AddComment("The connections among cluster-service peers")
	g.AddEdge(c0, c1, true)
	g.AddEdge(c0, c2, true)	
	g.AddEdge(c1, c0, true)	
	g.AddEdge(c1, c2, true)	
	g.AddEdge(c2, c0, true)
	g.AddEdge(c2, c1, true)
	g.AddNewLine()

	g.AddComment("The connections between cluster peers and their ipfs daemons")
	g.AddEdge(c0, i1, true)
	g.AddEdge(c1, i0, true)
	g.AddEdge(c2, i2, true)
	g.AddNewLine()

	g.AddComment("The swarm peer connections among ipfs daemons in the cluster")
	g.AddEdge(i0, i1, true)
	g.AddEdge(i0, i2, true)
	g.AddEdge(i1, i0, true)		
	g.AddEdge(i1, i2, true)
	g.AddEdge(i2, i0, true)
	g.AddEdge(i2, i1, true)

	err := g.WriteDot()
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if s != basicGraph {
		t.Errorf("unexpected output \n%s\n", s)
		t.Errorf("The expected output \n%s\n", basicGraph)
	}
}

