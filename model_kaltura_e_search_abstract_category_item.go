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

// KalturaESearchAbstractCategoryItem `abstract`
type KalturaESearchAbstractCategoryItem struct {
	KalturaESearchCategoryBaseItem
}

// NewKalturaESearchAbstractCategoryItem instantiates a new KalturaESearchAbstractCategoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchAbstractCategoryItem() *KalturaESearchAbstractCategoryItem {
	this := KalturaESearchAbstractCategoryItem{}
	return &this
}

// NewKalturaESearchAbstractCategoryItemWithDefaults instantiates a new KalturaESearchAbstractCategoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchAbstractCategoryItemWithDefaults() *KalturaESearchAbstractCategoryItem {
	this := KalturaESearchAbstractCategoryItem{}
	return &this
}

func (o KalturaESearchAbstractCategoryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchCategoryBaseItem, errKalturaESearchCategoryBaseItem := json.Marshal(o.KalturaESearchCategoryBaseItem)
	if errKalturaESearchCategoryBaseItem != nil {
		return []byte{}, errKalturaESearchCategoryBaseItem
	}
	errKalturaESearchCategoryBaseItem = json.Unmarshal([]byte(serializedKalturaESearchCategoryBaseItem), &toSerialize)
	if errKalturaESearchCategoryBaseItem != nil {
		return []byte{}, errKalturaESearchCategoryBaseItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchAbstractCategoryItem struct {
	value *KalturaESearchAbstractCategoryItem
	isSet bool
}

func (v NullableKalturaESearchAbstractCategoryItem) Get() *KalturaESearchAbstractCategoryItem {
	return v.value
}

func (v *NullableKalturaESearchAbstractCategoryItem) Set(val *KalturaESearchAbstractCategoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchAbstractCategoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchAbstractCategoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchAbstractCategoryItem(val *KalturaESearchAbstractCategoryItem) *NullableKalturaESearchAbstractCategoryItem {
	return &NullableKalturaESearchAbstractCategoryItem{value: val, isSet: true}
}

func (v NullableKalturaESearchAbstractCategoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchAbstractCategoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


