package main

import (
	"testing"
)

func TestGetOdai(t *testing.T) {

	cache := OdaiCache{}
	result := cache.GetOdai("服", "靴")

	if result.OdaiType != "服" {
		t.Fatal("failed test")
	}

	if result.OdaiSubtype != "靴" {
		t.Fatal("failed test")
	}
}

func TestIsOdaiExist(t *testing.T) {

	testOdaiType := "TestType"
	testOdaiSubtype := "TestSubtype"
	testOdai := "TestOdaiExist"

	cache := OdaiCache{}
	if cache.IsOdaiExist(testOdaiType, testOdaiSubtype, testOdai) {
		t.Fatal("Odai should not exist.")
	}

	cache.AddOdai(testOdaiType, testOdaiSubtype, testOdai)

	if !cache.IsOdaiExist(testOdaiType, testOdaiSubtype, testOdai) {
		t.Fatal("Odai should exist.")
	}

}

func TestAddOdai(t *testing.T) {

	testOdaiType := "TestType"
	testOdaiSubtype := "TestSubtype"
	testOdai := "TestAddOdai"

	cache := OdaiCache{}
	cache.AddOdai(testOdaiType, testOdaiSubtype, testOdai)

	result := cache.GetOdai(testOdaiType, testOdaiSubtype)

	if result.OdaiType != testOdaiType {
		t.Fatal("failed test")
	}

	if result.OdaiSubtype != testOdaiSubtype {
		t.Fatal("failed test")
	}

	cache.AddOdai(testOdaiType, testOdaiSubtype, testOdai)

	validOdaiCount := 0
	for _, odai := range result.OdaiList {
		if odai == testOdai {
			validOdaiCount++
		}
	}

	if validOdaiCount < 1 {
		t.Fatal("Odai not found.")
	} else if validOdaiCount > 1 {
		t.Fatal("Too many odai found.")
	}
}
