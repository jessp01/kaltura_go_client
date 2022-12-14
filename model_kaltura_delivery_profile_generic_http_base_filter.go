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

// KalturaDeliveryProfileGenericHttpBaseFilter `abstract`
type KalturaDeliveryProfileGenericHttpBaseFilter struct {
	KalturaDeliveryProfileFilter
}

// NewKalturaDeliveryProfileGenericHttpBaseFilter instantiates a new KalturaDeliveryProfileGenericHttpBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileGenericHttpBaseFilter() *KalturaDeliveryProfileGenericHttpBaseFilter {
	this := KalturaDeliveryProfileGenericHttpBaseFilter{}
	return &this
}

// NewKalturaDeliveryProfileGenericHttpBaseFilterWithDefaults instantiates a new KalturaDeliveryProfileGenericHttpBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileGenericHttpBaseFilterWithDefaults() *KalturaDeliveryProfileGenericHttpBaseFilter {
	this := KalturaDeliveryProfileGenericHttpBaseFilter{}
	return &this
}

func (o KalturaDeliveryProfileGenericHttpBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDeliveryProfileFilter, errKalturaDeliveryProfileFilter := json.Marshal(o.KalturaDeliveryProfileFilter)
	if errKalturaDeliveryProfileFilter != nil {
		return []byte{}, errKalturaDeliveryProfileFilter
	}
	errKalturaDeliveryProfileFilter = json.Unmarshal([]byte(serializedKalturaDeliveryProfileFilter), &toSerialize)
	if errKalturaDeliveryProfileFilter != nil {
		return []byte{}, errKalturaDeliveryProfileFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfileGenericHttpBaseFilter struct {
	value *KalturaDeliveryProfileGenericHttpBaseFilter
	isSet bool
}

func (v NullableKalturaDeliveryProfileGenericHttpBaseFilter) Get() *KalturaDeliveryProfileGenericHttpBaseFilter {
	return v.value
}

func (v *NullableKalturaDeliveryProfileGenericHttpBaseFilter) Set(val *KalturaDeliveryProfileGenericHttpBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileGenericHttpBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileGenericHttpBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileGenericHttpBaseFilter(val *KalturaDeliveryProfileGenericHttpBaseFilter) *NullableKalturaDeliveryProfileGenericHttpBaseFilter {
	return &NullableKalturaDeliveryProfileGenericHttpBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileGenericHttpBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileGenericHttpBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


