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

// KalturaDistributionProviderFilter struct for KalturaDistributionProviderFilter
type KalturaDistributionProviderFilter struct {
	KalturaDistributionProviderBaseFilter
}

// NewKalturaDistributionProviderFilter instantiates a new KalturaDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDistributionProviderFilter() *KalturaDistributionProviderFilter {
	this := KalturaDistributionProviderFilter{}
	return &this
}

// NewKalturaDistributionProviderFilterWithDefaults instantiates a new KalturaDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDistributionProviderFilterWithDefaults() *KalturaDistributionProviderFilter {
	this := KalturaDistributionProviderFilter{}
	return &this
}

func (o KalturaDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionProviderBaseFilter, errKalturaDistributionProviderBaseFilter := json.Marshal(o.KalturaDistributionProviderBaseFilter)
	if errKalturaDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaDistributionProviderBaseFilter
	}
	errKalturaDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaDistributionProviderBaseFilter), &toSerialize)
	if errKalturaDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDistributionProviderFilter struct {
	value *KalturaDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaDistributionProviderFilter) Get() *KalturaDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaDistributionProviderFilter) Set(val *KalturaDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDistributionProviderFilter(val *KalturaDistributionProviderFilter) *NullableKalturaDistributionProviderFilter {
	return &NullableKalturaDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


