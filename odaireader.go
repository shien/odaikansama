package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type OdaiCache struct {
	Data []Odai
}

type Odai struct {
	OdaiType    string
	OdaiSubtype string
	OdaiList    []string
}

func (cache *OdaiCache) GetOdai(OdaiType string, OdaiSubtype string) Odai {

	length := len(cache.Data)
	for i := 0; i < length; i++ {
		if cache.Data[i].OdaiType == OdaiType && cache.Data[i].OdaiSubtype == OdaiSubtype {
			return cache.Data[i]
		}
	}

	var odai Odai
	odai.OdaiList = ReadFile(OdaiType, OdaiSubtype)
	odai.OdaiType = OdaiType
	odai.OdaiSubtype = OdaiSubtype

	cache.Data = append(cache.Data, odai)
	return odai
}

func ReadFile(OdaiType string, OdaiSubtype string) []string {

	fp, err := os.OpenFile("odai.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)
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
		if (record[0] == OdaiType) && (record[1] == OdaiSubtype) {
			result = append(result, record[2])
		}
	}
	return result
}
