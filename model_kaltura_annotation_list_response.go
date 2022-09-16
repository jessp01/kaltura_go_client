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

// KalturaAnnotationListResponse struct for KalturaAnnotationListResponse
type KalturaAnnotationListResponse struct {
	KalturaListResponse
}

// NewKalturaAnnotationListResponse instantiates a new KalturaAnnotationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAnnotationListResponse() *KalturaAnnotationListResponse {
	this := KalturaAnnotationListResponse{}
	return &this
}

// NewKalturaAnnotationListResponseWithDefaults instantiates a new KalturaAnnotationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAnnotationListResponseWithDefaults() *KalturaAnnotationListResponse {
	this := KalturaAnnotationListResponse{}
	return &this
}

func (o KalturaAnnotationListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAnnotationListResponse struct {
	value *KalturaAnnotationListResponse
	isSet bool
}

func (v NullableKalturaAnnotationListResponse) Get() *KalturaAnnotationListResponse {
	return v.value
}

func (v *NullableKalturaAnnotationListResponse) Set(val *KalturaAnnotationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAnnotationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAnnotationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAnnotationListResponse(val *KalturaAnnotationListResponse) *NullableKalturaAnnotationListResponse {
	return &NullableKalturaAnnotationListResponse{value: val, isSet: true}
}

func (v NullableKalturaAnnotationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAnnotationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


