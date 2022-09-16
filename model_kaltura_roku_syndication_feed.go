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

// KalturaRokuSyndicationFeed struct for KalturaRokuSyndicationFeed
type KalturaRokuSyndicationFeed struct {
	KalturaConstantXsltSyndicationFeed
}

// NewKalturaRokuSyndicationFeed instantiates a new KalturaRokuSyndicationFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaRokuSyndicationFeed() *KalturaRokuSyndicationFeed {
	this := KalturaRokuSyndicationFeed{}
	return &this
}

// NewKalturaRokuSyndicationFeedWithDefaults instantiates a new KalturaRokuSyndicationFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaRokuSyndicationFeedWithDefaults() *KalturaRokuSyndicationFeed {
	this := KalturaRokuSyndicationFeed{}
	return &this
}

func (o KalturaRokuSyndicationFeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConstantXsltSyndicationFeed, errKalturaConstantXsltSyndicationFeed := json.Marshal(o.KalturaConstantXsltSyndicationFeed)
	if errKalturaConstantXsltSyndicationFeed != nil {
		return []byte{}, errKalturaConstantXsltSyndicationFeed
	}
	errKalturaConstantXsltSyndicationFeed = json.Unmarshal([]byte(serializedKalturaConstantXsltSyndicationFeed), &toSerialize)
	if errKalturaConstantXsltSyndicationFeed != nil {
		return []byte{}, errKalturaConstantXsltSyndicationFeed
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaRokuSyndicationFeed struct {
	value *KalturaRokuSyndicationFeed
	isSet bool
}

func (v NullableKalturaRokuSyndicationFeed) Get() *KalturaRokuSyndicationFeed {
	return v.value
}

func (v *NullableKalturaRokuSyndicationFeed) Set(val *KalturaRokuSyndicationFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaRokuSyndicationFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaRokuSyndicationFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaRokuSyndicationFeed(val *KalturaRokuSyndicationFeed) *NullableKalturaRokuSyndicationFeed {
	return &NullableKalturaRokuSyndicationFeed{value: val, isSet: true}
}

func (v NullableKalturaRokuSyndicationFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaRokuSyndicationFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


