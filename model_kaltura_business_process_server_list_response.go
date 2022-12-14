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

// KalturaBusinessProcessServerListResponse struct for KalturaBusinessProcessServerListResponse
type KalturaBusinessProcessServerListResponse struct {
	KalturaListResponse
}

// NewKalturaBusinessProcessServerListResponse instantiates a new KalturaBusinessProcessServerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBusinessProcessServerListResponse() *KalturaBusinessProcessServerListResponse {
	this := KalturaBusinessProcessServerListResponse{}
	return &this
}

// NewKalturaBusinessProcessServerListResponseWithDefaults instantiates a new KalturaBusinessProcessServerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBusinessProcessServerListResponseWithDefaults() *KalturaBusinessProcessServerListResponse {
	this := KalturaBusinessProcessServerListResponse{}
	return &this
}

func (o KalturaBusinessProcessServerListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBusinessProcessServerListResponse struct {
	value *KalturaBusinessProcessServerListResponse
	isSet bool
}

func (v NullableKalturaBusinessProcessServerListResponse) Get() *KalturaBusinessProcessServerListResponse {
	return v.value
}

func (v *NullableKalturaBusinessProcessServerListResponse) Set(val *KalturaBusinessProcessServerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBusinessProcessServerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBusinessProcessServerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBusinessProcessServerListResponse(val *KalturaBusinessProcessServerListResponse) *NullableKalturaBusinessProcessServerListResponse {
	return &NullableKalturaBusinessProcessServerListResponse{value: val, isSet: true}
}

func (v NullableKalturaBusinessProcessServerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBusinessProcessServerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


