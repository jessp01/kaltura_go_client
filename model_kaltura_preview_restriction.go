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

// KalturaPreviewRestriction struct for KalturaPreviewRestriction
type KalturaPreviewRestriction struct {
	KalturaSessionRestriction
}

// NewKalturaPreviewRestriction instantiates a new KalturaPreviewRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPreviewRestriction() *KalturaPreviewRestriction {
	this := KalturaPreviewRestriction{}
	return &this
}

// NewKalturaPreviewRestrictionWithDefaults instantiates a new KalturaPreviewRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPreviewRestrictionWithDefaults() *KalturaPreviewRestriction {
	this := KalturaPreviewRestriction{}
	return &this
}

func (o KalturaPreviewRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSessionRestriction, errKalturaSessionRestriction := json.Marshal(o.KalturaSessionRestriction)
	if errKalturaSessionRestriction != nil {
		return []byte{}, errKalturaSessionRestriction
	}
	errKalturaSessionRestriction = json.Unmarshal([]byte(serializedKalturaSessionRestriction), &toSerialize)
	if errKalturaSessionRestriction != nil {
		return []byte{}, errKalturaSessionRestriction
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPreviewRestriction struct {
	value *KalturaPreviewRestriction
	isSet bool
}

func (v NullableKalturaPreviewRestriction) Get() *KalturaPreviewRestriction {
	return v.value
}

func (v *NullableKalturaPreviewRestriction) Set(val *KalturaPreviewRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPreviewRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPreviewRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPreviewRestriction(val *KalturaPreviewRestriction) *NullableKalturaPreviewRestriction {
	return &NullableKalturaPreviewRestriction{value: val, isSet: true}
}

func (v NullableKalturaPreviewRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPreviewRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

