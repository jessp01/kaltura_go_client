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

// KalturaMediaFlavorParams struct for KalturaMediaFlavorParams
type KalturaMediaFlavorParams struct {
	KalturaFlavorParams
}

// NewKalturaMediaFlavorParams instantiates a new KalturaMediaFlavorParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMediaFlavorParams() *KalturaMediaFlavorParams {
	this := KalturaMediaFlavorParams{}
	return &this
}

// NewKalturaMediaFlavorParamsWithDefaults instantiates a new KalturaMediaFlavorParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMediaFlavorParamsWithDefaults() *KalturaMediaFlavorParams {
	this := KalturaMediaFlavorParams{}
	return &this
}

func (o KalturaMediaFlavorParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorParams, errKalturaFlavorParams := json.Marshal(o.KalturaFlavorParams)
	if errKalturaFlavorParams != nil {
		return []byte{}, errKalturaFlavorParams
	}
	errKalturaFlavorParams = json.Unmarshal([]byte(serializedKalturaFlavorParams), &toSerialize)
	if errKalturaFlavorParams != nil {
		return []byte{}, errKalturaFlavorParams
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMediaFlavorParams struct {
	value *KalturaMediaFlavorParams
	isSet bool
}

func (v NullableKalturaMediaFlavorParams) Get() *KalturaMediaFlavorParams {
	return v.value
}

func (v *NullableKalturaMediaFlavorParams) Set(val *KalturaMediaFlavorParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMediaFlavorParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMediaFlavorParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMediaFlavorParams(val *KalturaMediaFlavorParams) *NullableKalturaMediaFlavorParams {
	return &NullableKalturaMediaFlavorParams{value: val, isSet: true}
}

func (v NullableKalturaMediaFlavorParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMediaFlavorParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

