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

// MediaCountRequest struct for MediaCountRequest
type MediaCountRequest struct {
	Filter *KalturaMediaEntryFilter `json:"filter,omitempty"`
}

// NewMediaCountRequest instantiates a new MediaCountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaCountRequest() *MediaCountRequest {
	this := MediaCountRequest{}
	return &this
}

// NewMediaCountRequestWithDefaults instantiates a new MediaCountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaCountRequestWithDefaults() *MediaCountRequest {
	this := MediaCountRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MediaCountRequest) GetFilter() KalturaMediaEntryFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaMediaEntryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaCountRequest) GetFilterOk() (*KalturaMediaEntryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MediaCountRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaMediaEntryFilter and assigns it to the Filter field.
func (o *MediaCountRequest) SetFilter(v KalturaMediaEntryFilter) {
	o.Filter = &v
}

func (o MediaCountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableMediaCountRequest struct {
	value *MediaCountRequest
	isSet bool
}

func (v NullableMediaCountRequest) Get() *MediaCountRequest {
	return v.value
}

func (v *NullableMediaCountRequest) Set(val *MediaCountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaCountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaCountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaCountRequest(val *MediaCountRequest) *NullableMediaCountRequest {
	return &NullableMediaCountRequest{value: val, isSet: true}
}

func (v NullableMediaCountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaCountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


