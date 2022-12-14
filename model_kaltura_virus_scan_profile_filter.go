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

// KalturaVirusScanProfileFilter struct for KalturaVirusScanProfileFilter
type KalturaVirusScanProfileFilter struct {
	KalturaVirusScanProfileBaseFilter
}

// NewKalturaVirusScanProfileFilter instantiates a new KalturaVirusScanProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirusScanProfileFilter() *KalturaVirusScanProfileFilter {
	this := KalturaVirusScanProfileFilter{}
	return &this
}

// NewKalturaVirusScanProfileFilterWithDefaults instantiates a new KalturaVirusScanProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirusScanProfileFilterWithDefaults() *KalturaVirusScanProfileFilter {
	this := KalturaVirusScanProfileFilter{}
	return &this
}

func (o KalturaVirusScanProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVirusScanProfileBaseFilter, errKalturaVirusScanProfileBaseFilter := json.Marshal(o.KalturaVirusScanProfileBaseFilter)
	if errKalturaVirusScanProfileBaseFilter != nil {
		return []byte{}, errKalturaVirusScanProfileBaseFilter
	}
	errKalturaVirusScanProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaVirusScanProfileBaseFilter), &toSerialize)
	if errKalturaVirusScanProfileBaseFilter != nil {
		return []byte{}, errKalturaVirusScanProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVirusScanProfileFilter struct {
	value *KalturaVirusScanProfileFilter
	isSet bool
}

func (v NullableKalturaVirusScanProfileFilter) Get() *KalturaVirusScanProfileFilter {
	return v.value
}

func (v *NullableKalturaVirusScanProfileFilter) Set(val *KalturaVirusScanProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirusScanProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirusScanProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirusScanProfileFilter(val *KalturaVirusScanProfileFilter) *NullableKalturaVirusScanProfileFilter {
	return &NullableKalturaVirusScanProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaVirusScanProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirusScanProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


