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

// KalturaAttachmentServeOptions struct for KalturaAttachmentServeOptions
type KalturaAttachmentServeOptions struct {
	KalturaAssetServeOptions
}

// NewKalturaAttachmentServeOptions instantiates a new KalturaAttachmentServeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttachmentServeOptions() *KalturaAttachmentServeOptions {
	this := KalturaAttachmentServeOptions{}
	return &this
}

// NewKalturaAttachmentServeOptionsWithDefaults instantiates a new KalturaAttachmentServeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttachmentServeOptionsWithDefaults() *KalturaAttachmentServeOptions {
	this := KalturaAttachmentServeOptions{}
	return &this
}

func (o KalturaAttachmentServeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAssetServeOptions, errKalturaAssetServeOptions := json.Marshal(o.KalturaAssetServeOptions)
	if errKalturaAssetServeOptions != nil {
		return []byte{}, errKalturaAssetServeOptions
	}
	errKalturaAssetServeOptions = json.Unmarshal([]byte(serializedKalturaAssetServeOptions), &toSerialize)
	if errKalturaAssetServeOptions != nil {
		return []byte{}, errKalturaAssetServeOptions
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAttachmentServeOptions struct {
	value *KalturaAttachmentServeOptions
	isSet bool
}

func (v NullableKalturaAttachmentServeOptions) Get() *KalturaAttachmentServeOptions {
	return v.value
}

func (v *NullableKalturaAttachmentServeOptions) Set(val *KalturaAttachmentServeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttachmentServeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttachmentServeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttachmentServeOptions(val *KalturaAttachmentServeOptions) *NullableKalturaAttachmentServeOptions {
	return &NullableKalturaAttachmentServeOptions{value: val, isSet: true}
}

func (v NullableKalturaAttachmentServeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttachmentServeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


