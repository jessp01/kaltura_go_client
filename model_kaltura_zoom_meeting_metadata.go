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

// KalturaZoomMeetingMetadata struct for KalturaZoomMeetingMetadata
type KalturaZoomMeetingMetadata struct {
	AccountId *string `json:"accountId,omitempty"`
	HostId *string `json:"hostId,omitempty"`
	MeetingId *string `json:"meetingId,omitempty"`
	MeetingStartTime *string `json:"meetingStartTime,omitempty"`
	Topic *string `json:"topic,omitempty"`
	// Enum Type: `KalturaRecordingType`
	Type *int32 `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewKalturaZoomMeetingMetadata instantiates a new KalturaZoomMeetingMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaZoomMeetingMetadata() *KalturaZoomMeetingMetadata {
	this := KalturaZoomMeetingMetadata{}
	return &this
}

// NewKalturaZoomMeetingMetadataWithDefaults instantiates a new KalturaZoomMeetingMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaZoomMeetingMetadataWithDefaults() *KalturaZoomMeetingMetadata {
	this := KalturaZoomMeetingMetadata{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *KalturaZoomMeetingMetadata) SetAccountId(v string) {
	o.AccountId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *KalturaZoomMeetingMetadata) SetHostId(v string) {
	o.HostId = &v
}

// GetMeetingId returns the MeetingId field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetMeetingId() string {
	if o == nil || o.MeetingId == nil {
		var ret string
		return ret
	}
	return *o.MeetingId
}

// GetMeetingIdOk returns a tuple with the MeetingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetMeetingIdOk() (*string, bool) {
	if o == nil || o.MeetingId == nil {
		return nil, false
	}
	return o.MeetingId, true
}

// HasMeetingId returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasMeetingId() bool {
	if o != nil && o.MeetingId != nil {
		return true
	}

	return false
}

// SetMeetingId gets a reference to the given string and assigns it to the MeetingId field.
func (o *KalturaZoomMeetingMetadata) SetMeetingId(v string) {
	o.MeetingId = &v
}

// GetMeetingStartTime returns the MeetingStartTime field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetMeetingStartTime() string {
	if o == nil || o.MeetingStartTime == nil {
		var ret string
		return ret
	}
	return *o.MeetingStartTime
}

// GetMeetingStartTimeOk returns a tuple with the MeetingStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetMeetingStartTimeOk() (*string, bool) {
	if o == nil || o.MeetingStartTime == nil {
		return nil, false
	}
	return o.MeetingStartTime, true
}

// HasMeetingStartTime returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasMeetingStartTime() bool {
	if o != nil && o.MeetingStartTime != nil {
		return true
	}

	return false
}

// SetMeetingStartTime gets a reference to the given string and assigns it to the MeetingStartTime field.
func (o *KalturaZoomMeetingMetadata) SetMeetingStartTime(v string) {
	o.MeetingStartTime = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *KalturaZoomMeetingMetadata) SetTopic(v string) {
	o.Topic = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *KalturaZoomMeetingMetadata) SetType(v int32) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *KalturaZoomMeetingMetadata) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomMeetingMetadata) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *KalturaZoomMeetingMetadata) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *KalturaZoomMeetingMetadata) SetUuid(v string) {
	o.Uuid = &v
}

func (o KalturaZoomMeetingMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.HostId != nil {
		toSerialize["hostId"] = o.HostId
	}
	if o.MeetingId != nil {
		toSerialize["meetingId"] = o.MeetingId
	}
	if o.MeetingStartTime != nil {
		toSerialize["meetingStartTime"] = o.MeetingStartTime
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaZoomMeetingMetadata struct {
	value *KalturaZoomMeetingMetadata
	isSet bool
}

func (v NullableKalturaZoomMeetingMetadata) Get() *KalturaZoomMeetingMetadata {
	return v.value
}

func (v *NullableKalturaZoomMeetingMetadata) Set(val *KalturaZoomMeetingMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaZoomMeetingMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaZoomMeetingMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaZoomMeetingMetadata(val *KalturaZoomMeetingMetadata) *NullableKalturaZoomMeetingMetadata {
	return &NullableKalturaZoomMeetingMetadata{value: val, isSet: true}
}

func (v NullableKalturaZoomMeetingMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaZoomMeetingMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


