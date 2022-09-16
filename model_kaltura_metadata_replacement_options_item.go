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

// KalturaMetadataReplacementOptionsItem Advanced metadata configuration for entry replacement process
type KalturaMetadataReplacementOptionsItem struct {
	KalturaPluginReplacementOptionsItem
}

// NewKalturaMetadataReplacementOptionsItem instantiates a new KalturaMetadataReplacementOptionsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMetadataReplacementOptionsItem() *KalturaMetadataReplacementOptionsItem {
	this := KalturaMetadataReplacementOptionsItem{}
	return &this
}

// NewKalturaMetadataReplacementOptionsItemWithDefaults instantiates a new KalturaMetadataReplacementOptionsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMetadataReplacementOptionsItemWithDefaults() *KalturaMetadataReplacementOptionsItem {
	this := KalturaMetadataReplacementOptionsItem{}
	return &this
}

func (o KalturaMetadataReplacementOptionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPluginReplacementOptionsItem, errKalturaPluginReplacementOptionsItem := json.Marshal(o.KalturaPluginReplacementOptionsItem)
	if errKalturaPluginReplacementOptionsItem != nil {
		return []byte{}, errKalturaPluginReplacementOptionsItem
	}
	errKalturaPluginReplacementOptionsItem = json.Unmarshal([]byte(serializedKalturaPluginReplacementOptionsItem), &toSerialize)
	if errKalturaPluginReplacementOptionsItem != nil {
		return []byte{}, errKalturaPluginReplacementOptionsItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMetadataReplacementOptionsItem struct {
	value *KalturaMetadataReplacementOptionsItem
	isSet bool
}

func (v NullableKalturaMetadataReplacementOptionsItem) Get() *KalturaMetadataReplacementOptionsItem {
	return v.value
}

func (v *NullableKalturaMetadataReplacementOptionsItem) Set(val *KalturaMetadataReplacementOptionsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMetadataReplacementOptionsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMetadataReplacementOptionsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMetadataReplacementOptionsItem(val *KalturaMetadataReplacementOptionsItem) *NullableKalturaMetadataReplacementOptionsItem {
	return &NullableKalturaMetadataReplacementOptionsItem{value: val, isSet: true}
}

func (v NullableKalturaMetadataReplacementOptionsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMetadataReplacementOptionsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


