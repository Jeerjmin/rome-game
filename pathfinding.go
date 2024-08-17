package main

import (
	"container/heap"
	"math"
)

// Position - структура для хранения координат на сетке
type Position struct {
	X, Y int
}

// Node - узел для A*
type Node struct {
	Position
	F, G, H int
	Parent  *Node
}

// PriorityQueue - очередь с приоритетом для A*
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// FindPath - функция для поиска пути с использованием A*
func FindPath(start, goal Position, gridWidth, gridHeight int) []Position {
	pq := &PriorityQueue{}
	heap.Init(pq)

	startNode := &Node{Position: start, G: 0, H: heuristic(start, goal)}
	startNode.F = startNode.G + startNode.H
	heap.Push(pq, startNode)

	cameFrom := make(map[Position]*Node)
	cameFrom[start] = startNode

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		// Если достигли цели, то восстанавливаем путь
		if current.X == goal.X && current.Y == goal.Y {
			var path []Position
			for current != nil {
				path = append([]Position{current.Position}, path...)
				current = current.Parent
			}
			return path
		}

		// Проверяем соседей
		for _, neighbor := range getNeighbors(current.Position, gridWidth, gridHeight) {
			neighborNode := &Node{Position: neighbor, G: current.G + 1, H: heuristic(neighbor, goal)}
			neighborNode.F = neighborNode.G + neighborNode.H
			neighborNode.Parent = current

			if _, ok := cameFrom[neighbor]; !ok {
				cameFrom[neighbor] = neighborNode
				heap.Push(pq, neighborNode)
			}
		}
	}

	return nil // Путь не найден
}

func heuristic(a, b Position) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func getNeighbors(pos Position, gridWidth, gridHeight int) []Position {
	var neighbors []Position

	directions := []Position{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}

	for _, dir := range directions {
		neighbor := Position{pos.X + dir.X, pos.Y + dir.Y}

		if neighbor.X >= 0 && neighbor.X < gridWidth && neighbor.Y >= 0 && neighbor.Y < gridHeight {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
