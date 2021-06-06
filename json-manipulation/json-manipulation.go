package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Data []struct {
	InventoryID int64  `json:"inventory_id"`
	Name        string `json:"name"`
	Placement   struct {
		Name   string `json:"name"`
		RoomID int64  `json:"room_id"`
	} `json:"placement"`
	PurchasedAt int64    `json:"purchased_at"`
	Tags        []string `json:"tags"`
	Type        string   `json:"type"`
}

func main() {
	data := getData()
	soal1(data)
	fmt.Println()
	soal2(data)
	fmt.Println()
	soal3(data)
	fmt.Println()
	soal4(data)
	fmt.Println()
	soal5(data)
}

func getData() Data {
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	//unmarshal to Data struct
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Data
	json.Unmarshal(byteValue, &data)
	return data
}

func soal1(data Data) Data {
	fmt.Println("1. Find items in the Meeting Room.")
	var item Data
	i := 0
	for _, v := range data {
		if v.Placement.Name == "Meeting Room" { // find by room name
			item = append(item, v)
			fmt.Println(v)
			i++
		}
	}
	return item
}

func soal2(data Data) Data {
	fmt.Println("2. Find all electronic devices.")
	var item Data
	i := 0
	for _, v := range data {
		if v.Type == "electronic" { // find by Type
			item = append(item, v)
			fmt.Println(v)
			i++
		}
	}
	return item
}

func soal3(data Data) Data {
	fmt.Println("3. Find all the furniture.")
	var item Data
	i := 0
	for _, v := range data {
		if v.Type == "furniture" { // find by Type
			item = append(item, v)
			fmt.Println(v)
			i++
		}
	}
	return item
}

func soal4(data Data) Data {
	fmt.Println("4. Find all items were purchased on 16 Januari 2020.")
	var item Data
	i := 0
	for _, v := range data {
		// get/convert date
		date := time.Date(2020, time.January, 16, 0, 0, 0, 0, time.UTC)
		vDate := time.Unix(v.PurchasedAt, 0)
		// compare if date == vDate
		if date.Year() == vDate.Year() && date.Month() == vDate.Month() && date.Day() == vDate.Day() {
			item = append(item, v)
			fmt.Print(v)
			fmt.Println(" ->", vDate)
			i++
		}
	}

	return item
}

func soal5(data Data) Data {
	fmt.Println("5. Find all items with brown color.")
	var item Data
	i := 0
	for _, v := range data {
		if isIn(v.Tags, "brown") { // find by Tags, linear search
			item = append(item, v)
			fmt.Println(v)
			i++
		}
	}
	return item
}

func isIn(slice []string, i string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, i) {
			return true
		}
	}
	return false
}
