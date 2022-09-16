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

// KalturaVirtualEventFilter struct for KalturaVirtualEventFilter
type KalturaVirtualEventFilter struct {
	KalturaVirtualEventBaseFilter
}

// NewKalturaVirtualEventFilter instantiates a new KalturaVirtualEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirtualEventFilter() *KalturaVirtualEventFilter {
	this := KalturaVirtualEventFilter{}
	return &this
}

// NewKalturaVirtualEventFilterWithDefaults instantiates a new KalturaVirtualEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirtualEventFilterWithDefaults() *KalturaVirtualEventFilter {
	this := KalturaVirtualEventFilter{}
	return &this
}

func (o KalturaVirtualEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVirtualEventBaseFilter, errKalturaVirtualEventBaseFilter := json.Marshal(o.KalturaVirtualEventBaseFilter)
	if errKalturaVirtualEventBaseFilter != nil {
		return []byte{}, errKalturaVirtualEventBaseFilter
	}
	errKalturaVirtualEventBaseFilter = json.Unmarshal([]byte(serializedKalturaVirtualEventBaseFilter), &toSerialize)
	if errKalturaVirtualEventBaseFilter != nil {
		return []byte{}, errKalturaVirtualEventBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVirtualEventFilter struct {
	value *KalturaVirtualEventFilter
	isSet bool
}

func (v NullableKalturaVirtualEventFilter) Get() *KalturaVirtualEventFilter {
	return v.value
}

func (v *NullableKalturaVirtualEventFilter) Set(val *KalturaVirtualEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirtualEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirtualEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirtualEventFilter(val *KalturaVirtualEventFilter) *NullableKalturaVirtualEventFilter {
	return &NullableKalturaVirtualEventFilter{value: val, isSet: true}
}

func (v NullableKalturaVirtualEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirtualEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

