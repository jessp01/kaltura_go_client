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

// KalturaAuthenticatedCondition struct for KalturaAuthenticatedCondition
type KalturaAuthenticatedCondition struct {
	KalturaCondition
}

// NewKalturaAuthenticatedCondition instantiates a new KalturaAuthenticatedCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAuthenticatedCondition() *KalturaAuthenticatedCondition {
	this := KalturaAuthenticatedCondition{}
	return &this
}

// NewKalturaAuthenticatedConditionWithDefaults instantiates a new KalturaAuthenticatedCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAuthenticatedConditionWithDefaults() *KalturaAuthenticatedCondition {
	this := KalturaAuthenticatedCondition{}
	return &this
}

func (o KalturaAuthenticatedCondition) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAuthenticatedCondition struct {
	value *KalturaAuthenticatedCondition
	isSet bool
}

func (v NullableKalturaAuthenticatedCondition) Get() *KalturaAuthenticatedCondition {
	return v.value
}

func (v *NullableKalturaAuthenticatedCondition) Set(val *KalturaAuthenticatedCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAuthenticatedCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAuthenticatedCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAuthenticatedCondition(val *KalturaAuthenticatedCondition) *NullableKalturaAuthenticatedCondition {
	return &NullableKalturaAuthenticatedCondition{value: val, isSet: true}
}

func (v NullableKalturaAuthenticatedCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAuthenticatedCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


