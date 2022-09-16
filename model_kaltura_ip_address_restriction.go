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

// KalturaIpAddressRestriction struct for KalturaIpAddressRestriction
type KalturaIpAddressRestriction struct {
	KalturaBaseRestriction
}

// NewKalturaIpAddressRestriction instantiates a new KalturaIpAddressRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaIpAddressRestriction() *KalturaIpAddressRestriction {
	this := KalturaIpAddressRestriction{}
	return &this
}

// NewKalturaIpAddressRestrictionWithDefaults instantiates a new KalturaIpAddressRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaIpAddressRestrictionWithDefaults() *KalturaIpAddressRestriction {
	this := KalturaIpAddressRestriction{}
	return &this
}

func (o KalturaIpAddressRestriction) MarshalJSON() ([]byte, error) {
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

type NullableKalturaIpAddressRestriction struct {
	value *KalturaIpAddressRestriction
	isSet bool
}

func (v NullableKalturaIpAddressRestriction) Get() *KalturaIpAddressRestriction {
	return v.value
}

func (v *NullableKalturaIpAddressRestriction) Set(val *KalturaIpAddressRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaIpAddressRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaIpAddressRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaIpAddressRestriction(val *KalturaIpAddressRestriction) *NullableKalturaIpAddressRestriction {
	return &NullableKalturaIpAddressRestriction{value: val, isSet: true}
}

func (v NullableKalturaIpAddressRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaIpAddressRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


