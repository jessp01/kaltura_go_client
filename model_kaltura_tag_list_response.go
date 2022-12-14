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

// KalturaTagListResponse struct for KalturaTagListResponse
type KalturaTagListResponse struct {
	KalturaListResponse
}

// NewKalturaTagListResponse instantiates a new KalturaTagListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTagListResponse() *KalturaTagListResponse {
	this := KalturaTagListResponse{}
	return &this
}

// NewKalturaTagListResponseWithDefaults instantiates a new KalturaTagListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTagListResponseWithDefaults() *KalturaTagListResponse {
	this := KalturaTagListResponse{}
	return &this
}

func (o KalturaTagListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaListResponse, errKalturaListResponse := json.Marshal(o.KalturaListResponse)
	if errKalturaListResponse != nil {
		return []byte{}, errKalturaListResponse
	}
	errKalturaListResponse = json.Unmarshal([]byte(serializedKalturaListResponse), &toSerialize)
	if errKalturaListResponse != nil {
		return []byte{}, errKalturaListResponse
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaTagListResponse struct {
	value *KalturaTagListResponse
	isSet bool
}

func (v NullableKalturaTagListResponse) Get() *KalturaTagListResponse {
	return v.value
}

func (v *NullableKalturaTagListResponse) Set(val *KalturaTagListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTagListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTagListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTagListResponse(val *KalturaTagListResponse) *NullableKalturaTagListResponse {
	return &NullableKalturaTagListResponse{value: val, isSet: true}
}

func (v NullableKalturaTagListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTagListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


