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

// KalturaUserBaseFilter `abstract`
type KalturaUserBaseFilter struct {
	KalturaBaseUserFilter
}

// NewKalturaUserBaseFilter instantiates a new KalturaUserBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUserBaseFilter() *KalturaUserBaseFilter {
	this := KalturaUserBaseFilter{}
	return &this
}

// NewKalturaUserBaseFilterWithDefaults instantiates a new KalturaUserBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUserBaseFilterWithDefaults() *KalturaUserBaseFilter {
	this := KalturaUserBaseFilter{}
	return &this
}

func (o KalturaUserBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseUserFilter, errKalturaBaseUserFilter := json.Marshal(o.KalturaBaseUserFilter)
	if errKalturaBaseUserFilter != nil {
		return []byte{}, errKalturaBaseUserFilter
	}
	errKalturaBaseUserFilter = json.Unmarshal([]byte(serializedKalturaBaseUserFilter), &toSerialize)
	if errKalturaBaseUserFilter != nil {
		return []byte{}, errKalturaBaseUserFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUserBaseFilter struct {
	value *KalturaUserBaseFilter
	isSet bool
}

func (v NullableKalturaUserBaseFilter) Get() *KalturaUserBaseFilter {
	return v.value
}

func (v *NullableKalturaUserBaseFilter) Set(val *KalturaUserBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUserBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUserBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUserBaseFilter(val *KalturaUserBaseFilter) *NullableKalturaUserBaseFilter {
	return &NullableKalturaUserBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaUserBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUserBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


