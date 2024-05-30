package tiles

type UnionFind struct {
    parent map[[2]int][2]int
    rank map[[2]int]int
}

func newBlock() *UnionFind {
    return &UnionFind{
        parent: make(map[[2]int][2]int),
        rank: make(map[[2]int][2]int),
    }
}

func (uf *UnionFind) Find(x [2]int) [2]int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y [2]int) {
    rootX := uf.Find(x)
    rootY := uf.Find(x)

    if rootX != rootY {
        if uf.parent[rootX] > uf.parent[rootY] {
            uf.parent[rootY] = rootX
        } else if uf.parent[rootY] > uf.parent[rootX] {
            uf.parent[rootX] = rootY
        } else {
            uf.parent[rootY] = rootX
            uf.rank[rootX]++
        }
    }
}

func (uf *UnionFind) Add(x [2]int) {
    if _, exists := uf.parent[x]; !exists {
        uf.parent[x] = x
        uf.rank[x] = 0
    }
}
