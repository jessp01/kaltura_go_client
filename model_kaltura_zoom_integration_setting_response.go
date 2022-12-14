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

// KalturaZoomIntegrationSettingResponse struct for KalturaZoomIntegrationSettingResponse
type KalturaZoomIntegrationSettingResponse struct {
	KalturaListResponse
}

// NewKalturaZoomIntegrationSettingResponse instantiates a new KalturaZoomIntegrationSettingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaZoomIntegrationSettingResponse() *KalturaZoomIntegrationSettingResponse {
	this := KalturaZoomIntegrationSettingResponse{}
	return &this
}

// NewKalturaZoomIntegrationSettingResponseWithDefaults instantiates a new KalturaZoomIntegrationSettingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaZoomIntegrationSettingResponseWithDefaults() *KalturaZoomIntegrationSettingResponse {
	this := KalturaZoomIntegrationSettingResponse{}
	return &this
}

func (o KalturaZoomIntegrationSettingResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaZoomIntegrationSettingResponse struct {
	value *KalturaZoomIntegrationSettingResponse
	isSet bool
}

func (v NullableKalturaZoomIntegrationSettingResponse) Get() *KalturaZoomIntegrationSettingResponse {
	return v.value
}

func (v *NullableKalturaZoomIntegrationSettingResponse) Set(val *KalturaZoomIntegrationSettingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaZoomIntegrationSettingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaZoomIntegrationSettingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaZoomIntegrationSettingResponse(val *KalturaZoomIntegrationSettingResponse) *NullableKalturaZoomIntegrationSettingResponse {
	return &NullableKalturaZoomIntegrationSettingResponse{value: val, isSet: true}
}

func (v NullableKalturaZoomIntegrationSettingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaZoomIntegrationSettingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


