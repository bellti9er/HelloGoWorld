package main

import (
	_"errors"
	"fmt"
)

// func divide(x, y float64) (float64, error) {
// 	if y == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}

// 	return x / y, nil
// }

// func main() {
// 	result, err := divide(10, 2.3)
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 		return
// 	}

// 	fmt.Println("Result: ", result)
// 	return
// }

func main() {
	_, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println("Faild to pring: ", err)
	}
}

// learning point

/*
nil은 Go에서 "zero value"를 나타내며, 포인터, 함수, 인터페이스, 슬라이스 등의 타입에 대해 초기화되지 않은 값을 의미 

var result float64
var err error
result, err = divide(10, 0) 이렇게 타입을 명시하고 값을 선언하고 초기화하는 것이 더 좋지 않을까

에러 메시지를 출력하고 종료하려면 fmt.Println과 os.Exit을 사용 할 수 있음

_는 빈 식별자라고 하며, 변수를 무시할 때 사용. 
여러 개의 반환 값을 받을 때, 특정 값이 필요하지 않다면 _로 무시할 수 있을 듯
*/