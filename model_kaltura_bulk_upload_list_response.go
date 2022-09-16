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

// KalturaBulkUploadListResponse struct for KalturaBulkUploadListResponse
type KalturaBulkUploadListResponse struct {
	KalturaListResponse
}

// NewKalturaBulkUploadListResponse instantiates a new KalturaBulkUploadListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadListResponse() *KalturaBulkUploadListResponse {
	this := KalturaBulkUploadListResponse{}
	return &this
}

// NewKalturaBulkUploadListResponseWithDefaults instantiates a new KalturaBulkUploadListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadListResponseWithDefaults() *KalturaBulkUploadListResponse {
	this := KalturaBulkUploadListResponse{}
	return &this
}

func (o KalturaBulkUploadListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBulkUploadListResponse struct {
	value *KalturaBulkUploadListResponse
	isSet bool
}

func (v NullableKalturaBulkUploadListResponse) Get() *KalturaBulkUploadListResponse {
	return v.value
}

func (v *NullableKalturaBulkUploadListResponse) Set(val *KalturaBulkUploadListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadListResponse(val *KalturaBulkUploadListResponse) *NullableKalturaBulkUploadListResponse {
	return &NullableKalturaBulkUploadListResponse{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


