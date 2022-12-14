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

// KalturaDeliveryProfileListResponse struct for KalturaDeliveryProfileListResponse
type KalturaDeliveryProfileListResponse struct {
	KalturaListResponse
}

// NewKalturaDeliveryProfileListResponse instantiates a new KalturaDeliveryProfileListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileListResponse() *KalturaDeliveryProfileListResponse {
	this := KalturaDeliveryProfileListResponse{}
	return &this
}

// NewKalturaDeliveryProfileListResponseWithDefaults instantiates a new KalturaDeliveryProfileListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileListResponseWithDefaults() *KalturaDeliveryProfileListResponse {
	this := KalturaDeliveryProfileListResponse{}
	return &this
}

func (o KalturaDeliveryProfileListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaDeliveryProfileListResponse struct {
	value *KalturaDeliveryProfileListResponse
	isSet bool
}

func (v NullableKalturaDeliveryProfileListResponse) Get() *KalturaDeliveryProfileListResponse {
	return v.value
}

func (v *NullableKalturaDeliveryProfileListResponse) Set(val *KalturaDeliveryProfileListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileListResponse(val *KalturaDeliveryProfileListResponse) *NullableKalturaDeliveryProfileListResponse {
	return &NullableKalturaDeliveryProfileListResponse{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


