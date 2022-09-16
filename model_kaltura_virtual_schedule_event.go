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

// KalturaVirtualScheduleEvent struct for KalturaVirtualScheduleEvent
type KalturaVirtualScheduleEvent struct {
	KalturaScheduleEvent
}

// NewKalturaVirtualScheduleEvent instantiates a new KalturaVirtualScheduleEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirtualScheduleEvent() *KalturaVirtualScheduleEvent {
	this := KalturaVirtualScheduleEvent{}
	return &this
}

// NewKalturaVirtualScheduleEventWithDefaults instantiates a new KalturaVirtualScheduleEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirtualScheduleEventWithDefaults() *KalturaVirtualScheduleEvent {
	this := KalturaVirtualScheduleEvent{}
	return &this
}

func (o KalturaVirtualScheduleEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduleEvent, errKalturaScheduleEvent := json.Marshal(o.KalturaScheduleEvent)
	if errKalturaScheduleEvent != nil {
		return []byte{}, errKalturaScheduleEvent
	}
	errKalturaScheduleEvent = json.Unmarshal([]byte(serializedKalturaScheduleEvent), &toSerialize)
	if errKalturaScheduleEvent != nil {
		return []byte{}, errKalturaScheduleEvent
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVirtualScheduleEvent struct {
	value *KalturaVirtualScheduleEvent
	isSet bool
}

func (v NullableKalturaVirtualScheduleEvent) Get() *KalturaVirtualScheduleEvent {
	return v.value
}

func (v *NullableKalturaVirtualScheduleEvent) Set(val *KalturaVirtualScheduleEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirtualScheduleEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirtualScheduleEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirtualScheduleEvent(val *KalturaVirtualScheduleEvent) *NullableKalturaVirtualScheduleEvent {
	return &NullableKalturaVirtualScheduleEvent{value: val, isSet: true}
}

func (v NullableKalturaVirtualScheduleEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirtualScheduleEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

