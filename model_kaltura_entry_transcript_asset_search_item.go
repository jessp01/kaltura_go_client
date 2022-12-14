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

// KalturaEntryTranscriptAssetSearchItem struct for KalturaEntryTranscriptAssetSearchItem
type KalturaEntryTranscriptAssetSearchItem struct {
	KalturaSearchItem
}

// NewKalturaEntryTranscriptAssetSearchItem instantiates a new KalturaEntryTranscriptAssetSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryTranscriptAssetSearchItem() *KalturaEntryTranscriptAssetSearchItem {
	this := KalturaEntryTranscriptAssetSearchItem{}
	return &this
}

// NewKalturaEntryTranscriptAssetSearchItemWithDefaults instantiates a new KalturaEntryTranscriptAssetSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryTranscriptAssetSearchItemWithDefaults() *KalturaEntryTranscriptAssetSearchItem {
	this := KalturaEntryTranscriptAssetSearchItem{}
	return &this
}

func (o KalturaEntryTranscriptAssetSearchItem) MarshalJSON() ([]byte, error) {
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

type NullableKalturaEntryTranscriptAssetSearchItem struct {
	value *KalturaEntryTranscriptAssetSearchItem
	isSet bool
}

func (v NullableKalturaEntryTranscriptAssetSearchItem) Get() *KalturaEntryTranscriptAssetSearchItem {
	return v.value
}

func (v *NullableKalturaEntryTranscriptAssetSearchItem) Set(val *KalturaEntryTranscriptAssetSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryTranscriptAssetSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryTranscriptAssetSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryTranscriptAssetSearchItem(val *KalturaEntryTranscriptAssetSearchItem) *NullableKalturaEntryTranscriptAssetSearchItem {
	return &NullableKalturaEntryTranscriptAssetSearchItem{value: val, isSet: true}
}

func (v NullableKalturaEntryTranscriptAssetSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryTranscriptAssetSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


