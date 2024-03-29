// Copyright 2017 Inca Roads LLC. All rights reserved.
// Use of this source code is governed by licenses granted by the
// copyright holder including that found in the LICENSE file.

// Package ttdefs defines the core Safecast API data structures
package safecast

// Loc is Device Location Data
type Loc struct {
	Lat         *float64 `json:"loc_lat,omitempty"`
	Lon         *float64 `json:"loc_lon,omitempty"`
	Alt         *float64 `json:"loc_alt,omitempty"`
	MotionBegan *string  `json:"loc_when_motion_began,omitempty"`
	Olc         *string  `json:"loc_olc,omitempty"`
	LocName     *string  `json:"loc_name,omitempty"`
	LocCountry  *string  `json:"loc_country,omitempty"`
	LocZone     *string  `json:"loc_zone,omitempty"`
}

// Track is a device tracking structure
type Track struct {
	Lat      *float64 `json:"track_lat,omitempty"`
	Lon      *float64 `json:"track_lon,omitempty"`
	Distance *float64 `json:"track_distance,omitempty"`
	Seconds  *uint32  `json:"track_seconds,omitempty"`
	Velocity *float64 `json:"track_velocity,omitempty"`
	Bearing  *float64 `json:"track_bearing,omitempty"`
}

// Env is Device Basic Environmental Data
type Env struct {
	Temp  *float64 `json:"env_temp,omitempty"`
	Humid *float64 `json:"env_humid,omitempty"`
	Press *float64 `json:"env_press,omitempty"`
}

// Bat is Device Battery Performance Data
type Bat struct {
	Voltage  *float64 `json:"bat_voltage,omitempty"`
	Current  *float64 `json:"bat_current,omitempty"`
	Charge   *float64 `json:"bat_charge,omitempty"`
	Charging *bool    `json:"bat_charging,omitempty"`
	Line     *bool    `json:"bat_line,omitempty"`
}

// Lnd is support for LND Geiger Tubes
type Lnd struct {
	// Unshielded LND 7318
	U7318 *float64 `json:"lnd_7318u,omitempty"`
	// Shielded LND 7318
	C7318 *float64 `json:"lnd_7318c,omitempty"`
	// Energy-compensated LND 7128
	EC7128 *float64 `json:"lnd_7128ec,omitempty"`
	// Unshielded LND 712
	U712 *float64 `json:"lnd_712u,omitempty"`
	// Water-attenuated LND LND 78017
	W78017 *float64 `json:"lnd_78017w,omitempty"`
	// Device-computed uSV/h
	USv *float64 `json:"lnd_usv,omitempty"`
}

// AqiUSEPAHumidity (PM is not humidity-adjusted according to US)
const AqiUSEPAHumidity string = "us-epa-humidity"

// AqiCF1 (PM is "cf1" (which means no conversion))
const AqiCF1 string = "cf-1"

// AqiCFATM (PM is "atmostpheric" conversion factor)
const AqiCFATM string = "cf-atm"

// AqiCFEN481 (PM is European Standard EN-481)
const AqiCFEN481 string = "cf-en481"

// AqiLevelGood (golint)
const AqiLevelGood string = "good"

// AqiLevelModerate (golint)
const AqiLevelModerate string = "moderate"

// AqiLevelUnhealthyIfSensitive (unhealthy for sensitive groups)
const AqiLevelUnhealthyIfSensitive string = "unhealthy-if-sensitive"

// AqiLevelUnhealthy (golint)
const AqiLevelUnhealthy string = "unhealthy"

// AqiLevelVeryUnhealthy (golint)
const AqiLevelVeryUnhealthy string = "very-unhealthy"

// AqiLevelHazardous (golint)
const AqiLevelHazardous string = "hazardous"

// AqiLevelVeryHazardous (golint)
const AqiLevelVeryHazardous string = "very-hazardous"

// Pms is for Plantower Air Sensor Data
type Pms struct {
	Pm01_0     *float64 `json:"pms_pm01_0,omitempty"`
	Pm02_5     *float64 `json:"pms_pm02_5,omitempty"`
	Pm10_0     *float64 `json:"pms_pm10_0,omitempty"`
	Std01_0    *float64 `json:"pms_std01_0,omitempty"`
	Std02_5    *float64 `json:"pms_std02_5,omitempty"`
	Std10_0    *float64 `json:"pms_std10_0,omitempty"`
	Count00_30 *uint32  `json:"pms_c00_30,omitempty"`
	Count00_50 *uint32  `json:"pms_c00_50,omitempty"`
	Count01_00 *uint32  `json:"pms_c01_00,omitempty"`
	Count02_50 *uint32  `json:"pms_c02_50,omitempty"`
	Count05_00 *uint32  `json:"pms_c05_00,omitempty"`
	Count10_00 *uint32  `json:"pms_c10_00,omitempty"`
	CountSecs  *uint32  `json:"pms_csecs,omitempty"`
	Samples    *uint32  `json:"pms_csamples,omitempty"`
	Pm01_0cf1  *float64 `json:"pms_pm01_0_cf1,omitempty"`
	Pm02_5cf1  *float64 `json:"pms_pm02_5_cf1,omitempty"`
	Pm10_0cf1  *float64 `json:"pms_pm10_0_cf1,omitempty"`
	Model      *string  `json:"pms_model,omitempty"`
	AqiNotes   *string  `json:"pms_aqi_notes,omitempty"`
	AqiLevel   *string  `json:"pms_aqi_level,omitempty"`
	AqiPm      *float64 `json:"pms_aqi_pm,omitempty"`
	Aqi        *uint32  `json:"pms_aqi,omitempty"`
}

// Pms2 is for an auxiliary Plantower Air Sensor Data
type Pms2 struct {
	Pm01_0     *float64 `json:"pms2_pm01_0,omitempty"`
	Pm02_5     *float64 `json:"pms2_pm02_5,omitempty"`
	Pm10_0     *float64 `json:"pms2_pm10_0,omitempty"`
	Std01_0    *float64 `json:"pms2_std01_0,omitempty"`
	Std02_5    *float64 `json:"pms2_std02_5,omitempty"`
	Std10_0    *float64 `json:"pms2_std10_0,omitempty"`
	Count00_30 *uint32  `json:"pms2_c00_30,omitempty"`
	Count00_50 *uint32  `json:"pms2_c00_50,omitempty"`
	Count01_00 *uint32  `json:"pms2_c01_00,omitempty"`
	Count02_50 *uint32  `json:"pms2_c02_50,omitempty"`
	Count05_00 *uint32  `json:"pms2_c05_00,omitempty"`
	Count10_00 *uint32  `json:"pms2_c10_00,omitempty"`
	CountSecs  *uint32  `json:"pms2_csecs,omitempty"`
	Samples    *uint32  `json:"pms2_csamples,omitempty"`
	Pm01_0cf1  *float64 `json:"pms2_pm01_0_cf1,omitempty"`
	Pm02_5cf1  *float64 `json:"pms2_pm02_5_cf1,omitempty"`
	Pm10_0cf1  *float64 `json:"pms2_pm10_0_cf1,omitempty"`
	Model      *string  `json:"pms2_model,omitempty"`
	AqiNotes   *string  `json:"pms2_aqi_notes,omitempty"`
	AqiLevel   *string  `json:"pms2_aqi_level,omitempty"`
	AqiPm      *float64 `json:"pms2_aqi_pm,omitempty"`
	Aqi        *uint32  `json:"pms2_aqi,omitempty"`
}

// Opc is for Alphasense OPC-N2 Air Sensor Data
type Opc struct {
	Pm01_0     *float64 `json:"opc_pm01_0,omitempty"`
	Pm02_5     *float64 `json:"opc_pm02_5,omitempty"`
	Pm10_0     *float64 `json:"opc_pm10_0,omitempty"`
	Std01_0    *float64 `json:"opc_std01_0,omitempty"`
	Std02_5    *float64 `json:"opc_std02_5,omitempty"`
	Std10_0    *float64 `json:"opc_std10_0,omitempty"`
	Count00_38 *uint32  `json:"opc_c00_38,omitempty"`
	Count00_54 *uint32  `json:"opc_c00_54,omitempty"`
	Count01_00 *uint32  `json:"opc_c01_00,omitempty"`
	Count02_10 *uint32  `json:"opc_c02_10,omitempty"`
	Count05_00 *uint32  `json:"opc_c05_00,omitempty"`
	Count10_00 *uint32  `json:"opc_c10_00,omitempty"`
	CountSecs  *uint32  `json:"opc_csecs,omitempty"`
	Samples    *uint32  `json:"opc_csamples,omitempty"`
	Pm01_0cf1  *float64 `json:"opc,omitempty"`
	Pm02_5cf1  *float64 `json:"opc_pm02_5_cf1,omitempty"`
	Pm10_0cf1  *float64 `json:"opc_pm10_0_cf1,omitempty"`
	AqiNotes   *string  `json:"opc_aqi_notes,omitempty"`
	AqiLevel   *string  `json:"opc_aqi_level,omitempty"`
	AqiPm      *float64 `json:"opc_aqi_pm,omitempty"`
	Aqi        *uint32  `json:"opc_aqi,omitempty"`
}

// Dev contains General Device Statistics
type Dev struct {
	Test                  *bool    `json:"dev_test,omitempty"`
	Indoors               *bool    `json:"dev_indoors,omitempty"`
	Motion                *bool    `json:"dev_motion,omitempty"`
	DeviceLabel           *string  `json:"dev_label,omitempty"`
	UptimeMinutes         *uint32  `json:"dev_uptime,omitempty"`
	AppVersion            *string  `json:"dev_firmware,omitempty"`
	DeviceParams          *string  `json:"dev_cfgdev,omitempty"`
	ServiceParams         *string  `json:"dev_cfgsvc,omitempty"`
	TtnParams             *string  `json:"dev_cfgttn,omitempty"`
	GpsParams             *string  `json:"dev_cfggps,omitempty"`
	SensorParams          *string  `json:"dev_cfgsen,omitempty"`
	TransmittedBytes      *uint32  `json:"dev_transmitted_bytes,omitempty"`
	ReceivedBytes         *uint32  `json:"dev_received_bytes,omitempty"`
	CommsResets           *uint32  `json:"dev_comms_resets,omitempty"`
	CommsFails            *uint32  `json:"dev_comms_failures,omitempty"`
	CommsPowerFails       *uint32  `json:"dev_comms_power_fails,omitempty"`
	DeviceRestarts        *uint32  `json:"dev_restarts,omitempty"`
	MotionEvents          *uint32  `json:"dev_motion_events,omitempty"`
	Oneshots              *uint32  `json:"dev_oneshots,omitempty"`
	OneshotSeconds        *uint32  `json:"dev_oneshot_seconds,omitempty"`
	Iccid                 *string  `json:"dev_iccid,omitempty"`
	Cpsi                  *string  `json:"dev_cpsi,omitempty"`
	Dfu                   *string  `json:"dev_dfu,omitempty"`
	FreeMem               *uint32  `json:"dev_free_memory,omitempty"`
	NTPCount              *uint32  `json:"dev_ntp_count,omitempty"`
	LastFailure           *string  `json:"dev_last_failure,omitempty"`
	Status                *string  `json:"dev_status,omitempty"`
	ModuleLora            *string  `json:"dev_module_lora,omitempty"`
	ModuleFona            *string  `json:"dev_module_fona,omitempty"`
	Temp                  *float64 `json:"dev_temp,omitempty"`
	Humid                 *float64 `json:"dev_humid,omitempty"`
	Press                 *float64 `json:"dev_press,omitempty"`
	Rat                   *string  `json:"dev_rat,omitempty"`
	Bars                  *uint32  `json:"dev_bars,omitempty"`
	CommsAntFails         *uint32  `json:"dev_comms_ant_fails,omitempty"`
	OvercurrentEvents     *uint32  `json:"dev_overcurrent_events,omitempty"`
	Seqno                 *uint32  `json:"dev_seqno,omitempty"`
	Moved                 *string  `json:"dev_moved,omitempty"`
	Orientation           *string  `json:"dev_orientation,omitempty"`
	Dashboard             *string  `json:"dev_dashboard,omitempty"`
	ErrorsMtu             *uint32  `json:"dev_err_mtu,omitempty"`
	ErrorsOpc             *uint32  `json:"dev_err_opc,omitempty"`
	ErrorsPms             *uint32  `json:"dev_err_pms,omitempty"`
	ErrorsPms2            *uint32  `json:"dev_err_pms2,omitempty"`
	ErrorsBme0            *uint32  `json:"dev_err_bme0,omitempty"`
	ErrorsBme1            *uint32  `json:"dev_err_bme1,omitempty"`
	ErrorsLora            *uint32  `json:"dev_err_lora,omitempty"`
	ErrorsFona            *uint32  `json:"dev_err_fona,omitempty"`
	ErrorsGeiger          *uint32  `json:"dev_err_geiger,omitempty"`
	ErrorsMax01           *uint32  `json:"dev_err_max01,omitempty"`
	ErrorsUgps            *uint32  `json:"dev_err_ugps,omitempty"`
	ErrorsTwi             *uint32  `json:"dev_err_twi,omitempty"`
	ErrorsTwiInfo         *string  `json:"dev_err_twi_info,omitempty"`
	ErrorsLis             *uint32  `json:"dev_err_lis,omitempty"`
	ErrorsSpi             *uint32  `json:"dev_err_spi,omitempty"`
	ErrorsConnectLora     *uint32  `json:"dev_err_con_lora,omitempty"`
	ErrorsConnectFona     *uint32  `json:"dev_err_con_fona,omitempty"`
	ErrorsConnectWireless *uint32  `json:"dev_err_con_wireless,omitempty"`
	ErrorsConnectData     *uint32  `json:"dev_err_con_data,omitempty"`
	ErrorsConnectService  *uint32  `json:"dev_err_con_service,omitempty"`
	ErrorsConnectGateway  *uint32  `json:"dev_err_con_gateway,omitempty"`
}

// Gateway is Lora home gateway-supplied metadata
type Gateway struct {
	ReceivedAt *string  `json:"gateway_received,omitempty"`
	SNR        *float64 `json:"gateway_lora_snr,omitempty"`
	Lat        *float64 `json:"gateway_loc_lat,omitempty"`
	Lon        *float64 `json:"gateway_loc_lon,omitempty"`
	Alt        *float64 `json:"gateway_loc_alt,omitempty"`
}

// Service contains service metadata
type Service struct {
	UploadedAt *string `json:"service_uploaded,omitempty"`
	Transport  *string `json:"service_transport,omitempty"`
	Handler    *string `json:"service_handler,omitempty"`
}

// Note that this structure has been designed so that we could convert, at a later date,
// to a structured JSON out put by modifying these definitions by changing this of this form:
//    *Location `json:",omitempty"`
// to this form, using the data type as the fiel name and specifying a json field name..
//	  Location *Location `json:"location,omitempty"`

// SafecastData is our primary in-memory data structure for a Safecast message
type SafecastData struct {

	// This new device ID has replaced DeviceID as our "master" device identifier
	// because of the fact that DeviceID is only 32-bits and will eventually
	// have conflicts.
	DeviceUID string `json:"device_urn,omitempty"`

	// This string identifies, at a high level, what type of device this is, so
	// that from an experience perspective we know roughly what kind of UI is
	// appropriate for this device's data.  The format of this string is
	// to be either just a text string ("pointcast" or anything resembling
	// URN rules ("product:com.coca-cola.vending-machine.v3")
	DeviceClass string `json:"device_class,omitempty"`

	// An optional device serial number string which is DeviceClass-specific
	DeviceSN string `json:"device_sn,omitempty"`

	// An optional contact for the device itself
	DeviceContactName  string `json:"device_contact_name,omitempty"`
	DeviceContactOrg   string `json:"device_contact_org,omitempty"`
	DeviceContactRole  string `json:"device_contact_role,omitempty"`
	DeviceContactEmail string `json:"device_contact_email,omitempty"`

	// Optional native JSON information to retain in persistent storage for the event
	Native *map[string]interface{} `json:"native,omitempty"`

	// Data generated by the device itself and untouched in transit
	DeviceID   uint32  `json:"device,omitempty"`
	CapturedAt *string `json:"when_captured,omitempty"`
	*Loc       `json:",omitempty"`
	*Env       `json:",omitempty"`
	*Bat       `json:",omitempty"`
	*Lnd       `json:",omitempty"`
	*Pms       `json:",omitempty"`
	*Pms2      `json:",omitempty"`
	*Opc       `json:",omitempty"`
	*Dev       `json:",omitempty"`
	*Track     `json:",omitempty"`

	// Metadata added as the above is being
	*Gateway `json:",omitempty"`
	*Service `json:",omitempty"`
}
