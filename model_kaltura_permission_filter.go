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

// KalturaPermissionFilter struct for KalturaPermissionFilter
type KalturaPermissionFilter struct {
	KalturaPermissionBaseFilter
}

// NewKalturaPermissionFilter instantiates a new KalturaPermissionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPermissionFilter() *KalturaPermissionFilter {
	this := KalturaPermissionFilter{}
	return &this
}

// NewKalturaPermissionFilterWithDefaults instantiates a new KalturaPermissionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPermissionFilterWithDefaults() *KalturaPermissionFilter {
	this := KalturaPermissionFilter{}
	return &this
}

func (o KalturaPermissionFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPermissionBaseFilter, errKalturaPermissionBaseFilter := json.Marshal(o.KalturaPermissionBaseFilter)
	if errKalturaPermissionBaseFilter != nil {
		return []byte{}, errKalturaPermissionBaseFilter
	}
	errKalturaPermissionBaseFilter = json.Unmarshal([]byte(serializedKalturaPermissionBaseFilter), &toSerialize)
	if errKalturaPermissionBaseFilter != nil {
		return []byte{}, errKalturaPermissionBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPermissionFilter struct {
	value *KalturaPermissionFilter
	isSet bool
}

func (v NullableKalturaPermissionFilter) Get() *KalturaPermissionFilter {
	return v.value
}

func (v *NullableKalturaPermissionFilter) Set(val *KalturaPermissionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPermissionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPermissionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPermissionFilter(val *KalturaPermissionFilter) *NullableKalturaPermissionFilter {
	return &NullableKalturaPermissionFilter{value: val, isSet: true}
}

func (v NullableKalturaPermissionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPermissionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


