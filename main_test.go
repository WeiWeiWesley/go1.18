package main

import (
	"reflect"
	"sort"
	"testing"
)

// 需額外定義  test cases struct
type testCase[T comparable] struct {
	name string
	a    T
	b    T
	want T
}

// 需定義執行泛型的 method
func runTestCases[T ~int | float64](t *testing.T, cases []testCase[T]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add error")
			}
		})
	}
}

// Test func
func TestAdd2(t *testing.T) {
	intTestCases := []testCase[int]{
		{
			name: "int 1",
			a:    1,
			b:    1,
			want: 2,
		},
		{
			name: "int 2",
			a:    10,
			b:    10,
			want: 20,
		},
	}

	floatCases := []testCase[float64]{
		{
			name: "float 1",
			a:    1,
			b:    2.2,
			want: 3.2,
		},
		{
			name: "float 2",
			a:    0.1,
			b:    0.5,
			want: 0.6,
		},
	}

	runTestCases(t, intTestCases)

	runTestCases(t, floatCases)
}

//FuzzInArray 模糊測試
func FuzzAdd(f *testing.F) {

	testA := map[int]int{
		1:  2,
		3:  4,
		50: 100,
		33: 66,
	}

	for a, b := range testA {
		f.Add(a, b) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, a int, b int) {

		res := Add(a, b)

		if a != res-b {
			t.Errorf("%v + %v = %v but %v != %v - %v ", a, b, res, a, res, b)
		}

	})
}

// func FuzzStrInArray(f *testing.F) {
// 	f.Fuzz(func(t *testing.T, arr []string, find string) {
// 		res, found := StrInArray(arr, find)

// 		if !res {
// 			t.Errorf("%s not found\n", find)
// 		}

// 		if found != find {
// 			t.Errorf("want %s but get %s", found, find)
// 		}
// 	})
// }

func BenchmarkBubbleSortInterface(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BubbleSortInterface(sort.IntSlice([]int{7, 3, 2, 9, 1, 3}))
	}

}

func BenchmarkBubbleSortGenerics(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BubbleSortGeneric([]int{7, 3, 2, 9, 1, 3})
	}

}

func BenchmarkBubbleSortInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BubbleSortInt([]int{7, 3, 2, 9, 1, 3})
	}

}
