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

// KalturaTranscriptAsset struct for KalturaTranscriptAsset
type KalturaTranscriptAsset struct {
	KalturaAttachmentAsset
}

// NewKalturaTranscriptAsset instantiates a new KalturaTranscriptAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTranscriptAsset() *KalturaTranscriptAsset {
	this := KalturaTranscriptAsset{}
	return &this
}

// NewKalturaTranscriptAssetWithDefaults instantiates a new KalturaTranscriptAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTranscriptAssetWithDefaults() *KalturaTranscriptAsset {
	this := KalturaTranscriptAsset{}
	return &this
}

func (o KalturaTranscriptAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAttachmentAsset, errKalturaAttachmentAsset := json.Marshal(o.KalturaAttachmentAsset)
	if errKalturaAttachmentAsset != nil {
		return []byte{}, errKalturaAttachmentAsset
	}
	errKalturaAttachmentAsset = json.Unmarshal([]byte(serializedKalturaAttachmentAsset), &toSerialize)
	if errKalturaAttachmentAsset != nil {
		return []byte{}, errKalturaAttachmentAsset
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaTranscriptAsset struct {
	value *KalturaTranscriptAsset
	isSet bool
}

func (v NullableKalturaTranscriptAsset) Get() *KalturaTranscriptAsset {
	return v.value
}

func (v *NullableKalturaTranscriptAsset) Set(val *KalturaTranscriptAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTranscriptAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTranscriptAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTranscriptAsset(val *KalturaTranscriptAsset) *NullableKalturaTranscriptAsset {
	return &NullableKalturaTranscriptAsset{value: val, isSet: true}
}

func (v NullableKalturaTranscriptAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTranscriptAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


