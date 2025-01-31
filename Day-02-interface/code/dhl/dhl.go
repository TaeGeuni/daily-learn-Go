package dhl

import "fmt"

type DhlSender struct{}

func (d *DhlSender) Send(parcel string) {
	fmt.Printf("DHL에서 택배 %s를 보냅니다.\n", parcel)
}
