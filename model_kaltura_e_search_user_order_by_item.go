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

// KalturaESearchUserOrderByItem struct for KalturaESearchUserOrderByItem
type KalturaESearchUserOrderByItem struct {
	KalturaESearchOrderByItem
}

// NewKalturaESearchUserOrderByItem instantiates a new KalturaESearchUserOrderByItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchUserOrderByItem() *KalturaESearchUserOrderByItem {
	this := KalturaESearchUserOrderByItem{}
	return &this
}

// NewKalturaESearchUserOrderByItemWithDefaults instantiates a new KalturaESearchUserOrderByItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchUserOrderByItemWithDefaults() *KalturaESearchUserOrderByItem {
	this := KalturaESearchUserOrderByItem{}
	return &this
}

func (o KalturaESearchUserOrderByItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchOrderByItem, errKalturaESearchOrderByItem := json.Marshal(o.KalturaESearchOrderByItem)
	if errKalturaESearchOrderByItem != nil {
		return []byte{}, errKalturaESearchOrderByItem
	}
	errKalturaESearchOrderByItem = json.Unmarshal([]byte(serializedKalturaESearchOrderByItem), &toSerialize)
	if errKalturaESearchOrderByItem != nil {
		return []byte{}, errKalturaESearchOrderByItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchUserOrderByItem struct {
	value *KalturaESearchUserOrderByItem
	isSet bool
}

func (v NullableKalturaESearchUserOrderByItem) Get() *KalturaESearchUserOrderByItem {
	return v.value
}

func (v *NullableKalturaESearchUserOrderByItem) Set(val *KalturaESearchUserOrderByItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchUserOrderByItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchUserOrderByItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchUserOrderByItem(val *KalturaESearchUserOrderByItem) *NullableKalturaESearchUserOrderByItem {
	return &NullableKalturaESearchUserOrderByItem{value: val, isSet: true}
}

func (v NullableKalturaESearchUserOrderByItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchUserOrderByItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


