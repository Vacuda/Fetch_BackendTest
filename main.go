package main

import (
	//"fmt"
	//"reflect"
	"time"
	"os"
	"encoding/json"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* STRUCTS */

type Item struct {
	ShortDescription string
	Price            float32
}

type Receipt struct{
	ID 				int
	Retailer 		string
	PurchaseDate	time.Time
	PurchaseTime	time.Time
	Total			float32
	Items        	[]*Item
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

//map to hold receipts
var Receipts = make(map[int]*Receipt)

//rudementary counter to create UniqueIDs
var UniqueID_Counter int = 1

/* MAIN */

func main(){

	ConfigureAndRun_LocalServer()

}

/* SERVER INFORMATION */

func ConfigureAndRun_LocalServer(){

	//use gin to create a router
	router := gin.Default()

	//ENDPOINT: Returns ID of all receipts in examples folder
	router.GET("/receipts/process", handleProcessReceipts)
	// (this needs to a GET because we are not actually receiving the json through http.  It's just in memory)

	//ENDPOINT: Returns points of receipt at this ID
	router.GET("/receipts/:id/points", handleGetPointsTotal)

	//This starts a local server
	//Use command: curl http://localhost/{and the above routes} to access service
	router.Run("localhost:8080")

	/*
		COMMANDS TO COPY FOR SEPARATE TERMINAL
		--------------------------------------
		curl http://localhost:8080/receipts/process
		curl http://localhost:8080/receipts/       /points
		
	*/

}

/* HANDLER FUNCTION */

func handleProcessReceipts(c *gin.Context){

	gatherReceipts_FromExamplesFolder()

	//make slice of ids, empty but capacity set at the length of the map
	ids := make([]int, 0, len(Receipts))

	//loop Receipts map
	for key := range Receipts {
		//fill ids slice
		ids = append(ids, key)
	}

	//format ids into JSON and send to client
	c.IndentedJSON(http.StatusOK, ids)

}

func handleGetPointsTotal(c *gin.Context){

	//get id
	id,_ := strconv.Atoi(c.Param("id"))








	c.IndentedJSON(http.StatusOK, Receipts[id])

}


/* UTILITIES */

func gatherReceipts_FromExamplesFolder(){

	/* This enables you to put new receipts in the examples folder and they will be processed */

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

		//put into Receipts map
		Receipts[this_receipt.ID] = this_receipt
	}
}

func Process_Raw_Receipt(rec *Receipt_raw) (*Receipt) {

	//make new receipt
	var receipt *Receipt = new(Receipt)

	//create unique ID for this receipt
	receipt.ID = Get_UniqueID()

	//process values for each member
	receipt.Retailer = rec.Retailer
	receipt.PurchaseDate,_ = time.Parse("YYYY-MM-DD", rec.PurchaseDate)			//@@@@ This doesn't work
	receipt.PurchaseTime,_ = time.Parse("HH:MM", rec.PurchaseTime)				//@@@@ This doesn't work
	receipt.Total = Parse_String_ToFloat32(rec.Total)
	receipt.Items = make([]*Item, 0)

	//loop rec.Items
	for _, item := range rec.Items {

		//make new Item
		var new_item = new(Item)

		//alter values
		new_item.ShortDescription = item.ShortDescription
		new_item.Price = Parse_String_ToFloat32(item.Price)

		//add to .Items
		receipt.Items = append(receipt.Items, new_item)
	}

	return receipt
}

func Parse_String_ToFloat32(s string) (float32){

	//get float64
	var num float64
	num,_ = strconv.ParseFloat(s, 32)

	//return float32
	return float32(num)
}

func Get_UniqueID()(int){

	//store current ID available
	var int_to_return int = UniqueID_Counter

	//increment for next use
	UniqueID_Counter++

	//return stored value
	return int_to_return
}

