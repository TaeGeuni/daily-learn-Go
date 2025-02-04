//Go 1.22 버전부터 루프 변수(i)의 스코프가 변경되었기 때문에 해당사항 없음
//두 함수 모두 같은 결과

package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("CaptureLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
