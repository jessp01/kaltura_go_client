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

// KalturaGroupUserListResponse struct for KalturaGroupUserListResponse
type KalturaGroupUserListResponse struct {
	KalturaListResponse
}

// NewKalturaGroupUserListResponse instantiates a new KalturaGroupUserListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGroupUserListResponse() *KalturaGroupUserListResponse {
	this := KalturaGroupUserListResponse{}
	return &this
}

// NewKalturaGroupUserListResponseWithDefaults instantiates a new KalturaGroupUserListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGroupUserListResponseWithDefaults() *KalturaGroupUserListResponse {
	this := KalturaGroupUserListResponse{}
	return &this
}

func (o KalturaGroupUserListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaGroupUserListResponse struct {
	value *KalturaGroupUserListResponse
	isSet bool
}

func (v NullableKalturaGroupUserListResponse) Get() *KalturaGroupUserListResponse {
	return v.value
}

func (v *NullableKalturaGroupUserListResponse) Set(val *KalturaGroupUserListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGroupUserListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGroupUserListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGroupUserListResponse(val *KalturaGroupUserListResponse) *NullableKalturaGroupUserListResponse {
	return &NullableKalturaGroupUserListResponse{value: val, isSet: true}
}

func (v NullableKalturaGroupUserListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGroupUserListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

