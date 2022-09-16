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

// KalturaSessionRestriction struct for KalturaSessionRestriction
type KalturaSessionRestriction struct {
	KalturaBaseRestriction
}

// NewKalturaSessionRestriction instantiates a new KalturaSessionRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSessionRestriction() *KalturaSessionRestriction {
	this := KalturaSessionRestriction{}
	return &this
}

// NewKalturaSessionRestrictionWithDefaults instantiates a new KalturaSessionRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSessionRestrictionWithDefaults() *KalturaSessionRestriction {
	this := KalturaSessionRestriction{}
	return &this
}

func (o KalturaSessionRestriction) MarshalJSON() ([]byte, error) {
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

type NullableKalturaSessionRestriction struct {
	value *KalturaSessionRestriction
	isSet bool
}

func (v NullableKalturaSessionRestriction) Get() *KalturaSessionRestriction {
	return v.value
}

func (v *NullableKalturaSessionRestriction) Set(val *KalturaSessionRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSessionRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSessionRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSessionRestriction(val *KalturaSessionRestriction) *NullableKalturaSessionRestriction {
	return &NullableKalturaSessionRestriction{value: val, isSet: true}
}

func (v NullableKalturaSessionRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSessionRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

