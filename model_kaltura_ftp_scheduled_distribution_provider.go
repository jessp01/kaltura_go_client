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

// KalturaFtpScheduledDistributionProvider struct for KalturaFtpScheduledDistributionProvider
type KalturaFtpScheduledDistributionProvider struct {
	KalturaFtpDistributionProvider
}

// NewKalturaFtpScheduledDistributionProvider instantiates a new KalturaFtpScheduledDistributionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFtpScheduledDistributionProvider() *KalturaFtpScheduledDistributionProvider {
	this := KalturaFtpScheduledDistributionProvider{}
	return &this
}

// NewKalturaFtpScheduledDistributionProviderWithDefaults instantiates a new KalturaFtpScheduledDistributionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFtpScheduledDistributionProviderWithDefaults() *KalturaFtpScheduledDistributionProvider {
	this := KalturaFtpScheduledDistributionProvider{}
	return &this
}

func (o KalturaFtpScheduledDistributionProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFtpDistributionProvider, errKalturaFtpDistributionProvider := json.Marshal(o.KalturaFtpDistributionProvider)
	if errKalturaFtpDistributionProvider != nil {
		return []byte{}, errKalturaFtpDistributionProvider
	}
	errKalturaFtpDistributionProvider = json.Unmarshal([]byte(serializedKalturaFtpDistributionProvider), &toSerialize)
	if errKalturaFtpDistributionProvider != nil {
		return []byte{}, errKalturaFtpDistributionProvider
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFtpScheduledDistributionProvider struct {
	value *KalturaFtpScheduledDistributionProvider
	isSet bool
}

func (v NullableKalturaFtpScheduledDistributionProvider) Get() *KalturaFtpScheduledDistributionProvider {
	return v.value
}

func (v *NullableKalturaFtpScheduledDistributionProvider) Set(val *KalturaFtpScheduledDistributionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFtpScheduledDistributionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFtpScheduledDistributionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFtpScheduledDistributionProvider(val *KalturaFtpScheduledDistributionProvider) *NullableKalturaFtpScheduledDistributionProvider {
	return &NullableKalturaFtpScheduledDistributionProvider{value: val, isSet: true}
}

func (v NullableKalturaFtpScheduledDistributionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFtpScheduledDistributionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

