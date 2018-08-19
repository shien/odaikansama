package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

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

func (cache *OdaiCache) IsOdaiExist(OdaiType string, OdaiSubtype string, Odai string) bool {

	length := len(cache.Data)

	for c := 0; c < length; c++ {
		if cache.Data[c].OdaiType == OdaiType && cache.Data[c].OdaiSubtype == OdaiSubtype {
			for _, n := range cache.Data[c].OdaiList {
				if Odai == n {
					return true
				}
			}
		}
	}
	return false
}

func (cache *OdaiCache) AddOdai(OdaiType string, OdaiSubtype string, newOdai string) {

	if !cache.IsOdaiExist(OdaiType, OdaiSubtype, newOdai) {
		WriteFile(OdaiType, OdaiSubtype, newOdai)

		length := len(cache.Data)
		odaiTypeExist := false
		for i := 0; i < length; i++ {
			if cache.Data[i].OdaiType == OdaiType && cache.Data[i].OdaiSubtype == OdaiSubtype {
				cache.Data[i].OdaiList = append(cache.Data[i].OdaiList, newOdai)
				odaiTypeExist = true
				break
			}
		}

		if !odaiTypeExist {
			var odai Odai
			odai.OdaiType = OdaiType
			odai.OdaiSubtype = OdaiSubtype
			odai.OdaiList = append(odai.OdaiList, newOdai)
			cache.Data = append(cache.Data, odai)
		}
	}
}

func WriteFile(OdaiType string, OdaiSubtype string, Odai string) {
	fp, err := os.OpenFile("odai.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	_, err2 := fp.WriteString(OdaiType + "," + OdaiSubtype + "," + Odai + "\n")
	if err2 != nil {
		panic(err2)
	}
}
