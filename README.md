# go-dot
> Simple dot file writer in go

## Example use
```
import (
       os
       
       dot "github.com/zenground0/go-dot"
       )

g := dot.NewGraph("MyGraph")
v1 := dot.Vertex()
v1.ID = "V1"
v1.Label = "Node1"
v1.Color = "blue2"
g.AddVertex(v1)
v2 := dot.Vertex()
v2.ID = "V2"
v2.Label = "Node2"
v2.Color = "blue1"
g.AddVertex(v2)

g.AddDirectedEdge(v1, v2)
g.WriteDot(os.Stdout)
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
Currently attribute support is barebones.  Only attributes "label" and "color" are implemented.  PRs welcome to extend this or add other features.  Check the issues for ideas and proposals.

## License
MIT © Protocol Labs, Inc