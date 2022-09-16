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

// ReportGetCsvRequest struct for ReportGetCsvRequest
type ReportGetCsvRequest struct {
	ExcludedFields *string `json:"excludedFields,omitempty"`
	Id int32 `json:"id"`
	Params []KalturaKeyValue `json:"params,omitempty"`
}

// NewReportGetCsvRequest instantiates a new ReportGetCsvRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportGetCsvRequest(id int32) *ReportGetCsvRequest {
	this := ReportGetCsvRequest{}
	this.Id = id
	return &this
}

// NewReportGetCsvRequestWithDefaults instantiates a new ReportGetCsvRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportGetCsvRequestWithDefaults() *ReportGetCsvRequest {
	this := ReportGetCsvRequest{}
	return &this
}

// GetExcludedFields returns the ExcludedFields field value if set, zero value otherwise.
func (o *ReportGetCsvRequest) GetExcludedFields() string {
	if o == nil || o.ExcludedFields == nil {
		var ret string
		return ret
	}
	return *o.ExcludedFields
}

// GetExcludedFieldsOk returns a tuple with the ExcludedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetCsvRequest) GetExcludedFieldsOk() (*string, bool) {
	if o == nil || o.ExcludedFields == nil {
		return nil, false
	}
	return o.ExcludedFields, true
}

// HasExcludedFields returns a boolean if a field has been set.
func (o *ReportGetCsvRequest) HasExcludedFields() bool {
	if o != nil && o.ExcludedFields != nil {
		return true
	}

	return false
}

// SetExcludedFields gets a reference to the given string and assigns it to the ExcludedFields field.
func (o *ReportGetCsvRequest) SetExcludedFields(v string) {
	o.ExcludedFields = &v
}

// GetId returns the Id field value
func (o *ReportGetCsvRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportGetCsvRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReportGetCsvRequest) SetId(v int32) {
	o.Id = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ReportGetCsvRequest) GetParams() []KalturaKeyValue {
	if o == nil || o.Params == nil {
		var ret []KalturaKeyValue
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetCsvRequest) GetParamsOk() ([]KalturaKeyValue, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ReportGetCsvRequest) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given []KalturaKeyValue and assigns it to the Params field.
func (o *ReportGetCsvRequest) SetParams(v []KalturaKeyValue) {
	o.Params = v
}

func (o ReportGetCsvRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludedFields != nil {
		toSerialize["excludedFields"] = o.ExcludedFields
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableReportGetCsvRequest struct {
	value *ReportGetCsvRequest
	isSet bool
}

func (v NullableReportGetCsvRequest) Get() *ReportGetCsvRequest {
	return v.value
}

func (v *NullableReportGetCsvRequest) Set(val *ReportGetCsvRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportGetCsvRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportGetCsvRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportGetCsvRequest(val *ReportGetCsvRequest) *NullableReportGetCsvRequest {
	return &NullableReportGetCsvRequest{value: val, isSet: true}
}

func (v NullableReportGetCsvRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportGetCsvRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


