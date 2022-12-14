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

// KalturaUnicornDistributionProviderFilter struct for KalturaUnicornDistributionProviderFilter
type KalturaUnicornDistributionProviderFilter struct {
	KalturaUnicornDistributionProviderBaseFilter
}

// NewKalturaUnicornDistributionProviderFilter instantiates a new KalturaUnicornDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUnicornDistributionProviderFilter() *KalturaUnicornDistributionProviderFilter {
	this := KalturaUnicornDistributionProviderFilter{}
	return &this
}

// NewKalturaUnicornDistributionProviderFilterWithDefaults instantiates a new KalturaUnicornDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUnicornDistributionProviderFilterWithDefaults() *KalturaUnicornDistributionProviderFilter {
	this := KalturaUnicornDistributionProviderFilter{}
	return &this
}

func (o KalturaUnicornDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUnicornDistributionProviderBaseFilter, errKalturaUnicornDistributionProviderBaseFilter := json.Marshal(o.KalturaUnicornDistributionProviderBaseFilter)
	if errKalturaUnicornDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaUnicornDistributionProviderBaseFilter
	}
	errKalturaUnicornDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaUnicornDistributionProviderBaseFilter), &toSerialize)
	if errKalturaUnicornDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaUnicornDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUnicornDistributionProviderFilter struct {
	value *KalturaUnicornDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaUnicornDistributionProviderFilter) Get() *KalturaUnicornDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaUnicornDistributionProviderFilter) Set(val *KalturaUnicornDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUnicornDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUnicornDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUnicornDistributionProviderFilter(val *KalturaUnicornDistributionProviderFilter) *NullableKalturaUnicornDistributionProviderFilter {
	return &NullableKalturaUnicornDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaUnicornDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUnicornDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


