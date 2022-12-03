package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	lc, rc := 0, 0
	dfs(&res, "", n, lc, rc)
	return res
}

func dfs(res *[]string, path string, n, lc, rc int) {
	if rc > lc || rc > n || lc > n {
		return
	}
	if rc == lc && rc == n {
		*res = append(*res, path)
		return
	}
	dfs(res, path+"(", n, lc+1, rc)
	dfs(res, path+")", n, lc, rc+1)
}
