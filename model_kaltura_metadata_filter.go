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

// KalturaMetadataFilter struct for KalturaMetadataFilter
type KalturaMetadataFilter struct {
	KalturaMetadataBaseFilter
}

// NewKalturaMetadataFilter instantiates a new KalturaMetadataFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMetadataFilter() *KalturaMetadataFilter {
	this := KalturaMetadataFilter{}
	return &this
}

// NewKalturaMetadataFilterWithDefaults instantiates a new KalturaMetadataFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMetadataFilterWithDefaults() *KalturaMetadataFilter {
	this := KalturaMetadataFilter{}
	return &this
}

func (o KalturaMetadataFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMetadataBaseFilter, errKalturaMetadataBaseFilter := json.Marshal(o.KalturaMetadataBaseFilter)
	if errKalturaMetadataBaseFilter != nil {
		return []byte{}, errKalturaMetadataBaseFilter
	}
	errKalturaMetadataBaseFilter = json.Unmarshal([]byte(serializedKalturaMetadataBaseFilter), &toSerialize)
	if errKalturaMetadataBaseFilter != nil {
		return []byte{}, errKalturaMetadataBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMetadataFilter struct {
	value *KalturaMetadataFilter
	isSet bool
}

func (v NullableKalturaMetadataFilter) Get() *KalturaMetadataFilter {
	return v.value
}

func (v *NullableKalturaMetadataFilter) Set(val *KalturaMetadataFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMetadataFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMetadataFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMetadataFilter(val *KalturaMetadataFilter) *NullableKalturaMetadataFilter {
	return &NullableKalturaMetadataFilter{value: val, isSet: true}
}

func (v NullableKalturaMetadataFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMetadataFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


