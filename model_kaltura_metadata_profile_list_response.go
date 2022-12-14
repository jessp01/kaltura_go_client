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

// KalturaMetadataProfileListResponse struct for KalturaMetadataProfileListResponse
type KalturaMetadataProfileListResponse struct {
	KalturaListResponse
}

// NewKalturaMetadataProfileListResponse instantiates a new KalturaMetadataProfileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMetadataProfileListResponse() *KalturaMetadataProfileListResponse {
	this := KalturaMetadataProfileListResponse{}
	return &this
}

// NewKalturaMetadataProfileListResponseWithDefaults instantiates a new KalturaMetadataProfileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMetadataProfileListResponseWithDefaults() *KalturaMetadataProfileListResponse {
	this := KalturaMetadataProfileListResponse{}
	return &this
}

func (o KalturaMetadataProfileListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaMetadataProfileListResponse struct {
	value *KalturaMetadataProfileListResponse
	isSet bool
}

func (v NullableKalturaMetadataProfileListResponse) Get() *KalturaMetadataProfileListResponse {
	return v.value
}

func (v *NullableKalturaMetadataProfileListResponse) Set(val *KalturaMetadataProfileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMetadataProfileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMetadataProfileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMetadataProfileListResponse(val *KalturaMetadataProfileListResponse) *NullableKalturaMetadataProfileListResponse {
	return &NullableKalturaMetadataProfileListResponse{value: val, isSet: true}
}

func (v NullableKalturaMetadataProfileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMetadataProfileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


