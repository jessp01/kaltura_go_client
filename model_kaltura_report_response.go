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

// KalturaReportResponse struct for KalturaReportResponse
type KalturaReportResponse struct {
	Columns *string `json:"columns,omitempty"`
	Results []KalturaString `json:"results,omitempty"`
}

// NewKalturaReportResponse instantiates a new KalturaReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaReportResponse() *KalturaReportResponse {
	this := KalturaReportResponse{}
	return &this
}

// NewKalturaReportResponseWithDefaults instantiates a new KalturaReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaReportResponseWithDefaults() *KalturaReportResponse {
	this := KalturaReportResponse{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *KalturaReportResponse) GetColumns() string {
	if o == nil || o.Columns == nil {
		var ret string
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportResponse) GetColumnsOk() (*string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *KalturaReportResponse) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given string and assigns it to the Columns field.
func (o *KalturaReportResponse) SetColumns(v string) {
	o.Columns = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *KalturaReportResponse) GetResults() []KalturaString {
	if o == nil || o.Results == nil {
		var ret []KalturaString
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportResponse) GetResultsOk() ([]KalturaString, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *KalturaReportResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []KalturaString and assigns it to the Results field.
func (o *KalturaReportResponse) SetResults(v []KalturaString) {
	o.Results = v
}

func (o KalturaReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaReportResponse struct {
	value *KalturaReportResponse
	isSet bool
}

func (v NullableKalturaReportResponse) Get() *KalturaReportResponse {
	return v.value
}

func (v *NullableKalturaReportResponse) Set(val *KalturaReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaReportResponse(val *KalturaReportResponse) *NullableKalturaReportResponse {
	return &NullableKalturaReportResponse{value: val, isSet: true}
}

func (v NullableKalturaReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

