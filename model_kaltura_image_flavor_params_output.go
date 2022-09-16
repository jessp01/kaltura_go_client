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

// KalturaImageFlavorParamsOutput struct for KalturaImageFlavorParamsOutput
type KalturaImageFlavorParamsOutput struct {
	KalturaFlavorParamsOutput
}

// NewKalturaImageFlavorParamsOutput instantiates a new KalturaImageFlavorParamsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaImageFlavorParamsOutput() *KalturaImageFlavorParamsOutput {
	this := KalturaImageFlavorParamsOutput{}
	return &this
}

// NewKalturaImageFlavorParamsOutputWithDefaults instantiates a new KalturaImageFlavorParamsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaImageFlavorParamsOutputWithDefaults() *KalturaImageFlavorParamsOutput {
	this := KalturaImageFlavorParamsOutput{}
	return &this
}

func (o KalturaImageFlavorParamsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorParamsOutput, errKalturaFlavorParamsOutput := json.Marshal(o.KalturaFlavorParamsOutput)
	if errKalturaFlavorParamsOutput != nil {
		return []byte{}, errKalturaFlavorParamsOutput
	}
	errKalturaFlavorParamsOutput = json.Unmarshal([]byte(serializedKalturaFlavorParamsOutput), &toSerialize)
	if errKalturaFlavorParamsOutput != nil {
		return []byte{}, errKalturaFlavorParamsOutput
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaImageFlavorParamsOutput struct {
	value *KalturaImageFlavorParamsOutput
	isSet bool
}

func (v NullableKalturaImageFlavorParamsOutput) Get() *KalturaImageFlavorParamsOutput {
	return v.value
}

func (v *NullableKalturaImageFlavorParamsOutput) Set(val *KalturaImageFlavorParamsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaImageFlavorParamsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaImageFlavorParamsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaImageFlavorParamsOutput(val *KalturaImageFlavorParamsOutput) *NullableKalturaImageFlavorParamsOutput {
	return &NullableKalturaImageFlavorParamsOutput{value: val, isSet: true}
}

func (v NullableKalturaImageFlavorParamsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaImageFlavorParamsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

