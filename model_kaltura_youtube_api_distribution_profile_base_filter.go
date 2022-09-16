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

// KalturaYoutubeApiDistributionProfileBaseFilter `abstract`
type KalturaYoutubeApiDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaYoutubeApiDistributionProfileBaseFilter instantiates a new KalturaYoutubeApiDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaYoutubeApiDistributionProfileBaseFilter() *KalturaYoutubeApiDistributionProfileBaseFilter {
	this := KalturaYoutubeApiDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaYoutubeApiDistributionProfileBaseFilterWithDefaults instantiates a new KalturaYoutubeApiDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaYoutubeApiDistributionProfileBaseFilterWithDefaults() *KalturaYoutubeApiDistributionProfileBaseFilter {
	this := KalturaYoutubeApiDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaYoutubeApiDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaYoutubeApiDistributionProfileBaseFilter struct {
	value *KalturaYoutubeApiDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaYoutubeApiDistributionProfileBaseFilter) Get() *KalturaYoutubeApiDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaYoutubeApiDistributionProfileBaseFilter) Set(val *KalturaYoutubeApiDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaYoutubeApiDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaYoutubeApiDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaYoutubeApiDistributionProfileBaseFilter(val *KalturaYoutubeApiDistributionProfileBaseFilter) *NullableKalturaYoutubeApiDistributionProfileBaseFilter {
	return &NullableKalturaYoutubeApiDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaYoutubeApiDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaYoutubeApiDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


