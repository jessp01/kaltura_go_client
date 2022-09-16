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

// KalturaAccessControlListResponse struct for KalturaAccessControlListResponse
type KalturaAccessControlListResponse struct {
	KalturaListResponse
}

// NewKalturaAccessControlListResponse instantiates a new KalturaAccessControlListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAccessControlListResponse() *KalturaAccessControlListResponse {
	this := KalturaAccessControlListResponse{}
	return &this
}

// NewKalturaAccessControlListResponseWithDefaults instantiates a new KalturaAccessControlListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAccessControlListResponseWithDefaults() *KalturaAccessControlListResponse {
	this := KalturaAccessControlListResponse{}
	return &this
}

func (o KalturaAccessControlListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAccessControlListResponse struct {
	value *KalturaAccessControlListResponse
	isSet bool
}

func (v NullableKalturaAccessControlListResponse) Get() *KalturaAccessControlListResponse {
	return v.value
}

func (v *NullableKalturaAccessControlListResponse) Set(val *KalturaAccessControlListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAccessControlListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAccessControlListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAccessControlListResponse(val *KalturaAccessControlListResponse) *NullableKalturaAccessControlListResponse {
	return &NullableKalturaAccessControlListResponse{value: val, isSet: true}
}

func (v NullableKalturaAccessControlListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAccessControlListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


