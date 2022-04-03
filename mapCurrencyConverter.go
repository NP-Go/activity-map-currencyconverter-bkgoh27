package main

import "fmt"

//Declare Stuct here
type currency struct {
	currencyName string
	rate         float64
}

//Declare Map here

func main() {
	// map[]
	m := make(map[string]currency)
	fmt.Println(m)
	// map[code:{currencyName rate}]
	m["USD"] = currency{"US dollar", 1.1318}
	m["USD"] = currency{"Japanese yen", 121.05}
	m["GBP"] = currency{"Pound Sterling", 0.90630}
	m["CNY"] = currency{"Chinese yuan renminbi", 7.9944}
	m["SGD"] = currency{"Singapore dollar", 1.5743}
	m["MYR"] = currency{"Malaysian ringgit", 4.8390}
	fmt.Println(m)
}
