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

// KalturaCoordinatesContextField Represents the current request country context as calculated based on the IP address
type KalturaCoordinatesContextField struct {
	KalturaStringField
}

// NewKalturaCoordinatesContextField instantiates a new KalturaCoordinatesContextField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCoordinatesContextField() *KalturaCoordinatesContextField {
	this := KalturaCoordinatesContextField{}
	return &this
}

// NewKalturaCoordinatesContextFieldWithDefaults instantiates a new KalturaCoordinatesContextField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCoordinatesContextFieldWithDefaults() *KalturaCoordinatesContextField {
	this := KalturaCoordinatesContextField{}
	return &this
}

func (o KalturaCoordinatesContextField) MarshalJSON() ([]byte, error) {
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

type NullableKalturaCoordinatesContextField struct {
	value *KalturaCoordinatesContextField
	isSet bool
}

func (v NullableKalturaCoordinatesContextField) Get() *KalturaCoordinatesContextField {
	return v.value
}

func (v *NullableKalturaCoordinatesContextField) Set(val *KalturaCoordinatesContextField) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCoordinatesContextField) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCoordinatesContextField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCoordinatesContextField(val *KalturaCoordinatesContextField) *NullableKalturaCoordinatesContextField {
	return &NullableKalturaCoordinatesContextField{value: val, isSet: true}
}

func (v NullableKalturaCoordinatesContextField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCoordinatesContextField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

