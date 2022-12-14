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

// LiveStreamUnregisterMediaServerRequest struct for LiveStreamUnregisterMediaServerRequest
type LiveStreamUnregisterMediaServerRequest struct {
	EntryId string `json:"entryId"`
	Hostname string `json:"hostname"`
	MediaServerIndex string `json:"mediaServerIndex"`
}

// NewLiveStreamUnregisterMediaServerRequest instantiates a new LiveStreamUnregisterMediaServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStreamUnregisterMediaServerRequest(entryId string, hostname string, mediaServerIndex string) *LiveStreamUnregisterMediaServerRequest {
	this := LiveStreamUnregisterMediaServerRequest{}
	this.EntryId = entryId
	this.Hostname = hostname
	this.MediaServerIndex = mediaServerIndex
	return &this
}

// NewLiveStreamUnregisterMediaServerRequestWithDefaults instantiates a new LiveStreamUnregisterMediaServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamUnregisterMediaServerRequestWithDefaults() *LiveStreamUnregisterMediaServerRequest {
	this := LiveStreamUnregisterMediaServerRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *LiveStreamUnregisterMediaServerRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *LiveStreamUnregisterMediaServerRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *LiveStreamUnregisterMediaServerRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetHostname returns the Hostname field value
func (o *LiveStreamUnregisterMediaServerRequest) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *LiveStreamUnregisterMediaServerRequest) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *LiveStreamUnregisterMediaServerRequest) SetHostname(v string) {
	o.Hostname = v
}

// GetMediaServerIndex returns the MediaServerIndex field value
func (o *LiveStreamUnregisterMediaServerRequest) GetMediaServerIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaServerIndex
}

// GetMediaServerIndexOk returns a tuple with the MediaServerIndex field value
// and a boolean to check if the value has been set.
func (o *LiveStreamUnregisterMediaServerRequest) GetMediaServerIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaServerIndex, true
}

// SetMediaServerIndex sets field value
func (o *LiveStreamUnregisterMediaServerRequest) SetMediaServerIndex(v string) {
	o.MediaServerIndex = v
}

func (o LiveStreamUnregisterMediaServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["mediaServerIndex"] = o.MediaServerIndex
	}
	return json.Marshal(toSerialize)
}

type NullableLiveStreamUnregisterMediaServerRequest struct {
	value *LiveStreamUnregisterMediaServerRequest
	isSet bool
}

func (v NullableLiveStreamUnregisterMediaServerRequest) Get() *LiveStreamUnregisterMediaServerRequest {
	return v.value
}

func (v *NullableLiveStreamUnregisterMediaServerRequest) Set(val *LiveStreamUnregisterMediaServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStreamUnregisterMediaServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStreamUnregisterMediaServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStreamUnregisterMediaServerRequest(val *LiveStreamUnregisterMediaServerRequest) *NullableLiveStreamUnregisterMediaServerRequest {
	return &NullableLiveStreamUnregisterMediaServerRequest{value: val, isSet: true}
}

func (v NullableLiveStreamUnregisterMediaServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveStreamUnregisterMediaServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


