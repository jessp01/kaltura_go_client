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

// KalturaAssetDistributionPropertyCondition Defines the condition to match a property and value on core asset object (or one if its inherited objects)
type KalturaAssetDistributionPropertyCondition struct {
	KalturaAssetDistributionCondition
}

// NewKalturaAssetDistributionPropertyCondition instantiates a new KalturaAssetDistributionPropertyCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAssetDistributionPropertyCondition() *KalturaAssetDistributionPropertyCondition {
	this := KalturaAssetDistributionPropertyCondition{}
	return &this
}

// NewKalturaAssetDistributionPropertyConditionWithDefaults instantiates a new KalturaAssetDistributionPropertyCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetDistributionPropertyConditionWithDefaults() *KalturaAssetDistributionPropertyCondition {
	this := KalturaAssetDistributionPropertyCondition{}
	return &this
}

func (o KalturaAssetDistributionPropertyCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAssetDistributionCondition, errKalturaAssetDistributionCondition := json.Marshal(o.KalturaAssetDistributionCondition)
	if errKalturaAssetDistributionCondition != nil {
		return []byte{}, errKalturaAssetDistributionCondition
	}
	errKalturaAssetDistributionCondition = json.Unmarshal([]byte(serializedKalturaAssetDistributionCondition), &toSerialize)
	if errKalturaAssetDistributionCondition != nil {
		return []byte{}, errKalturaAssetDistributionCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAssetDistributionPropertyCondition struct {
	value *KalturaAssetDistributionPropertyCondition
	isSet bool
}

func (v NullableKalturaAssetDistributionPropertyCondition) Get() *KalturaAssetDistributionPropertyCondition {
	return v.value
}

func (v *NullableKalturaAssetDistributionPropertyCondition) Set(val *KalturaAssetDistributionPropertyCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAssetDistributionPropertyCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAssetDistributionPropertyCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAssetDistributionPropertyCondition(val *KalturaAssetDistributionPropertyCondition) *NullableKalturaAssetDistributionPropertyCondition {
	return &NullableKalturaAssetDistributionPropertyCondition{value: val, isSet: true}
}

func (v NullableKalturaAssetDistributionPropertyCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAssetDistributionPropertyCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


