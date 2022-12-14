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

// UserLoginRequest struct for UserLoginRequest
type UserLoginRequest struct {
	Expiry *int32 `json:"expiry,omitempty"`
	PartnerId int32 `json:"partnerId"`
	Password string `json:"password"`
	Privileges *string `json:"privileges,omitempty"`
	UserId string `json:"userId"`
}

// NewUserLoginRequest instantiates a new UserLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginRequest(partnerId int32, password string, userId string) *UserLoginRequest {
	this := UserLoginRequest{}
	this.PartnerId = partnerId
	this.Password = password
	var privileges string = "*"
	this.Privileges = &privileges
	this.UserId = userId
	return &this
}

// NewUserLoginRequestWithDefaults instantiates a new UserLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginRequestWithDefaults() *UserLoginRequest {
	this := UserLoginRequest{}
	var privileges string = "*"
	this.Privileges = &privileges
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *UserLoginRequest) GetExpiry() int32 {
	if o == nil || o.Expiry == nil {
		var ret int32
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginRequest) GetExpiryOk() (*int32, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *UserLoginRequest) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given int32 and assigns it to the Expiry field.
func (o *UserLoginRequest) SetExpiry(v int32) {
	o.Expiry = &v
}

// GetPartnerId returns the PartnerId field value
func (o *UserLoginRequest) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *UserLoginRequest) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *UserLoginRequest) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetPassword returns the Password field value
func (o *UserLoginRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserLoginRequest) SetPassword(v string) {
	o.Password = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *UserLoginRequest) GetPrivileges() string {
	if o == nil || o.Privileges == nil {
		var ret string
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginRequest) GetPrivilegesOk() (*string, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *UserLoginRequest) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given string and assigns it to the Privileges field.
func (o *UserLoginRequest) SetPrivileges(v string) {
	o.Privileges = &v
}

// GetUserId returns the UserId field value
func (o *UserLoginRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserLoginRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserLoginRequest) SetUserId(v string) {
	o.UserId = v
}

func (o UserLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["partnerId"] = o.PartnerId
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginRequest struct {
	value *UserLoginRequest
	isSet bool
}

func (v NullableUserLoginRequest) Get() *UserLoginRequest {
	return v.value
}

func (v *NullableUserLoginRequest) Set(val *UserLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginRequest(val *UserLoginRequest) *NullableUserLoginRequest {
	return &NullableUserLoginRequest{value: val, isSet: true}
}

func (v NullableUserLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


