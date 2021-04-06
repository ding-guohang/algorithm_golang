package main

/*
17.22. 单词转换
给定字典中的两个词，长度相等。写一个方法，把一个词转换成另一个词， 但是一次只能改变一个字符。每一步得到的新词都必须能在字典中找到。

编写一个程序，返回一个可能的转换序列。如有多个可能的转换序列，你可以返回任何一个。

示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
["hit","hot","dot","lot","log","cog"]
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: []

解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
*/

// 穷举吧..
func findLadders(beginWord string, endWord string, wordList []string) []string {
	if !contains(endWord, wordList) {
		return []string{}
	}

	wordList = removeDuplicate(beginWord, wordList)
	worstWords := make(map[string]bool)

	for idx, word := range wordList {
		if worstWords[word] {
			continue
		}
		if !canBeNext(beginWord, word) {
			continue
		}
		ret, found := doFind(word, endWord, remove(wordList, idx), []string{beginWord}, worstWords)
		if found {
			return ret
		}
	}

	return []string{}
}

func removeDuplicate(duplicateWord string, words []string) []string {
	m := make(map[string]struct{}, len(words))
	for _, w := range words {
		if duplicateWord == w {
			continue
		}

		m[w] = struct{}{}
	}
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func doFind(thisWord, endWord string, wordList, ret []string, worstWords map[string]bool) ([]string, bool) {
	ret = append(ret, thisWord)

	if thisWord == endWord {
		return ret, true
	}

	for idx, word := range wordList {
		if worstWords[word] {
			continue
		}
		if !canBeNext(thisWord, word) {
			continue
		}
		r, found := doFind(word, endWord, remove(wordList, idx), ret, worstWords)
		if found {
			return r, true
		}
	}

	// 说明这个单词到不了终点
	worstWords[thisWord] = true
	return ret, false
}

func remove(wordList []string, removeIdx int) []string {
	return append(wordList[:removeIdx], wordList[removeIdx+1:]...)
}

func contains(word string, wordList []string) bool {
	for _, x := range wordList {
		if x == word {
			return true
		}
	}
	return false
}

func canBeNext(a, b string) bool {
	differentCount := 0
	for idx, c := range a {
		if differentCount > 1 {
			break
		}

		if int(c) != int(b[idx]) {
			differentCount++
			continue
		}
	}
	return differentCount <= 1
}
