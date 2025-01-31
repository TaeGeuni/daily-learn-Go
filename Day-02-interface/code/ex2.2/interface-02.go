package main

import "interface/fedex"

//DHL 로 바꾸려면 기존 코드를 바꿔야함
//한가지가 바뀌는데 코드 여러곳을 바꿔야함
//이것을 해결하지위해 인터페이스를 사용

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := fedex.FedexSender{}

	SendBook("어린왕자", &sender)
	SendBook("그리스인 조르바", &sender)
}
