package main

import (
	"fmt"
	"sort"
)

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

func BubbleSortGeneric[T Ordered](x []T) {

	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func BubbleSortInterface(x sort.Interface) {
	n := x.Len()
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x.Less(i, i-1) {
				x.Swap(i, i-1)
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func BubbleSortInt(x []int) {

	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

//泛型泡泡排序範例
func bubbleSortExample() {

	//Interface example
	{
		intSlice := []int{9, 3, 6, 1, 7}
		floatSlice := []float64{3.3, 1.0, 9.3, 6.1}
		strSlice := []string{"d", "e", "b", "a", "c"}

		BubbleSortInterface(sort.IntSlice(intSlice))

		fmt.Println("interfae[int]", intSlice)

		BubbleSortInterface(sort.Float64Slice(floatSlice))

		fmt.Println("interfae[float]", floatSlice)

		BubbleSortInterface(sort.StringSlice(strSlice))

		fmt.Println("interfae[string]", strSlice)
	}

	fmt.Println()

	//Generics example
	{

		intSlice := []int{9, 3, 6, 1, 7}
		floatSlice := []float64{3.3, 1.0, 9.3, 6.1}
		strSlice := []string{"d", "e", "b", "a", "c"}

		BubbleSortGeneric(intSlice)

		fmt.Println("generics[int]", intSlice)

		BubbleSortGeneric(floatSlice)

		fmt.Println("generics[float]", floatSlice)

		BubbleSortGeneric(strSlice)

		fmt.Println("generics[string]", strSlice)

	}

}
