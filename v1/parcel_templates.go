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

import "github.com/orijtech/otils"

type ParcelTemplate otils.NullableString

// FedEx parcel templates
const (
	// 15.81 x 12.94 x 10.19 in
	FedEx10KgBox ParcelTemplate = "FedEx_Box_10kg"

	// 54.80 x 42.10 x 33.50 in
	FedEx25KgBox ParcelTemplate = "FedEx_Box_25kg"

	// 11.88 x 11.00 x 10.75 in
	FedExExtraLargeBox1 ParcelTemplate = "FedEx_Box_Extra_Large_1"

	// 15.75 x 14.13 x 6.00 in
	FedExExtraLargeBox2 ParcelTemplate = "FedEx_Box_Extra_Large_2"

	// 17.50 x 12.38 x 3.00 in
	FedExLargeBox1 ParcelTemplate = "FedEx_Box_Large_1"

	// 11.25 x 8.75 x 7.75 in
	FedExLargeBox2 ParcelTemplate = "FedEx_Box_Large_2"

	// 13.25 x 11.50 x 2.38 in
	FedExMediumBox1 ParcelTemplate = "FedEx_Box_Medium_1"

	// 11.25 x 8.75 x 4.38 in
	FedExMediumBox2 ParcelTemplate = "FedEx_Box_Medium_2"

	// 12.38 x 10.88 x 1.50 in
	FedExSmallBox1 ParcelTemplate = "FedEx_Box_Small_1"

	// 11.25 x 8.75 x 4.38 in
	FedExSmallBox2 ParcelTemplate = "FedEx_Box_Small_2"

	// 12.50 x 9.50 x 0.80 in
	FedExEnvelope ParcelTemplate = "FedEx_Envelope"

	// 11.75 x 14.75 x 2.00 in
	FedExPaddedPak ParcelTemplate = "FedEx_Padded_Pak"

	// 15.50 x 12.00 x 0.80 in
	FedExPak1 ParcelTemplate = "FedEx_Pak_1"

	// 12.75 x 10.25 x 0.80 in
	FedExPak2 ParcelTemplate = "FedEx_Pak_2"

	// 38.00 x 6.00 x 6.00 in
	FedExTube ParcelTemplate = "FedEx_Tube"

	// 17.50 x 20.75 x 2.00 in
	FedExXLPak ParcelTemplate = "FedEx_XL_Pak"
)

// UPS Templates
const (
	// 410.00 x 335.00 x 265.00 mm
	UPS10KgBox ParcelTemplate = "UPS_Box_10kg"

	// 484.00 x 433.00 x 350.00 mm
	UPS25KgBox ParcelTemplate = "UPS_Box_25kg"

	// 460.00 x 315.00 x 95.00 mm
	UPSExpressBox ParcelTemplate = "UPS_Express_Box"

	// 18.00 x 13.00 x 3.00 in
	UPSExpressLargeBox ParcelTemplate = "UPS_Express_Box_Large"

	// 15.00 x 11.00 x 3.00 in
	UPSExpressMediumBox ParcelTemplate = "UPS_Express_Box_Medium"

	// 13.00 x 11.00 x 2.00 in
	UPSExpressSmallBox ParcelTemplate = "UPS_Express_Box_Small"

	// 12.50 x 9.50 x 2.00 in
	UPSExpressEnvelope ParcelTemplate = "UPS_Express_Envelope"

	// 14.75 x 11.50 x 2.00 in
	UPSExpressHardPak ParcelTemplate = "UPS_Express_Hard_Pak"

	// 15.00 x 9.50 x 2.00 in
	UPSLegalEnvelope ParcelTemplate = "UPS_Express_Legal_Envelope"

	// 16.00 x 12.75 x 2.00 in
	UPSExpressPak ParcelTemplate = "UPS_Express_Pak"

	// 970.00 x 190.00 x 165.00 mm
	UPSExpressTube ParcelTemplate = "UPS_Express_Tube"

	// 17.25 x 12.75 x 2.00 in
	UPSLaboratoryPak ParcelTemplate = "UPS_Laboratory_Pak"

	// 0.00 x 0.00 x 0.00 in
	// BPM (Mail Innovations - Domestic & International)
	UPSMIBPM ParcelTemplate = "UPS_MI_BPM"

	// 0.00 x 0.00 x 0.00 in
	// BPM Flat (Mail Innovations - Domestic & International)
	UPSMIBPMFlat ParcelTemplate = "UPS_MI_BPM_Flat"

	// 0.00 x 0.00 x 0.00 in
	// BPM Parcel (Mail Innovations - Domestic & International)
	UPSMIBPMParcel ParcelTemplate = "UPS_MI_BPM_Parcel"

	// 0.00 x 0.00 x 0.00 in
	// First Class (Mail Innovations - Domestic only)
	UPSMIFirstClass ParcelTemplate = "UPS_MI_BPM_Parcel"

	// 0.00 x 0.00 x 0.00 in
	// Flat (Mail Innovations - Domestic only)
	UPSMIFlat ParcelTemplate = "UPS_MI_Flat"

	// 0.00 x 0.00 x 0.00 in
	// Irregular (Mail Innovations - Domestic only)
	UPSMIRegular ParcelTemplate = "UPS_MI_Irregular"

	// 0.00 x 0.00 x 0.00 in
	// Machinable (Mail Innovations - Domestic only)
	UPSMIMachinable ParcelTemplate = "UPS_MI_Machinable"

	// 0.00 x 0.00 x 0.00 in
	// Media Mail (Mail Innovations - Domestic only)
	UPSMIMediaMail ParcelTemplate = "UPS_MI_MEDIA_MAIL"

	// 0.00 x 0.00 x 0.00 in
	// Parcel Post (Mail Innovations - Domestic only)
	UPSMIParcelPost ParcelTemplate = "UPS_MI_Parcel_Post"

	// 0.00 x 0.00 x 0.00 in
	// Priority (Mail Innovations - Domestic only)
	UPSMIPriority ParcelTemplate = "UPS_MI_Priority"

	// 0.00 x 0.00 x 0.00 in
	// Standard Flat (Mail Innovations - Domestic only)
	UPSMIStandardFlat ParcelTemplate = "UPS_MI_Standard_Flat"

	// 14.75 x 11.00 x 2.00 in
	// Pad Pak
	UPSPadPak ParcelTemplate = "UPS_Pad_Pak"

	// 120.00 x 80.00 x 200.00 cm
	// Pallet
	UPSPallet ParcelTemplate = "UPS_Pallet"
)

// USPS parcel templates
const (
	// 12.50 x 9.50 x 0.75 in
	USPSFlatRateCardboardEnvelope ParcelTemplate = "USPS_FlatRateCardboardEnvelope"

	// 12.50 x 9.50 x 0.75 in
	USPSFlatRateEnvelope ParcelTemplate = "USPS_FlatRateEnvelope"

	// 10.00 x 7.00 x 0.75 in
	USPSFlatRateGiftCardEnvelope ParcelTemplate = "USPS_FlatRateGiftCardEnvelope"

	// 15.00 x 9.50 x 0.75 in
	USPSFlatRateLegalEnvelope ParcelTemplate = "USPS_FlatRateLegalEnvelope"

	// 12.50 x 9.50 x 1.00 in
	USPSFlatRatePaddedEnvelope ParcelTemplate = "USPS_FlatRatePaddedEnvelope"

	// 10.00 x 5.00 x 0.75 in
	USPSFlatRateWindowEnvelope ParcelTemplate = "USPS_FlatRateWindowEnvelope"

	// 0.00 x 0.00 x 0.00 in
	USPSIrregularParcel ParcelTemplate = "USPS_IrregularParcel"

	// 24.06 x 11.88 x 3.13 in
	USPSLargeFlatRateBoardGameBox ParcelTemplate = "USPS_LargeFlatRateBoardGameBox"

	// 12.25 x 12.25 x 6.00 in
	USPSLargeFlatRateBox ParcelTemplate = "USPS_LargeFlatRateBox"

	// 12.25 x 12.25 x 6.00 in
	USPSAPOFlatRateBox ParcelTemplate = "USPS_APOFlatRateBox"

	// 9.60 x 6.40 x 2.20 in
	USPSLargeVideoFlatRateBox ParcelTemplate = "USPS_LargeVideoFlatRateBox"

	// 11.25 x 8.75 x 6.00 in
	USPSMediumFlatRateBox1 ParcelTemplate = "USPS_MediumFlatRateBox1"

	// 14.00 x 12.00 x 3.50 in
	USPSMediumFlatRateBox2 ParcelTemplate = "USPS_MediumFlatRateBox2"

	// 10.13 x 7.13 x 5.00 in
	USPSRegionalRateBoxA1 ParcelTemplate = "USPS_RegionalRateBoxA1"

	// 13.06 x 11.06 x 2.50 in
	USPSRegionalRateBoxA2 ParcelTemplate = "USPS_RegionalRateBoxA2"

	// 12.25 x 10.50 x 5.50 in
	USPSRegionalRateBoxB1 ParcelTemplate = "USPS_RegionalRateBoxB1"

	// 16.25 x 14.50 x 3.00 in
	USPSRegionalRateBoxB2 ParcelTemplate = "USPS_RegionalRateBoxB2"

	// 8.69 x 5.44 x 1.75 in
	USPSSmallFlatRateBox ParcelTemplate = "USPS_SmallFlatRateBox"

	// 10.00 x 6.00 x 4.00 in
	USPSSmallFlatRateEnvelope ParcelTemplate = "USPS_SmallFlatRateEnvelope"
)

// DHL eCommerce parcel templates
const (
	// 10.00 x 10.00 x 10.00 in
	// Irregular Shipment
	DHLeCommerceIrregular ParcelTemplate = "DHLeC_Irregular"

	// 27.00 x 17.00 x 17.00 in
	// Flats
	DHLeCommerceFlats ParcelTemplate = "DHLeC_SM_Flats"
)
