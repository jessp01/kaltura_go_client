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

// KalturaContentDistributionSearchItem struct for KalturaContentDistributionSearchItem
type KalturaContentDistributionSearchItem struct {
	KalturaSearchItem
}

// NewKalturaContentDistributionSearchItem instantiates a new KalturaContentDistributionSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaContentDistributionSearchItem() *KalturaContentDistributionSearchItem {
	this := KalturaContentDistributionSearchItem{}
	return &this
}

// NewKalturaContentDistributionSearchItemWithDefaults instantiates a new KalturaContentDistributionSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaContentDistributionSearchItemWithDefaults() *KalturaContentDistributionSearchItem {
	this := KalturaContentDistributionSearchItem{}
	return &this
}

func (o KalturaContentDistributionSearchItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearchItem, errKalturaSearchItem := json.Marshal(o.KalturaSearchItem)
	if errKalturaSearchItem != nil {
		return []byte{}, errKalturaSearchItem
	}
	errKalturaSearchItem = json.Unmarshal([]byte(serializedKalturaSearchItem), &toSerialize)
	if errKalturaSearchItem != nil {
		return []byte{}, errKalturaSearchItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaContentDistributionSearchItem struct {
	value *KalturaContentDistributionSearchItem
	isSet bool
}

func (v NullableKalturaContentDistributionSearchItem) Get() *KalturaContentDistributionSearchItem {
	return v.value
}

func (v *NullableKalturaContentDistributionSearchItem) Set(val *KalturaContentDistributionSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaContentDistributionSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaContentDistributionSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaContentDistributionSearchItem(val *KalturaContentDistributionSearchItem) *NullableKalturaContentDistributionSearchItem {
	return &NullableKalturaContentDistributionSearchItem{value: val, isSet: true}
}

func (v NullableKalturaContentDistributionSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaContentDistributionSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


