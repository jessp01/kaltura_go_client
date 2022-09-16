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

// KalturaPermissionItemListResponse struct for KalturaPermissionItemListResponse
type KalturaPermissionItemListResponse struct {
	KalturaListResponse
}

// NewKalturaPermissionItemListResponse instantiates a new KalturaPermissionItemListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPermissionItemListResponse() *KalturaPermissionItemListResponse {
	this := KalturaPermissionItemListResponse{}
	return &this
}

// NewKalturaPermissionItemListResponseWithDefaults instantiates a new KalturaPermissionItemListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPermissionItemListResponseWithDefaults() *KalturaPermissionItemListResponse {
	this := KalturaPermissionItemListResponse{}
	return &this
}

func (o KalturaPermissionItemListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaPermissionItemListResponse struct {
	value *KalturaPermissionItemListResponse
	isSet bool
}

func (v NullableKalturaPermissionItemListResponse) Get() *KalturaPermissionItemListResponse {
	return v.value
}

func (v *NullableKalturaPermissionItemListResponse) Set(val *KalturaPermissionItemListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPermissionItemListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPermissionItemListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPermissionItemListResponse(val *KalturaPermissionItemListResponse) *NullableKalturaPermissionItemListResponse {
	return &NullableKalturaPermissionItemListResponse{value: val, isSet: true}
}

func (v NullableKalturaPermissionItemListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPermissionItemListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


