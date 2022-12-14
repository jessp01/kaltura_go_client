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

// KalturaRegistrationUserEntryFilter struct for KalturaRegistrationUserEntryFilter
type KalturaRegistrationUserEntryFilter struct {
	KalturaUserEntryFilter
}

// NewKalturaRegistrationUserEntryFilter instantiates a new KalturaRegistrationUserEntryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaRegistrationUserEntryFilter() *KalturaRegistrationUserEntryFilter {
	this := KalturaRegistrationUserEntryFilter{}
	return &this
}

// NewKalturaRegistrationUserEntryFilterWithDefaults instantiates a new KalturaRegistrationUserEntryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaRegistrationUserEntryFilterWithDefaults() *KalturaRegistrationUserEntryFilter {
	this := KalturaRegistrationUserEntryFilter{}
	return &this
}

func (o KalturaRegistrationUserEntryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUserEntryFilter, errKalturaUserEntryFilter := json.Marshal(o.KalturaUserEntryFilter)
	if errKalturaUserEntryFilter != nil {
		return []byte{}, errKalturaUserEntryFilter
	}
	errKalturaUserEntryFilter = json.Unmarshal([]byte(serializedKalturaUserEntryFilter), &toSerialize)
	if errKalturaUserEntryFilter != nil {
		return []byte{}, errKalturaUserEntryFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaRegistrationUserEntryFilter struct {
	value *KalturaRegistrationUserEntryFilter
	isSet bool
}

func (v NullableKalturaRegistrationUserEntryFilter) Get() *KalturaRegistrationUserEntryFilter {
	return v.value
}

func (v *NullableKalturaRegistrationUserEntryFilter) Set(val *KalturaRegistrationUserEntryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaRegistrationUserEntryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaRegistrationUserEntryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaRegistrationUserEntryFilter(val *KalturaRegistrationUserEntryFilter) *NullableKalturaRegistrationUserEntryFilter {
	return &NullableKalturaRegistrationUserEntryFilter{value: val, isSet: true}
}

func (v NullableKalturaRegistrationUserEntryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaRegistrationUserEntryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


