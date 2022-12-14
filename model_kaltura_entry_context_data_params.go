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

// KalturaEntryContextDataParams Object which contains contextual entry-related data.
type KalturaEntryContextDataParams struct {
	KalturaAccessControlScope
}

// NewKalturaEntryContextDataParams instantiates a new KalturaEntryContextDataParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryContextDataParams() *KalturaEntryContextDataParams {
	this := KalturaEntryContextDataParams{}
	return &this
}

// NewKalturaEntryContextDataParamsWithDefaults instantiates a new KalturaEntryContextDataParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryContextDataParamsWithDefaults() *KalturaEntryContextDataParams {
	this := KalturaEntryContextDataParams{}
	return &this
}

func (o KalturaEntryContextDataParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAccessControlScope, errKalturaAccessControlScope := json.Marshal(o.KalturaAccessControlScope)
	if errKalturaAccessControlScope != nil {
		return []byte{}, errKalturaAccessControlScope
	}
	errKalturaAccessControlScope = json.Unmarshal([]byte(serializedKalturaAccessControlScope), &toSerialize)
	if errKalturaAccessControlScope != nil {
		return []byte{}, errKalturaAccessControlScope
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEntryContextDataParams struct {
	value *KalturaEntryContextDataParams
	isSet bool
}

func (v NullableKalturaEntryContextDataParams) Get() *KalturaEntryContextDataParams {
	return v.value
}

func (v *NullableKalturaEntryContextDataParams) Set(val *KalturaEntryContextDataParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryContextDataParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryContextDataParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryContextDataParams(val *KalturaEntryContextDataParams) *NullableKalturaEntryContextDataParams {
	return &NullableKalturaEntryContextDataParams{value: val, isSet: true}
}

func (v NullableKalturaEntryContextDataParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryContextDataParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


