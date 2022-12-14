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

// KalturaDropFolderFileListResponse struct for KalturaDropFolderFileListResponse
type KalturaDropFolderFileListResponse struct {
	KalturaListResponse
}

// NewKalturaDropFolderFileListResponse instantiates a new KalturaDropFolderFileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDropFolderFileListResponse() *KalturaDropFolderFileListResponse {
	this := KalturaDropFolderFileListResponse{}
	return &this
}

// NewKalturaDropFolderFileListResponseWithDefaults instantiates a new KalturaDropFolderFileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDropFolderFileListResponseWithDefaults() *KalturaDropFolderFileListResponse {
	this := KalturaDropFolderFileListResponse{}
	return &this
}

func (o KalturaDropFolderFileListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaDropFolderFileListResponse struct {
	value *KalturaDropFolderFileListResponse
	isSet bool
}

func (v NullableKalturaDropFolderFileListResponse) Get() *KalturaDropFolderFileListResponse {
	return v.value
}

func (v *NullableKalturaDropFolderFileListResponse) Set(val *KalturaDropFolderFileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDropFolderFileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDropFolderFileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDropFolderFileListResponse(val *KalturaDropFolderFileListResponse) *NullableKalturaDropFolderFileListResponse {
	return &NullableKalturaDropFolderFileListResponse{value: val, isSet: true}
}

func (v NullableKalturaDropFolderFileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDropFolderFileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


