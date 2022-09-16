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

// KalturaApiParameterPermissionItemBaseFilter `abstract`
type KalturaApiParameterPermissionItemBaseFilter struct {
	KalturaPermissionItemFilter
}

// NewKalturaApiParameterPermissionItemBaseFilter instantiates a new KalturaApiParameterPermissionItemBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaApiParameterPermissionItemBaseFilter() *KalturaApiParameterPermissionItemBaseFilter {
	this := KalturaApiParameterPermissionItemBaseFilter{}
	return &this
}

// NewKalturaApiParameterPermissionItemBaseFilterWithDefaults instantiates a new KalturaApiParameterPermissionItemBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaApiParameterPermissionItemBaseFilterWithDefaults() *KalturaApiParameterPermissionItemBaseFilter {
	this := KalturaApiParameterPermissionItemBaseFilter{}
	return &this
}

func (o KalturaApiParameterPermissionItemBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPermissionItemFilter, errKalturaPermissionItemFilter := json.Marshal(o.KalturaPermissionItemFilter)
	if errKalturaPermissionItemFilter != nil {
		return []byte{}, errKalturaPermissionItemFilter
	}
	errKalturaPermissionItemFilter = json.Unmarshal([]byte(serializedKalturaPermissionItemFilter), &toSerialize)
	if errKalturaPermissionItemFilter != nil {
		return []byte{}, errKalturaPermissionItemFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaApiParameterPermissionItemBaseFilter struct {
	value *KalturaApiParameterPermissionItemBaseFilter
	isSet bool
}

func (v NullableKalturaApiParameterPermissionItemBaseFilter) Get() *KalturaApiParameterPermissionItemBaseFilter {
	return v.value
}

func (v *NullableKalturaApiParameterPermissionItemBaseFilter) Set(val *KalturaApiParameterPermissionItemBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaApiParameterPermissionItemBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaApiParameterPermissionItemBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaApiParameterPermissionItemBaseFilter(val *KalturaApiParameterPermissionItemBaseFilter) *NullableKalturaApiParameterPermissionItemBaseFilter {
	return &NullableKalturaApiParameterPermissionItemBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaApiParameterPermissionItemBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaApiParameterPermissionItemBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

