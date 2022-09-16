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

// KalturaEntryDistributionListResponse struct for KalturaEntryDistributionListResponse
type KalturaEntryDistributionListResponse struct {
	KalturaListResponse
}

// NewKalturaEntryDistributionListResponse instantiates a new KalturaEntryDistributionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryDistributionListResponse() *KalturaEntryDistributionListResponse {
	this := KalturaEntryDistributionListResponse{}
	return &this
}

// NewKalturaEntryDistributionListResponseWithDefaults instantiates a new KalturaEntryDistributionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryDistributionListResponseWithDefaults() *KalturaEntryDistributionListResponse {
	this := KalturaEntryDistributionListResponse{}
	return &this
}

func (o KalturaEntryDistributionListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaEntryDistributionListResponse struct {
	value *KalturaEntryDistributionListResponse
	isSet bool
}

func (v NullableKalturaEntryDistributionListResponse) Get() *KalturaEntryDistributionListResponse {
	return v.value
}

func (v *NullableKalturaEntryDistributionListResponse) Set(val *KalturaEntryDistributionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryDistributionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryDistributionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryDistributionListResponse(val *KalturaEntryDistributionListResponse) *NullableKalturaEntryDistributionListResponse {
	return &NullableKalturaEntryDistributionListResponse{value: val, isSet: true}
}

func (v NullableKalturaEntryDistributionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryDistributionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

