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

// KalturaDeliveryProfileFilter struct for KalturaDeliveryProfileFilter
type KalturaDeliveryProfileFilter struct {
	KalturaDeliveryProfileBaseFilter
}

// NewKalturaDeliveryProfileFilter instantiates a new KalturaDeliveryProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileFilter() *KalturaDeliveryProfileFilter {
	this := KalturaDeliveryProfileFilter{}
	return &this
}

// NewKalturaDeliveryProfileFilterWithDefaults instantiates a new KalturaDeliveryProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileFilterWithDefaults() *KalturaDeliveryProfileFilter {
	this := KalturaDeliveryProfileFilter{}
	return &this
}

func (o KalturaDeliveryProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDeliveryProfileBaseFilter, errKalturaDeliveryProfileBaseFilter := json.Marshal(o.KalturaDeliveryProfileBaseFilter)
	if errKalturaDeliveryProfileBaseFilter != nil {
		return []byte{}, errKalturaDeliveryProfileBaseFilter
	}
	errKalturaDeliveryProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaDeliveryProfileBaseFilter), &toSerialize)
	if errKalturaDeliveryProfileBaseFilter != nil {
		return []byte{}, errKalturaDeliveryProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfileFilter struct {
	value *KalturaDeliveryProfileFilter
	isSet bool
}

func (v NullableKalturaDeliveryProfileFilter) Get() *KalturaDeliveryProfileFilter {
	return v.value
}

func (v *NullableKalturaDeliveryProfileFilter) Set(val *KalturaDeliveryProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileFilter(val *KalturaDeliveryProfileFilter) *NullableKalturaDeliveryProfileFilter {
	return &NullableKalturaDeliveryProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

