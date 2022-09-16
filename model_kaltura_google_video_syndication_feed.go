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

// KalturaGoogleVideoSyndicationFeed struct for KalturaGoogleVideoSyndicationFeed
type KalturaGoogleVideoSyndicationFeed struct {
	KalturaBaseSyndicationFeed
}

// NewKalturaGoogleVideoSyndicationFeed instantiates a new KalturaGoogleVideoSyndicationFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGoogleVideoSyndicationFeed() *KalturaGoogleVideoSyndicationFeed {
	this := KalturaGoogleVideoSyndicationFeed{}
	return &this
}

// NewKalturaGoogleVideoSyndicationFeedWithDefaults instantiates a new KalturaGoogleVideoSyndicationFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGoogleVideoSyndicationFeedWithDefaults() *KalturaGoogleVideoSyndicationFeed {
	this := KalturaGoogleVideoSyndicationFeed{}
	return &this
}

func (o KalturaGoogleVideoSyndicationFeed) MarshalJSON() ([]byte, error) {
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

type NullableKalturaGoogleVideoSyndicationFeed struct {
	value *KalturaGoogleVideoSyndicationFeed
	isSet bool
}

func (v NullableKalturaGoogleVideoSyndicationFeed) Get() *KalturaGoogleVideoSyndicationFeed {
	return v.value
}

func (v *NullableKalturaGoogleVideoSyndicationFeed) Set(val *KalturaGoogleVideoSyndicationFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGoogleVideoSyndicationFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGoogleVideoSyndicationFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGoogleVideoSyndicationFeed(val *KalturaGoogleVideoSyndicationFeed) *NullableKalturaGoogleVideoSyndicationFeed {
	return &NullableKalturaGoogleVideoSyndicationFeed{value: val, isSet: true}
}

func (v NullableKalturaGoogleVideoSyndicationFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGoogleVideoSyndicationFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

