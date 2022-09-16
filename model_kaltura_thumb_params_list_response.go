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

// KalturaThumbParamsListResponse struct for KalturaThumbParamsListResponse
type KalturaThumbParamsListResponse struct {
	KalturaListResponse
}

// NewKalturaThumbParamsListResponse instantiates a new KalturaThumbParamsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaThumbParamsListResponse() *KalturaThumbParamsListResponse {
	this := KalturaThumbParamsListResponse{}
	return &this
}

// NewKalturaThumbParamsListResponseWithDefaults instantiates a new KalturaThumbParamsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaThumbParamsListResponseWithDefaults() *KalturaThumbParamsListResponse {
	this := KalturaThumbParamsListResponse{}
	return &this
}

func (o KalturaThumbParamsListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaThumbParamsListResponse struct {
	value *KalturaThumbParamsListResponse
	isSet bool
}

func (v NullableKalturaThumbParamsListResponse) Get() *KalturaThumbParamsListResponse {
	return v.value
}

func (v *NullableKalturaThumbParamsListResponse) Set(val *KalturaThumbParamsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaThumbParamsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaThumbParamsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaThumbParamsListResponse(val *KalturaThumbParamsListResponse) *NullableKalturaThumbParamsListResponse {
	return &NullableKalturaThumbParamsListResponse{value: val, isSet: true}
}

func (v NullableKalturaThumbParamsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaThumbParamsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


