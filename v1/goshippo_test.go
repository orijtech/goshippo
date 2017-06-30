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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/orijtech/goshippo/v1"
	"github.com/orijtech/otils"

	"github.com/odeke-em/go-uuid"
)

func TestCreateAddress(t *testing.T) {
	client, err := goshippo.NewClient(token1)
	if err != nil {
		t.Fatalf("address client err: %v", err)
	}
	client.SetHTTPRoundTripper(&backend{route: createAddressRoute})

	tests := [...]struct {
		addr    *goshippo.Address
		wantErr bool
	}{
		0: {addr: &goshippo.Address{}, wantErr: true},
		1: {addr: &goshippo.Address{Purpose: "Testing"}},
	}

	for i, tt := range tests {
		addr, err := client.CreateAddress(tt.addr)
		if tt.wantErr {
			if err == nil {
				t.Errorf("#%d: want non-nil error", i)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: gotErr=%v", i, err)
			continue
		}
		if addr == nil {
			t.Errorf("#%d: expected non-nil address", i)
			continue
		}

		if reflect.DeepEqual(addr, blankAddress) {
			t.Errorf("#%d: expected a non-blank address", i)
		}
	}
}

const (
	addrID1 = "d799c2679e644279b59fe661ac8fa488"
)

func TestAddressByID(t *testing.T) {
	client, err := goshippo.NewClient(token1)
	if err != nil {
		t.Fatalf("address client err: %v", err)
	}
	client.SetHTTPRoundTripper(&backend{route: addressByIDRoute})

	tests := [...]struct {
		addrID  string
		wantErr bool
	}{
		0: {addrID: "", wantErr: true},
		1: {addrID: "        ", wantErr: true},
		2: {addrID: uuid.NewRandom().String(), wantErr: true}, // Unknown address
		3: {addrID: addrID1},
	}

	for i, tt := range tests {
		addr, err := client.AddressByID(tt.addrID)
		if tt.wantErr {
			if err == nil {
				t.Errorf("#%d: want non-nil error", i)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: gotErr=%v", i, err)
			continue
		}
		if addr == nil {
			t.Errorf("#%d: expected non-nil address", i)
			continue
		}

		if reflect.DeepEqual(addr, blankAddress) {
			t.Errorf("#%d: expected a non-blank address", i)
		}
	}
}

func TestListAddresses(t *testing.T) {
	client, err := goshippo.NewClient(token1)
	if err != nil {
		t.Fatalf("list address client err: %v", err)
	}
	client.SetHTTPRoundTripper(&backend{route: listAddressesRoute})

	tests := [...]struct {
		req         *goshippo.AddressListRequest
		wantErr     bool
		wantAtLeast uint64
	}{
		0: {req: nil, wantAtLeast: 6},
		1: {req: &goshippo.AddressListRequest{}, wantAtLeast: 6},
		2: {req: &goshippo.AddressListRequest{PageNumber: 10}, wantAtLeast: 0},
		3: {
			req: &goshippo.AddressListRequest{
				PageToken: "flux",
			}, wantErr: true}, // PageToken must be a valid URL
		4: {
			req: &goshippo.AddressListRequest{
				PageToken: "https://orijtech.com?aha=bypassed",
			}, wantErr: true, // PageToken must be a URL constructed from the baseURL
		},
		5: {
			req: &goshippo.AddressListRequest{
				PageToken: "https://api.goshippo.com/addresses?page=1",
			}, wantAtLeast: 1,
		},
		6: {
			req: &goshippo.AddressListRequest{
				PageToken: "https://api.goshippo.com.com/addresses?page=1",
			}, wantErr: true, // trickery with the api.goshippo.com.com a valid URL and host
		},
		7: {
			req: &goshippo.AddressListRequest{
				PageToken: "http://api.goshippo.com/addresses?page=1",
			}, wantErr: true, // trickery with the http scheme, want https
		},
	}

	for i, tt := range tests {
		if tt.req != nil {
			tt.req.ThrottleDurationMs = goshippo.NoThrottle
		}
		res, err := client.ListAddresses(tt.req)
		if tt.wantErr {
			if err == nil {
				t.Errorf("#%d: want non-nil error", i)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: gotErr=%v", i, err)
			continue
		}
		if res == nil {
			t.Errorf("#%d: expected non-nil listAddresses results", i)
			continue
		}

		pageCount := uint64(0)
		nonBlankAddresses := uint64(0)
		for page := range res.Pages {
			if page.Err != nil && tt.wantAtLeast > 0 {
				t.Errorf("#%d: pageNumber: %d err: %v\n", i, page.PageNumber, page.Err)
				continue
			}

			for _, addr := range page.Addresses {
				if !reflect.DeepEqual(addr, blankAddress) {
					nonBlankAddresses += 1
				}
			}
			pageCount += 1
		}

		if nonBlankAddresses < tt.wantAtLeast {
			t.Errorf("#%d: got nAddresses=%d wantAtLeast nAddresses=%d", i, nonBlankAddresses, tt.wantAtLeast)
		}
	}
}

func TestValidateAddress(t *testing.T) {
	client, err := goshippo.NewClient(token1)
	if err != nil {
		t.Fatalf("address client err: %v", err)
	}
	client.SetHTTPRoundTripper(&backend{route: validateAddressRoute})

	tests := [...]struct {
		addrID  string
		wantErr bool
	}{
		0: {addrID: "", wantErr: true},
		1: {addrID: "        ", wantErr: true},
		2: {addrID: uuid.NewRandom().String(), wantErr: true}, // Unknown address
		3: {addrID: addrID1},
	}

	for i, tt := range tests {
		addr, err := client.ValidateAddress(tt.addrID)
		if tt.wantErr {
			if err == nil {
				t.Errorf("#%d: want non-nil error", i)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: gotErr=%v", i, err)
			continue
		}
		if addr == nil {
			t.Errorf("#%d: expected non-nil address", i)
			continue
		}

		if reflect.DeepEqual(addr, blankAddress) {
			t.Errorf("#%d: expected a non-blank address", i)
		}
	}
}

func TestCreateParcel(t *testing.T) {
	client, err := goshippo.NewClient(token1)
	if err != nil {
		t.Fatalf("address client err: %v", err)
	}
	client.SetHTTPRoundTripper(&backend{route: createParcelRoute})

	tests := [...]struct {
		p       *goshippo.Parcel
		wantErr bool
	}{
		0: {p: &goshippo.Parcel{}, wantErr: true},
		1: {p: nil, wantErr: true},
		2: {
			p: &goshippo.Parcel{
				Length: 10.3,
				Width:  12.0,
				Height: 25,
				Weight: 99.2,

				DistanceUnit: goshippo.DistanceYard,
				MassUnit:     goshippo.MassKilogram,
			},
		},
	}

	for i, tt := range tests {
		parcel, err := client.CreateParcel(tt.p)
		if tt.wantErr {
			if err == nil {
				t.Errorf("#%d: want non-nil error", i)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: gotErr=%v", i, err)
			continue
		}
		if parcel == nil {
			t.Errorf("#%d: expected non-nil parcel", i)
			continue
		}

		if reflect.DeepEqual(parcel, blankParcel) {
			t.Errorf("#%d: expected a non-blank parcel", i)
		}
	}
}

const (
	createAddressRoute   = "/create-address"
	addressByIDRoute     = "/retrieve-address"
	validateAddressRoute = "/validate-address"
	listAddressesRoute   = "/list-addresses"

	createParcelRoute = "/create-parcel"
)

const (
	token1 = "token-1"
	token2 = "token-2"
)

var _ http.RoundTripper = (*backend)(nil)

func (b *backend) RoundTrip(req *http.Request) (*http.Response, error) {
	switch b.route {
	case createAddressRoute:
		return b.createAddressRoundTrip(req)
	case addressByIDRoute:
		return b.addressByIDRoundTrip(req)
	case validateAddressRoute:
		return b.validateAddressRoundTrip(req)
	case listAddressesRoute:
		return b.listAddressesRoundTrip(req)

	case createParcelRoute:
		return b.createParcelRoundTrip(req)
	default:
		return makeResp(fmt.Sprintf("%q unknown route", b.route), http.StatusNotFound), nil
	}
}

func removeBlanks(segs []string) (pruned []string) {
	for _, seg := range segs {
		seg = strings.TrimSpace(seg)
		if seg != "" {
			pruned = append(pruned, seg)
		}
	}
	return pruned
}

var blankAddress goshippo.Address

func (b *backend) listAddressesRoundTrip(req *http.Request) (*http.Response, error) {
	if badAuthResp, err := checkBadAuth(req, "GET"); badAuthResp != nil || err != nil {
		return badAuthResp, err
	}
	pathSplits := removeBlanks(strings.Split(req.URL.Path, "/"))
	if len(pathSplits) < 1 || pathSplits[0] != "addresses" {
		return makeResp("expecting a path of the form /addresses<?query=value>", http.StatusBadRequest), nil
	}

	query := req.URL.Query()

	var pageNumber int
	var err error
	if pageStr := query.Get("page"); pageStr != "" {
		pageNumber, err = strconv.Atoi(query.Get("page"))
	}
	if err != nil {
		return makeResp(err.Error(), http.StatusBadRequest), nil
	}

	srcPath := fmt.Sprintf("./testdata/address-list-%d.json", pageNumber)
	return respFromFile(srcPath)
}

func (b *backend) validateAddressRoundTrip(req *http.Request) (*http.Response, error) {
	if badAuthResp, err := checkBadAuth(req, "GET"); badAuthResp != nil || err != nil {
		return badAuthResp, err
	}
	pathSplits := removeBlanks(strings.Split(req.URL.Path, "/"))
	if len(pathSplits) != 3 || pathSplits[0] != "addresses" || pathSplits[2] != "validate" {
		return makeResp("expecting a path of the form /addresses/<addressID>/validate", http.StatusBadRequest), nil
	}
	addressID := pathSplits[1]
	srcPath := fmt.Sprintf("./testdata/address-validation-%s.json", addressID)

	return respFromFile(srcPath)
}

func (b *backend) addressByIDRoundTrip(req *http.Request) (*http.Response, error) {
	if badAuthResp, err := checkBadAuth(req, "GET"); badAuthResp != nil || err != nil {
		return badAuthResp, err
	}
	pathSplits := removeBlanks(strings.Split(req.URL.Path, "/"))
	if len(pathSplits) != 2 || pathSplits[0] != "addresses" {
		return makeResp("expecting a path of the form /addresses/<addressID>", http.StatusBadRequest), nil
	}
	addressID := pathSplits[len(pathSplits)-1]
	srcPath := fmt.Sprintf("./testdata/address-%s.json", addressID)

	return respFromFile(srcPath)
}

func (b *backend) createAddressRoundTrip(req *http.Request) (*http.Response, error) {
	if badAuthResp, err := checkBadAuth(req, "POST"); badAuthResp != nil || err != nil {
		return badAuthResp, err
	}
	if req.Body == nil {
		return makeResp("expecting a non-nil body", http.StatusBadRequest), nil
	}
	defer req.Body.Close()

	if got, want := req.Header.Get("Content-Type"), "application/json"; got != want {
		return makeResp(fmt.Sprintf("got contentType=%q want=%q", got, want), http.StatusBadRequest), nil
	}

	slurp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return makeResp(err.Error(), http.StatusBadRequest), nil
	}

	addr := new(goshippo.Address)
	if err := json.Unmarshal(slurp, addr); err != nil {
		return makeResp(err.Error(), http.StatusBadRequest), nil
	}
	if *addr == blankAddress {
		return makeResp("expecting a non blank body", http.StatusBadRequest), nil
	}

	return respFromFile("./testdata/address-1.json")
}

var blankParcel goshippo.Parcel

func (b *backend) createParcelRoundTrip(req *http.Request) (*http.Response, error) {
	if badAuthResp, err := checkBadAuth(req, "POST"); badAuthResp != nil || err != nil {
		return badAuthResp, err
	}
	if req.Body == nil {
		return makeResp("expecting a non-nil body", http.StatusBadRequest), nil
	}
	defer req.Body.Close()

	if got, want := req.Header.Get("Content-Type"), "application/json"; got != want {
		return makeResp(fmt.Sprintf("got contentType=%q want=%q", got, want), http.StatusBadRequest), nil
	}

	slurp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return makeResp(err.Error(), http.StatusBadRequest), nil
	}

	parcel := new(goshippo.Parcel)
	if err := json.Unmarshal(slurp, parcel); err != nil {
		return makeResp(err.Error(), http.StatusBadRequest), nil
	}
	if reflect.DeepEqual(*parcel, blankParcel) {
		return makeResp("expecting a non blank body", http.StatusBadRequest), nil
	}

	return respFromFile("./testdata/parcel-1.json")
}

func respFromFile(path string) (*http.Response, error) {
	f, err := os.Open(path)
	if err != nil {
		return makeResp(err.Error(), http.StatusBadRequest), nil
	}
	res := makeResp("200 OK", http.StatusOK)
	res.Body = f
	return res, nil
}

func checkBadAuth(req *http.Request, wantMethod string) (*http.Response, error) {
	if req.Method != wantMethod {
		return makeResp(fmt.Sprintf("gotMethod=%q wantMethod=%q", req.Method, wantMethod), http.StatusMethodNotAllowed), nil
	}
	authKey := req.Header.Get("Authorization")
	splits := strings.Split(authKey, "ShippoToken")
	shippoToken := strings.TrimSpace(otils.FirstNonEmptyString(splits...))
	if shippoToken == "" {
		return makeResp(`expected auth token "ShippoToken"`, http.StatusUnauthorized), nil
	}
	switch shippoToken {
	case token1, token2:
		return nil, nil
	default:
		return makeResp("unauthorized token", http.StatusUnauthorized), nil
	}
}

func makeResp(status string, code int) *http.Response {
	return &http.Response{
		StatusCode: code,
		Header:     make(http.Header),
		Status:     status,
	}
}

type backend struct {
	route string
}
