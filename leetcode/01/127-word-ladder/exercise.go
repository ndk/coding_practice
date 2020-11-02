package exercise

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)

	end := -1

	mutants := map[string][]int{}
	for i, word := range wordList {
		if word == endWord {
			end = i
		}
		for p := range word {
			mutant := word[:p] + "*" + word[p+1:]
			mutants[mutant] = append(mutants[mutant], i)
		}
	}
	graph := map[int]map[int]struct{}{}
	for _, group := range mutants {
		for i := 0; i < len(group); i++ {
			for j := 0; j < len(group); j++ {
				if i == j {
					continue
				}
				source, target := group[i], group[j]
				targets, ok := graph[source]
				if !ok {
					targets = map[int]struct{}{}
				}
				targets[target] = struct{}{}
				graph[source] = targets
			}
		}
	}

	queue := []int{len(wordList) - 1}
	visited := map[int]struct{}{
		queue[0]: {},
	}
	steps := 0
	for len(queue) > 0 {
		steps++
		for size := len(queue); size > 0; size-- {
			current := queue[0]
			if current == end {
				return steps
			}
			queue = queue[1:]
			for next := range graph[current] {
				if _, ok := visited[next]; ok {
					continue
				}
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}

	return 0
}

///////////////////////////////////////////////////////////////////////////////
// The two best solutions on Leetcode at the moment

func ladderLength_leet1(beginWord string, endWord string, wordList []string) int {
	startQueue := list.New()
	endQueue := list.New()
	wordsIdx := make(map[string]int)
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))
	step := 0

	for i, v := range wordList {
		wordsIdx[v] = i
	}

	if idx, ok := wordsIdx[endWord]; ok {
		endUsed[idx] = true
	} else {
		return 0
	}

	startQueue.PushBack(beginWord)
	endQueue.PushBack(endWord)

	for startQueue.Len() > 0 {
		step++
		l := startQueue.Len()
		for i := 0; i < l; i = i + 1 {
			wordBytes := []byte(startQueue.Remove(startQueue.Front()).(string))
			for j := 0; j < len(wordBytes); j++ {
				o := wordBytes[j]
				for c := byte('a'); c <= 'z'; c++ {
					if c != o {
						wordBytes[j] = c
						if idx, ok := wordsIdx[string(wordBytes)]; ok {
							if endUsed[idx] {
								return step + 1
							} else {
								if !startUsed[idx] {
									startQueue.PushBack(wordList[idx])
									startUsed[idx] = true
								}
							}
						}
					}
				}
				wordBytes[j] = o
			}
		}
		if startQueue.Len() > endQueue.Len() {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}

func ladderLength_leet2(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]struct{})
	for _, word := range wordList {
		dict[word] = struct{}{}
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	beginSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	beginSet[beginWord] = struct{}{}
	endSet[endWord] = struct{}{}
	ladder := 1
	for len(beginSet) != 0 && len(endSet) != 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		nextSet := make(map[string]struct{})
		for word := range beginSet {
			for i := 0; i < len(word); i++ {
				for c := byte('a'); c <= byte('z'); c++ {
					next := []byte(word)
					next[i] = c
					if _, ok := endSet[string(next)]; ok {
						return ladder + 1
					}
					if _, ok := dict[string(next)]; ok {
						delete(dict, string(next))
						nextSet[string(next)] = struct{}{}
					}
				}
			}
		}
		ladder++
		beginSet = nextSet
	}
	return 0
}
