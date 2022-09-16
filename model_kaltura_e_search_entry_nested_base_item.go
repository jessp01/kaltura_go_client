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

// KalturaESearchEntryNestedBaseItem `abstract`
type KalturaESearchEntryNestedBaseItem struct {
	KalturaESearchEntryBaseNestedObject
}

// NewKalturaESearchEntryNestedBaseItem instantiates a new KalturaESearchEntryNestedBaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchEntryNestedBaseItem() *KalturaESearchEntryNestedBaseItem {
	this := KalturaESearchEntryNestedBaseItem{}
	return &this
}

// NewKalturaESearchEntryNestedBaseItemWithDefaults instantiates a new KalturaESearchEntryNestedBaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchEntryNestedBaseItemWithDefaults() *KalturaESearchEntryNestedBaseItem {
	this := KalturaESearchEntryNestedBaseItem{}
	return &this
}

func (o KalturaESearchEntryNestedBaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchEntryBaseNestedObject, errKalturaESearchEntryBaseNestedObject := json.Marshal(o.KalturaESearchEntryBaseNestedObject)
	if errKalturaESearchEntryBaseNestedObject != nil {
		return []byte{}, errKalturaESearchEntryBaseNestedObject
	}
	errKalturaESearchEntryBaseNestedObject = json.Unmarshal([]byte(serializedKalturaESearchEntryBaseNestedObject), &toSerialize)
	if errKalturaESearchEntryBaseNestedObject != nil {
		return []byte{}, errKalturaESearchEntryBaseNestedObject
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchEntryNestedBaseItem struct {
	value *KalturaESearchEntryNestedBaseItem
	isSet bool
}

func (v NullableKalturaESearchEntryNestedBaseItem) Get() *KalturaESearchEntryNestedBaseItem {
	return v.value
}

func (v *NullableKalturaESearchEntryNestedBaseItem) Set(val *KalturaESearchEntryNestedBaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchEntryNestedBaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchEntryNestedBaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchEntryNestedBaseItem(val *KalturaESearchEntryNestedBaseItem) *NullableKalturaESearchEntryNestedBaseItem {
	return &NullableKalturaESearchEntryNestedBaseItem{value: val, isSet: true}
}

func (v NullableKalturaESearchEntryNestedBaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchEntryNestedBaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


