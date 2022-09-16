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

// KalturaESearchGroupUserItem struct for KalturaESearchGroupUserItem
type KalturaESearchGroupUserItem struct {
	KalturaESearchAbstractUserItem
}

// NewKalturaESearchGroupUserItem instantiates a new KalturaESearchGroupUserItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchGroupUserItem() *KalturaESearchGroupUserItem {
	this := KalturaESearchGroupUserItem{}
	return &this
}

// NewKalturaESearchGroupUserItemWithDefaults instantiates a new KalturaESearchGroupUserItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchGroupUserItemWithDefaults() *KalturaESearchGroupUserItem {
	this := KalturaESearchGroupUserItem{}
	return &this
}

func (o KalturaESearchGroupUserItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchAbstractUserItem, errKalturaESearchAbstractUserItem := json.Marshal(o.KalturaESearchAbstractUserItem)
	if errKalturaESearchAbstractUserItem != nil {
		return []byte{}, errKalturaESearchAbstractUserItem
	}
	errKalturaESearchAbstractUserItem = json.Unmarshal([]byte(serializedKalturaESearchAbstractUserItem), &toSerialize)
	if errKalturaESearchAbstractUserItem != nil {
		return []byte{}, errKalturaESearchAbstractUserItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchGroupUserItem struct {
	value *KalturaESearchGroupUserItem
	isSet bool
}

func (v NullableKalturaESearchGroupUserItem) Get() *KalturaESearchGroupUserItem {
	return v.value
}

func (v *NullableKalturaESearchGroupUserItem) Set(val *KalturaESearchGroupUserItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchGroupUserItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchGroupUserItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchGroupUserItem(val *KalturaESearchGroupUserItem) *NullableKalturaESearchGroupUserItem {
	return &NullableKalturaESearchGroupUserItem{value: val, isSet: true}
}

func (v NullableKalturaESearchGroupUserItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchGroupUserItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

