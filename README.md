Go1.18
===

## 泛型

- Method 宣告
    ```go
    type MyInt int

    type MyFloat64 float64

    func Add[T ~int | int16 | float64](a, b T) T {
        return a + b
    }
    ```
    於函式可宣告 `T` 可接受的型態種類，如 `Add` 宣告型別 `T` 為 `int` 或 `float64`，且輸入的參數 a, b 型態為 `T`。波浪符號 `~` 用以表示泛型，表示可接受所有類似的基礎型別。以上例來說，`MyInt` 可被 `Add` 方法接受，`MyFloat64` 則否

- 呼叫方式
    ```go
    func main() {
        Add(3, 5)

        Add[int](3, 5)

        Add[int16](3, 5)

        Add[MyInt](3, 5)
    }
    ```
    可於方法後指定數入參數型別，但可以省略

- comparable
    ```go
    // comparable is an interface that is implemented by all comparable types
    // (booleans, numbers, strings, pointers, channels, arrays of comparable types,
    // structs whose fields are all comparable types).
    // The comparable interface may only be used as a type parameter constraint,
    // not as the type of a variable.
    type comparable interface{ comparable }
    ```
    介面 comparable 包含了各種可比較的基礎型別，用於宣告 struct & method 可接受型別。不能用於變數型別宣告

- Struct 宣告
    ```go
    type Array[v comparable] struct {
        Data []v
    }

    intArr := Array[int]{
		Data: []int{1, 2, 3},
	}

    strArr := Array[string]{
		Data: []string{"a", "b", "c"},
	}
    ```
    自定義泛型 struct 須於實體化時，指定基礎型別方可編譯

- Benchmark

    以[泡泡排序](./bubble_sort.go)為例，泛型的程式碼將可以優於使用 sort.Interface 版本，且極近似於指定型別效能。由 benchmark 結果可以看出，泛型除了可以簡化我們的程式碼外，在部分使用時機還可以有效的提升效能。（詳盡的效能解說可參考[Faster sorting with Go generics](https://eli.thegreenplace.net/2022/faster-sorting-with-go-generics/)）

    ```
    goarch: amd64
    cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
    BenchmarkBubbleSortInterface-4           7215160               161.9 ns/op
    BenchmarkBubbleSortGenerics-4           38881078                30.31 ns/op
    BenchmarkBubbleSortInt-4                37684952                30.22 ns/op
    PASS
    ok      command-line-arguments  3.753s
    ```

## 測試

- 泛型測試：
單元測試時由於匿名函數與匿名 struct 無法宣告泛型，會增加泛型方法單元測試上的困難，實作上建議額外寫一個小方法，來簡化整個單元測試所需要的程式碼。請參考 [main_test.go](./main_test.go#TestAdd2)

- 模糊測試：
一般單元測試採取，比較期望值(want)與實際結果(result)的方式進行，但人們能提供的測試案例有限，且在輸入測試案例時，難免發生案例本身錯誤的狀況。現在我們可以利用 Fuzzing 協助我們大量產生隨機測試參數，可一定程度增加測試的可靠性，測試方法採用 `f.Fuzz()`，參考 [main_test.go](./main_test.go#FuzzAdd)。 ，並於測試時加入參數 `-fuzz=Fuzz` 即可，如遇測試無法通過案例，會於 `./testdata` 產生所使用的 cases 紀錄，並且下次執行模糊測試時會優先採用此紀錄測試

    ```
    go test -v -run=FuzzAdd -fuzz=Fuzz *.go

    === FUZZ  FuzzAdd
    fuzz: elapsed: 0s, gathering baseline coverage: 0/4 completed
    fuzz: elapsed: 0s, gathering baseline coverage: 4/4 completed, now fuzzing with 4 workers
    fuzz: elapsed: 0s, execs: 58 (2037/sec), new interesting: 0 (total: 4)
    --- FAIL: FuzzAdd (0.03s)
        --- FAIL: FuzzAdd (0.00s)
            main_test.go:84: 203 + 4 expect=207 but get -1

        Failing input written to testdata/fuzz/FuzzAdd/2bc1e5ee96bcb54041df85910c132ab440c3c097af61aaa2143e01394188dc60
        To re-run:
        go test -run=FuzzAdd/2bc1e5ee96bcb54041df85910c132ab440c3c097af61aaa2143e01394188dc60
    FAIL
    exit status 1
    FAIL    command-line-arguments  0.037s
    ```

    ```
    tree ./testdata

    ./testdata
    └── fuzz
        └── FuzzAdd
            └── bebf82f9b04feaa1ea1a9d540f3032401ff90347424a30058ed433532eddfe82
    ```


## Workspaces

- go work

    Go1.18 前如果我們需要將特定 module 於本地替換的話，我們可能需要於 `go.mod` replace package path，但這樣的缺點是每當要推回遠端時，必須記得把 replace 的路徑拿掉，而現在我們有了 `go work` 可以替我們解決這個困擾。僅需將想取代的 modules 加入 `go.work`，並且 `echo 'go.work' >> .gitignore` 避免將此設定上到 remote repo 即可

    - Before go1.18

        ```go
        module main

        go 1.18

        require github.com/WeiWeiWesley/echo v0.0.0-20220512090346-ea286e617d40

        //go.mod
        replace (
            github.com/WeiWeiWesley/echo => ./echo
        )
        ```

    - After go1.18

        ```
        go work init ./echo
        ```

        ```go
        go 1.18

        use ./echo
        ```

