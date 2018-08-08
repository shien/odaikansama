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
