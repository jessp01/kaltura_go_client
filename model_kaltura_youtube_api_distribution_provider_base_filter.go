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

// KalturaYoutubeApiDistributionProviderBaseFilter `abstract`
type KalturaYoutubeApiDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaYoutubeApiDistributionProviderBaseFilter instantiates a new KalturaYoutubeApiDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaYoutubeApiDistributionProviderBaseFilter() *KalturaYoutubeApiDistributionProviderBaseFilter {
	this := KalturaYoutubeApiDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaYoutubeApiDistributionProviderBaseFilterWithDefaults instantiates a new KalturaYoutubeApiDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaYoutubeApiDistributionProviderBaseFilterWithDefaults() *KalturaYoutubeApiDistributionProviderBaseFilter {
	this := KalturaYoutubeApiDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaYoutubeApiDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaYoutubeApiDistributionProviderBaseFilter struct {
	value *KalturaYoutubeApiDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaYoutubeApiDistributionProviderBaseFilter) Get() *KalturaYoutubeApiDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaYoutubeApiDistributionProviderBaseFilter) Set(val *KalturaYoutubeApiDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaYoutubeApiDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaYoutubeApiDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaYoutubeApiDistributionProviderBaseFilter(val *KalturaYoutubeApiDistributionProviderBaseFilter) *NullableKalturaYoutubeApiDistributionProviderBaseFilter {
	return &NullableKalturaYoutubeApiDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaYoutubeApiDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaYoutubeApiDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


