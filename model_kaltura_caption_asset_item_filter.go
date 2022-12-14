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

// KalturaCaptionAssetItemFilter struct for KalturaCaptionAssetItemFilter
type KalturaCaptionAssetItemFilter struct {
	KalturaCaptionAssetFilter
}

// NewKalturaCaptionAssetItemFilter instantiates a new KalturaCaptionAssetItemFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCaptionAssetItemFilter() *KalturaCaptionAssetItemFilter {
	this := KalturaCaptionAssetItemFilter{}
	return &this
}

// NewKalturaCaptionAssetItemFilterWithDefaults instantiates a new KalturaCaptionAssetItemFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCaptionAssetItemFilterWithDefaults() *KalturaCaptionAssetItemFilter {
	this := KalturaCaptionAssetItemFilter{}
	return &this
}

func (o KalturaCaptionAssetItemFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCaptionAssetFilter, errKalturaCaptionAssetFilter := json.Marshal(o.KalturaCaptionAssetFilter)
	if errKalturaCaptionAssetFilter != nil {
		return []byte{}, errKalturaCaptionAssetFilter
	}
	errKalturaCaptionAssetFilter = json.Unmarshal([]byte(serializedKalturaCaptionAssetFilter), &toSerialize)
	if errKalturaCaptionAssetFilter != nil {
		return []byte{}, errKalturaCaptionAssetFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCaptionAssetItemFilter struct {
	value *KalturaCaptionAssetItemFilter
	isSet bool
}

func (v NullableKalturaCaptionAssetItemFilter) Get() *KalturaCaptionAssetItemFilter {
	return v.value
}

func (v *NullableKalturaCaptionAssetItemFilter) Set(val *KalturaCaptionAssetItemFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCaptionAssetItemFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCaptionAssetItemFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCaptionAssetItemFilter(val *KalturaCaptionAssetItemFilter) *NullableKalturaCaptionAssetItemFilter {
	return &NullableKalturaCaptionAssetItemFilter{value: val, isSet: true}
}

func (v NullableKalturaCaptionAssetItemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCaptionAssetItemFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


