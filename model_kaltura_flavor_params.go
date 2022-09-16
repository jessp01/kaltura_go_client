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

// KalturaFlavorParams struct for KalturaFlavorParams
type KalturaFlavorParams struct {
	KalturaAssetParams
}

// NewKalturaFlavorParams instantiates a new KalturaFlavorParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFlavorParams() *KalturaFlavorParams {
	this := KalturaFlavorParams{}
	return &this
}

// NewKalturaFlavorParamsWithDefaults instantiates a new KalturaFlavorParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFlavorParamsWithDefaults() *KalturaFlavorParams {
	this := KalturaFlavorParams{}
	return &this
}

func (o KalturaFlavorParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAssetParams, errKalturaAssetParams := json.Marshal(o.KalturaAssetParams)
	if errKalturaAssetParams != nil {
		return []byte{}, errKalturaAssetParams
	}
	errKalturaAssetParams = json.Unmarshal([]byte(serializedKalturaAssetParams), &toSerialize)
	if errKalturaAssetParams != nil {
		return []byte{}, errKalturaAssetParams
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFlavorParams struct {
	value *KalturaFlavorParams
	isSet bool
}

func (v NullableKalturaFlavorParams) Get() *KalturaFlavorParams {
	return v.value
}

func (v *NullableKalturaFlavorParams) Set(val *KalturaFlavorParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFlavorParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFlavorParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFlavorParams(val *KalturaFlavorParams) *NullableKalturaFlavorParams {
	return &NullableKalturaFlavorParams{value: val, isSet: true}
}

func (v NullableKalturaFlavorParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFlavorParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

