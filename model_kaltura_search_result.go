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

// KalturaSearchResult struct for KalturaSearchResult
type KalturaSearchResult struct {
	KalturaSearch
}

// NewKalturaSearchResult instantiates a new KalturaSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSearchResult() *KalturaSearchResult {
	this := KalturaSearchResult{}
	return &this
}

// NewKalturaSearchResultWithDefaults instantiates a new KalturaSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSearchResultWithDefaults() *KalturaSearchResult {
	this := KalturaSearchResult{}
	return &this
}

func (o KalturaSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearch, errKalturaSearch := json.Marshal(o.KalturaSearch)
	if errKalturaSearch != nil {
		return []byte{}, errKalturaSearch
	}
	errKalturaSearch = json.Unmarshal([]byte(serializedKalturaSearch), &toSerialize)
	if errKalturaSearch != nil {
		return []byte{}, errKalturaSearch
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSearchResult struct {
	value *KalturaSearchResult
	isSet bool
}

func (v NullableKalturaSearchResult) Get() *KalturaSearchResult {
	return v.value
}

func (v *NullableKalturaSearchResult) Set(val *KalturaSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSearchResult(val *KalturaSearchResult) *NullableKalturaSearchResult {
	return &NullableKalturaSearchResult{value: val, isSet: true}
}

func (v NullableKalturaSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

