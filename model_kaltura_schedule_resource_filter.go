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

// KalturaScheduleResourceFilter struct for KalturaScheduleResourceFilter
type KalturaScheduleResourceFilter struct {
	KalturaScheduleResourceBaseFilter
}

// NewKalturaScheduleResourceFilter instantiates a new KalturaScheduleResourceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduleResourceFilter() *KalturaScheduleResourceFilter {
	this := KalturaScheduleResourceFilter{}
	return &this
}

// NewKalturaScheduleResourceFilterWithDefaults instantiates a new KalturaScheduleResourceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduleResourceFilterWithDefaults() *KalturaScheduleResourceFilter {
	this := KalturaScheduleResourceFilter{}
	return &this
}

func (o KalturaScheduleResourceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduleResourceBaseFilter, errKalturaScheduleResourceBaseFilter := json.Marshal(o.KalturaScheduleResourceBaseFilter)
	if errKalturaScheduleResourceBaseFilter != nil {
		return []byte{}, errKalturaScheduleResourceBaseFilter
	}
	errKalturaScheduleResourceBaseFilter = json.Unmarshal([]byte(serializedKalturaScheduleResourceBaseFilter), &toSerialize)
	if errKalturaScheduleResourceBaseFilter != nil {
		return []byte{}, errKalturaScheduleResourceBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScheduleResourceFilter struct {
	value *KalturaScheduleResourceFilter
	isSet bool
}

func (v NullableKalturaScheduleResourceFilter) Get() *KalturaScheduleResourceFilter {
	return v.value
}

func (v *NullableKalturaScheduleResourceFilter) Set(val *KalturaScheduleResourceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduleResourceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduleResourceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduleResourceFilter(val *KalturaScheduleResourceFilter) *NullableKalturaScheduleResourceFilter {
	return &NullableKalturaScheduleResourceFilter{value: val, isSet: true}
}

func (v NullableKalturaScheduleResourceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduleResourceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


