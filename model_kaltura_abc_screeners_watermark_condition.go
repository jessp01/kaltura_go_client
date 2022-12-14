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

// KalturaABCScreenersWatermarkCondition struct for KalturaABCScreenersWatermarkCondition
type KalturaABCScreenersWatermarkCondition struct {
	KalturaCondition
}

// NewKalturaABCScreenersWatermarkCondition instantiates a new KalturaABCScreenersWatermarkCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaABCScreenersWatermarkCondition() *KalturaABCScreenersWatermarkCondition {
	this := KalturaABCScreenersWatermarkCondition{}
	return &this
}

// NewKalturaABCScreenersWatermarkConditionWithDefaults instantiates a new KalturaABCScreenersWatermarkCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaABCScreenersWatermarkConditionWithDefaults() *KalturaABCScreenersWatermarkCondition {
	this := KalturaABCScreenersWatermarkCondition{}
	return &this
}

func (o KalturaABCScreenersWatermarkCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCondition, errKalturaCondition := json.Marshal(o.KalturaCondition)
	if errKalturaCondition != nil {
		return []byte{}, errKalturaCondition
	}
	errKalturaCondition = json.Unmarshal([]byte(serializedKalturaCondition), &toSerialize)
	if errKalturaCondition != nil {
		return []byte{}, errKalturaCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaABCScreenersWatermarkCondition struct {
	value *KalturaABCScreenersWatermarkCondition
	isSet bool
}

func (v NullableKalturaABCScreenersWatermarkCondition) Get() *KalturaABCScreenersWatermarkCondition {
	return v.value
}

func (v *NullableKalturaABCScreenersWatermarkCondition) Set(val *KalturaABCScreenersWatermarkCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaABCScreenersWatermarkCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaABCScreenersWatermarkCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaABCScreenersWatermarkCondition(val *KalturaABCScreenersWatermarkCondition) *NullableKalturaABCScreenersWatermarkCondition {
	return &NullableKalturaABCScreenersWatermarkCondition{value: val, isSet: true}
}

func (v NullableKalturaABCScreenersWatermarkCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaABCScreenersWatermarkCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


