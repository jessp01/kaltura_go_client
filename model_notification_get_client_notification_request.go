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

// NotificationGetClientNotificationRequest struct for NotificationGetClientNotificationRequest
type NotificationGetClientNotificationRequest struct {
	EntryId string `json:"entryId"`
	Type int32 `json:"type"`
}

// NewNotificationGetClientNotificationRequest instantiates a new NotificationGetClientNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationGetClientNotificationRequest(entryId string, type_ int32) *NotificationGetClientNotificationRequest {
	this := NotificationGetClientNotificationRequest{}
	this.EntryId = entryId
	this.Type = type_
	return &this
}

// NewNotificationGetClientNotificationRequestWithDefaults instantiates a new NotificationGetClientNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationGetClientNotificationRequestWithDefaults() *NotificationGetClientNotificationRequest {
	this := NotificationGetClientNotificationRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *NotificationGetClientNotificationRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *NotificationGetClientNotificationRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *NotificationGetClientNotificationRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetType returns the Type field value
func (o *NotificationGetClientNotificationRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationGetClientNotificationRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationGetClientNotificationRequest) SetType(v int32) {
	o.Type = v
}

func (o NotificationGetClientNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationGetClientNotificationRequest struct {
	value *NotificationGetClientNotificationRequest
	isSet bool
}

func (v NullableNotificationGetClientNotificationRequest) Get() *NotificationGetClientNotificationRequest {
	return v.value
}

func (v *NullableNotificationGetClientNotificationRequest) Set(val *NotificationGetClientNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationGetClientNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationGetClientNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationGetClientNotificationRequest(val *NotificationGetClientNotificationRequest) *NullableNotificationGetClientNotificationRequest {
	return &NullableNotificationGetClientNotificationRequest{value: val, isSet: true}
}

func (v NullableNotificationGetClientNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationGetClientNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

