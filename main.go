package main

import "fmt"

type Report interface {
	Report() string
}

type MarketingReport struct {
	report string
}

func (r *MarketingReport) Report() string {
	return r.report
}

func ReportSend(report Report, email string) {
	fmt.Printf("%s 여기로 보냈습니다\n", email)
}

func main() {
	email := "btg0805@example.com"
	var data MarketingReport
	data.report = "안녕하세요"

	ReportSend(&data, email)
}
