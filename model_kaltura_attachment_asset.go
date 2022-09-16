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

// KalturaAttachmentAsset struct for KalturaAttachmentAsset
type KalturaAttachmentAsset struct {
	KalturaAsset
}

// NewKalturaAttachmentAsset instantiates a new KalturaAttachmentAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttachmentAsset() *KalturaAttachmentAsset {
	this := KalturaAttachmentAsset{}
	return &this
}

// NewKalturaAttachmentAssetWithDefaults instantiates a new KalturaAttachmentAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttachmentAssetWithDefaults() *KalturaAttachmentAsset {
	this := KalturaAttachmentAsset{}
	return &this
}

func (o KalturaAttachmentAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAsset, errKalturaAsset := json.Marshal(o.KalturaAsset)
	if errKalturaAsset != nil {
		return []byte{}, errKalturaAsset
	}
	errKalturaAsset = json.Unmarshal([]byte(serializedKalturaAsset), &toSerialize)
	if errKalturaAsset != nil {
		return []byte{}, errKalturaAsset
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAttachmentAsset struct {
	value *KalturaAttachmentAsset
	isSet bool
}

func (v NullableKalturaAttachmentAsset) Get() *KalturaAttachmentAsset {
	return v.value
}

func (v *NullableKalturaAttachmentAsset) Set(val *KalturaAttachmentAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttachmentAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttachmentAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttachmentAsset(val *KalturaAttachmentAsset) *NullableKalturaAttachmentAsset {
	return &NullableKalturaAttachmentAsset{value: val, isSet: true}
}

func (v NullableKalturaAttachmentAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttachmentAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

