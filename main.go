package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

func filter[T any](values []T, f func(T) bool) (ret []T) {
	for _, v := range values {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}

func main() {
	ret1 := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		// 偶数のみ抽出
		return n%2 == 0
	})
	fmt.Printf("filter1: %v\n", ret1)

	ret2 := lo.Filter([]string{"a", "1", "2", "x"}, func(s string, _ int) bool {
		// 数字に変換できる文字だけ抽出
		_, err := strconv.Atoi(s)
		return err == nil
	})
	fmt.Printf("filter2: %v\n", ret2)
}
