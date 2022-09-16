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

// KalturaVodScheduleEvent struct for KalturaVodScheduleEvent
type KalturaVodScheduleEvent struct {
	KalturaEntryScheduleEvent
}

// NewKalturaVodScheduleEvent instantiates a new KalturaVodScheduleEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVodScheduleEvent() *KalturaVodScheduleEvent {
	this := KalturaVodScheduleEvent{}
	return &this
}

// NewKalturaVodScheduleEventWithDefaults instantiates a new KalturaVodScheduleEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVodScheduleEventWithDefaults() *KalturaVodScheduleEvent {
	this := KalturaVodScheduleEvent{}
	return &this
}

func (o KalturaVodScheduleEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEntryScheduleEvent, errKalturaEntryScheduleEvent := json.Marshal(o.KalturaEntryScheduleEvent)
	if errKalturaEntryScheduleEvent != nil {
		return []byte{}, errKalturaEntryScheduleEvent
	}
	errKalturaEntryScheduleEvent = json.Unmarshal([]byte(serializedKalturaEntryScheduleEvent), &toSerialize)
	if errKalturaEntryScheduleEvent != nil {
		return []byte{}, errKalturaEntryScheduleEvent
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVodScheduleEvent struct {
	value *KalturaVodScheduleEvent
	isSet bool
}

func (v NullableKalturaVodScheduleEvent) Get() *KalturaVodScheduleEvent {
	return v.value
}

func (v *NullableKalturaVodScheduleEvent) Set(val *KalturaVodScheduleEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVodScheduleEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVodScheduleEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVodScheduleEvent(val *KalturaVodScheduleEvent) *NullableKalturaVodScheduleEvent {
	return &NullableKalturaVodScheduleEvent{value: val, isSet: true}
}

func (v NullableKalturaVodScheduleEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVodScheduleEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

