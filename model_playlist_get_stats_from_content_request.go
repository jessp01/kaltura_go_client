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

// PlaylistGetStatsFromContentRequest struct for PlaylistGetStatsFromContentRequest
type PlaylistGetStatsFromContentRequest struct {
	PlaylistContent string `json:"playlistContent"`
	PlaylistType int32 `json:"playlistType"`
}

// NewPlaylistGetStatsFromContentRequest instantiates a new PlaylistGetStatsFromContentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistGetStatsFromContentRequest(playlistContent string, playlistType int32) *PlaylistGetStatsFromContentRequest {
	this := PlaylistGetStatsFromContentRequest{}
	this.PlaylistContent = playlistContent
	this.PlaylistType = playlistType
	return &this
}

// NewPlaylistGetStatsFromContentRequestWithDefaults instantiates a new PlaylistGetStatsFromContentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistGetStatsFromContentRequestWithDefaults() *PlaylistGetStatsFromContentRequest {
	this := PlaylistGetStatsFromContentRequest{}
	return &this
}

// GetPlaylistContent returns the PlaylistContent field value
func (o *PlaylistGetStatsFromContentRequest) GetPlaylistContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaylistContent
}

// GetPlaylistContentOk returns a tuple with the PlaylistContent field value
// and a boolean to check if the value has been set.
func (o *PlaylistGetStatsFromContentRequest) GetPlaylistContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaylistContent, true
}

// SetPlaylistContent sets field value
func (o *PlaylistGetStatsFromContentRequest) SetPlaylistContent(v string) {
	o.PlaylistContent = v
}

// GetPlaylistType returns the PlaylistType field value
func (o *PlaylistGetStatsFromContentRequest) GetPlaylistType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlaylistType
}

// GetPlaylistTypeOk returns a tuple with the PlaylistType field value
// and a boolean to check if the value has been set.
func (o *PlaylistGetStatsFromContentRequest) GetPlaylistTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaylistType, true
}

// SetPlaylistType sets field value
func (o *PlaylistGetStatsFromContentRequest) SetPlaylistType(v int32) {
	o.PlaylistType = v
}

func (o PlaylistGetStatsFromContentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["playlistContent"] = o.PlaylistContent
	}
	if true {
		toSerialize["playlistType"] = o.PlaylistType
	}
	return json.Marshal(toSerialize)
}

type NullablePlaylistGetStatsFromContentRequest struct {
	value *PlaylistGetStatsFromContentRequest
	isSet bool
}

func (v NullablePlaylistGetStatsFromContentRequest) Get() *PlaylistGetStatsFromContentRequest {
	return v.value
}

func (v *NullablePlaylistGetStatsFromContentRequest) Set(val *PlaylistGetStatsFromContentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistGetStatsFromContentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistGetStatsFromContentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistGetStatsFromContentRequest(val *PlaylistGetStatsFromContentRequest) *NullablePlaylistGetStatsFromContentRequest {
	return &NullablePlaylistGetStatsFromContentRequest{value: val, isSet: true}
}

func (v NullablePlaylistGetStatsFromContentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistGetStatsFromContentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


