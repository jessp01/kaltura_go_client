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

// KalturaESearchCuePointItemData struct for KalturaESearchCuePointItemData
type KalturaESearchCuePointItemData struct {
	KalturaESearchItemData
}

// NewKalturaESearchCuePointItemData instantiates a new KalturaESearchCuePointItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchCuePointItemData() *KalturaESearchCuePointItemData {
	this := KalturaESearchCuePointItemData{}
	return &this
}

// NewKalturaESearchCuePointItemDataWithDefaults instantiates a new KalturaESearchCuePointItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchCuePointItemDataWithDefaults() *KalturaESearchCuePointItemData {
	this := KalturaESearchCuePointItemData{}
	return &this
}

func (o KalturaESearchCuePointItemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchItemData, errKalturaESearchItemData := json.Marshal(o.KalturaESearchItemData)
	if errKalturaESearchItemData != nil {
		return []byte{}, errKalturaESearchItemData
	}
	errKalturaESearchItemData = json.Unmarshal([]byte(serializedKalturaESearchItemData), &toSerialize)
	if errKalturaESearchItemData != nil {
		return []byte{}, errKalturaESearchItemData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchCuePointItemData struct {
	value *KalturaESearchCuePointItemData
	isSet bool
}

func (v NullableKalturaESearchCuePointItemData) Get() *KalturaESearchCuePointItemData {
	return v.value
}

func (v *NullableKalturaESearchCuePointItemData) Set(val *KalturaESearchCuePointItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchCuePointItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchCuePointItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchCuePointItemData(val *KalturaESearchCuePointItemData) *NullableKalturaESearchCuePointItemData {
	return &NullableKalturaESearchCuePointItemData{value: val, isSet: true}
}

func (v NullableKalturaESearchCuePointItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchCuePointItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


