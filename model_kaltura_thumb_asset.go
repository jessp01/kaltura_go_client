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

// KalturaThumbAsset struct for KalturaThumbAsset
type KalturaThumbAsset struct {
	KalturaAsset
}

// NewKalturaThumbAsset instantiates a new KalturaThumbAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaThumbAsset() *KalturaThumbAsset {
	this := KalturaThumbAsset{}
	return &this
}

// NewKalturaThumbAssetWithDefaults instantiates a new KalturaThumbAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaThumbAssetWithDefaults() *KalturaThumbAsset {
	this := KalturaThumbAsset{}
	return &this
}

func (o KalturaThumbAsset) MarshalJSON() ([]byte, error) {
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

type NullableKalturaThumbAsset struct {
	value *KalturaThumbAsset
	isSet bool
}

func (v NullableKalturaThumbAsset) Get() *KalturaThumbAsset {
	return v.value
}

func (v *NullableKalturaThumbAsset) Set(val *KalturaThumbAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaThumbAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaThumbAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaThumbAsset(val *KalturaThumbAsset) *NullableKalturaThumbAsset {
	return &NullableKalturaThumbAsset{value: val, isSet: true}
}

func (v NullableKalturaThumbAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaThumbAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


