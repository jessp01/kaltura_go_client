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

// KalturaCountryRestriction struct for KalturaCountryRestriction
type KalturaCountryRestriction struct {
	KalturaBaseRestriction
}

// NewKalturaCountryRestriction instantiates a new KalturaCountryRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCountryRestriction() *KalturaCountryRestriction {
	this := KalturaCountryRestriction{}
	return &this
}

// NewKalturaCountryRestrictionWithDefaults instantiates a new KalturaCountryRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCountryRestrictionWithDefaults() *KalturaCountryRestriction {
	this := KalturaCountryRestriction{}
	return &this
}

func (o KalturaCountryRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseRestriction, errKalturaBaseRestriction := json.Marshal(o.KalturaBaseRestriction)
	if errKalturaBaseRestriction != nil {
		return []byte{}, errKalturaBaseRestriction
	}
	errKalturaBaseRestriction = json.Unmarshal([]byte(serializedKalturaBaseRestriction), &toSerialize)
	if errKalturaBaseRestriction != nil {
		return []byte{}, errKalturaBaseRestriction
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCountryRestriction struct {
	value *KalturaCountryRestriction
	isSet bool
}

func (v NullableKalturaCountryRestriction) Get() *KalturaCountryRestriction {
	return v.value
}

func (v *NullableKalturaCountryRestriction) Set(val *KalturaCountryRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCountryRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCountryRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCountryRestriction(val *KalturaCountryRestriction) *NullableKalturaCountryRestriction {
	return &NullableKalturaCountryRestriction{value: val, isSet: true}
}

func (v NullableKalturaCountryRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCountryRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


