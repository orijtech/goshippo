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
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/orijtech/otils"
)

type Address struct {
	// Purpose is required for creating addresses.
	Purpose string `json:"object_purpose,omitempty"`

	// Complete is an output only variable that
	// is set only when the address contains
	// all the required values.
	Complete bool `json:"is_complete"`

	// CreatedAt is an output only variable that
	// records when this address was created.
	CreatedAt *time.Time `json:"object_created"`

	// UpdatedAt is an output only variable that
	// records when this address was last updated.
	UpdatedAt *time.Time `json:"object_updated"`

	// ID is the unique identifier of the given Address.
	// This ID is required to create a shipment object.
	ID string `json:"object_id"`

	// OwnerUsername  is the username of
	// the user who created the Address.
	OwnerUsername string `json:"object_owner"`

	// AddresseeName is required for purchase.
	AddresseeName string `json:"name"`

	// Company is the name of the addressee's company.
	Company string `json:"company"`

	// Street1 is the first street line, 35 character limit.
	// Usually street number and street name(except for DHL Germany).
	// Street1 is required for purchase.
	Street1 string `json:"street1"`

	// StreetNumber is the street number of the
	// addressed building. This field can be included
	// in Street1 for all carriers except for DHL Germany.
	StreetNumber string `json:"street_no"`

	// Street2 is the second street line, 35 character limit.
	Street2 string `json:"street2"`

	// Street3 is the second street line, 35 character limit.
	// This field is only accepted for USPS international
	// shipments, UPS domestic and UPS international shipments.
	Street3 string `json:"street3"`

	// City is required for purchases. When creating a Quote Address,
	// sending a city is optional but will yield more accurate Rates.
	// Please bear in mind that city names may be ambiguous(there are
	// 34 Springfields in the US). Pass in a state or a ZIP code,
	// if known, it will yield more accurate results.
	City string `json:"city"`

	// ZipCode is the postal code of an Address. When
	// creating a Quote Address, sending a ZIP is optional
	// but will yield more accurate rates.
	// ZipCode is required for purchases.
	ZipCode string `json:"zip"`

	// State is required for shipments from the United States and Canada
	// (most carriers only accept two character state abbreviations).
	// However, to receive more accurate quotes, passing it is generally
	// recommended.
	// State is required for purchase for some countries.
	State string `json:"state"`

	// Country is the ISO 3166-1-alpha-2 code. An example is "US" or "DE".
	// All accepted values can be found at http://www.iso.org/
	// Country is always required.
	Country string `json:"country"`

	// Phone allows carriers to call the receipient when
	// delivering the Parcel. This increases the probability
	// of delivery and helps to avoid accessorial charges
	// after a Parcel has been shipped.
	Phone string `json:"phone"`

	// Email is the email address of the contact person.
	// It has to be RFC3696/5321-compliant.
	Email string `json:"email"`

	// Residential indicates whether the address
	// provided is a residential address or not.
	Residential otils.NumericBool `json:"is_residential,omitempty"`

	// ShouldValidate if set to true indicates that the
	// GoShippo backend should validate the Address.
	ShouldValidate otils.NumericBool `json:"validate,omitempty"`

	// Metadata is an optional string of upto 100 characters
	// that can be filled with any additional information
	// that you want to attach to the object.
	Metadata string `json:"metadata,omitempty"`

	// InTestMode indicates whether the
	// object has been created in test mode.
	InTestMode bool `json:"test"`

	// ValidationResults contains information regarding if an
	// address had been validated or not. Also contains any
	// messages generated during validation.
	ValidationResults *ValidationResult `json:"validation_results,omitempty"`
}

var (
	errBlankPurpose = errors.New("purpose is required")

	errAlreadyClosed = errors.New("already closed channel")

	errEmptyAddressID = errors.New("expecting a non-empty addressID")

	errBlankAddressReceived = errors.New("received a blank address from the backend")
)

func (addr *Address) Validate() error {
	if addr == nil || strings.TrimSpace(addr.Purpose) == "" {
		return errBlankPurpose
	}
	return nil
}

var blankAddress Address

func (c *Client) CreateAddress(addr *Address) (*Address, error) {
	if err := addr.Validate(); err != nil {
		return nil, err
	}

	blob, err := json.Marshal(addr)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("%s/addresses/", baseURL)
	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doReqAndAddressByID(req)
}

func (c *Client) doReqAndAddressByID(req *http.Request) (*Address, error) {
	blob, _, err := c.doAuthAndReq(req)
	if err != nil {
		return nil, err
	}
	recvAddr := new(Address)
	if err := json.Unmarshal(blob, recvAddr); err != nil {
		return nil, err
	}
	if *recvAddr == blankAddress {
		return nil, errBlankAddressReceived
	}
	return recvAddr, nil
}

func (c *Client) AddressByID(addressID string) (*Address, error) {
	addressID = strings.TrimSpace(addressID)
	if addressID == "" {
		return nil, errEmptyAddressID
	}
	fullURL := fmt.Sprintf("%s/addresses/%s/", baseURL, addressID)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	return c.doReqAndAddressByID(req)
}

func (c *Client) ValidateAddress(addressID string) (*Address, error) {
	addressID = strings.TrimSpace(addressID)
	if addressID == "" {
		return nil, errEmptyAddressID
	}
	fullURL := fmt.Sprintf("%s/addresses/%s/validate/", baseURL, addressID)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	return c.doReqAndAddressByID(req)
}

type AddressPage struct {
	PageNumber uint64     `json:"page_number"`
	Addresses  []*Address `json:"addresses"`
	Err        error

	PreviousToken string `json:"previous_token"`
	NextToken     string `json:"next_token"`
}

type AddressesPager struct {
	Pages  <-chan *AddressPage
	Cancel func() error
}

func makeCanceler() (func() error, <-chan bool) {
	var closeOnce sync.Once
	cancelChan := make(chan bool)
	closeFn := func() error {
		var err error = errAlreadyClosed
		closeOnce.Do(func() {
			close(cancelChan)
			err = nil
		})
		return err
	}
	return closeFn, cancelChan
}

const (
	NoThrottle = -1
)

type AddressListRequest struct {
	MaxPages     uint64 `json:"max_pages"`
	LimitPerPage uint64 `json:"limit_per_page"`
	PageToken    string `json:"page_token"`
	PageNumber   uint64 `json:"page_number"`

	ThrottleDurationMs int64 `json:"throttle_duration_ms"`
}

type pager struct {
	Count      uint64 `json:"count"`
	Limit      uint64 `json:"limit"`
	PageNumber uint64 `json:"page"`
}

type addressesWrap struct {
	Count         uint64               `json:"count"`
	PreviousToken otils.NullableString `json:"previous"`
	NextToken     otils.NullableString `json:"next"`
	Addresses     []*Address           `json:"results"`
}

func pagerFromAddressListRequest(alReq *AddressListRequest) *pager {
	pg := &pager{Limit: alReq.LimitPerPage, PageNumber: alReq.PageNumber}
	if pg.PageNumber <= 0 {
		// PageNumbers are 1-based for goshippo
		pg.PageNumber = 1
	}
	return pg
}

var (
	parsedBaseURL, _ = url.Parse(baseURL)
)

func (c *Client) ListAddresses(alReq *AddressListRequest) (*AddressesPager, error) {
	if alReq == nil {
		alReq = new(AddressListRequest)
	}

	cancelFn, cancelChan := makeCanceler()
	throttleDurationMs := 150 * time.Millisecond
	if alReq.ThrottleDurationMs == NoThrottle {
		throttleDurationMs = 0
	} else if alReq.ThrottleDurationMs > 0 {
		throttleDurationMs = time.Duration(alReq.ThrottleDurationMs) * time.Millisecond
	}
	pagesChan := make(chan *AddressPage)

	maxPage := alReq.MaxPages
	pageExceeded := func(page uint64) bool {
		return maxPage > 0 && page >= maxPage
	}

	// pageNumbers are 1-based for goshippo
	pageNumber := uint64(1)

	var fullURL string
	if alReq.PageToken != "" {
		// GoShippo PageTokens are full URLs e.g
		// https://api.goshippo.com/addresses/?limit=1&page=2
		fullURL = alReq.PageToken
	} else {
		pg := pagerFromAddressListRequest(alReq)
		qv, err := otils.ToURLValues(pg)
		if err != nil {
			return nil, err
		}
		fullURL = fmt.Sprintf("%s/addresses/", baseURL)
		if len(qv) > 0 {
			fullURL += "?" + qv.Encode()
		}
	}

	// Ensure that fullURL parses upfront since insertion of tokens
	// could trip out if they aren't proper URLs.
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}

	var errsList []string
	if got, want := parsedURL.Host, parsedBaseURL.Host; got != want {
		errsList = append(errsList, fmt.Sprintf("gotHost=%q wantHost=%q", got, want))
	}
	if got, want := parsedURL.Scheme, parsedBaseURL.Scheme; got != want {
		errsList = append(errsList, fmt.Sprintf("gotScheme=%q wantSchem=%q", got, want))
	}
	if len(errsList) > 0 {
		return nil, errors.New(strings.Join(errsList, "\n"))
	}

	go func() {
		defer close(pagesChan)
		var previousToken, nextToken otils.NullableString

		for {
			page := &AddressPage{PageNumber: pageNumber}
			req, err := http.NewRequest("GET", fullURL, nil)
			if err != nil {
				page.Err = err
				pagesChan <- page
				return
			}

			blob, _, err := c.doAuthAndReq(req)
			if err != nil {
				page.Err = err
				pagesChan <- page
				return
			}

			pwrap := new(addressesWrap)
			if err := json.Unmarshal(blob, pwrap); err != nil {
				page.Err = err
				pagesChan <- page
				return
			}
			previousToken, nextToken = pwrap.PreviousToken, pwrap.NextToken
			page.PreviousToken = string(previousToken)
			page.NextToken = string(nextToken)

			page.Addresses = pwrap.Addresses[:]
			pagesChan <- page
			pageNumber += 1

			if pageExceeded(pageNumber) || len(pwrap.Addresses) == 0 || nextToken == "" {
				return
			}

			select {
			case <-time.After(throttleDurationMs):
			case <-cancelChan:
				return
			}

			// Now setting fullURL to be the next token because
			// GoShippo page tokens are literally fullURLs e.g:
			//  https://api.goshippo.com/addresses/?limit=1&page=2
			fullURL = string(nextToken)
		}
	}()

	return &AddressesPager{Cancel: cancelFn, Pages: pagesChan}, nil
}

type ValidationResult struct {
	Valid    bool                 `json:"is_valid"`
	Messages []*ValidationMessage `json:"messages"`
}

type ValidationMessage struct {
	Source string `json:"source"`
	Text   string `json:"text"`
	Code   string `json:"code"`
}
