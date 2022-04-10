package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) *bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.00,
	}

	return &b
}

func (b *bill) addItem(menu string, amount float64) {
	b.items[menu] = amount
	b.tip = 2.09
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) format() string {
	f := ""
	total := 0.

	f += fmt.Sprintf("Breakdown of bill (%v) \n", b.name)

	for k, v := range b.items {
		f += fmt.Sprintf("%-20v $%v \n", k+":", v)
		total += v
	}

	f += fmt.Sprintf("%-20v $%.2f \n", "Tip:", b.tip)
	f += fmt.Sprintf("%-20v $%.2f \n", "Total:", total+b.tip)
	return f
}

func (b bill) save() {

	data := []byte(b.format())
	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
