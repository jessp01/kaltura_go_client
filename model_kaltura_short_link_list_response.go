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

// KalturaShortLinkListResponse struct for KalturaShortLinkListResponse
type KalturaShortLinkListResponse struct {
	KalturaListResponse
}

// NewKalturaShortLinkListResponse instantiates a new KalturaShortLinkListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaShortLinkListResponse() *KalturaShortLinkListResponse {
	this := KalturaShortLinkListResponse{}
	return &this
}

// NewKalturaShortLinkListResponseWithDefaults instantiates a new KalturaShortLinkListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaShortLinkListResponseWithDefaults() *KalturaShortLinkListResponse {
	this := KalturaShortLinkListResponse{}
	return &this
}

func (o KalturaShortLinkListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaShortLinkListResponse struct {
	value *KalturaShortLinkListResponse
	isSet bool
}

func (v NullableKalturaShortLinkListResponse) Get() *KalturaShortLinkListResponse {
	return v.value
}

func (v *NullableKalturaShortLinkListResponse) Set(val *KalturaShortLinkListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaShortLinkListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaShortLinkListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaShortLinkListResponse(val *KalturaShortLinkListResponse) *NullableKalturaShortLinkListResponse {
	return &NullableKalturaShortLinkListResponse{value: val, isSet: true}
}

func (v NullableKalturaShortLinkListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaShortLinkListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


