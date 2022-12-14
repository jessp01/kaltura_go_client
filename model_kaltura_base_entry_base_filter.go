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

// KalturaBaseEntryBaseFilter `abstract`
type KalturaBaseEntryBaseFilter struct {
	KalturaRelatedFilter
}

// NewKalturaBaseEntryBaseFilter instantiates a new KalturaBaseEntryBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBaseEntryBaseFilter() *KalturaBaseEntryBaseFilter {
	this := KalturaBaseEntryBaseFilter{}
	return &this
}

// NewKalturaBaseEntryBaseFilterWithDefaults instantiates a new KalturaBaseEntryBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBaseEntryBaseFilterWithDefaults() *KalturaBaseEntryBaseFilter {
	this := KalturaBaseEntryBaseFilter{}
	return &this
}

func (o KalturaBaseEntryBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaRelatedFilter, errKalturaRelatedFilter := json.Marshal(o.KalturaRelatedFilter)
	if errKalturaRelatedFilter != nil {
		return []byte{}, errKalturaRelatedFilter
	}
	errKalturaRelatedFilter = json.Unmarshal([]byte(serializedKalturaRelatedFilter), &toSerialize)
	if errKalturaRelatedFilter != nil {
		return []byte{}, errKalturaRelatedFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBaseEntryBaseFilter struct {
	value *KalturaBaseEntryBaseFilter
	isSet bool
}

func (v NullableKalturaBaseEntryBaseFilter) Get() *KalturaBaseEntryBaseFilter {
	return v.value
}

func (v *NullableKalturaBaseEntryBaseFilter) Set(val *KalturaBaseEntryBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBaseEntryBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBaseEntryBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBaseEntryBaseFilter(val *KalturaBaseEntryBaseFilter) *NullableKalturaBaseEntryBaseFilter {
	return &NullableKalturaBaseEntryBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaBaseEntryBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBaseEntryBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


