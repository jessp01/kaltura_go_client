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

// KalturaInteractivity struct for KalturaInteractivity
type KalturaInteractivity struct {
	KalturaBaseInteractivity
}

// NewKalturaInteractivity instantiates a new KalturaInteractivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaInteractivity() *KalturaInteractivity {
	this := KalturaInteractivity{}
	return &this
}

// NewKalturaInteractivityWithDefaults instantiates a new KalturaInteractivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaInteractivityWithDefaults() *KalturaInteractivity {
	this := KalturaInteractivity{}
	return &this
}

func (o KalturaInteractivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseInteractivity, errKalturaBaseInteractivity := json.Marshal(o.KalturaBaseInteractivity)
	if errKalturaBaseInteractivity != nil {
		return []byte{}, errKalturaBaseInteractivity
	}
	errKalturaBaseInteractivity = json.Unmarshal([]byte(serializedKalturaBaseInteractivity), &toSerialize)
	if errKalturaBaseInteractivity != nil {
		return []byte{}, errKalturaBaseInteractivity
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaInteractivity struct {
	value *KalturaInteractivity
	isSet bool
}

func (v NullableKalturaInteractivity) Get() *KalturaInteractivity {
	return v.value
}

func (v *NullableKalturaInteractivity) Set(val *KalturaInteractivity) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaInteractivity) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaInteractivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaInteractivity(val *KalturaInteractivity) *NullableKalturaInteractivity {
	return &NullableKalturaInteractivity{value: val, isSet: true}
}

func (v NullableKalturaInteractivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaInteractivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

