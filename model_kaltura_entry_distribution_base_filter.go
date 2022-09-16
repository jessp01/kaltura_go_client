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

// KalturaEntryDistributionBaseFilter `abstract`
type KalturaEntryDistributionBaseFilter struct {
	KalturaRelatedFilter
}

// NewKalturaEntryDistributionBaseFilter instantiates a new KalturaEntryDistributionBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryDistributionBaseFilter() *KalturaEntryDistributionBaseFilter {
	this := KalturaEntryDistributionBaseFilter{}
	return &this
}

// NewKalturaEntryDistributionBaseFilterWithDefaults instantiates a new KalturaEntryDistributionBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryDistributionBaseFilterWithDefaults() *KalturaEntryDistributionBaseFilter {
	this := KalturaEntryDistributionBaseFilter{}
	return &this
}

func (o KalturaEntryDistributionBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaEntryDistributionBaseFilter struct {
	value *KalturaEntryDistributionBaseFilter
	isSet bool
}

func (v NullableKalturaEntryDistributionBaseFilter) Get() *KalturaEntryDistributionBaseFilter {
	return v.value
}

func (v *NullableKalturaEntryDistributionBaseFilter) Set(val *KalturaEntryDistributionBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryDistributionBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryDistributionBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryDistributionBaseFilter(val *KalturaEntryDistributionBaseFilter) *NullableKalturaEntryDistributionBaseFilter {
	return &NullableKalturaEntryDistributionBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaEntryDistributionBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryDistributionBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

