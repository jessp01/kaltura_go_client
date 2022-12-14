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

// LiveStreamCreatePeriodicSyncPointsRequest struct for LiveStreamCreatePeriodicSyncPointsRequest
type LiveStreamCreatePeriodicSyncPointsRequest struct {
	Duration int32 `json:"duration"`
	EntryId string `json:"entryId"`
	Interval int32 `json:"interval"`
}

// NewLiveStreamCreatePeriodicSyncPointsRequest instantiates a new LiveStreamCreatePeriodicSyncPointsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStreamCreatePeriodicSyncPointsRequest(duration int32, entryId string, interval int32) *LiveStreamCreatePeriodicSyncPointsRequest {
	this := LiveStreamCreatePeriodicSyncPointsRequest{}
	this.Duration = duration
	this.EntryId = entryId
	this.Interval = interval
	return &this
}

// NewLiveStreamCreatePeriodicSyncPointsRequestWithDefaults instantiates a new LiveStreamCreatePeriodicSyncPointsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamCreatePeriodicSyncPointsRequestWithDefaults() *LiveStreamCreatePeriodicSyncPointsRequest {
	this := LiveStreamCreatePeriodicSyncPointsRequest{}
	return &this
}

// GetDuration returns the Duration field value
func (o *LiveStreamCreatePeriodicSyncPointsRequest) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *LiveStreamCreatePeriodicSyncPointsRequest) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *LiveStreamCreatePeriodicSyncPointsRequest) SetDuration(v int32) {
	o.Duration = v
}

// GetEntryId returns the EntryId field value
func (o *LiveStreamCreatePeriodicSyncPointsRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *LiveStreamCreatePeriodicSyncPointsRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *LiveStreamCreatePeriodicSyncPointsRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetInterval returns the Interval field value
func (o *LiveStreamCreatePeriodicSyncPointsRequest) GetInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *LiveStreamCreatePeriodicSyncPointsRequest) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *LiveStreamCreatePeriodicSyncPointsRequest) SetInterval(v int32) {
	o.Interval = v
}

func (o LiveStreamCreatePeriodicSyncPointsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	return json.Marshal(toSerialize)
}

type NullableLiveStreamCreatePeriodicSyncPointsRequest struct {
	value *LiveStreamCreatePeriodicSyncPointsRequest
	isSet bool
}

func (v NullableLiveStreamCreatePeriodicSyncPointsRequest) Get() *LiveStreamCreatePeriodicSyncPointsRequest {
	return v.value
}

func (v *NullableLiveStreamCreatePeriodicSyncPointsRequest) Set(val *LiveStreamCreatePeriodicSyncPointsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStreamCreatePeriodicSyncPointsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStreamCreatePeriodicSyncPointsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStreamCreatePeriodicSyncPointsRequest(val *LiveStreamCreatePeriodicSyncPointsRequest) *NullableLiveStreamCreatePeriodicSyncPointsRequest {
	return &NullableLiveStreamCreatePeriodicSyncPointsRequest{value: val, isSet: true}
}

func (v NullableLiveStreamCreatePeriodicSyncPointsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveStreamCreatePeriodicSyncPointsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


