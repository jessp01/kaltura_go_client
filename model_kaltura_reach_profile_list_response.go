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

// KalturaReachProfileListResponse struct for KalturaReachProfileListResponse
type KalturaReachProfileListResponse struct {
	KalturaListResponse
}

// NewKalturaReachProfileListResponse instantiates a new KalturaReachProfileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaReachProfileListResponse() *KalturaReachProfileListResponse {
	this := KalturaReachProfileListResponse{}
	return &this
}

// NewKalturaReachProfileListResponseWithDefaults instantiates a new KalturaReachProfileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaReachProfileListResponseWithDefaults() *KalturaReachProfileListResponse {
	this := KalturaReachProfileListResponse{}
	return &this
}

func (o KalturaReachProfileListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaReachProfileListResponse struct {
	value *KalturaReachProfileListResponse
	isSet bool
}

func (v NullableKalturaReachProfileListResponse) Get() *KalturaReachProfileListResponse {
	return v.value
}

func (v *NullableKalturaReachProfileListResponse) Set(val *KalturaReachProfileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaReachProfileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaReachProfileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaReachProfileListResponse(val *KalturaReachProfileListResponse) *NullableKalturaReachProfileListResponse {
	return &NullableKalturaReachProfileListResponse{value: val, isSet: true}
}

func (v NullableKalturaReachProfileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaReachProfileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

