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

// KalturaSsoListResponse struct for KalturaSsoListResponse
type KalturaSsoListResponse struct {
	KalturaListResponse
}

// NewKalturaSsoListResponse instantiates a new KalturaSsoListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSsoListResponse() *KalturaSsoListResponse {
	this := KalturaSsoListResponse{}
	return &this
}

// NewKalturaSsoListResponseWithDefaults instantiates a new KalturaSsoListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSsoListResponseWithDefaults() *KalturaSsoListResponse {
	this := KalturaSsoListResponse{}
	return &this
}

func (o KalturaSsoListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaSsoListResponse struct {
	value *KalturaSsoListResponse
	isSet bool
}

func (v NullableKalturaSsoListResponse) Get() *KalturaSsoListResponse {
	return v.value
}

func (v *NullableKalturaSsoListResponse) Set(val *KalturaSsoListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSsoListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSsoListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSsoListResponse(val *KalturaSsoListResponse) *NullableKalturaSsoListResponse {
	return &NullableKalturaSsoListResponse{value: val, isSet: true}
}

func (v NullableKalturaSsoListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSsoListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


