package labs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ID int

type Graph struct {
	Edges map[ID]Edge
}

type Edge struct {
	From        []ID
	Coordinates []int
	To          []ID
	Direction   []string
	Rib         int
	Weight      int
}

func nextPoints(coords []int, fromId ID, size []int, field [][]string) [][]int {
	var points [][]int
	if coords[0]-1 >= 0 {
		switch field[coords[0]-1][coords[1]] {
		case ".":
			points = append(points, []int{coords[0] - 1, coords[1], int(fromId), 1, 8})
		case "W":
			points = append(points, []int{coords[0] - 1, coords[1], int(fromId), 2, 8})
		}
	}
	if coords[0]+1 < size[0] {
		switch field[coords[0]+1][coords[1]] {
		case ".":
			points = append(points, []int{coords[0] + 1, coords[1], int(fromId), 1, 2})
		case "W":
			points = append(points, []int{coords[0] + 1, coords[1], int(fromId), 2, 2})
		}
	}
	if coords[1]+1 < size[1] {
		switch field[coords[0]][coords[1]+1] {
		case ".":
			points = append(points, []int{coords[0], coords[1] + 1, int(fromId), 1, 6})
		case "W":
			points = append(points, []int{coords[0], coords[1] + 1, int(fromId), 2, 6})
		}
	}
	if coords[1]-1 >= 0 {
		switch field[coords[0]][coords[1]-1] {
		case ".":
			points = append(points, []int{coords[0], coords[1] - 1, int(fromId), 1, 4})
		case "W":
			points = append(points, []int{coords[0], coords[1] - 1, int(fromId), 2, 4})
		}
	}
	return points
}

func (graph *Graph) addEdge(point []int) {
	dict := map[int]string{
		8: "N",
		2: "S",
		4: "W",
		6: "E",
	}
	var edge Edge
	edge.Coordinates = []int{point[0], point[1]}
	edge.From = append(edge.From, ID(point[2]))
	edge.Rib = point[3]
	edge.Direction = append(edge.Direction, dict[point[4]])
	edge.Weight = -1
	graph.Edges[ID(len(graph.Edges))] = edge
}

func searchVisitedPoints(elemId ID, array []ID) bool {
	next := true
	for _, id := range array {
		if elemId == id {
			next = false
		}
	}
	return next
}

func searchCreatedEdges(point []int, graph Graph) bool {
	var edge Edge
	next := true
	for id, elem := range graph.Edges {
		edge = graph.Edges[id]
		if elem.Coordinates[0] == point[0] && elem.Coordinates[1] == point[1] {
			if id > ID(point[2]) {
				edge.From = append(edge.From, ID(point[2]))
			} else {
				edge.To = append(edge.To, ID(point[2]))
			}
			graph.Edges[id] = edge
			next = false
		}
	}
	return next
}

func (graph *Graph) DFS(startPoint ID, visited *[]bool) {
	(*visited)[startPoint] = true
	var edge Edge
	for _, id := range graph.Edges[startPoint].To {
		if !(*visited)[id] {

		}
		if graph.Edges[id].Weight == 0 {
			edge = graph.Edges[id]
			edge.Weight += graph.Edges[startPoint].Weight
			graph.Edges[id] = edge
		}
	}
	return
}

func Lab4C() {
	inputData := strings.Split(scanFile("src/Lab4C.txt"), ";")
	inputData = inputData[:len(inputData)-1]
	infoRowString := strings.Split(inputData[0], " ")
	var infoRow []int
	for _, elem := range infoRowString {
		numberToAdd, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		infoRow = append(infoRow, numberToAdd)
	}
	fieldSize, civilCoordinates, pointCoordinates :=
		[]int{infoRow[0], infoRow[1]}, []int{infoRow[2], infoRow[3]}, []int{infoRow[4], infoRow[5]}

	inputData = inputData[1:]
	var gameField [][]string
	for i := 0; i < fieldSize[0]; i++ {
		row := strings.Split(inputData[i], "")
		gameField = append(gameField, row)
	}

	graph := Graph{make(map[ID]Edge)}
	graph.Edges[0] = Edge{[]ID{}, []int{civilCoordinates[0] - 1, civilCoordinates[1] - 1},
		[]ID{}, []string{}, 0, -1}

	var visited []ID
	for len(graph.Edges) != len(visited) {
		for id, elem := range graph.Edges {
			if searchVisitedPoints(id, visited) {
				visited = append(visited, id)
				points := nextPoints(elem.Coordinates, id, fieldSize, gameField)
				for _, point := range points {
					if searchCreatedEdges(point, graph) {
						graph.addEdge(point)
					}
				}
			}
		}
	}

	var keys []int
	for key, _ := range graph.Edges {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)
	for key := range keys {
		fmt.Println(key, graph.Edges[ID(key)])
	}

	var weightsOfEdges []int
	var dfsVisited []bool

	for i := 0; i < len(graph.Edges); i++ {
		dfsVisited = append(dfsVisited, false)
		weightsOfEdges = append(weightsOfEdges, 0)
	}

	dfsVisited[0] = true
	for _, id := range graph.Edges[0].To {
		if !dfsVisited[id] && (graph.Edges[id].Weight == -1 || graph.Edges[id].Weight < -1) {

		}
	}

	fmt.Println(dfsVisited)

	fmt.Println(pointCoordinates)
}
