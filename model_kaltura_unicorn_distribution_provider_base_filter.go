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

// KalturaUnicornDistributionProviderBaseFilter `abstract`
type KalturaUnicornDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaUnicornDistributionProviderBaseFilter instantiates a new KalturaUnicornDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUnicornDistributionProviderBaseFilter() *KalturaUnicornDistributionProviderBaseFilter {
	this := KalturaUnicornDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaUnicornDistributionProviderBaseFilterWithDefaults instantiates a new KalturaUnicornDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUnicornDistributionProviderBaseFilterWithDefaults() *KalturaUnicornDistributionProviderBaseFilter {
	this := KalturaUnicornDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaUnicornDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionProviderFilter, errKalturaDistributionProviderFilter := json.Marshal(o.KalturaDistributionProviderFilter)
	if errKalturaDistributionProviderFilter != nil {
		return []byte{}, errKalturaDistributionProviderFilter
	}
	errKalturaDistributionProviderFilter = json.Unmarshal([]byte(serializedKalturaDistributionProviderFilter), &toSerialize)
	if errKalturaDistributionProviderFilter != nil {
		return []byte{}, errKalturaDistributionProviderFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUnicornDistributionProviderBaseFilter struct {
	value *KalturaUnicornDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaUnicornDistributionProviderBaseFilter) Get() *KalturaUnicornDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaUnicornDistributionProviderBaseFilter) Set(val *KalturaUnicornDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUnicornDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUnicornDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUnicornDistributionProviderBaseFilter(val *KalturaUnicornDistributionProviderBaseFilter) *NullableKalturaUnicornDistributionProviderBaseFilter {
	return &NullableKalturaUnicornDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaUnicornDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUnicornDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


