package main

/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] 连通网络的操作次数
 */

// Node 定义图上的节点
type Node struct {
	id      int     // 节点在图中的id
	visited bool    // 该节点在遍历过程中是否访问过
	nexts   []*Node // 节点的邻边
}

// @lc code=start
func makeConnected(n int, connections [][]int) int {

	graph := bulidGraph(n, connections)

	validConn, group := bfs(graph)
	redundancy := len(connections) - validConn
	if redundancy >= group-1 {
		return group - 1
	} else {
		return -1
	}
}

func bulidGraph(n int, connections [][]int) []*Node {
	graph := make([]*Node, n)
	for i := 0; i < n; i++ {
		graph[i] = &Node{
			id:    i,
			nexts: make([]*Node, 0),
		}
	}

	for _, connection := range connections {
		graph[connection[0]].nexts = append(graph[connection[0]].nexts, graph[connection[1]])
		// 因为是无向图，索引反过来的边也要加上
		graph[connection[1]].nexts = append(graph[connection[1]].nexts, graph[connection[0]])
	}

	return graph
}

func bfs(graph []*Node) (int, int) {
	validConn, group := 0, 0
	for i := 0; i < len(graph); i++ {
		// 如果当前节点没有被访问过，说明它是在一个新的group里
		if graph[i].visited == false {
			group++
		} else {
			continue
		}

		// 经典的bfs
		q := make([]*Node, 0)
		q = append(q, graph[i])
		// 只要进入队列，就标记为访问过
		graph[i].visited = true
		for len(q) > 0 {

			poll := q[0]
			q = q[1:]
			for _, node := range poll.nexts {
				if node.visited == false {
					q = append(q, node)
					// 只要进入队列，就标记为访问过
					node.visited = true

					// 如果是首次访问过，说明这条connection是有效的，需要累计起来
					validConn++

				}
			}
		}
	}
	return validConn, group
}

// @lc code=end
