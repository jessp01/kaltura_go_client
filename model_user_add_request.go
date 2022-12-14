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

// UserAddRequest struct for UserAddRequest
type UserAddRequest struct {
	User KalturaUser `json:"user"`
}

// NewUserAddRequest instantiates a new UserAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAddRequest(user KalturaUser) *UserAddRequest {
	this := UserAddRequest{}
	this.User = user
	return &this
}

// NewUserAddRequestWithDefaults instantiates a new UserAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAddRequestWithDefaults() *UserAddRequest {
	this := UserAddRequest{}
	return &this
}

// GetUser returns the User field value
func (o *UserAddRequest) GetUser() KalturaUser {
	if o == nil {
		var ret KalturaUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserAddRequest) GetUserOk() (*KalturaUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserAddRequest) SetUser(v KalturaUser) {
	o.User = v
}

func (o UserAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableUserAddRequest struct {
	value *UserAddRequest
	isSet bool
}

func (v NullableUserAddRequest) Get() *UserAddRequest {
	return v.value
}

func (v *NullableUserAddRequest) Set(val *UserAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAddRequest(val *UserAddRequest) *NullableUserAddRequest {
	return &NullableUserAddRequest{value: val, isSet: true}
}

func (v NullableUserAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


