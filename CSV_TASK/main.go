package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type product struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
}

func convertJSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var productList []product
	if err := json.NewDecoder(sourceFile).Decode(&productList); err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"product_id", "product_name", "price", "quantity"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Calculate the number of chunks to divide the data
	numChunks := 4
	chunkSize := (len(productList) + numChunks - 1) / numChunks

	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numChunks)

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(productList) {
			end = len(productList)
		}

		// Start a goroutine for processing each chunk
		go func(products []product) {
			defer wg.Done()
			for _, r := range products {
				csvRow := []string{
					fmt.Sprint(r.ProductID),
					r.ProductName,
					fmt.Sprint(r.Price),
					fmt.Sprint(r.Quantity),
				}
				if err := writer.Write(csvRow); err != nil {
					log.Println(err)
					return
				}
			}
		}(productList[start:end])
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return nil
}

func main() {
	if err := convertJSONToCSV("products.json", "data.csv"); err != nil {
		log.Fatal(err)
	}
}
