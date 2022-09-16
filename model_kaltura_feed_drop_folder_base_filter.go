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

// KalturaFeedDropFolderBaseFilter `abstract`
type KalturaFeedDropFolderBaseFilter struct {
	KalturaDropFolderFilter
}

// NewKalturaFeedDropFolderBaseFilter instantiates a new KalturaFeedDropFolderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFeedDropFolderBaseFilter() *KalturaFeedDropFolderBaseFilter {
	this := KalturaFeedDropFolderBaseFilter{}
	return &this
}

// NewKalturaFeedDropFolderBaseFilterWithDefaults instantiates a new KalturaFeedDropFolderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFeedDropFolderBaseFilterWithDefaults() *KalturaFeedDropFolderBaseFilter {
	this := KalturaFeedDropFolderBaseFilter{}
	return &this
}

func (o KalturaFeedDropFolderBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDropFolderFilter, errKalturaDropFolderFilter := json.Marshal(o.KalturaDropFolderFilter)
	if errKalturaDropFolderFilter != nil {
		return []byte{}, errKalturaDropFolderFilter
	}
	errKalturaDropFolderFilter = json.Unmarshal([]byte(serializedKalturaDropFolderFilter), &toSerialize)
	if errKalturaDropFolderFilter != nil {
		return []byte{}, errKalturaDropFolderFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFeedDropFolderBaseFilter struct {
	value *KalturaFeedDropFolderBaseFilter
	isSet bool
}

func (v NullableKalturaFeedDropFolderBaseFilter) Get() *KalturaFeedDropFolderBaseFilter {
	return v.value
}

func (v *NullableKalturaFeedDropFolderBaseFilter) Set(val *KalturaFeedDropFolderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFeedDropFolderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFeedDropFolderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFeedDropFolderBaseFilter(val *KalturaFeedDropFolderBaseFilter) *NullableKalturaFeedDropFolderBaseFilter {
	return &NullableKalturaFeedDropFolderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaFeedDropFolderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFeedDropFolderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


