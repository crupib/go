package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		fmt.Println(i, row)
		// now using index access
		if i == 8 {
			fmt.Println(row[0])
			break
		}
	}
}
