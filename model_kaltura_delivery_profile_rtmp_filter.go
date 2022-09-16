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

// KalturaDeliveryProfileRtmpFilter struct for KalturaDeliveryProfileRtmpFilter
type KalturaDeliveryProfileRtmpFilter struct {
	KalturaDeliveryProfileRtmpBaseFilter
}

// NewKalturaDeliveryProfileRtmpFilter instantiates a new KalturaDeliveryProfileRtmpFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileRtmpFilter() *KalturaDeliveryProfileRtmpFilter {
	this := KalturaDeliveryProfileRtmpFilter{}
	return &this
}

// NewKalturaDeliveryProfileRtmpFilterWithDefaults instantiates a new KalturaDeliveryProfileRtmpFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileRtmpFilterWithDefaults() *KalturaDeliveryProfileRtmpFilter {
	this := KalturaDeliveryProfileRtmpFilter{}
	return &this
}

func (o KalturaDeliveryProfileRtmpFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDeliveryProfileRtmpBaseFilter, errKalturaDeliveryProfileRtmpBaseFilter := json.Marshal(o.KalturaDeliveryProfileRtmpBaseFilter)
	if errKalturaDeliveryProfileRtmpBaseFilter != nil {
		return []byte{}, errKalturaDeliveryProfileRtmpBaseFilter
	}
	errKalturaDeliveryProfileRtmpBaseFilter = json.Unmarshal([]byte(serializedKalturaDeliveryProfileRtmpBaseFilter), &toSerialize)
	if errKalturaDeliveryProfileRtmpBaseFilter != nil {
		return []byte{}, errKalturaDeliveryProfileRtmpBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfileRtmpFilter struct {
	value *KalturaDeliveryProfileRtmpFilter
	isSet bool
}

func (v NullableKalturaDeliveryProfileRtmpFilter) Get() *KalturaDeliveryProfileRtmpFilter {
	return v.value
}

func (v *NullableKalturaDeliveryProfileRtmpFilter) Set(val *KalturaDeliveryProfileRtmpFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileRtmpFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileRtmpFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileRtmpFilter(val *KalturaDeliveryProfileRtmpFilter) *NullableKalturaDeliveryProfileRtmpFilter {
	return &NullableKalturaDeliveryProfileRtmpFilter{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileRtmpFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileRtmpFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

