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

// KalturaPlaylistBaseFilter `abstract`
type KalturaPlaylistBaseFilter struct {
	KalturaBaseEntryFilter
}

// NewKalturaPlaylistBaseFilter instantiates a new KalturaPlaylistBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlaylistBaseFilter() *KalturaPlaylistBaseFilter {
	this := KalturaPlaylistBaseFilter{}
	return &this
}

// NewKalturaPlaylistBaseFilterWithDefaults instantiates a new KalturaPlaylistBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlaylistBaseFilterWithDefaults() *KalturaPlaylistBaseFilter {
	this := KalturaPlaylistBaseFilter{}
	return &this
}

func (o KalturaPlaylistBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseEntryFilter, errKalturaBaseEntryFilter := json.Marshal(o.KalturaBaseEntryFilter)
	if errKalturaBaseEntryFilter != nil {
		return []byte{}, errKalturaBaseEntryFilter
	}
	errKalturaBaseEntryFilter = json.Unmarshal([]byte(serializedKalturaBaseEntryFilter), &toSerialize)
	if errKalturaBaseEntryFilter != nil {
		return []byte{}, errKalturaBaseEntryFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlaylistBaseFilter struct {
	value *KalturaPlaylistBaseFilter
	isSet bool
}

func (v NullableKalturaPlaylistBaseFilter) Get() *KalturaPlaylistBaseFilter {
	return v.value
}

func (v *NullableKalturaPlaylistBaseFilter) Set(val *KalturaPlaylistBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlaylistBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlaylistBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlaylistBaseFilter(val *KalturaPlaylistBaseFilter) *NullableKalturaPlaylistBaseFilter {
	return &NullableKalturaPlaylistBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaPlaylistBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlaylistBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


