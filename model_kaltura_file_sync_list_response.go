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

// KalturaFileSyncListResponse struct for KalturaFileSyncListResponse
type KalturaFileSyncListResponse struct {
	KalturaListResponse
}

// NewKalturaFileSyncListResponse instantiates a new KalturaFileSyncListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFileSyncListResponse() *KalturaFileSyncListResponse {
	this := KalturaFileSyncListResponse{}
	return &this
}

// NewKalturaFileSyncListResponseWithDefaults instantiates a new KalturaFileSyncListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFileSyncListResponseWithDefaults() *KalturaFileSyncListResponse {
	this := KalturaFileSyncListResponse{}
	return &this
}

func (o KalturaFileSyncListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaFileSyncListResponse struct {
	value *KalturaFileSyncListResponse
	isSet bool
}

func (v NullableKalturaFileSyncListResponse) Get() *KalturaFileSyncListResponse {
	return v.value
}

func (v *NullableKalturaFileSyncListResponse) Set(val *KalturaFileSyncListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFileSyncListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFileSyncListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFileSyncListResponse(val *KalturaFileSyncListResponse) *NullableKalturaFileSyncListResponse {
	return &NullableKalturaFileSyncListResponse{value: val, isSet: true}
}

func (v NullableKalturaFileSyncListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFileSyncListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


