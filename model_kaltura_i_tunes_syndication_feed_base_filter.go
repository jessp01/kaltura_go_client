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

// KalturaITunesSyndicationFeedBaseFilter `abstract`
type KalturaITunesSyndicationFeedBaseFilter struct {
	KalturaBaseSyndicationFeedFilter
}

// NewKalturaITunesSyndicationFeedBaseFilter instantiates a new KalturaITunesSyndicationFeedBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaITunesSyndicationFeedBaseFilter() *KalturaITunesSyndicationFeedBaseFilter {
	this := KalturaITunesSyndicationFeedBaseFilter{}
	return &this
}

// NewKalturaITunesSyndicationFeedBaseFilterWithDefaults instantiates a new KalturaITunesSyndicationFeedBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaITunesSyndicationFeedBaseFilterWithDefaults() *KalturaITunesSyndicationFeedBaseFilter {
	this := KalturaITunesSyndicationFeedBaseFilter{}
	return &this
}

func (o KalturaITunesSyndicationFeedBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaITunesSyndicationFeedBaseFilter struct {
	value *KalturaITunesSyndicationFeedBaseFilter
	isSet bool
}

func (v NullableKalturaITunesSyndicationFeedBaseFilter) Get() *KalturaITunesSyndicationFeedBaseFilter {
	return v.value
}

func (v *NullableKalturaITunesSyndicationFeedBaseFilter) Set(val *KalturaITunesSyndicationFeedBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaITunesSyndicationFeedBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaITunesSyndicationFeedBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaITunesSyndicationFeedBaseFilter(val *KalturaITunesSyndicationFeedBaseFilter) *NullableKalturaITunesSyndicationFeedBaseFilter {
	return &NullableKalturaITunesSyndicationFeedBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaITunesSyndicationFeedBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaITunesSyndicationFeedBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


