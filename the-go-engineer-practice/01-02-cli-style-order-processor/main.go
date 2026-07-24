/*
Requirements:
	1.	Parse input ("ITEM:QTY")
	2.	Validate:
	•	invalid format → error
	•	unknown product → error
	3.	Apply pricing logic (reuse your function)
	4.	Return total price

	Might need:
	type Order struct {
		Code string
		Qty  int
}
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var productPrices = map[string]float64{
	"PEN":         20.99,
	"PENCIL":      10.99,
	"ERASER":      5.0,
	"BOOK":        50.0,
	"MARKER":      25.0,
	"WHITE_BOARD": 123.0,
}

type Order struct {
	Code string
	Qty  int
}

func normalize(s string) string {
	return strings.TrimSuffix(s, "_SALE")
}

func parseOrder(input string) (Order, error) {
	// res := strings.Split(input, ":")
	res := strings.SplitN(input, ":", 2)

	if len(res) != 2 {
		return Order{}, fmt.Errorf("invalid format: %s", input)
	}

	qty, err := strconv.Atoi(res[1])

	if err != nil {
		return Order{}, fmt.Errorf("invalid number: %s", res[1])
	}

	return Order{Code: res[0], Qty: qty}, nil

}

func calculatePricing(item string) (float64, error) {
	base := normalize(item)

	price, ok := productPrices[base]

	if !ok {
		return 0, fmt.Errorf("product not found: %s", item)
	}

	if base != item {
		return price * 0.90, nil
	}

	return price, nil
}

func ProcessOrders(input []string) (float64, error) {
	var total float64

	for _, val := range input {
		// parse the input
		order, err := parseOrder(val)

		if err != nil {
			return 0, err
		}

		// apply pricing logic
		price, err := calculatePricing(order.Code)

		if err != nil {
			return 0, err
		}

		// return total price
		total += price * float64(order.Qty)
	}

	return total, nil
}

func main() {
	orders := []string{
		"PEN:2",
		"PENCIL_SALE:3",
		"BOOK:1",
		// "Random",
	}

	cost, err := ProcessOrders(orders)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Total cost is %.2f\n", cost)
}
