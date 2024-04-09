package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func _() {
	//opens the file
	file, err := os.Open("customers-100.csv")
	//checs the error
	if err != nil {
		log.Fatal("Error Reading the File")
	}
	//closes the file
	defer file.Close()
	reader:=csv.NewReader(file)
	records,err:=reader.ReadAll()
	if err !=nil{
		fmt.Println("Error reading records")
	}
	for _,eachrecord :=range records{
		fmt.Println(eachrecord)
	}

	

}
