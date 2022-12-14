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

// KalturaEntryServerNodeFilter struct for KalturaEntryServerNodeFilter
type KalturaEntryServerNodeFilter struct {
	KalturaEntryServerNodeBaseFilter
}

// NewKalturaEntryServerNodeFilter instantiates a new KalturaEntryServerNodeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryServerNodeFilter() *KalturaEntryServerNodeFilter {
	this := KalturaEntryServerNodeFilter{}
	return &this
}

// NewKalturaEntryServerNodeFilterWithDefaults instantiates a new KalturaEntryServerNodeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryServerNodeFilterWithDefaults() *KalturaEntryServerNodeFilter {
	this := KalturaEntryServerNodeFilter{}
	return &this
}

func (o KalturaEntryServerNodeFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEntryServerNodeBaseFilter, errKalturaEntryServerNodeBaseFilter := json.Marshal(o.KalturaEntryServerNodeBaseFilter)
	if errKalturaEntryServerNodeBaseFilter != nil {
		return []byte{}, errKalturaEntryServerNodeBaseFilter
	}
	errKalturaEntryServerNodeBaseFilter = json.Unmarshal([]byte(serializedKalturaEntryServerNodeBaseFilter), &toSerialize)
	if errKalturaEntryServerNodeBaseFilter != nil {
		return []byte{}, errKalturaEntryServerNodeBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEntryServerNodeFilter struct {
	value *KalturaEntryServerNodeFilter
	isSet bool
}

func (v NullableKalturaEntryServerNodeFilter) Get() *KalturaEntryServerNodeFilter {
	return v.value
}

func (v *NullableKalturaEntryServerNodeFilter) Set(val *KalturaEntryServerNodeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryServerNodeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryServerNodeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryServerNodeFilter(val *KalturaEntryServerNodeFilter) *NullableKalturaEntryServerNodeFilter {
	return &NullableKalturaEntryServerNodeFilter{value: val, isSet: true}
}

func (v NullableKalturaEntryServerNodeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryServerNodeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


