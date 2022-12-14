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

// KalturaStorageProfileListResponse struct for KalturaStorageProfileListResponse
type KalturaStorageProfileListResponse struct {
	KalturaListResponse
}

// NewKalturaStorageProfileListResponse instantiates a new KalturaStorageProfileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaStorageProfileListResponse() *KalturaStorageProfileListResponse {
	this := KalturaStorageProfileListResponse{}
	return &this
}

// NewKalturaStorageProfileListResponseWithDefaults instantiates a new KalturaStorageProfileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaStorageProfileListResponseWithDefaults() *KalturaStorageProfileListResponse {
	this := KalturaStorageProfileListResponse{}
	return &this
}

func (o KalturaStorageProfileListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaStorageProfileListResponse struct {
	value *KalturaStorageProfileListResponse
	isSet bool
}

func (v NullableKalturaStorageProfileListResponse) Get() *KalturaStorageProfileListResponse {
	return v.value
}

func (v *NullableKalturaStorageProfileListResponse) Set(val *KalturaStorageProfileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaStorageProfileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaStorageProfileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaStorageProfileListResponse(val *KalturaStorageProfileListResponse) *NullableKalturaStorageProfileListResponse {
	return &NullableKalturaStorageProfileListResponse{value: val, isSet: true}
}

func (v NullableKalturaStorageProfileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaStorageProfileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


