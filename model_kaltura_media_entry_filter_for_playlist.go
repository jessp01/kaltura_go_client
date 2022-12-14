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

// KalturaMediaEntryFilterForPlaylist struct for KalturaMediaEntryFilterForPlaylist
type KalturaMediaEntryFilterForPlaylist struct {
	KalturaMediaEntryFilter
}

// NewKalturaMediaEntryFilterForPlaylist instantiates a new KalturaMediaEntryFilterForPlaylist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMediaEntryFilterForPlaylist() *KalturaMediaEntryFilterForPlaylist {
	this := KalturaMediaEntryFilterForPlaylist{}
	return &this
}

// NewKalturaMediaEntryFilterForPlaylistWithDefaults instantiates a new KalturaMediaEntryFilterForPlaylist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMediaEntryFilterForPlaylistWithDefaults() *KalturaMediaEntryFilterForPlaylist {
	this := KalturaMediaEntryFilterForPlaylist{}
	return &this
}

func (o KalturaMediaEntryFilterForPlaylist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMediaEntryFilter, errKalturaMediaEntryFilter := json.Marshal(o.KalturaMediaEntryFilter)
	if errKalturaMediaEntryFilter != nil {
		return []byte{}, errKalturaMediaEntryFilter
	}
	errKalturaMediaEntryFilter = json.Unmarshal([]byte(serializedKalturaMediaEntryFilter), &toSerialize)
	if errKalturaMediaEntryFilter != nil {
		return []byte{}, errKalturaMediaEntryFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMediaEntryFilterForPlaylist struct {
	value *KalturaMediaEntryFilterForPlaylist
	isSet bool
}

func (v NullableKalturaMediaEntryFilterForPlaylist) Get() *KalturaMediaEntryFilterForPlaylist {
	return v.value
}

func (v *NullableKalturaMediaEntryFilterForPlaylist) Set(val *KalturaMediaEntryFilterForPlaylist) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMediaEntryFilterForPlaylist) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMediaEntryFilterForPlaylist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMediaEntryFilterForPlaylist(val *KalturaMediaEntryFilterForPlaylist) *NullableKalturaMediaEntryFilterForPlaylist {
	return &NullableKalturaMediaEntryFilterForPlaylist{value: val, isSet: true}
}

func (v NullableKalturaMediaEntryFilterForPlaylist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMediaEntryFilterForPlaylist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


