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

// EntryDistributionServeReturnedDataRequest struct for EntryDistributionServeReturnedDataRequest
type EntryDistributionServeReturnedDataRequest struct {
	ActionType int32 `json:"actionType"`
	Id int32 `json:"id"`
}

// NewEntryDistributionServeReturnedDataRequest instantiates a new EntryDistributionServeReturnedDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryDistributionServeReturnedDataRequest(actionType int32, id int32) *EntryDistributionServeReturnedDataRequest {
	this := EntryDistributionServeReturnedDataRequest{}
	this.ActionType = actionType
	this.Id = id
	return &this
}

// NewEntryDistributionServeReturnedDataRequestWithDefaults instantiates a new EntryDistributionServeReturnedDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryDistributionServeReturnedDataRequestWithDefaults() *EntryDistributionServeReturnedDataRequest {
	this := EntryDistributionServeReturnedDataRequest{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *EntryDistributionServeReturnedDataRequest) GetActionType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *EntryDistributionServeReturnedDataRequest) GetActionTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *EntryDistributionServeReturnedDataRequest) SetActionType(v int32) {
	o.ActionType = v
}

// GetId returns the Id field value
func (o *EntryDistributionServeReturnedDataRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntryDistributionServeReturnedDataRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntryDistributionServeReturnedDataRequest) SetId(v int32) {
	o.Id = v
}

func (o EntryDistributionServeReturnedDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actionType"] = o.ActionType
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEntryDistributionServeReturnedDataRequest struct {
	value *EntryDistributionServeReturnedDataRequest
	isSet bool
}

func (v NullableEntryDistributionServeReturnedDataRequest) Get() *EntryDistributionServeReturnedDataRequest {
	return v.value
}

func (v *NullableEntryDistributionServeReturnedDataRequest) Set(val *EntryDistributionServeReturnedDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryDistributionServeReturnedDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryDistributionServeReturnedDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryDistributionServeReturnedDataRequest(val *EntryDistributionServeReturnedDataRequest) *NullableEntryDistributionServeReturnedDataRequest {
	return &NullableEntryDistributionServeReturnedDataRequest{value: val, isSet: true}
}

func (v NullableEntryDistributionServeReturnedDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryDistributionServeReturnedDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

