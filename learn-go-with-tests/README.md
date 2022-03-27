# [learn-go-with-tests](https://github.com/quii/learn-go-with-tests)
> - Korean: https://miryang.gitbook.io/learn-go-with-tests

- [1_integers](https://github.com/quii/learn-go-with-tests/blob/main/integers.md)
    - Go의 소스 파일은 각 폴더에 하나의 package만 가질 수 있다.
    - `ExampleAdd`
        - `godoc -http=:6060` > `curl localhost:6060/pkg`
> `go get -v golang.org/x/tools/cmd/godoc`
- [2_iteration](https://github.com/quii/learn-go-with-tests/blob/main/iteration.md)
    - for, range
    - `BenchmarkRepeat`
        - `go test -bench=.`
- [3_arrays](https://github.com/quii/learn-go-with-tests/blob/main/arrays-and-slices.md)
    - 배열, 슬라이스
    - `go test -cover`
    - `reflect.DeepEqual()`이 왜 유용하며, 왜 코드의 type-safety를 저해할 수 있는가
    - [여기 예시가 있다](https://go.dev/play/p/bTrRmYfNYCp), 배열을 자르고 슬라이스를 변경하는 것이 본래의 배열에 어떠한 영향을 주는지 알 수 있다.(하지만 복제된 슬라이스는 본래의 배열에 영향을 주지 못한다.)
    - [또 다른 예시가 있다](https://go.dev/play/p/Poth8JS28sc), 매우 큰 슬라이스를 자르고 난 후 그들을 복제해두면 좋은 이유에 대해서 알 수 있다.
- [4_structs](https://github.com/quii/learn-go-with-tests/blob/main/structs-methods-and-interfaces.md)
    - struct, method, interface
    - [테이블 기반 테스트](https://github.com/golang/go/wiki/TableDrivenTests)
        - 테이블 기반 테스트에서 특정 테스트만 실행하기: `go test -run TestArea/Rectangle`
