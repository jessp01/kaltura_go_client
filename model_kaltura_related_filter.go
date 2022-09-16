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

// KalturaRelatedFilter `abstract`
type KalturaRelatedFilter struct {
	KalturaFilter
}

// NewKalturaRelatedFilter instantiates a new KalturaRelatedFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaRelatedFilter() *KalturaRelatedFilter {
	this := KalturaRelatedFilter{}
	return &this
}

// NewKalturaRelatedFilterWithDefaults instantiates a new KalturaRelatedFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaRelatedFilterWithDefaults() *KalturaRelatedFilter {
	this := KalturaRelatedFilter{}
	return &this
}

func (o KalturaRelatedFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaRelatedFilter struct {
	value *KalturaRelatedFilter
	isSet bool
}

func (v NullableKalturaRelatedFilter) Get() *KalturaRelatedFilter {
	return v.value
}

func (v *NullableKalturaRelatedFilter) Set(val *KalturaRelatedFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaRelatedFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaRelatedFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaRelatedFilter(val *KalturaRelatedFilter) *NullableKalturaRelatedFilter {
	return &NullableKalturaRelatedFilter{value: val, isSet: true}
}

func (v NullableKalturaRelatedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaRelatedFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

