package main

import (
	"fmt"
	"reflect"

	"github.com/WeiWeiWesley/echo"
)

type MyInt int

type MyFloat64 float64

type Array[v comparable] struct {
	Data []v
}

func main() {

	//basic usage
	basicExample()

	//Search in Array
	// searchInArr()

	//bubble sort example
	// bubbleSortExample()

	//workspace example
	// workExample()

}

func Add[T ~int | int16 | float64](a, b T) T {
	fmt.Printf("%v, ", reflect.TypeOf(a))

	return a + b
}

func StrInArray(arr []string, find string) (bool, string) {

	for i := range arr {
		if arr[i] == find {
			return true, find
		}
	}

	return false, ""
}

func InArray[T comparable](arr []T, find T) (ok bool, result T) {

	for i := range arr {
		if arr[i] == find {
			ok = true
			result = find
			return
		}
	}

	return
}

func (a *Array[v]) InArray(find v) (bool, v) {

	for i := range a.Data {
		if a.Data[i] == find {
			return true, find
		}
	}

	return false, find
}

//泛型基本範例
func basicExample() {
	//int
	fmt.Println(Add(3, 5))

	//int16
	fmt.Println(Add[int16](3, 5))

	//float64
	fmt.Println(Add(3.3, 5.5))

	//~int
	a := MyInt(3)
	b := MyInt(5)
	fmt.Println(Add(a, b))

	//!float64
	// fmt.Println(Add[MyFloat64](3.3, 5.5))
}

//泛型基本範例2
func searchInArr() {

	fmt.Println(InArray([]string{"x", "y", "z"}, "P"))

	fmt.Println(InArray([]int{1, 2, 3}, 3))

	intArr := Array[int]{
		Data: []int{1, 2, 3},
	}

	strArr := Array[string]{
		Data: []string{"a", "b", "c"},
	}

	fmt.Println(intArr.InArray(1))

	fmt.Println(strArr.InArray("e"))
}

//work module 範例
func workExample() {
	echo.Echo()
}
