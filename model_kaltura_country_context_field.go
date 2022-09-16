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

// KalturaCountryContextField Represents the current request country context as calculated based on the IP address
type KalturaCountryContextField struct {
	KalturaStringField
}

// NewKalturaCountryContextField instantiates a new KalturaCountryContextField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCountryContextField() *KalturaCountryContextField {
	this := KalturaCountryContextField{}
	return &this
}

// NewKalturaCountryContextFieldWithDefaults instantiates a new KalturaCountryContextField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCountryContextFieldWithDefaults() *KalturaCountryContextField {
	this := KalturaCountryContextField{}
	return &this
}

func (o KalturaCountryContextField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaStringField, errKalturaStringField := json.Marshal(o.KalturaStringField)
	if errKalturaStringField != nil {
		return []byte{}, errKalturaStringField
	}
	errKalturaStringField = json.Unmarshal([]byte(serializedKalturaStringField), &toSerialize)
	if errKalturaStringField != nil {
		return []byte{}, errKalturaStringField
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCountryContextField struct {
	value *KalturaCountryContextField
	isSet bool
}

func (v NullableKalturaCountryContextField) Get() *KalturaCountryContextField {
	return v.value
}

func (v *NullableKalturaCountryContextField) Set(val *KalturaCountryContextField) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCountryContextField) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCountryContextField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCountryContextField(val *KalturaCountryContextField) *NullableKalturaCountryContextField {
	return &NullableKalturaCountryContextField{value: val, isSet: true}
}

func (v NullableKalturaCountryContextField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCountryContextField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


