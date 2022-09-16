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

// MediaUpdateRequest struct for MediaUpdateRequest
type MediaUpdateRequest struct {
	EntryId string `json:"entryId"`
	MediaEntry KalturaMediaEntry `json:"mediaEntry"`
}

// NewMediaUpdateRequest instantiates a new MediaUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaUpdateRequest(entryId string, mediaEntry KalturaMediaEntry) *MediaUpdateRequest {
	this := MediaUpdateRequest{}
	this.EntryId = entryId
	this.MediaEntry = mediaEntry
	return &this
}

// NewMediaUpdateRequestWithDefaults instantiates a new MediaUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaUpdateRequestWithDefaults() *MediaUpdateRequest {
	this := MediaUpdateRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *MediaUpdateRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *MediaUpdateRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *MediaUpdateRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetMediaEntry returns the MediaEntry field value
func (o *MediaUpdateRequest) GetMediaEntry() KalturaMediaEntry {
	if o == nil {
		var ret KalturaMediaEntry
		return ret
	}

	return o.MediaEntry
}

// GetMediaEntryOk returns a tuple with the MediaEntry field value
// and a boolean to check if the value has been set.
func (o *MediaUpdateRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaEntry, true
}

// SetMediaEntry sets field value
func (o *MediaUpdateRequest) SetMediaEntry(v KalturaMediaEntry) {
	o.MediaEntry = v
}

func (o MediaUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["mediaEntry"] = o.MediaEntry
	}
	return json.Marshal(toSerialize)
}

type NullableMediaUpdateRequest struct {
	value *MediaUpdateRequest
	isSet bool
}

func (v NullableMediaUpdateRequest) Get() *MediaUpdateRequest {
	return v.value
}

func (v *NullableMediaUpdateRequest) Set(val *MediaUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaUpdateRequest(val *MediaUpdateRequest) *NullableMediaUpdateRequest {
	return &NullableMediaUpdateRequest{value: val, isSet: true}
}

func (v NullableMediaUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


