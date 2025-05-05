package main

type Pair struct {
	A int
	B int
}

// can use 10*first + second as a unique int key to optmize
func numEquivDominoPairs(dominoes [][]int) int {

	domMap := map[Pair]int{}
	var pair Pair
	res := 0

	for i, _ := range dominoes {

		if dominoes[i][0] > dominoes[i][1] {
			pair = Pair{dominoes[i][1], dominoes[i][0]}
		} else {
			pair = Pair{dominoes[i][0], dominoes[i][1]}
		}

		res += domMap[pair]
		domMap[pair]++
	}
	return res
}
