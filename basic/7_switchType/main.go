package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Interface type의 x와 타입 T에 대하여 x.(T)로 표현했을 때,
// 이는 x가 nil이 아니며, x는 T 타입에 속한다는 점을 확인(assert)하는 것으로 이러한 표현을 "Type Assertion"이라 부른다.
// 만약 x가 nil 이거나 x의 타입이 T가 아니라면, 런타임 에러가 발생할 것이고, x가 T 타입인 경우는 T 타입의 x를 리턴한다.
// 즉, 아래 예제에서 변수 j는 a.(int)로부터 int형 변수 j가 된다.
func test1() {
	var a interface{} = 1

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	println(i) // 포인터주소 출력
	println(j) // 1 출력
}

func main() {
	do(21)
	do("hello")
	do(true)
	test1()
}
