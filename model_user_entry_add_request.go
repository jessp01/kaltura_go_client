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

// UserEntryAddRequest struct for UserEntryAddRequest
type UserEntryAddRequest struct {
	UserEntry KalturaUserEntry `json:"userEntry"`
}

// NewUserEntryAddRequest instantiates a new UserEntryAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEntryAddRequest(userEntry KalturaUserEntry) *UserEntryAddRequest {
	this := UserEntryAddRequest{}
	this.UserEntry = userEntry
	return &this
}

// NewUserEntryAddRequestWithDefaults instantiates a new UserEntryAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEntryAddRequestWithDefaults() *UserEntryAddRequest {
	this := UserEntryAddRequest{}
	return &this
}

// GetUserEntry returns the UserEntry field value
func (o *UserEntryAddRequest) GetUserEntry() KalturaUserEntry {
	if o == nil {
		var ret KalturaUserEntry
		return ret
	}

	return o.UserEntry
}

// GetUserEntryOk returns a tuple with the UserEntry field value
// and a boolean to check if the value has been set.
func (o *UserEntryAddRequest) GetUserEntryOk() (*KalturaUserEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEntry, true
}

// SetUserEntry sets field value
func (o *UserEntryAddRequest) SetUserEntry(v KalturaUserEntry) {
	o.UserEntry = v
}

func (o UserEntryAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userEntry"] = o.UserEntry
	}
	return json.Marshal(toSerialize)
}

type NullableUserEntryAddRequest struct {
	value *UserEntryAddRequest
	isSet bool
}

func (v NullableUserEntryAddRequest) Get() *UserEntryAddRequest {
	return v.value
}

func (v *NullableUserEntryAddRequest) Set(val *UserEntryAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEntryAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEntryAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEntryAddRequest(val *UserEntryAddRequest) *NullableUserEntryAddRequest {
	return &NullableUserEntryAddRequest{value: val, isSet: true}
}

func (v NullableUserEntryAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEntryAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


