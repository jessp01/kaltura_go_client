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

// KalturaQuickPlayDistributionProviderFilter struct for KalturaQuickPlayDistributionProviderFilter
type KalturaQuickPlayDistributionProviderFilter struct {
	KalturaQuickPlayDistributionProviderBaseFilter
}

// NewKalturaQuickPlayDistributionProviderFilter instantiates a new KalturaQuickPlayDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuickPlayDistributionProviderFilter() *KalturaQuickPlayDistributionProviderFilter {
	this := KalturaQuickPlayDistributionProviderFilter{}
	return &this
}

// NewKalturaQuickPlayDistributionProviderFilterWithDefaults instantiates a new KalturaQuickPlayDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuickPlayDistributionProviderFilterWithDefaults() *KalturaQuickPlayDistributionProviderFilter {
	this := KalturaQuickPlayDistributionProviderFilter{}
	return &this
}

func (o KalturaQuickPlayDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaQuickPlayDistributionProviderBaseFilter, errKalturaQuickPlayDistributionProviderBaseFilter := json.Marshal(o.KalturaQuickPlayDistributionProviderBaseFilter)
	if errKalturaQuickPlayDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaQuickPlayDistributionProviderBaseFilter
	}
	errKalturaQuickPlayDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaQuickPlayDistributionProviderBaseFilter), &toSerialize)
	if errKalturaQuickPlayDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaQuickPlayDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuickPlayDistributionProviderFilter struct {
	value *KalturaQuickPlayDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaQuickPlayDistributionProviderFilter) Get() *KalturaQuickPlayDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaQuickPlayDistributionProviderFilter) Set(val *KalturaQuickPlayDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuickPlayDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuickPlayDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuickPlayDistributionProviderFilter(val *KalturaQuickPlayDistributionProviderFilter) *NullableKalturaQuickPlayDistributionProviderFilter {
	return &NullableKalturaQuickPlayDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaQuickPlayDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuickPlayDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


