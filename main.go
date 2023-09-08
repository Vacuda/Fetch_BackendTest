package main

import (
	"fmt"
	"time"
)

/* STRUCTS */
type Receipt struct{
	Retailer		string
	PurchaseDate	time.Time
	PurchaseTime	time.Time
	Total			float32
	Items			[]Item
}

type Item struct{
	ShortDescription	string
	Price				float32
}

/* MAIN */

func main(){
	var message string = "Hello, World."
	fmt.Println(message)
	time.Now()
}

/* UTILITIES */

