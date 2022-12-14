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

// KalturaScheduleEventBaseFilter `abstract`
type KalturaScheduleEventBaseFilter struct {
	KalturaRelatedFilter
}

// NewKalturaScheduleEventBaseFilter instantiates a new KalturaScheduleEventBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduleEventBaseFilter() *KalturaScheduleEventBaseFilter {
	this := KalturaScheduleEventBaseFilter{}
	return &this
}

// NewKalturaScheduleEventBaseFilterWithDefaults instantiates a new KalturaScheduleEventBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduleEventBaseFilterWithDefaults() *KalturaScheduleEventBaseFilter {
	this := KalturaScheduleEventBaseFilter{}
	return &this
}

func (o KalturaScheduleEventBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaScheduleEventBaseFilter struct {
	value *KalturaScheduleEventBaseFilter
	isSet bool
}

func (v NullableKalturaScheduleEventBaseFilter) Get() *KalturaScheduleEventBaseFilter {
	return v.value
}

func (v *NullableKalturaScheduleEventBaseFilter) Set(val *KalturaScheduleEventBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduleEventBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduleEventBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduleEventBaseFilter(val *KalturaScheduleEventBaseFilter) *NullableKalturaScheduleEventBaseFilter {
	return &NullableKalturaScheduleEventBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaScheduleEventBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduleEventBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


