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

// KalturaFtpDistributionProviderFilter struct for KalturaFtpDistributionProviderFilter
type KalturaFtpDistributionProviderFilter struct {
	KalturaFtpDistributionProviderBaseFilter
}

// NewKalturaFtpDistributionProviderFilter instantiates a new KalturaFtpDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFtpDistributionProviderFilter() *KalturaFtpDistributionProviderFilter {
	this := KalturaFtpDistributionProviderFilter{}
	return &this
}

// NewKalturaFtpDistributionProviderFilterWithDefaults instantiates a new KalturaFtpDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFtpDistributionProviderFilterWithDefaults() *KalturaFtpDistributionProviderFilter {
	this := KalturaFtpDistributionProviderFilter{}
	return &this
}

func (o KalturaFtpDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFtpDistributionProviderBaseFilter, errKalturaFtpDistributionProviderBaseFilter := json.Marshal(o.KalturaFtpDistributionProviderBaseFilter)
	if errKalturaFtpDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaFtpDistributionProviderBaseFilter
	}
	errKalturaFtpDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaFtpDistributionProviderBaseFilter), &toSerialize)
	if errKalturaFtpDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaFtpDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFtpDistributionProviderFilter struct {
	value *KalturaFtpDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaFtpDistributionProviderFilter) Get() *KalturaFtpDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaFtpDistributionProviderFilter) Set(val *KalturaFtpDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFtpDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFtpDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFtpDistributionProviderFilter(val *KalturaFtpDistributionProviderFilter) *NullableKalturaFtpDistributionProviderFilter {
	return &NullableKalturaFtpDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaFtpDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFtpDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


