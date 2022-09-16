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

// KalturaFeatureStatusListResponse struct for KalturaFeatureStatusListResponse
type KalturaFeatureStatusListResponse struct {
	KalturaListResponse
}

// NewKalturaFeatureStatusListResponse instantiates a new KalturaFeatureStatusListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFeatureStatusListResponse() *KalturaFeatureStatusListResponse {
	this := KalturaFeatureStatusListResponse{}
	return &this
}

// NewKalturaFeatureStatusListResponseWithDefaults instantiates a new KalturaFeatureStatusListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFeatureStatusListResponseWithDefaults() *KalturaFeatureStatusListResponse {
	this := KalturaFeatureStatusListResponse{}
	return &this
}

func (o KalturaFeatureStatusListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaFeatureStatusListResponse struct {
	value *KalturaFeatureStatusListResponse
	isSet bool
}

func (v NullableKalturaFeatureStatusListResponse) Get() *KalturaFeatureStatusListResponse {
	return v.value
}

func (v *NullableKalturaFeatureStatusListResponse) Set(val *KalturaFeatureStatusListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFeatureStatusListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFeatureStatusListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFeatureStatusListResponse(val *KalturaFeatureStatusListResponse) *NullableKalturaFeatureStatusListResponse {
	return &NullableKalturaFeatureStatusListResponse{value: val, isSet: true}
}

func (v NullableKalturaFeatureStatusListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFeatureStatusListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


