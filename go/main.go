package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

// 文字列を整数の配列に変換する関数
func convertStringToIntArray(m string) []int {
	// 中身を埋める
	arr1 := strings.Split(m, ",")
	arr2 := make([]int, len(arr1))
	for n := range arr1 {
		arr2[n], _ = strconv.Atoi(arr1[n])
	}
	return arr2
}

// 各整数の出現回数を数える関数
func countNumberFrequency(a []int) map[int]int {
	// 中身を埋める
	m := make(map[int]int)
	for _, n := range a {
		m[n] += 1
	}
	return map1
}

// 整数の合計をsにする方法が何通りあるかを数える関数
func countCardCombinations(a []int, s int) int {
	// 中身を埋める
	Answer := 0
	var fact func(i int, sum_left int)

	fact = func(i int, sum_left int) {
		if i == len(a) {
			if sum_left == 0 {
				Answer++
			}
			return
		}

		// a[i]を選択する場合
		if sum_left >= a[i] {
			fact(i+1, sum_left-a[i])
		}

		// a[i]を選択しない場合
		fact(i+1, sum_left)
	}

	fact(0, s)

	if Answer == 0 {
		return -1
	} else {
		return Answer
	}
}

// map[int]intのkeyとvalueをkeyの昇順に出力する関数
func printMapKeyAndValue(m map[int]int) {
	// 中身を埋める
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func main() {
	var m string
	var s int

	fmt.Scan(&m)
	fmt.Scan(&s)

	a := convertStringToIntArray(m)
	frequencyCount := countNumberFrequency(a)
	combinationCount := countCardCombinations(a, s)

	printMapKeyAndValue(frequencyCount)
	fmt.Println(combinationCount)
}