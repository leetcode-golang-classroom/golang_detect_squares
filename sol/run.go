package sol

import "fmt"

func RunDetectSquares(commands []string, points [][]int) []string {
	ds := Constructor()
	lenCommands := len(commands)
	res := []string{"null"}
	for pos := 1; pos < lenCommands; pos++ {
		switch commands[pos] {
		case "add":
			ds.Add(points[pos])
			res = append(res, "null")
		case "count":
			res = append(res, fmt.Sprintf("%d", ds.Count(points[pos])))
		}
	}
	return res
}
