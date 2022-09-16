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

// KalturaTimedThumbAssetBaseFilter `abstract`
type KalturaTimedThumbAssetBaseFilter struct {
	KalturaThumbAssetFilter
}

// NewKalturaTimedThumbAssetBaseFilter instantiates a new KalturaTimedThumbAssetBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTimedThumbAssetBaseFilter() *KalturaTimedThumbAssetBaseFilter {
	this := KalturaTimedThumbAssetBaseFilter{}
	return &this
}

// NewKalturaTimedThumbAssetBaseFilterWithDefaults instantiates a new KalturaTimedThumbAssetBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTimedThumbAssetBaseFilterWithDefaults() *KalturaTimedThumbAssetBaseFilter {
	this := KalturaTimedThumbAssetBaseFilter{}
	return &this
}

func (o KalturaTimedThumbAssetBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaThumbAssetFilter, errKalturaThumbAssetFilter := json.Marshal(o.KalturaThumbAssetFilter)
	if errKalturaThumbAssetFilter != nil {
		return []byte{}, errKalturaThumbAssetFilter
	}
	errKalturaThumbAssetFilter = json.Unmarshal([]byte(serializedKalturaThumbAssetFilter), &toSerialize)
	if errKalturaThumbAssetFilter != nil {
		return []byte{}, errKalturaThumbAssetFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaTimedThumbAssetBaseFilter struct {
	value *KalturaTimedThumbAssetBaseFilter
	isSet bool
}

func (v NullableKalturaTimedThumbAssetBaseFilter) Get() *KalturaTimedThumbAssetBaseFilter {
	return v.value
}

func (v *NullableKalturaTimedThumbAssetBaseFilter) Set(val *KalturaTimedThumbAssetBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTimedThumbAssetBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTimedThumbAssetBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTimedThumbAssetBaseFilter(val *KalturaTimedThumbAssetBaseFilter) *NullableKalturaTimedThumbAssetBaseFilter {
	return &NullableKalturaTimedThumbAssetBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaTimedThumbAssetBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTimedThumbAssetBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

