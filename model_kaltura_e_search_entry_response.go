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

// KalturaESearchEntryResponse struct for KalturaESearchEntryResponse
type KalturaESearchEntryResponse struct {
	KalturaESearchResponse
}

// NewKalturaESearchEntryResponse instantiates a new KalturaESearchEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchEntryResponse() *KalturaESearchEntryResponse {
	this := KalturaESearchEntryResponse{}
	return &this
}

// NewKalturaESearchEntryResponseWithDefaults instantiates a new KalturaESearchEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchEntryResponseWithDefaults() *KalturaESearchEntryResponse {
	this := KalturaESearchEntryResponse{}
	return &this
}

func (o KalturaESearchEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchResponse, errKalturaESearchResponse := json.Marshal(o.KalturaESearchResponse)
	if errKalturaESearchResponse != nil {
		return []byte{}, errKalturaESearchResponse
	}
	errKalturaESearchResponse = json.Unmarshal([]byte(serializedKalturaESearchResponse), &toSerialize)
	if errKalturaESearchResponse != nil {
		return []byte{}, errKalturaESearchResponse
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchEntryResponse struct {
	value *KalturaESearchEntryResponse
	isSet bool
}

func (v NullableKalturaESearchEntryResponse) Get() *KalturaESearchEntryResponse {
	return v.value
}

func (v *NullableKalturaESearchEntryResponse) Set(val *KalturaESearchEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchEntryResponse(val *KalturaESearchEntryResponse) *NullableKalturaESearchEntryResponse {
	return &NullableKalturaESearchEntryResponse{value: val, isSet: true}
}

func (v NullableKalturaESearchEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


