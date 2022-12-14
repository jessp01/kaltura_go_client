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

// KalturaUrlTokenizerAkamaiSecureHd struct for KalturaUrlTokenizerAkamaiSecureHd
type KalturaUrlTokenizerAkamaiSecureHd struct {
	KalturaUrlTokenizer
}

// NewKalturaUrlTokenizerAkamaiSecureHd instantiates a new KalturaUrlTokenizerAkamaiSecureHd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUrlTokenizerAkamaiSecureHd() *KalturaUrlTokenizerAkamaiSecureHd {
	this := KalturaUrlTokenizerAkamaiSecureHd{}
	return &this
}

// NewKalturaUrlTokenizerAkamaiSecureHdWithDefaults instantiates a new KalturaUrlTokenizerAkamaiSecureHd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUrlTokenizerAkamaiSecureHdWithDefaults() *KalturaUrlTokenizerAkamaiSecureHd {
	this := KalturaUrlTokenizerAkamaiSecureHd{}
	return &this
}

func (o KalturaUrlTokenizerAkamaiSecureHd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUrlTokenizer, errKalturaUrlTokenizer := json.Marshal(o.KalturaUrlTokenizer)
	if errKalturaUrlTokenizer != nil {
		return []byte{}, errKalturaUrlTokenizer
	}
	errKalturaUrlTokenizer = json.Unmarshal([]byte(serializedKalturaUrlTokenizer), &toSerialize)
	if errKalturaUrlTokenizer != nil {
		return []byte{}, errKalturaUrlTokenizer
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUrlTokenizerAkamaiSecureHd struct {
	value *KalturaUrlTokenizerAkamaiSecureHd
	isSet bool
}

func (v NullableKalturaUrlTokenizerAkamaiSecureHd) Get() *KalturaUrlTokenizerAkamaiSecureHd {
	return v.value
}

func (v *NullableKalturaUrlTokenizerAkamaiSecureHd) Set(val *KalturaUrlTokenizerAkamaiSecureHd) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUrlTokenizerAkamaiSecureHd) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUrlTokenizerAkamaiSecureHd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUrlTokenizerAkamaiSecureHd(val *KalturaUrlTokenizerAkamaiSecureHd) *NullableKalturaUrlTokenizerAkamaiSecureHd {
	return &NullableKalturaUrlTokenizerAkamaiSecureHd{value: val, isSet: true}
}

func (v NullableKalturaUrlTokenizerAkamaiSecureHd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUrlTokenizerAkamaiSecureHd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


