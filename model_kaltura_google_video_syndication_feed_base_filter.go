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

// KalturaGoogleVideoSyndicationFeedBaseFilter `abstract`
type KalturaGoogleVideoSyndicationFeedBaseFilter struct {
	KalturaBaseSyndicationFeedFilter
}

// NewKalturaGoogleVideoSyndicationFeedBaseFilter instantiates a new KalturaGoogleVideoSyndicationFeedBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGoogleVideoSyndicationFeedBaseFilter() *KalturaGoogleVideoSyndicationFeedBaseFilter {
	this := KalturaGoogleVideoSyndicationFeedBaseFilter{}
	return &this
}

// NewKalturaGoogleVideoSyndicationFeedBaseFilterWithDefaults instantiates a new KalturaGoogleVideoSyndicationFeedBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGoogleVideoSyndicationFeedBaseFilterWithDefaults() *KalturaGoogleVideoSyndicationFeedBaseFilter {
	this := KalturaGoogleVideoSyndicationFeedBaseFilter{}
	return &this
}

func (o KalturaGoogleVideoSyndicationFeedBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseSyndicationFeedFilter, errKalturaBaseSyndicationFeedFilter := json.Marshal(o.KalturaBaseSyndicationFeedFilter)
	if errKalturaBaseSyndicationFeedFilter != nil {
		return []byte{}, errKalturaBaseSyndicationFeedFilter
	}
	errKalturaBaseSyndicationFeedFilter = json.Unmarshal([]byte(serializedKalturaBaseSyndicationFeedFilter), &toSerialize)
	if errKalturaBaseSyndicationFeedFilter != nil {
		return []byte{}, errKalturaBaseSyndicationFeedFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaGoogleVideoSyndicationFeedBaseFilter struct {
	value *KalturaGoogleVideoSyndicationFeedBaseFilter
	isSet bool
}

func (v NullableKalturaGoogleVideoSyndicationFeedBaseFilter) Get() *KalturaGoogleVideoSyndicationFeedBaseFilter {
	return v.value
}

func (v *NullableKalturaGoogleVideoSyndicationFeedBaseFilter) Set(val *KalturaGoogleVideoSyndicationFeedBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGoogleVideoSyndicationFeedBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGoogleVideoSyndicationFeedBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGoogleVideoSyndicationFeedBaseFilter(val *KalturaGoogleVideoSyndicationFeedBaseFilter) *NullableKalturaGoogleVideoSyndicationFeedBaseFilter {
	return &NullableKalturaGoogleVideoSyndicationFeedBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaGoogleVideoSyndicationFeedBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGoogleVideoSyndicationFeedBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


