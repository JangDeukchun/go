package main

func main() {
	// 변수는 go 키워드 var 를 사용하여 선언한다. var 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다. 예를 들어 아래는 a라는 정수 변수를 선언한 것이다.
	var a int

	// 변수 선언문에서 변수 초깃값을 할당할 수도 있다. 즉, float32 타입의 변수 f 에 11.0 이라는 초기값을 할당하기 위해서는 아래와 같이 쓸 수 있다.
	var f float32 = 11.

	// 일단 선언된 변수는 그 뒤의 문장에서 해당 타입의 값을 할당할 수 있다.
	a = 10
	f = 12.0

	// 만약 선언된 변수가 go 프로그램 내에서 사용되지 않는다면, 에러를 발생시킨다. 따라서 사용되지 않는 변수는 프로그램에서 삭제된더, 동일한 타입의 변수가 복수 개 있을 경우, 변수들을 나열하고 마지막에 타입을 한번만 지정할 수 있다.
	var i, j, k int

	// f복수 변수들이 선언된 상황에서 초기값을 지정할 수 있다. 초기값은 순서대로 변수에 할당된다. 예를 들어, 아래 코드의 경우는 isms 1, j는 2, k는 3이 할당된다.
	var i, j, k int = 1, 2, 3

	// 변수를 선언하면서 초기값을 지정하지 않으면, go는 zero value를 기본적으로 할당한다. 즉 숫자형에는 0, bool 타입에는 false, 그리고 string 형에는 "" 빈문자열을 할당한다.
	// 변수를 선언하는 또 다른 방식으로는 short Assignment Statement (:=) 를 사용할 수 있다. 즉, var i = 1 을 쓰는 대신 i:=1 이라고 var를 생략하고 사용할 수 있다. 하지만 이러한 표현은 함수 내에서만 사용할 수 있다. 함수 밖에서는 var를 사용해야 한다.

	// 상수는 go 키워드 const를 사용하여 선언한다 const 키워드 뒤에 상수명을 적고, 그 뒤에 상수타입, 그리고 상수값을 할당한다.
	const c int = 10
	const s string = "hi"

	// go에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용된다. 또 여러개의 상수를 묶어서 지정할 수 있다.
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "Amerclan Express"
	)

	// 상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다. 이 경우 iota가 지정된 apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.
}
