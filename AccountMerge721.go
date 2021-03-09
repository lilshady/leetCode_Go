package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {

	result := make([][]string, 0)
	graph, reverseIndex := buildGraph(accounts)

	visited := make(map[string]bool)

	for key := range graph {

		if _, ok := visited[key]; ok {
			continue
		}
		visited[key] = true
		emails := make([]string, 0)
		dfs(&graph, key, &visited, &emails)
		sort.Strings(emails)
		emails = append([]string{reverseIndex[key]}, emails...)
		result = append(result, emails)
	}
	return result
}

func dfs(graph *map[string]map[string]bool, key string, visited *map[string]bool, emails *[]string) {
	*emails = append(*emails, key)
	for neighbor := range (*graph)[key] {

		if _, ok := (*visited)[neighbor]; ok {
			continue
		}
		(*visited)[neighbor] = true
		dfs(graph, neighbor, visited, emails)
	}
}

func buildGraph(accounts [][]string) (map[string]map[string]bool, map[string]string) {

	graph := make(map[string]map[string]bool) // email -> neighbor nodes
	reverseIndex := make(map[string]string)
	for _, account := range accounts {
		for i, s := range account {
			if i == 0 {
				continue
			}
			reverseIndex[s] = account[0]
			if _ , ok := graph[s]; !ok {
				graph[s] = make(map[string]bool)
			}
			if i > 1 {
				graph[account[i - 1]][s] = true
				graph[s][account[i - 1]] = true
			}
		}
	}
	return graph, reverseIndex
}


func main() {
	result := accountsMerge([][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},{"John", "johnnybravo@mail.com"},{"Mary", "mary@mail.com"},{"John", "johnsmith@mail.com", "john_newyork@mail.com"}})
    fmt.Printf("the result is %+v", result)
}