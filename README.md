# go-dot
> Simple dot file writer in go

## Example use
```
import (
       os
       
       dot "github.com/zenground0/go-dot"
)

g := dot.NewGraph("MyGraph", os.Stdout)
var v1 dot.VertexDescription
v1.ID = "V1"
v1.Label = "Node1"
v1.Color = "blue2"
g.AddVertex(&v1)
var v2 dot.VertexDescription
v2.ID = "V2"
v2.Label = "Node2"
v2.Color = "blue1"
g.AddVertex(&v2)

g.AddEdge(v1, v2, true)
g.WriteDot()
```
output:
```
digraph MyGraph {
	V1 [label="Node1" color="blue2"]
	V2 [label="Node2" color="blue1"]
	V1 -> V2
}
```
## Contribute
Currently attribute support is barebones.  Only attributes "label" and "color" are implemented.  PRs welcome to extend this or add other features.  Check the issues for proposals in flight, or this reference http://www.graphviz.org/pdf/dotguide.pdf for new ideas.

## License
MIT © Protocol Labs, Inc