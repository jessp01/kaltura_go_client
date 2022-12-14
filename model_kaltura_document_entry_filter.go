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

// KalturaDocumentEntryFilter struct for KalturaDocumentEntryFilter
type KalturaDocumentEntryFilter struct {
	KalturaDocumentEntryBaseFilter
}

// NewKalturaDocumentEntryFilter instantiates a new KalturaDocumentEntryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDocumentEntryFilter() *KalturaDocumentEntryFilter {
	this := KalturaDocumentEntryFilter{}
	return &this
}

// NewKalturaDocumentEntryFilterWithDefaults instantiates a new KalturaDocumentEntryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDocumentEntryFilterWithDefaults() *KalturaDocumentEntryFilter {
	this := KalturaDocumentEntryFilter{}
	return &this
}

func (o KalturaDocumentEntryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDocumentEntryBaseFilter, errKalturaDocumentEntryBaseFilter := json.Marshal(o.KalturaDocumentEntryBaseFilter)
	if errKalturaDocumentEntryBaseFilter != nil {
		return []byte{}, errKalturaDocumentEntryBaseFilter
	}
	errKalturaDocumentEntryBaseFilter = json.Unmarshal([]byte(serializedKalturaDocumentEntryBaseFilter), &toSerialize)
	if errKalturaDocumentEntryBaseFilter != nil {
		return []byte{}, errKalturaDocumentEntryBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDocumentEntryFilter struct {
	value *KalturaDocumentEntryFilter
	isSet bool
}

func (v NullableKalturaDocumentEntryFilter) Get() *KalturaDocumentEntryFilter {
	return v.value
}

func (v *NullableKalturaDocumentEntryFilter) Set(val *KalturaDocumentEntryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDocumentEntryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDocumentEntryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDocumentEntryFilter(val *KalturaDocumentEntryFilter) *NullableKalturaDocumentEntryFilter {
	return &NullableKalturaDocumentEntryFilter{value: val, isSet: true}
}

func (v NullableKalturaDocumentEntryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDocumentEntryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


