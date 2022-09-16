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

// KalturaAttUverseDistributionProviderFilter struct for KalturaAttUverseDistributionProviderFilter
type KalturaAttUverseDistributionProviderFilter struct {
	KalturaAttUverseDistributionProviderBaseFilter
}

// NewKalturaAttUverseDistributionProviderFilter instantiates a new KalturaAttUverseDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttUverseDistributionProviderFilter() *KalturaAttUverseDistributionProviderFilter {
	this := KalturaAttUverseDistributionProviderFilter{}
	return &this
}

// NewKalturaAttUverseDistributionProviderFilterWithDefaults instantiates a new KalturaAttUverseDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttUverseDistributionProviderFilterWithDefaults() *KalturaAttUverseDistributionProviderFilter {
	this := KalturaAttUverseDistributionProviderFilter{}
	return &this
}

func (o KalturaAttUverseDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAttUverseDistributionProviderBaseFilter, errKalturaAttUverseDistributionProviderBaseFilter := json.Marshal(o.KalturaAttUverseDistributionProviderBaseFilter)
	if errKalturaAttUverseDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaAttUverseDistributionProviderBaseFilter
	}
	errKalturaAttUverseDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaAttUverseDistributionProviderBaseFilter), &toSerialize)
	if errKalturaAttUverseDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaAttUverseDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAttUverseDistributionProviderFilter struct {
	value *KalturaAttUverseDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaAttUverseDistributionProviderFilter) Get() *KalturaAttUverseDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaAttUverseDistributionProviderFilter) Set(val *KalturaAttUverseDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttUverseDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttUverseDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttUverseDistributionProviderFilter(val *KalturaAttUverseDistributionProviderFilter) *NullableKalturaAttUverseDistributionProviderFilter {
	return &NullableKalturaAttUverseDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaAttUverseDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttUverseDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

