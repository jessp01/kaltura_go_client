/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatsReportDeviceCapabilitiesRequest struct for StatsReportDeviceCapabilitiesRequest
type StatsReportDeviceCapabilitiesRequest struct {
	Data string `json:"data"`
}

// NewStatsReportDeviceCapabilitiesRequest instantiates a new StatsReportDeviceCapabilitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsReportDeviceCapabilitiesRequest(data string) *StatsReportDeviceCapabilitiesRequest {
	this := StatsReportDeviceCapabilitiesRequest{}
	this.Data = data
	return &this
}

// NewStatsReportDeviceCapabilitiesRequestWithDefaults instantiates a new StatsReportDeviceCapabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsReportDeviceCapabilitiesRequestWithDefaults() *StatsReportDeviceCapabilitiesRequest {
	this := StatsReportDeviceCapabilitiesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *StatsReportDeviceCapabilitiesRequest) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StatsReportDeviceCapabilitiesRequest) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StatsReportDeviceCapabilitiesRequest) SetData(v string) {
	o.Data = v
}

func (o StatsReportDeviceCapabilitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStatsReportDeviceCapabilitiesRequest struct {
	value *StatsReportDeviceCapabilitiesRequest
	isSet bool
}

func (v NullableStatsReportDeviceCapabilitiesRequest) Get() *StatsReportDeviceCapabilitiesRequest {
	return v.value
}

func (v *NullableStatsReportDeviceCapabilitiesRequest) Set(val *StatsReportDeviceCapabilitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsReportDeviceCapabilitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsReportDeviceCapabilitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsReportDeviceCapabilitiesRequest(val *StatsReportDeviceCapabilitiesRequest) *NullableStatsReportDeviceCapabilitiesRequest {
	return &NullableStatsReportDeviceCapabilitiesRequest{value: val, isSet: true}
}

func (v NullableStatsReportDeviceCapabilitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsReportDeviceCapabilitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

