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

// KalturaTvinciDistributionProviderBaseFilter `abstract`
type KalturaTvinciDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaTvinciDistributionProviderBaseFilter instantiates a new KalturaTvinciDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTvinciDistributionProviderBaseFilter() *KalturaTvinciDistributionProviderBaseFilter {
	this := KalturaTvinciDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaTvinciDistributionProviderBaseFilterWithDefaults instantiates a new KalturaTvinciDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTvinciDistributionProviderBaseFilterWithDefaults() *KalturaTvinciDistributionProviderBaseFilter {
	this := KalturaTvinciDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaTvinciDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaTvinciDistributionProviderBaseFilter struct {
	value *KalturaTvinciDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaTvinciDistributionProviderBaseFilter) Get() *KalturaTvinciDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaTvinciDistributionProviderBaseFilter) Set(val *KalturaTvinciDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTvinciDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTvinciDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTvinciDistributionProviderBaseFilter(val *KalturaTvinciDistributionProviderBaseFilter) *NullableKalturaTvinciDistributionProviderBaseFilter {
	return &NullableKalturaTvinciDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaTvinciDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTvinciDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

