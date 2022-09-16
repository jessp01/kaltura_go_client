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

// KalturaScheduledTaskProfileFilter struct for KalturaScheduledTaskProfileFilter
type KalturaScheduledTaskProfileFilter struct {
	KalturaScheduledTaskProfileBaseFilter
}

// NewKalturaScheduledTaskProfileFilter instantiates a new KalturaScheduledTaskProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduledTaskProfileFilter() *KalturaScheduledTaskProfileFilter {
	this := KalturaScheduledTaskProfileFilter{}
	return &this
}

// NewKalturaScheduledTaskProfileFilterWithDefaults instantiates a new KalturaScheduledTaskProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduledTaskProfileFilterWithDefaults() *KalturaScheduledTaskProfileFilter {
	this := KalturaScheduledTaskProfileFilter{}
	return &this
}

func (o KalturaScheduledTaskProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduledTaskProfileBaseFilter, errKalturaScheduledTaskProfileBaseFilter := json.Marshal(o.KalturaScheduledTaskProfileBaseFilter)
	if errKalturaScheduledTaskProfileBaseFilter != nil {
		return []byte{}, errKalturaScheduledTaskProfileBaseFilter
	}
	errKalturaScheduledTaskProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaScheduledTaskProfileBaseFilter), &toSerialize)
	if errKalturaScheduledTaskProfileBaseFilter != nil {
		return []byte{}, errKalturaScheduledTaskProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScheduledTaskProfileFilter struct {
	value *KalturaScheduledTaskProfileFilter
	isSet bool
}

func (v NullableKalturaScheduledTaskProfileFilter) Get() *KalturaScheduledTaskProfileFilter {
	return v.value
}

func (v *NullableKalturaScheduledTaskProfileFilter) Set(val *KalturaScheduledTaskProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduledTaskProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduledTaskProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduledTaskProfileFilter(val *KalturaScheduledTaskProfileFilter) *NullableKalturaScheduledTaskProfileFilter {
	return &NullableKalturaScheduledTaskProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaScheduledTaskProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduledTaskProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

