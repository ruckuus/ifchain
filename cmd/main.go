package main

import (
	"github.com/ruckuus/ifchain/internal/ifchain"
	"github.com/ruckuus/ifchain/pkg/dummydb"
)

func main() {
	obj := &ifchain.Object{
		Name:    "User",
		Version: "1.0",
	}

	store := dummydb.NewDB("db")
	cache := dummydb.NewDB("cache")

	svc := ifchain.NewService(store, cache)
	svc.Save(obj)
}
