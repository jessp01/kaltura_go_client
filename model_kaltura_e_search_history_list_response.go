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

// KalturaESearchHistoryListResponse struct for KalturaESearchHistoryListResponse
type KalturaESearchHistoryListResponse struct {
	KalturaListResponse
}

// NewKalturaESearchHistoryListResponse instantiates a new KalturaESearchHistoryListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchHistoryListResponse() *KalturaESearchHistoryListResponse {
	this := KalturaESearchHistoryListResponse{}
	return &this
}

// NewKalturaESearchHistoryListResponseWithDefaults instantiates a new KalturaESearchHistoryListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchHistoryListResponseWithDefaults() *KalturaESearchHistoryListResponse {
	this := KalturaESearchHistoryListResponse{}
	return &this
}

func (o KalturaESearchHistoryListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaESearchHistoryListResponse struct {
	value *KalturaESearchHistoryListResponse
	isSet bool
}

func (v NullableKalturaESearchHistoryListResponse) Get() *KalturaESearchHistoryListResponse {
	return v.value
}

func (v *NullableKalturaESearchHistoryListResponse) Set(val *KalturaESearchHistoryListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchHistoryListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchHistoryListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchHistoryListResponse(val *KalturaESearchHistoryListResponse) *NullableKalturaESearchHistoryListResponse {
	return &NullableKalturaESearchHistoryListResponse{value: val, isSet: true}
}

func (v NullableKalturaESearchHistoryListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchHistoryListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


