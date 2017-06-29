package goshippo

type ServiceLevel string

const (
	// Priority Mail
	USPSPriority ServiceLevel = "usps_priority"

	// Priority Mail Express
	USPSPriorityExpress ServiceLevel = "usps_priority_express"

	// First Class mail/package
	USPSFirstClass ServiceLevel = "usps_first"

	// Parcel Select
	USPSParcelSelect ServiceLevel = "usps_parcel_select"

	// Media Mail, only for existing Shippo
	// customers with grandfathered Media Mail option.
	USPSMediaMail ServiceLevel = "usps_media_mail"

	// Media Mail, only for existing Shippo
	// customers with grandfathered Media Mail option.
	USPSPriorityMailInternational ServiceLevel = "usps_priority_mail_international"
	USPSPriorityMailExpress       ServiceLevel = "usps_priority_mail_express"
	USPSFirstClassPackage         ServiceLevel = "usps_first_class_package"
)

// Fedex
const (
	FedexGround                   ServiceLevel = "fedex_ground"
	FedexHomeDelivery             ServiceLevel = "fedex_home_delivery"
	FedexSmartPost                ServiceLevel = "fedex_smart_post"
	Fedex2Day                     ServiceLevel = "fedex_2_day"
	Fedex2DayAM                   ServiceLevel = "fedex_2_day_am"
	FedexExpressSaver             ServiceLevel = "fedex_express_saver"
	FedexStandardOvernight        ServiceLevel = "fedex_standard_overnight"
	FedexPriorityOvernight        ServiceLevel = "fedex_priority_overnight"
	FedexFirstOvernight           ServiceLevel = "fedex_priority_overnight"
	FedexInternationalEconomy     ServiceLevel = "fedex_international_economy"
	FedexInternationalPriority    ServiceLevel = "fedex_international_priority"
	FedexInternationalFirst       ServiceLevel = "fedex_international_first"
	FedexEuropeFirstInternational ServiceLevel = "fedex_europe_first_international"
)

// UPS
const (
	UPSStandard                ServiceLevel = "ups_standard"
	UPSGround                  ServiceLevel = "ups_ground"
	UPSSaver                   ServiceLevel = "ups_saver"
	UPS3DaySelect              ServiceLevel = "ups_3_day_select"
	UPSSecondDayAir            ServiceLevel = "ups_second_day_air"
	UPSSecondDayAirAM          ServiceLevel = "ups_second_day_air_am"
	UPSNextDayAir              ServiceLevel = "ups_next_day_air"
	UPSNextDayAirSaver         ServiceLevel = "ups_next_day_air_saver"
	UPSNextDayAirEarlyAM       ServiceLevel = "ups_next_day_air_early_am"
	UPSMailInnovationsDomestic ServiceLevel = "ups_mail_innovations_domestic"
	UPSSurePost                ServiceLevel = "ups_surepost"
	UPSSurePostLightweight     ServiceLevel = "ups_surepost_lightweight"
	UPSExpress                 ServiceLevel = "ups_express"
	UPSExpressPlus             ServiceLevel = "ups_express_plus"
	UPSExpedited               ServiceLevel = "ups_expedited"
)

// Asendia
const (
	AsendiaUSPriorityTracked                      ServiceLevel = "asendia_us_priority_tracked"
	AsendiaUSPriorityInternational                ServiceLevel = "asendia_us_priority_international"
	AsendiaUSPriorityInternationalPriorityAirmail ServiceLevel = "asendia_us_priority_international_priority_airmail"
	AsendiaUSPriorityInternationalSurfaceAirlift  ServiceLevel = "asendia_us_priority_international_surface_airlift"
	AsendiaUSPriorityMailInternational            ServiceLevel = "asendia_us_priority_mail_international"
	AsendiaUSPriorityMailExpressInternational     ServiceLevel = "asendia_us_priority_mail_express_international"

	AsendiaUSEPacket ServiceLevel = "asendia_us_epacket"
	AsendiaUSOther   ServiceLevel = "asendia_us_other"
)

// AustraliaPost
const (
	AustraliaPostExpressPost ServiceLevel = "australia_post_express_post"
	AustraliaPostParcelPost  ServiceLevel = "australia_post_parcel_post"
)

// CanadaPost
const (
	CanadaPostRegularParcel               ServiceLevel = "canada_post_regular_parcel"
	CanadaPostExpeditedParcel             ServiceLevel = "canada_post_expedited_parcel"
	CanadaPostPriority                    ServiceLevel = "canada_post_priority"
	CanadaPostXpressPost                  ServiceLevel = "canada_post_xpress_post"
	CanadaPostXpressPostInternational     ServiceLevel = "canada_post_xpress_post_international"
	CanadaPostXpressPostUSA               ServiceLevel = "canada_post_xpress_post_usa"
	CanadaPostTrackedPacketUSA            ServiceLevel = "canada_post_tracked_packet_usa"
	CanadaPostTrackedPacketInternational  ServiceLevel = "canada_post_tracked_packet_international"
	CanadaPostSmallPacketInternationalAir ServiceLevel = "canada_post_small_packet_international_air"
)

// Purolator
const (
	PurolatorGround                     ServiceLevel = "purolator_ground"
	PurolatorGround9AM                  ServiceLevel = "purolator_ground9_am"
	PurolatorGround1030AM               ServiceLevel = "purolator_ground1030_am"
	PurolatorGroundDistribution         ServiceLevel = "purolator_ground_distribution"
	PurolatorGroundEvening              ServiceLevel = "purolator_ground_evening"
	PurolatorGroundUS                   ServiceLevel = "purolator_ground_us"
	PurolatorExpress                    ServiceLevel = "purolator_express"
	PurolatorExpress9AM                 ServiceLevel = "purolator_express9_am"
	PurolatorExpress1030AM              ServiceLevel = "purolator_express1030_am"
	PurolatorExpressEvening             ServiceLevel = "purolator_express_evening"
	PurolatorExpressUS9AM               ServiceLevel = "purolator_express_us9_am"
	PurolatorExpressUS1030AM            ServiceLevel = "purolator_express_us1030_am"
	PurolatorExpressUS12PM              ServiceLevel = "purolator_express_us1200"
	PurolatorExpressInternational       ServiceLevel = "purolator_express_international"
	PurolatorExpressInternational9AM    ServiceLevel = "purolator_express_international9_am"
	PurolatorExpressInternational1030AM ServiceLevel = "purolator_express_international1030_am"
	PurolatorExpressInternational12PM   ServiceLevel = "purolator_express_international1200"
)

// DHL Express
const (
	DHLExpressDomesticExpressDoc        ServiceLevel = "dhl_express_domestic_express_doc"
	DHLExpressEconomySelectDoc          ServiceLevel = "dhl_express_economy_select_doc"
	DHLExpressWorldwideNonDoc           ServiceLevel = "dhl_express_worldwide_nondoc"
	DHLExpressWorldwideDoc              ServiceLevel = "dhl_express_worldwide_doc"
	DHLExpressWorldwide                 ServiceLevel = "dhl_express_worldwide"
	DHLExpressBreakBulkExpressDoc       ServiceLevel = "dhl_express_break_bulk_express_doc"
	DHLExpressBreakBulkExpress9AMNonDoc ServiceLevel = "dhl_express_break_bulk_express_9_00_nondoc"
	DHLExpressEconomySelectNonDoc       ServiceLevel = "dhl_express_economy_select_nondoc"
	DHLExpressBreakBulkEconomyDoc       ServiceLevel = "dhl_express_break_bulk_economy_doc"
	DHLExpress9AMDoc                    ServiceLevel = "dhl_express_express_9_00_doc"
	DHLExpress1030AMDoc                 ServiceLevel = "dhl_express_express_10_30_doc"
	DHLExpress1030AMNonDoc              ServiceLevel = "dhl_express_express_10_30_nondoc"
	DHLExpress12PMDoc                   ServiceLevel = "dhl_express_express_12_00_doc"
	DHLExpressEuropackNonDoc            ServiceLevel = "dhl_express_europack_nondoc"
	DHLExpressEnvelopeDoc               ServiceLevel = "dhl_express_express_envelope_doc"
	DHLExpress12PMNonDoc                ServiceLevel = "dhl_express_express_12_00_nondoc"
	DHLExpressWorldwideB2CDoc           ServiceLevel = "dhl_express_worldwide_b2c_doc"
	DHLExpressWorldwideB2CNonDoc        ServiceLevel = "dhl_express_worldwide_b2c_nondoc"
	DHLExpressEasyNonDoc                ServiceLevel = "dhl_express_express_easy_nondoc"
)

// DHL eCommerce
const (
	DHLeCommerceMarketingParcelExpedited      ServiceLevel = "dhl_ecommerce_marketing_parcel_expedited"
	DHLeCommerceParcelInternationalExpedited  ServiceLevel = "dhl_ecommerce_globalmail_business_ips"
	DHLeCommerceGlobalMailBusinessStandard    ServiceLevel = "dhl_ecommerce_parcel_international_direct"
	DHLeCommerceParcelsExpeditedMax           ServiceLevel = "dhl_ecommerce_parcels_expedited_max"
	DHLeCommerceBoundedPrintedMatterGround    ServiceLevel = "dhl_ecommerce_bpm_ground"
	DHLeCommercePriorityExpedited             ServiceLevel = "dhl_ecommerce_priority_expedited"
	DHLeCommerceGlobalMailPacketPriority      ServiceLevel = "dhl_ecommerce_globalmail_packet_ipa"
	DHLeCommerceGlobalMailPacketStandard      ServiceLevel = "dhl_ecommerce_globalmail_packet_isal"
	DHLeCommerceMarketingParcelGround         ServiceLevel = "dhl_ecommerce_marketing_parcel_ground"
	DHLeCommerceFirstClassParcelExpedited     ServiceLevel = "dhl_ecommerce_first_class_parcel_expedited"
	DHLeCommerceParcelInternationalStandard   ServiceLevel = "dhl_ecommerce_globalmail_business_priority"
	DHLeCommerceParcelsExpedited              ServiceLevel = "dhl_ecommerce_parcels_expedited"
	DHLeCommerceParcelInternationalDirect     ServiceLevel = "dhl_ecommerce_globalmail_business_isal"
	DHLeCommerceParcelPlusExpeditedMax        ServiceLevel = "dhl_ecommerce_parcel_plus_expedited_max"
	DHLeCommerceGlobalMailPacketIPA           ServiceLevel = "dhl_ecommerce_globalmail_packet_plus"
	DHLeCommerceParcelsGround                 ServiceLevel = "dhl_ecommerce_parcels_ground"
	DHLeCommerceExpedited                     ServiceLevel = "dhl_ecommerce_expedited"
	DHLeCommerceParcelPlusGround              ServiceLevel = "dhl_ecommerce_parcel_plus_ground"
	DHLeCommerceGlobalMailBusinessISAL        ServiceLevel = "dhl_ecommerce_parcel_international_standard"
	DHLeCommerceBoundedPrintedMatterExpedited ServiceLevel = "dhl_ecommerce_bpm_expedited"
	DHLeCommerceGlobalMailBusinessIPA         ServiceLevel = "dhl_ecommerce_parcel_international_expedited"
	DHLeCommerceGlobalMailPacketISAL          ServiceLevel = "dhl_ecommerce_globalmail_packet_priority"
	DHLeCommerceEasyReturnLight               ServiceLevel = "dhl_ecommerce_easy_return_light"
	DHLeCommerceParcelPlusExpedited           ServiceLevel = "dhl_ecommerce_parcel_plus_expedited"
	DHLeCommerceGlobalMailPacketPlus          ServiceLevel = "dhl_ecommerce_globalmail_business_standard"
	DHLeCommerceGround                        ServiceLevel = "dhl_ecommerce_ground"
	DHLeCommerceGlobalMailBusinessPriority    ServiceLevel = "dhl_ecommerce_globalmail_packet_standard"
)

// DHL Germany
const (
	DHLGermanyPaketBusiness     ServiceLevel = "dhl_paket_business"
	DHLGermanyWeltpaketBusiness ServiceLevel = "dhl_weltpaket_business"
	DHLGermanyEuropaketBusiness ServiceLevel = "dhl_europaket_business"
)

// Deutsche Post
const (
	DeutschePostPostkarte     ServiceLevel = "deutsche_post_postkarte"
	DeutschePostStandardbrief ServiceLevel = "deutsche_post_standardbrief"
	DeutschePostKompakbrief   ServiceLevel = "deutsche_post_kompakbrief"
	DeutschePostGrossbrief    ServiceLevel = "deutsche_post_grossbrief"
	DeutschePostMaxibrief     ServiceLevel = "deutsche_post_maxibrief"
	DeutschePostMaxibriefPlus ServiceLevel = "deutsche_post_maxibrief_plus"
)

// GLS Germany
const (
	GLSGermanyBusinessParcel ServiceLevel = "gls_deutschland_business_parcel"
)

// GLS France
const (
	GLSFranceBusinessParcel ServiceLevel = "gls_france_business_parcel"
)

// Mondial Relay
const (
	MondialRelayPointRelais ServiceLevel = "mondial_relay_pointrelais"
)

// Parcelforce
const (
	ParcelForceExpress48 ServiceLevel = "parcelforce_express48"
	ParcelForceExpress24 ServiceLevel = "parcelforce_express24"
	ParcelForceExpressAM ServiceLevel = "parcelforce_expressam"
)

// RR Donnelley
const (
	RRDonnelleyDomesticEconomyParcel                ServiceLevel = "rr_donnelley_domestic_economy_parcel"
	RRDonnelleyDomesticPriorityParcel               ServiceLevel = "rr_donnelley_domestic_priority_parcel"
	RRDonnelleyDomesticParcelBPM                    ServiceLevel = "rr_donnelley_domestic_parcel_bpm"
	RRDonnelleyDomesticPriorityParcelBPM            ServiceLevel = "rr_donnelley_priority_domestic_priority_parcel_bpm"
	RRDonnelleyInternationalPriorityParcelDelcon    ServiceLevel = "rr_donnelley_priority_parcel_delcon"
	RRDonnelleyInternationalPriorityParcelNonDelcon ServiceLevel = "rr_donnelley_priority_parcel_nondelcon"
	RRDonnelleyEconomyParcel                        ServiceLevel = "rr_donnelley_economy_parcel"
	RRDonnelleyInternationalParcelAirmail           ServiceLevel = "rr_donnelley_ipa"
	RRDonnelleyInternationalCourier                 ServiceLevel = "rr_donnelley_courier"
	RRDonnelleyInternationalSurfaceAirLift          ServiceLevel = "rr_donnelley_isal"
	RRDonnelleyePacket                              ServiceLevel = "rr_donnelley_epacket"
	RRDonnelleyPriorityMailInternational            ServiceLevel = "rr_donnelley_pmi"
	RRDonnelleyExpressMailInternational             ServiceLevel = "rr_donnelley_emi"
)

// Newgistics
const (
	NewgisticsParcelSelectLightweight ServiceLevel = "newgistics_parcel_select_lightweight"
	NewgisticsParcelSelect            ServiceLevel = "newgistics_parcel_select"
	NewgisticsPriorityMail            ServiceLevel = "newgistics_priority_mail"
	NewgisticsFirstClassMail          ServiceLevel = "newgistics_first_class_mail"
)

// OnTrac
const (
	OnTracGround      ServiceLevel = "ontrac_ground"
	OnTracSunriseGold ServiceLevel = "ontrac_sunrise_gold"
	OnTracSunrise     ServiceLevel = "ontrac_sunrise"
)

// Lasership
const (
	LasershipRoutedDelivery ServiceLevel = "lasership_routed_delivery"
)

// UberRush
const (
	UberRushOnDemand ServiceLevel = "uber_on_demand"
)

// Hermes UK
const (
	HermesUKParcelShop   ServiceLevel = "hermes_uk_parcelshop"
	HermesUKHomeDelivery ServiceLevel = "hermes_uk_home_delivery"
)
