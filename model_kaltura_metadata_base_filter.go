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

// KalturaMetadataBaseFilter `abstract`
type KalturaMetadataBaseFilter struct {
	KalturaRelatedFilter
}

// NewKalturaMetadataBaseFilter instantiates a new KalturaMetadataBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMetadataBaseFilter() *KalturaMetadataBaseFilter {
	this := KalturaMetadataBaseFilter{}
	return &this
}

// NewKalturaMetadataBaseFilterWithDefaults instantiates a new KalturaMetadataBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMetadataBaseFilterWithDefaults() *KalturaMetadataBaseFilter {
	this := KalturaMetadataBaseFilter{}
	return &this
}

func (o KalturaMetadataBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaMetadataBaseFilter struct {
	value *KalturaMetadataBaseFilter
	isSet bool
}

func (v NullableKalturaMetadataBaseFilter) Get() *KalturaMetadataBaseFilter {
	return v.value
}

func (v *NullableKalturaMetadataBaseFilter) Set(val *KalturaMetadataBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMetadataBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMetadataBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMetadataBaseFilter(val *KalturaMetadataBaseFilter) *NullableKalturaMetadataBaseFilter {
	return &NullableKalturaMetadataBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaMetadataBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMetadataBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

