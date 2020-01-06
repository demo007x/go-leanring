package main

import "fmt"

type Income interface {
	//  计算并返回项目的收入
	calculate() int
	// 返回项目的名称
	source() string
}

// 定义项目的结构体
type FixedBilling struct {
	projectName string // 项目名称
	biddedAmount int   // 项目中投标的金额
}

//表示项目 Time and Material。
type TimeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

type Advertisement struct {
	adName string
	CPC int
	noOfClicks int
}

func (ad Advertisement) source() string {
	return ad.adName
}

func (ad Advertisement) calculate() int {
	return ad.CPC * ad.noOfClicks
}

// 计算项目中总金额
func calculateNetIncome(ic []Income)  {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}

	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	pro1 := FixedBilling{
		projectName:  "pro1",
		biddedAmount: 5000,
	}
	pro2 := FixedBilling{
		projectName:  "pro1",
		biddedAmount: 5000,
	}
	pro3 := TimeAndMaterial{
		projectName: "pro2",
		noOfHours:   160,
		hourlyRate:  25,
	}

	ad := Advertisement{
		adName:     "ad1",
		CPC:        100,
		noOfClicks: 1000,
	}

	incomeStreams := []Income{pro1, pro2, pro3, ad}
	calculateNetIncome(incomeStreams)
}