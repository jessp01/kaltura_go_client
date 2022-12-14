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

// KalturaESearchCategoryParams struct for KalturaESearchCategoryParams
type KalturaESearchCategoryParams struct {
	KalturaESearchParams
}

// NewKalturaESearchCategoryParams instantiates a new KalturaESearchCategoryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchCategoryParams() *KalturaESearchCategoryParams {
	this := KalturaESearchCategoryParams{}
	return &this
}

// NewKalturaESearchCategoryParamsWithDefaults instantiates a new KalturaESearchCategoryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchCategoryParamsWithDefaults() *KalturaESearchCategoryParams {
	this := KalturaESearchCategoryParams{}
	return &this
}

func (o KalturaESearchCategoryParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchParams, errKalturaESearchParams := json.Marshal(o.KalturaESearchParams)
	if errKalturaESearchParams != nil {
		return []byte{}, errKalturaESearchParams
	}
	errKalturaESearchParams = json.Unmarshal([]byte(serializedKalturaESearchParams), &toSerialize)
	if errKalturaESearchParams != nil {
		return []byte{}, errKalturaESearchParams
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchCategoryParams struct {
	value *KalturaESearchCategoryParams
	isSet bool
}

func (v NullableKalturaESearchCategoryParams) Get() *KalturaESearchCategoryParams {
	return v.value
}

func (v *NullableKalturaESearchCategoryParams) Set(val *KalturaESearchCategoryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchCategoryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchCategoryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchCategoryParams(val *KalturaESearchCategoryParams) *NullableKalturaESearchCategoryParams {
	return &NullableKalturaESearchCategoryParams{value: val, isSet: true}
}

func (v NullableKalturaESearchCategoryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchCategoryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


