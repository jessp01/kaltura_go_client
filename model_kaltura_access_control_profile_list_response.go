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

// KalturaAccessControlProfileListResponse struct for KalturaAccessControlProfileListResponse
type KalturaAccessControlProfileListResponse struct {
	KalturaListResponse
}

// NewKalturaAccessControlProfileListResponse instantiates a new KalturaAccessControlProfileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAccessControlProfileListResponse() *KalturaAccessControlProfileListResponse {
	this := KalturaAccessControlProfileListResponse{}
	return &this
}

// NewKalturaAccessControlProfileListResponseWithDefaults instantiates a new KalturaAccessControlProfileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAccessControlProfileListResponseWithDefaults() *KalturaAccessControlProfileListResponse {
	this := KalturaAccessControlProfileListResponse{}
	return &this
}

func (o KalturaAccessControlProfileListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAccessControlProfileListResponse struct {
	value *KalturaAccessControlProfileListResponse
	isSet bool
}

func (v NullableKalturaAccessControlProfileListResponse) Get() *KalturaAccessControlProfileListResponse {
	return v.value
}

func (v *NullableKalturaAccessControlProfileListResponse) Set(val *KalturaAccessControlProfileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAccessControlProfileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAccessControlProfileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAccessControlProfileListResponse(val *KalturaAccessControlProfileListResponse) *NullableKalturaAccessControlProfileListResponse {
	return &NullableKalturaAccessControlProfileListResponse{value: val, isSet: true}
}

func (v NullableKalturaAccessControlProfileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAccessControlProfileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


