package main

import "fmt"

type NodeID int

type Node struct {
    item NodeID
}

type Edge struct {
    from, to NodeID
    wt int
}

type Graph struct {
	nodes map[NodeID]*Node
	edges map[NodeID]map[NodeID]*Edge
}

type BFS struct {
    g *Graph
    processed map[NodeID]bool
    processQ *NodeList
    searchPath *NodeList
}

type DFS struct {
    g *Graph
    processed map[NodeID]bool
    topoSortStack *NodeStack  //Stack
}

func NewGraph() *Graph {
    return &Graph {
        nodes: make(map[NodeID]*Node),
        edges: make(map[NodeID]map[NodeID]*Edge),
    }
}

func (g *Graph) AddEdge(f, t NodeID, w int) {
    // Add from node, if does not exist
    if _, ok := g.nodes[f]; !ok {
        g.nodes[f] = &Node{item: f}
    }
    if _, ok := g.nodes[t]; !ok {
        g.nodes[t] = &Node{item: t}
    }
    // Add Edge
    edges, ok := g.edges[f]
    if !ok {
        edges = make(map[NodeID]*Edge)
        g.edges[f] = edges
    }
    if _, ok := edges[t]; !ok {
        edges[t] = &Edge{from: f, to: t, wt: w}
    }
}

func (g *Graph) Print() {
	fmt.Println("------ Graph -------")
	fmt.Printf("Graph: Nodes[%d]", len(g.nodes))
	for nk := range g.nodes {
		emap := g.edges[nk]
		fmt.Printf("\n  [%d]: Edges[%d]: ", nk, len(emap))
		for _, edge := range emap {
			fmt.Printf("%d:%d ", edge.from, edge.to)
		}
	}
	fmt.Println("\n--------")
}

func (bfs *BFS) Search(start, item NodeID) bool {
    sn := bfs.g.nodes[start]
    bfs.processQ.Enqueue(sn)

    for bfs.processQ.Size() > 0 {
        n := bfs.processQ.Dequeue()

        // Mark it processed
        bfs.processed[n.item] = true
        bfs.searchPath.Enqueue(n)
        // Search of item found
        if n.item == item {
            return true
        }
        // Enque edges, if not found already
        if edges, ok := bfs.g.edges[n.item]; ok {
            for t, _ := range edges {
                if _, ok := bfs.processed[t]; !ok {
                    bfs.processed[t] = false
                    tn := bfs.g.nodes[t]
                    bfs.processQ.Enqueue(tn)
                }
            }
        }
    }

    // Looked thru all nodes BFS, not found
    return false
}

func (dfs *DFS) Walk(nid NodeID) {
    // Node discovered
    dfs.processed[nid] = false

    // Process all the edges
    if edges, ok := dfs.g.edges[nid]; ok {
        for t, _ := range edges {
            // If a Undiscovered node is encountered
            if _, ok := dfs.processed[t]; !ok {
                dfs.Walk(t)
            }
        }
    }

    // Node is processed
    dfs.processed[nid] = true
    dfs.topoSortStack.Push(dfs.g.nodes[nid])
}

func bfs_search_wrapper(gr *Graph, start, item NodeID) {
    fmt.Println("====== BFS ========")
    gr.Print()
    bfs := &BFS {
        g: gr,
        processed: make(map[NodeID]bool),
        processQ: NewNodeList(),
        searchPath: NewNodeList(),
    }
    result := bfs.Search(start, item)
    if result {
        fmt.Printf("Path between %d and %d\n", start, item)
        bfs.searchPath.Print()
    } else {
        fmt.Printf("No path between %d and %d\n", start, item)
    }
}

func dfs_topological_sort_wrapper(gr *Graph, start NodeID) {
    fmt.Println("====== DFS ========")
    gr.Print()
    dfs := &DFS {
        g: gr,
        processed: make(map[NodeID]bool),
        topoSortStack: NewNodeStack(),
    }
    dfs.Walk(start)
    fmt.Printf("Topological Sort starting %d\n", start)
    dfs.topoSortStack.Print()
}

func test_graph() {
    test_bfs_item()
}

func test_bfs_item() {
	g1 := NewGraph()
	g1.AddEdge(1, 2, 0)
	g1.AddEdge(1, 5, 0)
	g1.AddEdge(1, 6, 0)
	g1.AddEdge(2, 3, 0)
	g1.AddEdge(2, 5, 0)
	g1.AddEdge(3, 4, 0)
	g1.AddEdge(5, 4, 0)

	// BFS search
    bfs_search_wrapper(g1, 1, 4)
    bfs_search_wrapper(g1, 1, 8)
    bfs_search_wrapper(g1, 2, 4)

    // DFS topological sort
    dfs_topological_sort_wrapper(g1, 1)
}
