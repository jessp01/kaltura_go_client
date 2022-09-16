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

// KalturaEvalBooleanField Evaluates PHP statement, depends on the execution context
type KalturaEvalBooleanField struct {
	KalturaBooleanField
}

// NewKalturaEvalBooleanField instantiates a new KalturaEvalBooleanField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEvalBooleanField() *KalturaEvalBooleanField {
	this := KalturaEvalBooleanField{}
	return &this
}

// NewKalturaEvalBooleanFieldWithDefaults instantiates a new KalturaEvalBooleanField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEvalBooleanFieldWithDefaults() *KalturaEvalBooleanField {
	this := KalturaEvalBooleanField{}
	return &this
}

func (o KalturaEvalBooleanField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBooleanField, errKalturaBooleanField := json.Marshal(o.KalturaBooleanField)
	if errKalturaBooleanField != nil {
		return []byte{}, errKalturaBooleanField
	}
	errKalturaBooleanField = json.Unmarshal([]byte(serializedKalturaBooleanField), &toSerialize)
	if errKalturaBooleanField != nil {
		return []byte{}, errKalturaBooleanField
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEvalBooleanField struct {
	value *KalturaEvalBooleanField
	isSet bool
}

func (v NullableKalturaEvalBooleanField) Get() *KalturaEvalBooleanField {
	return v.value
}

func (v *NullableKalturaEvalBooleanField) Set(val *KalturaEvalBooleanField) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEvalBooleanField) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEvalBooleanField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEvalBooleanField(val *KalturaEvalBooleanField) *NullableKalturaEvalBooleanField {
	return &NullableKalturaEvalBooleanField{value: val, isSet: true}
}

func (v NullableKalturaEvalBooleanField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEvalBooleanField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

