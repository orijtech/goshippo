package goshippo_test

import (
	"fmt"
	"log"

	"github.com/orijtech/goshippo/v1"
)

func Example_CreateAddress() {
	client, err := goshippo.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	addr, err := client.CreateAddress(&goshippo.Address{
		AddresseeName: "Orijtech Bot4",
		Purpose:       "This is a test",

		Company: "orijtech",
		Country: "US",
		Street1: "215 Clayton St.",
		City:    "San Francisco",
		State:   "CA",
		Phone:   "+1-555-341-9393",
		Email:   "bot4@orijtech.com",

		Metadata: "Second one",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created address: %+v\n", addr)
}

func Example_AddressByID() {
	client, err := goshippo.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	addr, err := client.AddressByID("7556f514e2ae4b468e215d7e04ca6277")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Here is the address: %+v\n", addr)
}

func Example_ValidateAddress() {
	client, err := goshippo.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	addr, err := client.ValidateAddress("7556f514e2ae4b468e215d7e04ca6277")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Here is the address: %+v\n", addr)
}

func Example_ListAddresses() {
	client, err := goshippo.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.ListAddresses(&goshippo.AddressListRequest{PageNumber: 1})
	if err != nil {
		log.Fatal(err)
	}

	for page := range res.Pages {
		if page.Err != nil {
			fmt.Printf("%#d: err: %v\n", page.PageNumber, page.Err)
			continue
		}
		for i, addr := range page.Addresses {
			fmt.Printf("#%d: address: %+v\n", i, addr)
		}
	}
}
