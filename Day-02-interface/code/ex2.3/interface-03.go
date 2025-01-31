package main

import (
	"interface/dhl"
	"interface/fedex"
)

// ❶ Sender 인터페이스를 만들었습니다.
type Sender interface {
	Send(parcel string)
}

// ❷ Sender 인터페이스를 입력으로 받습니다.
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// ❸ DHL 전송 객체, Fedex 전송 객체 모두 SendBook 인수로 사용할 수 있습니다.
	// DHL 전송 객체를 만듭니다.
	dhlSender := &dhl.DhlSender{}
	SendBook("어린 왕자", dhlSender)
	SendBook("그리스인 조르바", dhlSender)

	// Fedex 전송 객체를 만듭니다.
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}
