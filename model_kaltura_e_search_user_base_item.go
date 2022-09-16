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

// KalturaESearchUserBaseItem `abstract`
type KalturaESearchUserBaseItem struct {
	KalturaESearchBaseItem
}

// NewKalturaESearchUserBaseItem instantiates a new KalturaESearchUserBaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchUserBaseItem() *KalturaESearchUserBaseItem {
	this := KalturaESearchUserBaseItem{}
	return &this
}

// NewKalturaESearchUserBaseItemWithDefaults instantiates a new KalturaESearchUserBaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchUserBaseItemWithDefaults() *KalturaESearchUserBaseItem {
	this := KalturaESearchUserBaseItem{}
	return &this
}

func (o KalturaESearchUserBaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchBaseItem, errKalturaESearchBaseItem := json.Marshal(o.KalturaESearchBaseItem)
	if errKalturaESearchBaseItem != nil {
		return []byte{}, errKalturaESearchBaseItem
	}
	errKalturaESearchBaseItem = json.Unmarshal([]byte(serializedKalturaESearchBaseItem), &toSerialize)
	if errKalturaESearchBaseItem != nil {
		return []byte{}, errKalturaESearchBaseItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchUserBaseItem struct {
	value *KalturaESearchUserBaseItem
	isSet bool
}

func (v NullableKalturaESearchUserBaseItem) Get() *KalturaESearchUserBaseItem {
	return v.value
}

func (v *NullableKalturaESearchUserBaseItem) Set(val *KalturaESearchUserBaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchUserBaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchUserBaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchUserBaseItem(val *KalturaESearchUserBaseItem) *NullableKalturaESearchUserBaseItem {
	return &NullableKalturaESearchUserBaseItem{value: val, isSet: true}
}

func (v NullableKalturaESearchUserBaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchUserBaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

