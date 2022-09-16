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

// KalturaVirusScanProfileBaseFilter `abstract`
type KalturaVirusScanProfileBaseFilter struct {
	KalturaFilter
}

// NewKalturaVirusScanProfileBaseFilter instantiates a new KalturaVirusScanProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirusScanProfileBaseFilter() *KalturaVirusScanProfileBaseFilter {
	this := KalturaVirusScanProfileBaseFilter{}
	return &this
}

// NewKalturaVirusScanProfileBaseFilterWithDefaults instantiates a new KalturaVirusScanProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirusScanProfileBaseFilterWithDefaults() *KalturaVirusScanProfileBaseFilter {
	this := KalturaVirusScanProfileBaseFilter{}
	return &this
}

func (o KalturaVirusScanProfileBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaVirusScanProfileBaseFilter struct {
	value *KalturaVirusScanProfileBaseFilter
	isSet bool
}

func (v NullableKalturaVirusScanProfileBaseFilter) Get() *KalturaVirusScanProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaVirusScanProfileBaseFilter) Set(val *KalturaVirusScanProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirusScanProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirusScanProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirusScanProfileBaseFilter(val *KalturaVirusScanProfileBaseFilter) *NullableKalturaVirusScanProfileBaseFilter {
	return &NullableKalturaVirusScanProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaVirusScanProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirusScanProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


