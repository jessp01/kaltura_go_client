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

// KalturaApiParameterPermissionItem struct for KalturaApiParameterPermissionItem
type KalturaApiParameterPermissionItem struct {
	KalturaPermissionItem
}

// NewKalturaApiParameterPermissionItem instantiates a new KalturaApiParameterPermissionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaApiParameterPermissionItem() *KalturaApiParameterPermissionItem {
	this := KalturaApiParameterPermissionItem{}
	return &this
}

// NewKalturaApiParameterPermissionItemWithDefaults instantiates a new KalturaApiParameterPermissionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaApiParameterPermissionItemWithDefaults() *KalturaApiParameterPermissionItem {
	this := KalturaApiParameterPermissionItem{}
	return &this
}

func (o KalturaApiParameterPermissionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPermissionItem, errKalturaPermissionItem := json.Marshal(o.KalturaPermissionItem)
	if errKalturaPermissionItem != nil {
		return []byte{}, errKalturaPermissionItem
	}
	errKalturaPermissionItem = json.Unmarshal([]byte(serializedKalturaPermissionItem), &toSerialize)
	if errKalturaPermissionItem != nil {
		return []byte{}, errKalturaPermissionItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaApiParameterPermissionItem struct {
	value *KalturaApiParameterPermissionItem
	isSet bool
}

func (v NullableKalturaApiParameterPermissionItem) Get() *KalturaApiParameterPermissionItem {
	return v.value
}

func (v *NullableKalturaApiParameterPermissionItem) Set(val *KalturaApiParameterPermissionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaApiParameterPermissionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaApiParameterPermissionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaApiParameterPermissionItem(val *KalturaApiParameterPermissionItem) *NullableKalturaApiParameterPermissionItem {
	return &NullableKalturaApiParameterPermissionItem{value: val, isSet: true}
}

func (v NullableKalturaApiParameterPermissionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaApiParameterPermissionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

