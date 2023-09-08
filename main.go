package main

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
	"log"
)

/* STRUCTS */
type Item struct{
	ShortDescription	string
	Price				float32
}

type Receipt struct{
	Retailer 		string 		`json:"retailer"`	
	PurchaseDate	time.Time	`json:"purchaseDate"`	
	PurchaseTime	time.Time	`json:"purchaseTime"`	
	Total			float32		`json:"total"`	
	Items			[]Item		`json:"items"`	
}


/* DATA */

var Receipts []Receipt = make([]Receipt, 0)

/* MAIN */

func main(){

	var this_receipt Receipt

	lava, err := os.ReadFile("examples/morning-receipt.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	json.Unmarshal(lava, &this_receipt)

	var message string = this_receipt.Retailer
	fmt.Println(message)
}

/* UTILITIES */

