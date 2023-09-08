package main

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
	"log"
)

/* STRUCTS */

type Receipt struct{
	Retailer 		string 		`json:"retailer"`	
	PurchaseDate	time.Time	`json:"purchaseDate"`	
	PurchaseTime	time.Time	`json:"purchaseTime"`	
	Total			float32		`json:"total"`	
	Items        	[]struct {
		ShortDescription string `json:"shortDescription"`
		Price            float32 `json:"price"`
	} 	
}

type Receipt_raw struct{
	Retailer 		string 		`json:"retailer"`	
	PurchaseDate	string		`json:"purchaseDate"`	
	PurchaseTime	string		`json:"purchaseTime"`	
	Total			string		`json:"total"`	
	Items        	[]struct {
		ShortDescription string `json:"shortDescription"`
		Price            string `json:"price"`
	} 	
}

	


/* DATA */

var Receipts []*Receipt = make([]*Receipt, 0)

/* MAIN */

func main(){
	gatherReceipts_FromExamplesFolder()



	//fmt.Printf("%s", this_receipt.Items[0].Price)
	fmt.Println(Receipts[0].Retailer)

	//var message string = this_receipt.Items[0].ShortDescription
	//fmt.Println(message)
}

/* UTILITIES */

func gatherReceipts_FromExamplesFolder(){

	//read directory of examples
	file, err := os.Open("examples/")

	//err check
	if err != nil {
		log.Fatal(err)
	}

	//ensure order
	defer file.Close()

	//get filenames
	ListOf_ExampleReceipt_FileNames,_ := file.Readdirnames(0)

	//loop list of receipt file names
	for _, name := range ListOf_ExampleReceipt_FileNames {
		
		//read json file
		json_receipt, err := os.ReadFile("examples/" + name)
		
		//err check
		if err != nil {
			log.Fatal(err)
			return
		}
		
		//get raw receipt
		var this_raw_receipt Receipt_raw
		json.Unmarshal(json_receipt, &this_raw_receipt)

		//process receipt
		var this_receipt *Receipt = Process_Raw_Receipt(&this_raw_receipt)

		//put into Receipts
		Receipts = append(Receipts, this_receipt)

	}




}

func Process_Raw_Receipt(rec *Receipt_raw) (f *Receipt) {

	var receipt *Receipt = new(Receipt)

	receipt.Retailer = rec.Retailer
	

	return receipt

}

