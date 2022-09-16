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

// KalturaScheduleEventResourceListResponse struct for KalturaScheduleEventResourceListResponse
type KalturaScheduleEventResourceListResponse struct {
	KalturaListResponse
}

// NewKalturaScheduleEventResourceListResponse instantiates a new KalturaScheduleEventResourceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduleEventResourceListResponse() *KalturaScheduleEventResourceListResponse {
	this := KalturaScheduleEventResourceListResponse{}
	return &this
}

// NewKalturaScheduleEventResourceListResponseWithDefaults instantiates a new KalturaScheduleEventResourceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduleEventResourceListResponseWithDefaults() *KalturaScheduleEventResourceListResponse {
	this := KalturaScheduleEventResourceListResponse{}
	return &this
}

func (o KalturaScheduleEventResourceListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaScheduleEventResourceListResponse struct {
	value *KalturaScheduleEventResourceListResponse
	isSet bool
}

func (v NullableKalturaScheduleEventResourceListResponse) Get() *KalturaScheduleEventResourceListResponse {
	return v.value
}

func (v *NullableKalturaScheduleEventResourceListResponse) Set(val *KalturaScheduleEventResourceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduleEventResourceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduleEventResourceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduleEventResourceListResponse(val *KalturaScheduleEventResourceListResponse) *NullableKalturaScheduleEventResourceListResponse {
	return &NullableKalturaScheduleEventResourceListResponse{value: val, isSet: true}
}

func (v NullableKalturaScheduleEventResourceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduleEventResourceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


