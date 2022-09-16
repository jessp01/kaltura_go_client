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

// KalturaTubeMogulSyndicationFeedBaseFilter `abstract`
type KalturaTubeMogulSyndicationFeedBaseFilter struct {
	KalturaBaseSyndicationFeedFilter
}

// NewKalturaTubeMogulSyndicationFeedBaseFilter instantiates a new KalturaTubeMogulSyndicationFeedBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTubeMogulSyndicationFeedBaseFilter() *KalturaTubeMogulSyndicationFeedBaseFilter {
	this := KalturaTubeMogulSyndicationFeedBaseFilter{}
	return &this
}

// NewKalturaTubeMogulSyndicationFeedBaseFilterWithDefaults instantiates a new KalturaTubeMogulSyndicationFeedBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTubeMogulSyndicationFeedBaseFilterWithDefaults() *KalturaTubeMogulSyndicationFeedBaseFilter {
	this := KalturaTubeMogulSyndicationFeedBaseFilter{}
	return &this
}

func (o KalturaTubeMogulSyndicationFeedBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaTubeMogulSyndicationFeedBaseFilter struct {
	value *KalturaTubeMogulSyndicationFeedBaseFilter
	isSet bool
}

func (v NullableKalturaTubeMogulSyndicationFeedBaseFilter) Get() *KalturaTubeMogulSyndicationFeedBaseFilter {
	return v.value
}

func (v *NullableKalturaTubeMogulSyndicationFeedBaseFilter) Set(val *KalturaTubeMogulSyndicationFeedBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTubeMogulSyndicationFeedBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTubeMogulSyndicationFeedBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTubeMogulSyndicationFeedBaseFilter(val *KalturaTubeMogulSyndicationFeedBaseFilter) *NullableKalturaTubeMogulSyndicationFeedBaseFilter {
	return &NullableKalturaTubeMogulSyndicationFeedBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaTubeMogulSyndicationFeedBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTubeMogulSyndicationFeedBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

