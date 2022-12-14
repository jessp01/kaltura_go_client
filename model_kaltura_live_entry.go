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

// KalturaLiveEntry `abstract`
type KalturaLiveEntry struct {
	KalturaMediaEntry
}

// NewKalturaLiveEntry instantiates a new KalturaLiveEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveEntry() *KalturaLiveEntry {
	this := KalturaLiveEntry{}
	return &this
}

// NewKalturaLiveEntryWithDefaults instantiates a new KalturaLiveEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveEntryWithDefaults() *KalturaLiveEntry {
	this := KalturaLiveEntry{}
	return &this
}

func (o KalturaLiveEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMediaEntry, errKalturaMediaEntry := json.Marshal(o.KalturaMediaEntry)
	if errKalturaMediaEntry != nil {
		return []byte{}, errKalturaMediaEntry
	}
	errKalturaMediaEntry = json.Unmarshal([]byte(serializedKalturaMediaEntry), &toSerialize)
	if errKalturaMediaEntry != nil {
		return []byte{}, errKalturaMediaEntry
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveEntry struct {
	value *KalturaLiveEntry
	isSet bool
}

func (v NullableKalturaLiveEntry) Get() *KalturaLiveEntry {
	return v.value
}

func (v *NullableKalturaLiveEntry) Set(val *KalturaLiveEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveEntry(val *KalturaLiveEntry) *NullableKalturaLiveEntry {
	return &NullableKalturaLiveEntry{value: val, isSet: true}
}

func (v NullableKalturaLiveEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


