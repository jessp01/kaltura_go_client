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

// KalturaModerationFlagListResponse struct for KalturaModerationFlagListResponse
type KalturaModerationFlagListResponse struct {
	KalturaListResponse
}

// NewKalturaModerationFlagListResponse instantiates a new KalturaModerationFlagListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaModerationFlagListResponse() *KalturaModerationFlagListResponse {
	this := KalturaModerationFlagListResponse{}
	return &this
}

// NewKalturaModerationFlagListResponseWithDefaults instantiates a new KalturaModerationFlagListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaModerationFlagListResponseWithDefaults() *KalturaModerationFlagListResponse {
	this := KalturaModerationFlagListResponse{}
	return &this
}

func (o KalturaModerationFlagListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaModerationFlagListResponse struct {
	value *KalturaModerationFlagListResponse
	isSet bool
}

func (v NullableKalturaModerationFlagListResponse) Get() *KalturaModerationFlagListResponse {
	return v.value
}

func (v *NullableKalturaModerationFlagListResponse) Set(val *KalturaModerationFlagListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaModerationFlagListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaModerationFlagListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaModerationFlagListResponse(val *KalturaModerationFlagListResponse) *NullableKalturaModerationFlagListResponse {
	return &NullableKalturaModerationFlagListResponse{value: val, isSet: true}
}

func (v NullableKalturaModerationFlagListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaModerationFlagListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

