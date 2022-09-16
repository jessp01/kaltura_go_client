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

// KalturaSynacorHboDistributionProviderBaseFilter `abstract`
type KalturaSynacorHboDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaSynacorHboDistributionProviderBaseFilter instantiates a new KalturaSynacorHboDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSynacorHboDistributionProviderBaseFilter() *KalturaSynacorHboDistributionProviderBaseFilter {
	this := KalturaSynacorHboDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaSynacorHboDistributionProviderBaseFilterWithDefaults instantiates a new KalturaSynacorHboDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSynacorHboDistributionProviderBaseFilterWithDefaults() *KalturaSynacorHboDistributionProviderBaseFilter {
	this := KalturaSynacorHboDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaSynacorHboDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaSynacorHboDistributionProviderBaseFilter struct {
	value *KalturaSynacorHboDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaSynacorHboDistributionProviderBaseFilter) Get() *KalturaSynacorHboDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaSynacorHboDistributionProviderBaseFilter) Set(val *KalturaSynacorHboDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSynacorHboDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSynacorHboDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSynacorHboDistributionProviderBaseFilter(val *KalturaSynacorHboDistributionProviderBaseFilter) *NullableKalturaSynacorHboDistributionProviderBaseFilter {
	return &NullableKalturaSynacorHboDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaSynacorHboDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSynacorHboDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

