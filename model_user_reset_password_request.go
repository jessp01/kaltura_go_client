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

// UserResetPasswordRequest struct for UserResetPasswordRequest
type UserResetPasswordRequest struct {
	Email string `json:"email"`
	LinkType *string `json:"linkType,omitempty"`
}

// NewUserResetPasswordRequest instantiates a new UserResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResetPasswordRequest(email string) *UserResetPasswordRequest {
	this := UserResetPasswordRequest{}
	this.Email = email
	return &this
}

// NewUserResetPasswordRequestWithDefaults instantiates a new UserResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResetPasswordRequestWithDefaults() *UserResetPasswordRequest {
	this := UserResetPasswordRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserResetPasswordRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserResetPasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserResetPasswordRequest) SetEmail(v string) {
	o.Email = v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *UserResetPasswordRequest) GetLinkType() string {
	if o == nil || o.LinkType == nil {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResetPasswordRequest) GetLinkTypeOk() (*string, bool) {
	if o == nil || o.LinkType == nil {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *UserResetPasswordRequest) HasLinkType() bool {
	if o != nil && o.LinkType != nil {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *UserResetPasswordRequest) SetLinkType(v string) {
	o.LinkType = &v
}

func (o UserResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.LinkType != nil {
		toSerialize["linkType"] = o.LinkType
	}
	return json.Marshal(toSerialize)
}

type NullableUserResetPasswordRequest struct {
	value *UserResetPasswordRequest
	isSet bool
}

func (v NullableUserResetPasswordRequest) Get() *UserResetPasswordRequest {
	return v.value
}

func (v *NullableUserResetPasswordRequest) Set(val *UserResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResetPasswordRequest(val *UserResetPasswordRequest) *NullableUserResetPasswordRequest {
	return &NullableUserResetPasswordRequest{value: val, isSet: true}
}

func (v NullableUserResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


