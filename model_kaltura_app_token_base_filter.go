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

// KalturaAppTokenBaseFilter `abstract`
type KalturaAppTokenBaseFilter struct {
	KalturaFilter
}

// NewKalturaAppTokenBaseFilter instantiates a new KalturaAppTokenBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAppTokenBaseFilter() *KalturaAppTokenBaseFilter {
	this := KalturaAppTokenBaseFilter{}
	return &this
}

// NewKalturaAppTokenBaseFilterWithDefaults instantiates a new KalturaAppTokenBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAppTokenBaseFilterWithDefaults() *KalturaAppTokenBaseFilter {
	this := KalturaAppTokenBaseFilter{}
	return &this
}

func (o KalturaAppTokenBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFilter, errKalturaFilter := json.Marshal(o.KalturaFilter)
	if errKalturaFilter != nil {
		return []byte{}, errKalturaFilter
	}
	errKalturaFilter = json.Unmarshal([]byte(serializedKalturaFilter), &toSerialize)
	if errKalturaFilter != nil {
		return []byte{}, errKalturaFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAppTokenBaseFilter struct {
	value *KalturaAppTokenBaseFilter
	isSet bool
}

func (v NullableKalturaAppTokenBaseFilter) Get() *KalturaAppTokenBaseFilter {
	return v.value
}

func (v *NullableKalturaAppTokenBaseFilter) Set(val *KalturaAppTokenBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAppTokenBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAppTokenBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAppTokenBaseFilter(val *KalturaAppTokenBaseFilter) *NullableKalturaAppTokenBaseFilter {
	return &NullableKalturaAppTokenBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaAppTokenBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAppTokenBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

