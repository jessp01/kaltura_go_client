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

// KalturaLiveParams struct for KalturaLiveParams
type KalturaLiveParams struct {
	KalturaFlavorParams
}

// NewKalturaLiveParams instantiates a new KalturaLiveParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveParams() *KalturaLiveParams {
	this := KalturaLiveParams{}
	return &this
}

// NewKalturaLiveParamsWithDefaults instantiates a new KalturaLiveParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveParamsWithDefaults() *KalturaLiveParams {
	this := KalturaLiveParams{}
	return &this
}

func (o KalturaLiveParams) MarshalJSON() ([]byte, error) {
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

type NullableKalturaLiveParams struct {
	value *KalturaLiveParams
	isSet bool
}

func (v NullableKalturaLiveParams) Get() *KalturaLiveParams {
	return v.value
}

func (v *NullableKalturaLiveParams) Set(val *KalturaLiveParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveParams(val *KalturaLiveParams) *NullableKalturaLiveParams {
	return &NullableKalturaLiveParams{value: val, isSet: true}
}

func (v NullableKalturaLiveParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

