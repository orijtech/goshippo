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

package goshippo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"time"
)

type Parcel struct {
	State ParcelState `json:"object_state"`

	// Date and time of Parcel creation.
	CreatedAt *time.Time `json:"object_created"`

	// Date and time of last Parcel update. Since you
	// cannot update Parcels after they were created,
	// this timestamp reflects when the Parcel was changed
	// by GoShippo's systems for the last time e.g during
	// sorting the dimensions given.
	UpdatedAt *time.Time `json:"object_updated"`

	// ID is an output only variable created by GoShippo's
	// backend, for the given Parcel object. This ID is
	// required when creating a Shipment.
	ID string `json:"object_id"`

	// The username of the user who created the Parcel.
	OwnerUsername string `json:"object_owner"`

	// Length of the Parcel. Upto 6 digits in front and 4 digits
	// after the decimal separator are accepted.
	Length float64 `json:"length,string"`

	// Width of the Parcel. Upto 6 digits in front and 4 digits
	// after the decimal separator are accepted.
	Width float64 `json:"width,string"`

	// Height of the Parcel. Upto 6 digits in front and 4 digits
	// after the decimal separator are accepted.
	Height float64 `json:"height,string"`

	// The unit used for length, width and height dimensions.
	DistanceUnit DistanceUnit `json:"distance_unit"`

	Weight float64 `json:"weight,string"`

	// The unit used for mass/"weight"
	MassUnit MassUnit `json:"mass_unit"`

	// ParcelTemplate if set dictates that the parcel dimensions
	// have to be sent, but will not be used for the Rate generation.
	// The dimensions below will instead be used. The parcel weight
	// is not affected by the use of a template.
	ParcelTemplate ParcelTemplate `json:"template"`

	// Metadata is a string of upto 100 characters that
	// can be filled with any additional information that
	// you want to attach to the object.
	Metadata string `json:"metadata"`

	Extra *ParcelExtra `json:"extra"`
}

var blankParcel Parcel

func (c *Client) CreateParcel(parcel *Parcel) (*Parcel, error) {
	if err := parcel.Validate(); err != nil {
		return nil, err
	}

	blob, err := json.Marshal(parcel)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/parcels/", baseURL)
	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	blob, _, err = c.doAuthAndReq(req)
	if err != nil {
		return nil, err
	}
	recvParcel := new(Parcel)
	if err := json.Unmarshal(blob, recvParcel); err != nil {
		return nil, err
	}
	if reflect.DeepEqual(*recvParcel, blankParcel) {
		return nil, errBlankParcelFromServer
	}
	return recvParcel, nil
}

func zeroOrNegativeFloat64(f64 float64) bool {
	return math.Abs(f64-0.0) <= 0.0
}

var (
	errBlankLength       = errors.New("expecting length > 0.0")
	errBlankWidth        = errors.New("expecting width > 0.0")
	errBlankHeight       = errors.New("expecting height > 0.0")
	errBlankDistanceUnit = errors.New("expecting a non-blank distance unit")
	errBlankWeight       = errors.New("expecting weight > 0.0")
	errBlankMassUnit     = errors.New("expecting a non-blank mass unit")

	errBlankParcelFromServer = errors.New("got back a blank parcel from the server")
)

func (p *Parcel) Validate() error {
	if p == nil || zeroOrNegativeFloat64(p.Length) {
		return errBlankLength
	}
	if zeroOrNegativeFloat64(p.Width) {
		return errBlankWidth
	}
	if zeroOrNegativeFloat64(p.Height) {
		return errBlankHeight
	}
	if p.DistanceUnit == "" {
		return errBlankDistanceUnit
	}
	if zeroOrNegativeFloat64(p.Weight) {
		return errBlankWeight
	}
	if p.MassUnit == "" {
		return errBlankMassUnit
	}
	return nil
}

type ParcelState string

type DistanceUnit string

const (
	DistanceCentimetre DistanceUnit = "cm"
	DistanceInch       DistanceUnit = "in"
	DistanceFoot       DistanceUnit = "ft"
	DistanceMillimetre DistanceUnit = "mm"
	DistanceYard       DistanceUnit = "yd"
)

type MassUnit string

const (
	MassGram     MassUnit = "g"
	MassOunce    MassUnit = "oz"
	MassPound    MassUnit = "lb"
	MassKilogram MassUnit = "kg"
)

type PaymentMethod string

const (
	// SecuredFunds include money orders, certified cheques
	// and others (see UPS and FedEx for details).
	PaymentSecuredFunds PaymentMethod = "SECURED_FUNDS"
	PaymentCash         PaymentMethod = "CASH"
	PaymentAny          PaymentMethod = "ANY"
)

type ParcelExtra struct {
	CollectionOnDelivery *CollectionOnDelivery `json:"COD"`

	// Insurance specifies collection on
	// delivery details (UPS and FedEx only).
	Insurance *Insurance `json:"insurance"`
}

type CollectionOnDelivery struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency"`

	// If no payment method is set, it defaults to PaymentAny
	PaymentMethod string `json:"payment_method"`
}

type Insurance struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency"`
	Provider     string `json:"provider"`
	Content      string `json:"content"`
}
