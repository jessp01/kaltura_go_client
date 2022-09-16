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

// KalturaApiActionPermissionItemFilter struct for KalturaApiActionPermissionItemFilter
type KalturaApiActionPermissionItemFilter struct {
	KalturaApiActionPermissionItemBaseFilter
}

// NewKalturaApiActionPermissionItemFilter instantiates a new KalturaApiActionPermissionItemFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaApiActionPermissionItemFilter() *KalturaApiActionPermissionItemFilter {
	this := KalturaApiActionPermissionItemFilter{}
	return &this
}

// NewKalturaApiActionPermissionItemFilterWithDefaults instantiates a new KalturaApiActionPermissionItemFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaApiActionPermissionItemFilterWithDefaults() *KalturaApiActionPermissionItemFilter {
	this := KalturaApiActionPermissionItemFilter{}
	return &this
}

func (o KalturaApiActionPermissionItemFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaApiActionPermissionItemBaseFilter, errKalturaApiActionPermissionItemBaseFilter := json.Marshal(o.KalturaApiActionPermissionItemBaseFilter)
	if errKalturaApiActionPermissionItemBaseFilter != nil {
		return []byte{}, errKalturaApiActionPermissionItemBaseFilter
	}
	errKalturaApiActionPermissionItemBaseFilter = json.Unmarshal([]byte(serializedKalturaApiActionPermissionItemBaseFilter), &toSerialize)
	if errKalturaApiActionPermissionItemBaseFilter != nil {
		return []byte{}, errKalturaApiActionPermissionItemBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaApiActionPermissionItemFilter struct {
	value *KalturaApiActionPermissionItemFilter
	isSet bool
}

func (v NullableKalturaApiActionPermissionItemFilter) Get() *KalturaApiActionPermissionItemFilter {
	return v.value
}

func (v *NullableKalturaApiActionPermissionItemFilter) Set(val *KalturaApiActionPermissionItemFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaApiActionPermissionItemFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaApiActionPermissionItemFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaApiActionPermissionItemFilter(val *KalturaApiActionPermissionItemFilter) *NullableKalturaApiActionPermissionItemFilter {
	return &NullableKalturaApiActionPermissionItemFilter{value: val, isSet: true}
}

func (v NullableKalturaApiActionPermissionItemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaApiActionPermissionItemFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


