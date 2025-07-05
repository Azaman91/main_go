package main

import (
	"fmt"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func getTopWords(wordMap map[string]int, n int) []string {
	// 1. Преобразуем мапу в слайс структур WordCount
	var wordCounts []WordCount
	for word, count := range wordMap {
		wordCounts = append(wordCounts, WordCount{word, count})
	}

	// 2. Сортируем слайс по убыванию частоты
	sort.Slice(wordCounts, func(i, j int) bool {
		// Если частоты равны, сортируем по алфавиту
		if wordCounts[i].Count == wordCounts[j].Count {
			return wordCounts[i].Word < wordCounts[j].Word
		}
		return wordCounts[i].Count > wordCounts[j].Count
	})

	// 3. Берем первые N элементов
	var topWords []string
	for i := 0; i < n && i < len(wordCounts); i++ {
		topWords = append(topWords, wordCounts[i].Word)
	}

	return topWords
}
func AnalyzeText(text string) {
	lower := strings.ToLower(text)
	p := strings.Map(func(r rune) rune {
		if r == '.' || r == '!' || r == '?' || r == ',' {
			return ' '
		}
		return r
	}, lower)
	words := strings.Fields(p)

	mapa := make(map[string]bool)

	for _, num := range words {
		mapa[num] = true
	}

	wordCount := make(map[string]int)
	for _, num := range words {
		wordCount[num]++
	}

	var clovo string
	max := 0
	for word, count := range wordCount {
		if count > max || (count == max && word < clovo) {
			clovo = word
			max = count
		}
	}

	fmt.Printf("Количество слов: %d\n", len(words))
	fmt.Printf("Количество уникальных слов: %d\n", len(mapa))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", clovo, max)
	fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
	top5 := getTopWords(wordCount, 5)
	for _, num := range top5 {
		fmt.Printf("\"%s\": %d раз\n", num, wordCount[num])
	}
}
