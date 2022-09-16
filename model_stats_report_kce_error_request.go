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

// StatsReportKceErrorRequest struct for StatsReportKceErrorRequest
type StatsReportKceErrorRequest struct {
	KalturaCEError KalturaCEError `json:"kalturaCEError"`
}

// NewStatsReportKceErrorRequest instantiates a new StatsReportKceErrorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsReportKceErrorRequest(kalturaCEError KalturaCEError) *StatsReportKceErrorRequest {
	this := StatsReportKceErrorRequest{}
	this.KalturaCEError = kalturaCEError
	return &this
}

// NewStatsReportKceErrorRequestWithDefaults instantiates a new StatsReportKceErrorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsReportKceErrorRequestWithDefaults() *StatsReportKceErrorRequest {
	this := StatsReportKceErrorRequest{}
	return &this
}

// GetKalturaCEError returns the KalturaCEError field value
func (o *StatsReportKceErrorRequest) GetKalturaCEError() KalturaCEError {
	if o == nil {
		var ret KalturaCEError
		return ret
	}

	return o.KalturaCEError
}

// GetKalturaCEErrorOk returns a tuple with the KalturaCEError field value
// and a boolean to check if the value has been set.
func (o *StatsReportKceErrorRequest) GetKalturaCEErrorOk() (*KalturaCEError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KalturaCEError, true
}

// SetKalturaCEError sets field value
func (o *StatsReportKceErrorRequest) SetKalturaCEError(v KalturaCEError) {
	o.KalturaCEError = v
}

func (o StatsReportKceErrorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kalturaCEError"] = o.KalturaCEError
	}
	return json.Marshal(toSerialize)
}

type NullableStatsReportKceErrorRequest struct {
	value *StatsReportKceErrorRequest
	isSet bool
}

func (v NullableStatsReportKceErrorRequest) Get() *StatsReportKceErrorRequest {
	return v.value
}

func (v *NullableStatsReportKceErrorRequest) Set(val *StatsReportKceErrorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsReportKceErrorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsReportKceErrorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsReportKceErrorRequest(val *StatsReportKceErrorRequest) *NullableStatsReportKceErrorRequest {
	return &NullableStatsReportKceErrorRequest{value: val, isSet: true}
}

func (v NullableStatsReportKceErrorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsReportKceErrorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

