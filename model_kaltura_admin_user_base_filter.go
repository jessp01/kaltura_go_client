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

// KalturaAdminUserBaseFilter `abstract`
type KalturaAdminUserBaseFilter struct {
	KalturaUserFilter
}

// NewKalturaAdminUserBaseFilter instantiates a new KalturaAdminUserBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAdminUserBaseFilter() *KalturaAdminUserBaseFilter {
	this := KalturaAdminUserBaseFilter{}
	return &this
}

// NewKalturaAdminUserBaseFilterWithDefaults instantiates a new KalturaAdminUserBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAdminUserBaseFilterWithDefaults() *KalturaAdminUserBaseFilter {
	this := KalturaAdminUserBaseFilter{}
	return &this
}

func (o KalturaAdminUserBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUserFilter, errKalturaUserFilter := json.Marshal(o.KalturaUserFilter)
	if errKalturaUserFilter != nil {
		return []byte{}, errKalturaUserFilter
	}
	errKalturaUserFilter = json.Unmarshal([]byte(serializedKalturaUserFilter), &toSerialize)
	if errKalturaUserFilter != nil {
		return []byte{}, errKalturaUserFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAdminUserBaseFilter struct {
	value *KalturaAdminUserBaseFilter
	isSet bool
}

func (v NullableKalturaAdminUserBaseFilter) Get() *KalturaAdminUserBaseFilter {
	return v.value
}

func (v *NullableKalturaAdminUserBaseFilter) Set(val *KalturaAdminUserBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAdminUserBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAdminUserBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAdminUserBaseFilter(val *KalturaAdminUserBaseFilter) *NullableKalturaAdminUserBaseFilter {
	return &NullableKalturaAdminUserBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaAdminUserBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAdminUserBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

