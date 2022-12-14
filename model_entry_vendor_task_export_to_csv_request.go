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

// EntryVendorTaskExportToCsvRequest struct for EntryVendorTaskExportToCsvRequest
type EntryVendorTaskExportToCsvRequest struct {
	Filter KalturaEntryVendorTaskFilter `json:"filter"`
}

// NewEntryVendorTaskExportToCsvRequest instantiates a new EntryVendorTaskExportToCsvRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryVendorTaskExportToCsvRequest(filter KalturaEntryVendorTaskFilter) *EntryVendorTaskExportToCsvRequest {
	this := EntryVendorTaskExportToCsvRequest{}
	this.Filter = filter
	return &this
}

// NewEntryVendorTaskExportToCsvRequestWithDefaults instantiates a new EntryVendorTaskExportToCsvRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryVendorTaskExportToCsvRequestWithDefaults() *EntryVendorTaskExportToCsvRequest {
	this := EntryVendorTaskExportToCsvRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *EntryVendorTaskExportToCsvRequest) GetFilter() KalturaEntryVendorTaskFilter {
	if o == nil {
		var ret KalturaEntryVendorTaskFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *EntryVendorTaskExportToCsvRequest) GetFilterOk() (*KalturaEntryVendorTaskFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *EntryVendorTaskExportToCsvRequest) SetFilter(v KalturaEntryVendorTaskFilter) {
	o.Filter = v
}

func (o EntryVendorTaskExportToCsvRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableEntryVendorTaskExportToCsvRequest struct {
	value *EntryVendorTaskExportToCsvRequest
	isSet bool
}

func (v NullableEntryVendorTaskExportToCsvRequest) Get() *EntryVendorTaskExportToCsvRequest {
	return v.value
}

func (v *NullableEntryVendorTaskExportToCsvRequest) Set(val *EntryVendorTaskExportToCsvRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryVendorTaskExportToCsvRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryVendorTaskExportToCsvRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryVendorTaskExportToCsvRequest(val *EntryVendorTaskExportToCsvRequest) *NullableEntryVendorTaskExportToCsvRequest {
	return &NullableEntryVendorTaskExportToCsvRequest{value: val, isSet: true}
}

func (v NullableEntryVendorTaskExportToCsvRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryVendorTaskExportToCsvRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


