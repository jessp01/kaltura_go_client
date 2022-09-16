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

// KalturaScpDropFolderBaseFilter `abstract`
type KalturaScpDropFolderBaseFilter struct {
	KalturaSshDropFolderFilter
}

// NewKalturaScpDropFolderBaseFilter instantiates a new KalturaScpDropFolderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScpDropFolderBaseFilter() *KalturaScpDropFolderBaseFilter {
	this := KalturaScpDropFolderBaseFilter{}
	return &this
}

// NewKalturaScpDropFolderBaseFilterWithDefaults instantiates a new KalturaScpDropFolderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScpDropFolderBaseFilterWithDefaults() *KalturaScpDropFolderBaseFilter {
	this := KalturaScpDropFolderBaseFilter{}
	return &this
}

func (o KalturaScpDropFolderBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSshDropFolderFilter, errKalturaSshDropFolderFilter := json.Marshal(o.KalturaSshDropFolderFilter)
	if errKalturaSshDropFolderFilter != nil {
		return []byte{}, errKalturaSshDropFolderFilter
	}
	errKalturaSshDropFolderFilter = json.Unmarshal([]byte(serializedKalturaSshDropFolderFilter), &toSerialize)
	if errKalturaSshDropFolderFilter != nil {
		return []byte{}, errKalturaSshDropFolderFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScpDropFolderBaseFilter struct {
	value *KalturaScpDropFolderBaseFilter
	isSet bool
}

func (v NullableKalturaScpDropFolderBaseFilter) Get() *KalturaScpDropFolderBaseFilter {
	return v.value
}

func (v *NullableKalturaScpDropFolderBaseFilter) Set(val *KalturaScpDropFolderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScpDropFolderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScpDropFolderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScpDropFolderBaseFilter(val *KalturaScpDropFolderBaseFilter) *NullableKalturaScpDropFolderBaseFilter {
	return &NullableKalturaScpDropFolderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaScpDropFolderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScpDropFolderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

