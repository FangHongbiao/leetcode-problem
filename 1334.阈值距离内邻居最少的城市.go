import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=1334 lang=golang
 *
 * [1334] 阈值距离内邻居最少的城市
 */

// @lc code=start

const INF int = 0xffff

func findTheCity1(n int, edges [][]int, distanceThreshold int) int {
	// 获取任意两点之间的最短路径信息
	d := floyd(n, edges)

	fmt.Println(d)
	ans := 0
	cand := n
	for i := 0; i < n; i++ {
		// 统计邻居个数
		cnt := 0
		for j := 0; j < n; j++ {
			if d[i][j] != INF && d[i][j] <= distanceThreshold {
				cnt++
			}
		}
		// 检查是否是当前最小的
		if cnt <= cand {
			cand = cnt
			ans = i
		}
	}
	return ans
}

func floyd(n int, edges [][]int) [][]int {

	// 初始化dp 数组
	d := make([][]int, n)

	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
	}

	// 初始化为 INF， 表示不连通. 注意: i->i, 要初始化为0，自己到自己无代价
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = INF
			}
		}
	}

	for _, edge := range edges {
		// 要当成无向图
		d[edge[0]][edge[1]] = edge[2]
		d[edge[1]][edge[0]] = edge[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}

	return d
}

// Dijkstra

// An edge is something we manage in a priority queue.
type edge struct {
	node   int
	weight int
	index  int
}

// A PriorityQueue implements heap.Interface and holds edges.
type PriorityQueue []*edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].weight > pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	edge := x.(*edge)
	edge.index = n
	*pq = append(*pq, edge)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	old[n-1] = nil  // avoid memory leak
	edge.index = -1 // for safety
	*pq = old[0 : n-1]
	return edge
}

// update modifies the priority and value of an edge in the queue.
func (pq *PriorityQueue) update(edge *edge, node int, weight int) {
	edge.node = node
	edge.weight = weight
	heap.Fix(pq, edge.index)
}

type graph struct {
	nodes map[int][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[int][]edge)}
}

func (g *graph) addEdge(origin int, destiny int, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node int) []edge {
	return g.nodes[node]
}

func (g *graph) getEdgeWeight(src int, dst int) int {
	edges := g.nodes[src]
	for edge := edges {
		if edge.node == dst {
			return edge.weight
		}
	}
	return INF
}

func newHeap(n int) PriorityQueue {
	return make(PriorityQueue, n)
}

func (g *graph) dijkstra(origin int, n int) []int {
	visited := make(map[int]bool)
	dist := make([]int, n)
	ans := make([]int, n)
	dist[origin] = 0

	h := newHeap(n)

	heap.Init(&h)

	for h.Len() > 0 {
		// Find the nearest yet to visit node
		u := heap.Pop(&h).(edge)

		if visited[u.node] {
			continue
		}

		ans[u.node] = u.weight
		for _, e := range g.getEdges(u.node) {

			v := e.node
			alt := dist[u.node] + graph.getEdgeWeight(u.node, v.node)
			if alt < dist[v.node] {
				if !visited[e.node] { 
					dist[v.node] = alt
					h.update(e, e.node, alt)
				}
				
			}
		}
	}

	return ans
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	graph := newGraph()
	for _, edge := range edges {
		// 要当成无向图
		graph.addEdge(edge[0], edge[1], edge[2])
		graph.addEdge(edge[1], edge[0], edge[2])
	}

	ans := 0
	cand := n
	for i := 0; i < n; i++ {
		// 统计邻居个数
		neigs := graph.dijkstra(i, n)
		for _, neig := range neigs {
			if neig != INF {
				cnt++
			}
		}
		// 检查是否是当前最小的
		if cnt <= cand {
			cand = cnt
			ans = i
		}
	}
	return ans
}

// @lc code=end
