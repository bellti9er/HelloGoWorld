package main

import "fmt"

func main() {
	//
	// var name string = "jongbeom"  // 또는 name := "jongbeom"
	var age int = 26 // 또는 age := 26

	const pi float64 = 3.14

	// name = "jongbeom"
	age = 26

	// if-else
	if age > 18 { 
		fmt.Println("Adult")
	} else { 
		fmt.Println("Minor")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// learning point

/* 
1. Go는 변수를 선언해놓고 사용하지 않으면 컴파일 시 error raise, 아마 코드를 깔끔하고 효율성있게 유지할 수 있을 것 같다.

2. Go 또한 정적 타입언어로 변수를 선언할 때 타입을 지정할 수 있어보임. 
다만 타입을 지정하지 않고 := 를 사용하여 변수를 선언하면, 컴파일러가 초기화 값의 타입을 추론하기 떄문에 가급적 명시하는 것이 좋을 것 같다.
*/