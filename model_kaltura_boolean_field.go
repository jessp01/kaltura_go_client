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

// KalturaBooleanField `abstract`  A boolean representation to return evaluated dynamic value
type KalturaBooleanField struct {
	KalturaBooleanValue
}

// NewKalturaBooleanField instantiates a new KalturaBooleanField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBooleanField() *KalturaBooleanField {
	this := KalturaBooleanField{}
	return &this
}

// NewKalturaBooleanFieldWithDefaults instantiates a new KalturaBooleanField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBooleanFieldWithDefaults() *KalturaBooleanField {
	this := KalturaBooleanField{}
	return &this
}

func (o KalturaBooleanField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBooleanValue, errKalturaBooleanValue := json.Marshal(o.KalturaBooleanValue)
	if errKalturaBooleanValue != nil {
		return []byte{}, errKalturaBooleanValue
	}
	errKalturaBooleanValue = json.Unmarshal([]byte(serializedKalturaBooleanValue), &toSerialize)
	if errKalturaBooleanValue != nil {
		return []byte{}, errKalturaBooleanValue
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBooleanField struct {
	value *KalturaBooleanField
	isSet bool
}

func (v NullableKalturaBooleanField) Get() *KalturaBooleanField {
	return v.value
}

func (v *NullableKalturaBooleanField) Set(val *KalturaBooleanField) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBooleanField) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBooleanField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBooleanField(val *KalturaBooleanField) *NullableKalturaBooleanField {
	return &NullableKalturaBooleanField{value: val, isSet: true}
}

func (v NullableKalturaBooleanField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBooleanField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


