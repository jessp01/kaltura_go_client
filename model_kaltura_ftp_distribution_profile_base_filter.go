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

// KalturaFtpDistributionProfileBaseFilter `abstract`
type KalturaFtpDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaFtpDistributionProfileBaseFilter instantiates a new KalturaFtpDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFtpDistributionProfileBaseFilter() *KalturaFtpDistributionProfileBaseFilter {
	this := KalturaFtpDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaFtpDistributionProfileBaseFilterWithDefaults instantiates a new KalturaFtpDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFtpDistributionProfileBaseFilterWithDefaults() *KalturaFtpDistributionProfileBaseFilter {
	this := KalturaFtpDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaFtpDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConfigurableDistributionProfileFilter, errKalturaConfigurableDistributionProfileFilter := json.Marshal(o.KalturaConfigurableDistributionProfileFilter)
	if errKalturaConfigurableDistributionProfileFilter != nil {
		return []byte{}, errKalturaConfigurableDistributionProfileFilter
	}
	errKalturaConfigurableDistributionProfileFilter = json.Unmarshal([]byte(serializedKalturaConfigurableDistributionProfileFilter), &toSerialize)
	if errKalturaConfigurableDistributionProfileFilter != nil {
		return []byte{}, errKalturaConfigurableDistributionProfileFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFtpDistributionProfileBaseFilter struct {
	value *KalturaFtpDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaFtpDistributionProfileBaseFilter) Get() *KalturaFtpDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaFtpDistributionProfileBaseFilter) Set(val *KalturaFtpDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFtpDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFtpDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFtpDistributionProfileBaseFilter(val *KalturaFtpDistributionProfileBaseFilter) *NullableKalturaFtpDistributionProfileBaseFilter {
	return &NullableKalturaFtpDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaFtpDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFtpDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


