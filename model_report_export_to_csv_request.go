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

// ReportExportToCsvRequest struct for ReportExportToCsvRequest
type ReportExportToCsvRequest struct {
	Params KalturaReportExportParams `json:"params"`
}

// NewReportExportToCsvRequest instantiates a new ReportExportToCsvRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportExportToCsvRequest(params KalturaReportExportParams) *ReportExportToCsvRequest {
	this := ReportExportToCsvRequest{}
	this.Params = params
	return &this
}

// NewReportExportToCsvRequestWithDefaults instantiates a new ReportExportToCsvRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportExportToCsvRequestWithDefaults() *ReportExportToCsvRequest {
	this := ReportExportToCsvRequest{}
	return &this
}

// GetParams returns the Params field value
func (o *ReportExportToCsvRequest) GetParams() KalturaReportExportParams {
	if o == nil {
		var ret KalturaReportExportParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ReportExportToCsvRequest) GetParamsOk() (*KalturaReportExportParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ReportExportToCsvRequest) SetParams(v KalturaReportExportParams) {
	o.Params = v
}

func (o ReportExportToCsvRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableReportExportToCsvRequest struct {
	value *ReportExportToCsvRequest
	isSet bool
}

func (v NullableReportExportToCsvRequest) Get() *ReportExportToCsvRequest {
	return v.value
}

func (v *NullableReportExportToCsvRequest) Set(val *ReportExportToCsvRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportExportToCsvRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportExportToCsvRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportExportToCsvRequest(val *ReportExportToCsvRequest) *NullableReportExportToCsvRequest {
	return &NullableReportExportToCsvRequest{value: val, isSet: true}
}

func (v NullableReportExportToCsvRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportExportToCsvRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

