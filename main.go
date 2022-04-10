package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		p, _ := getInput("Item price: ", reader)

		price, err := strconv.ParseFloat(p, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, price)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "t":
		t, _ := getInput("Enter tip amount ($): ", reader)

		tip, err := strconv.ParseFloat(t, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(tip)
		fmt.Println("Tip updated  ", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You chose to save the bill - ", b.name)
		fmt.Println(b)

	default:
		fmt.Println("that was not a valid option")
		promptOptions(b)
	}
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("    ✓ Bill has been created - ", b.name)

	return b
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
