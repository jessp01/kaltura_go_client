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

// KalturaDropFolderBaseFilter `abstract`
type KalturaDropFolderBaseFilter struct {
	KalturaFilter
}

// NewKalturaDropFolderBaseFilter instantiates a new KalturaDropFolderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDropFolderBaseFilter() *KalturaDropFolderBaseFilter {
	this := KalturaDropFolderBaseFilter{}
	return &this
}

// NewKalturaDropFolderBaseFilterWithDefaults instantiates a new KalturaDropFolderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDropFolderBaseFilterWithDefaults() *KalturaDropFolderBaseFilter {
	this := KalturaDropFolderBaseFilter{}
	return &this
}

func (o KalturaDropFolderBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFilter, errKalturaFilter := json.Marshal(o.KalturaFilter)
	if errKalturaFilter != nil {
		return []byte{}, errKalturaFilter
	}
	errKalturaFilter = json.Unmarshal([]byte(serializedKalturaFilter), &toSerialize)
	if errKalturaFilter != nil {
		return []byte{}, errKalturaFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDropFolderBaseFilter struct {
	value *KalturaDropFolderBaseFilter
	isSet bool
}

func (v NullableKalturaDropFolderBaseFilter) Get() *KalturaDropFolderBaseFilter {
	return v.value
}

func (v *NullableKalturaDropFolderBaseFilter) Set(val *KalturaDropFolderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDropFolderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDropFolderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDropFolderBaseFilter(val *KalturaDropFolderBaseFilter) *NullableKalturaDropFolderBaseFilter {
	return &NullableKalturaDropFolderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaDropFolderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDropFolderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


