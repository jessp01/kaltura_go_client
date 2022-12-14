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

// KalturaUserRoleListResponse struct for KalturaUserRoleListResponse
type KalturaUserRoleListResponse struct {
	KalturaListResponse
}

// NewKalturaUserRoleListResponse instantiates a new KalturaUserRoleListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUserRoleListResponse() *KalturaUserRoleListResponse {
	this := KalturaUserRoleListResponse{}
	return &this
}

// NewKalturaUserRoleListResponseWithDefaults instantiates a new KalturaUserRoleListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUserRoleListResponseWithDefaults() *KalturaUserRoleListResponse {
	this := KalturaUserRoleListResponse{}
	return &this
}

func (o KalturaUserRoleListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaUserRoleListResponse struct {
	value *KalturaUserRoleListResponse
	isSet bool
}

func (v NullableKalturaUserRoleListResponse) Get() *KalturaUserRoleListResponse {
	return v.value
}

func (v *NullableKalturaUserRoleListResponse) Set(val *KalturaUserRoleListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUserRoleListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUserRoleListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUserRoleListResponse(val *KalturaUserRoleListResponse) *NullableKalturaUserRoleListResponse {
	return &NullableKalturaUserRoleListResponse{value: val, isSet: true}
}

func (v NullableKalturaUserRoleListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUserRoleListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


