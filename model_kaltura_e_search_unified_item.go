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

// KalturaESearchUnifiedItem struct for KalturaESearchUnifiedItem
type KalturaESearchUnifiedItem struct {
	KalturaESearchAbstractEntryItem
}

// NewKalturaESearchUnifiedItem instantiates a new KalturaESearchUnifiedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchUnifiedItem() *KalturaESearchUnifiedItem {
	this := KalturaESearchUnifiedItem{}
	return &this
}

// NewKalturaESearchUnifiedItemWithDefaults instantiates a new KalturaESearchUnifiedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchUnifiedItemWithDefaults() *KalturaESearchUnifiedItem {
	this := KalturaESearchUnifiedItem{}
	return &this
}

func (o KalturaESearchUnifiedItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchAbstractEntryItem, errKalturaESearchAbstractEntryItem := json.Marshal(o.KalturaESearchAbstractEntryItem)
	if errKalturaESearchAbstractEntryItem != nil {
		return []byte{}, errKalturaESearchAbstractEntryItem
	}
	errKalturaESearchAbstractEntryItem = json.Unmarshal([]byte(serializedKalturaESearchAbstractEntryItem), &toSerialize)
	if errKalturaESearchAbstractEntryItem != nil {
		return []byte{}, errKalturaESearchAbstractEntryItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchUnifiedItem struct {
	value *KalturaESearchUnifiedItem
	isSet bool
}

func (v NullableKalturaESearchUnifiedItem) Get() *KalturaESearchUnifiedItem {
	return v.value
}

func (v *NullableKalturaESearchUnifiedItem) Set(val *KalturaESearchUnifiedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchUnifiedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchUnifiedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchUnifiedItem(val *KalturaESearchUnifiedItem) *NullableKalturaESearchUnifiedItem {
	return &NullableKalturaESearchUnifiedItem{value: val, isSet: true}
}

func (v NullableKalturaESearchUnifiedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchUnifiedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


