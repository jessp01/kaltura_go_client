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

// UserUpdateRequest struct for UserUpdateRequest
type UserUpdateRequest struct {
	User KalturaUser `json:"user"`
	UserId string `json:"userId"`
}

// NewUserUpdateRequest instantiates a new UserUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateRequest(user KalturaUser, userId string) *UserUpdateRequest {
	this := UserUpdateRequest{}
	this.User = user
	this.UserId = userId
	return &this
}

// NewUserUpdateRequestWithDefaults instantiates a new UserUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateRequestWithDefaults() *UserUpdateRequest {
	this := UserUpdateRequest{}
	return &this
}

// GetUser returns the User field value
func (o *UserUpdateRequest) GetUser() KalturaUser {
	if o == nil {
		var ret KalturaUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserUpdateRequest) GetUserOk() (*KalturaUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserUpdateRequest) SetUser(v KalturaUser) {
	o.User = v
}

// GetUserId returns the UserId field value
func (o *UserUpdateRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserUpdateRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserUpdateRequest) SetUserId(v string) {
	o.UserId = v
}

func (o UserUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableUserUpdateRequest struct {
	value *UserUpdateRequest
	isSet bool
}

func (v NullableUserUpdateRequest) Get() *UserUpdateRequest {
	return v.value
}

func (v *NullableUserUpdateRequest) Set(val *UserUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateRequest(val *UserUpdateRequest) *NullableUserUpdateRequest {
	return &NullableUserUpdateRequest{value: val, isSet: true}
}

func (v NullableUserUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


