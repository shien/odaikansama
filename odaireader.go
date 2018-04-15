package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

const (
	clothing = "服"
	theme    = "テーマ"
)

func GetOdaiList(odai_type string, odai_subtype string) []string {

	fp, err := os.OpenFile("odai.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	r.Comma = ','

	result := []string{}

	for {
		record, read_err := r.Read()
		if read_err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(read_err)
		}
		if (record[0] == odai_type) && (record[1] == odai_subtype) {
			result = append(result, record[2])
		}
	}
	return result
}

//func main() {
//	fmt.Println(GetOdaiList("服", "靴"))
//}
