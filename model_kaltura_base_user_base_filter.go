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

// KalturaBaseUserBaseFilter `abstract`
type KalturaBaseUserBaseFilter struct {
	KalturaRelatedFilter
}

// NewKalturaBaseUserBaseFilter instantiates a new KalturaBaseUserBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBaseUserBaseFilter() *KalturaBaseUserBaseFilter {
	this := KalturaBaseUserBaseFilter{}
	return &this
}

// NewKalturaBaseUserBaseFilterWithDefaults instantiates a new KalturaBaseUserBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBaseUserBaseFilterWithDefaults() *KalturaBaseUserBaseFilter {
	this := KalturaBaseUserBaseFilter{}
	return &this
}

func (o KalturaBaseUserBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBaseUserBaseFilter struct {
	value *KalturaBaseUserBaseFilter
	isSet bool
}

func (v NullableKalturaBaseUserBaseFilter) Get() *KalturaBaseUserBaseFilter {
	return v.value
}

func (v *NullableKalturaBaseUserBaseFilter) Set(val *KalturaBaseUserBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBaseUserBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBaseUserBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBaseUserBaseFilter(val *KalturaBaseUserBaseFilter) *NullableKalturaBaseUserBaseFilter {
	return &NullableKalturaBaseUserBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaBaseUserBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBaseUserBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


