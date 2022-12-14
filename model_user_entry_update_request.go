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

// UserEntryUpdateRequest struct for UserEntryUpdateRequest
type UserEntryUpdateRequest struct {
	Id int32 `json:"id"`
	UserEntry KalturaUserEntry `json:"userEntry"`
}

// NewUserEntryUpdateRequest instantiates a new UserEntryUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEntryUpdateRequest(id int32, userEntry KalturaUserEntry) *UserEntryUpdateRequest {
	this := UserEntryUpdateRequest{}
	this.Id = id
	this.UserEntry = userEntry
	return &this
}

// NewUserEntryUpdateRequestWithDefaults instantiates a new UserEntryUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEntryUpdateRequestWithDefaults() *UserEntryUpdateRequest {
	this := UserEntryUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UserEntryUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserEntryUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserEntryUpdateRequest) SetId(v int32) {
	o.Id = v
}

// GetUserEntry returns the UserEntry field value
func (o *UserEntryUpdateRequest) GetUserEntry() KalturaUserEntry {
	if o == nil {
		var ret KalturaUserEntry
		return ret
	}

	return o.UserEntry
}

// GetUserEntryOk returns a tuple with the UserEntry field value
// and a boolean to check if the value has been set.
func (o *UserEntryUpdateRequest) GetUserEntryOk() (*KalturaUserEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEntry, true
}

// SetUserEntry sets field value
func (o *UserEntryUpdateRequest) SetUserEntry(v KalturaUserEntry) {
	o.UserEntry = v
}

func (o UserEntryUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["userEntry"] = o.UserEntry
	}
	return json.Marshal(toSerialize)
}

type NullableUserEntryUpdateRequest struct {
	value *UserEntryUpdateRequest
	isSet bool
}

func (v NullableUserEntryUpdateRequest) Get() *UserEntryUpdateRequest {
	return v.value
}

func (v *NullableUserEntryUpdateRequest) Set(val *UserEntryUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEntryUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEntryUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEntryUpdateRequest(val *UserEntryUpdateRequest) *NullableUserEntryUpdateRequest {
	return &NullableUserEntryUpdateRequest{value: val, isSet: true}
}

func (v NullableUserEntryUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEntryUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


