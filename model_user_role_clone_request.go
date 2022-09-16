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

// UserRoleCloneRequest struct for UserRoleCloneRequest
type UserRoleCloneRequest struct {
	UserRoleId int32 `json:"userRoleId"`
}

// NewUserRoleCloneRequest instantiates a new UserRoleCloneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRoleCloneRequest(userRoleId int32) *UserRoleCloneRequest {
	this := UserRoleCloneRequest{}
	this.UserRoleId = userRoleId
	return &this
}

// NewUserRoleCloneRequestWithDefaults instantiates a new UserRoleCloneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRoleCloneRequestWithDefaults() *UserRoleCloneRequest {
	this := UserRoleCloneRequest{}
	return &this
}

// GetUserRoleId returns the UserRoleId field value
func (o *UserRoleCloneRequest) GetUserRoleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserRoleId
}

// GetUserRoleIdOk returns a tuple with the UserRoleId field value
// and a boolean to check if the value has been set.
func (o *UserRoleCloneRequest) GetUserRoleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserRoleId, true
}

// SetUserRoleId sets field value
func (o *UserRoleCloneRequest) SetUserRoleId(v int32) {
	o.UserRoleId = v
}

func (o UserRoleCloneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userRoleId"] = o.UserRoleId
	}
	return json.Marshal(toSerialize)
}

type NullableUserRoleCloneRequest struct {
	value *UserRoleCloneRequest
	isSet bool
}

func (v NullableUserRoleCloneRequest) Get() *UserRoleCloneRequest {
	return v.value
}

func (v *NullableUserRoleCloneRequest) Set(val *UserRoleCloneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRoleCloneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRoleCloneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRoleCloneRequest(val *UserRoleCloneRequest) *NullableUserRoleCloneRequest {
	return &NullableUserRoleCloneRequest{value: val, isSet: true}
}

func (v NullableUserRoleCloneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRoleCloneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


