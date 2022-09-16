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

// ScheduleEventGetConflictsRequest struct for ScheduleEventGetConflictsRequest
type ScheduleEventGetConflictsRequest struct {
	ResourceIds string `json:"resourceIds"`
	ScheduleEvent KalturaScheduleEvent `json:"scheduleEvent"`
	ScheduleEventConflictType *int32 `json:"scheduleEventConflictType,omitempty"`
	ScheduleEventIdToIgnore *string `json:"scheduleEventIdToIgnore,omitempty"`
}

// NewScheduleEventGetConflictsRequest instantiates a new ScheduleEventGetConflictsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleEventGetConflictsRequest(resourceIds string, scheduleEvent KalturaScheduleEvent) *ScheduleEventGetConflictsRequest {
	this := ScheduleEventGetConflictsRequest{}
	this.ResourceIds = resourceIds
	this.ScheduleEvent = scheduleEvent
	return &this
}

// NewScheduleEventGetConflictsRequestWithDefaults instantiates a new ScheduleEventGetConflictsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleEventGetConflictsRequestWithDefaults() *ScheduleEventGetConflictsRequest {
	this := ScheduleEventGetConflictsRequest{}
	return &this
}

// GetResourceIds returns the ResourceIds field value
func (o *ScheduleEventGetConflictsRequest) GetResourceIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *ScheduleEventGetConflictsRequest) GetResourceIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceIds, true
}

// SetResourceIds sets field value
func (o *ScheduleEventGetConflictsRequest) SetResourceIds(v string) {
	o.ResourceIds = v
}

// GetScheduleEvent returns the ScheduleEvent field value
func (o *ScheduleEventGetConflictsRequest) GetScheduleEvent() KalturaScheduleEvent {
	if o == nil {
		var ret KalturaScheduleEvent
		return ret
	}

	return o.ScheduleEvent
}

// GetScheduleEventOk returns a tuple with the ScheduleEvent field value
// and a boolean to check if the value has been set.
func (o *ScheduleEventGetConflictsRequest) GetScheduleEventOk() (*KalturaScheduleEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleEvent, true
}

// SetScheduleEvent sets field value
func (o *ScheduleEventGetConflictsRequest) SetScheduleEvent(v KalturaScheduleEvent) {
	o.ScheduleEvent = v
}

// GetScheduleEventConflictType returns the ScheduleEventConflictType field value if set, zero value otherwise.
func (o *ScheduleEventGetConflictsRequest) GetScheduleEventConflictType() int32 {
	if o == nil || o.ScheduleEventConflictType == nil {
		var ret int32
		return ret
	}
	return *o.ScheduleEventConflictType
}

// GetScheduleEventConflictTypeOk returns a tuple with the ScheduleEventConflictType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleEventGetConflictsRequest) GetScheduleEventConflictTypeOk() (*int32, bool) {
	if o == nil || o.ScheduleEventConflictType == nil {
		return nil, false
	}
	return o.ScheduleEventConflictType, true
}

// HasScheduleEventConflictType returns a boolean if a field has been set.
func (o *ScheduleEventGetConflictsRequest) HasScheduleEventConflictType() bool {
	if o != nil && o.ScheduleEventConflictType != nil {
		return true
	}

	return false
}

// SetScheduleEventConflictType gets a reference to the given int32 and assigns it to the ScheduleEventConflictType field.
func (o *ScheduleEventGetConflictsRequest) SetScheduleEventConflictType(v int32) {
	o.ScheduleEventConflictType = &v
}

// GetScheduleEventIdToIgnore returns the ScheduleEventIdToIgnore field value if set, zero value otherwise.
func (o *ScheduleEventGetConflictsRequest) GetScheduleEventIdToIgnore() string {
	if o == nil || o.ScheduleEventIdToIgnore == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEventIdToIgnore
}

// GetScheduleEventIdToIgnoreOk returns a tuple with the ScheduleEventIdToIgnore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleEventGetConflictsRequest) GetScheduleEventIdToIgnoreOk() (*string, bool) {
	if o == nil || o.ScheduleEventIdToIgnore == nil {
		return nil, false
	}
	return o.ScheduleEventIdToIgnore, true
}

// HasScheduleEventIdToIgnore returns a boolean if a field has been set.
func (o *ScheduleEventGetConflictsRequest) HasScheduleEventIdToIgnore() bool {
	if o != nil && o.ScheduleEventIdToIgnore != nil {
		return true
	}

	return false
}

// SetScheduleEventIdToIgnore gets a reference to the given string and assigns it to the ScheduleEventIdToIgnore field.
func (o *ScheduleEventGetConflictsRequest) SetScheduleEventIdToIgnore(v string) {
	o.ScheduleEventIdToIgnore = &v
}

func (o ScheduleEventGetConflictsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceIds"] = o.ResourceIds
	}
	if true {
		toSerialize["scheduleEvent"] = o.ScheduleEvent
	}
	if o.ScheduleEventConflictType != nil {
		toSerialize["scheduleEventConflictType"] = o.ScheduleEventConflictType
	}
	if o.ScheduleEventIdToIgnore != nil {
		toSerialize["scheduleEventIdToIgnore"] = o.ScheduleEventIdToIgnore
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleEventGetConflictsRequest struct {
	value *ScheduleEventGetConflictsRequest
	isSet bool
}

func (v NullableScheduleEventGetConflictsRequest) Get() *ScheduleEventGetConflictsRequest {
	return v.value
}

func (v *NullableScheduleEventGetConflictsRequest) Set(val *ScheduleEventGetConflictsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleEventGetConflictsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleEventGetConflictsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleEventGetConflictsRequest(val *ScheduleEventGetConflictsRequest) *NullableScheduleEventGetConflictsRequest {
	return &NullableScheduleEventGetConflictsRequest{value: val, isSet: true}
}

func (v NullableScheduleEventGetConflictsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleEventGetConflictsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

