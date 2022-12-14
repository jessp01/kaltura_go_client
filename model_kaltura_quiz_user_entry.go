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

// KalturaQuizUserEntry struct for KalturaQuizUserEntry
type KalturaQuizUserEntry struct {
	KalturaUserEntry
}

// NewKalturaQuizUserEntry instantiates a new KalturaQuizUserEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuizUserEntry() *KalturaQuizUserEntry {
	this := KalturaQuizUserEntry{}
	return &this
}

// NewKalturaQuizUserEntryWithDefaults instantiates a new KalturaQuizUserEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuizUserEntryWithDefaults() *KalturaQuizUserEntry {
	this := KalturaQuizUserEntry{}
	return &this
}

func (o KalturaQuizUserEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUserEntry, errKalturaUserEntry := json.Marshal(o.KalturaUserEntry)
	if errKalturaUserEntry != nil {
		return []byte{}, errKalturaUserEntry
	}
	errKalturaUserEntry = json.Unmarshal([]byte(serializedKalturaUserEntry), &toSerialize)
	if errKalturaUserEntry != nil {
		return []byte{}, errKalturaUserEntry
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuizUserEntry struct {
	value *KalturaQuizUserEntry
	isSet bool
}

func (v NullableKalturaQuizUserEntry) Get() *KalturaQuizUserEntry {
	return v.value
}

func (v *NullableKalturaQuizUserEntry) Set(val *KalturaQuizUserEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuizUserEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuizUserEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuizUserEntry(val *KalturaQuizUserEntry) *NullableKalturaQuizUserEntry {
	return &NullableKalturaQuizUserEntry{value: val, isSet: true}
}

func (v NullableKalturaQuizUserEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuizUserEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


