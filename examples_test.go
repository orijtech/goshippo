// Copyright 2017 orijtech. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
