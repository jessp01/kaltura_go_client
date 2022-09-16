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

// KalturaITunesSyndicationFeed struct for KalturaITunesSyndicationFeed
type KalturaITunesSyndicationFeed struct {
	KalturaBaseSyndicationFeed
}

// NewKalturaITunesSyndicationFeed instantiates a new KalturaITunesSyndicationFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaITunesSyndicationFeed() *KalturaITunesSyndicationFeed {
	this := KalturaITunesSyndicationFeed{}
	return &this
}

// NewKalturaITunesSyndicationFeedWithDefaults instantiates a new KalturaITunesSyndicationFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaITunesSyndicationFeedWithDefaults() *KalturaITunesSyndicationFeed {
	this := KalturaITunesSyndicationFeed{}
	return &this
}

func (o KalturaITunesSyndicationFeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseSyndicationFeed, errKalturaBaseSyndicationFeed := json.Marshal(o.KalturaBaseSyndicationFeed)
	if errKalturaBaseSyndicationFeed != nil {
		return []byte{}, errKalturaBaseSyndicationFeed
	}
	errKalturaBaseSyndicationFeed = json.Unmarshal([]byte(serializedKalturaBaseSyndicationFeed), &toSerialize)
	if errKalturaBaseSyndicationFeed != nil {
		return []byte{}, errKalturaBaseSyndicationFeed
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaITunesSyndicationFeed struct {
	value *KalturaITunesSyndicationFeed
	isSet bool
}

func (v NullableKalturaITunesSyndicationFeed) Get() *KalturaITunesSyndicationFeed {
	return v.value
}

func (v *NullableKalturaITunesSyndicationFeed) Set(val *KalturaITunesSyndicationFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaITunesSyndicationFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaITunesSyndicationFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaITunesSyndicationFeed(val *KalturaITunesSyndicationFeed) *NullableKalturaITunesSyndicationFeed {
	return &NullableKalturaITunesSyndicationFeed{value: val, isSet: true}
}

func (v NullableKalturaITunesSyndicationFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaITunesSyndicationFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

