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

// UserLoginByLoginIdRequest struct for UserLoginByLoginIdRequest
type UserLoginByLoginIdRequest struct {
	Expiry *int32 `json:"expiry,omitempty"`
	LoginId string `json:"loginId"`
	Otp *string `json:"otp,omitempty"`
	PartnerId *int32 `json:"partnerId,omitempty"`
	Password string `json:"password"`
	Privileges *string `json:"privileges,omitempty"`
}

// NewUserLoginByLoginIdRequest instantiates a new UserLoginByLoginIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginByLoginIdRequest(loginId string, password string) *UserLoginByLoginIdRequest {
	this := UserLoginByLoginIdRequest{}
	this.LoginId = loginId
	this.Password = password
	var privileges string = "*"
	this.Privileges = &privileges
	return &this
}

// NewUserLoginByLoginIdRequestWithDefaults instantiates a new UserLoginByLoginIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginByLoginIdRequestWithDefaults() *UserLoginByLoginIdRequest {
	this := UserLoginByLoginIdRequest{}
	var privileges string = "*"
	this.Privileges = &privileges
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *UserLoginByLoginIdRequest) GetExpiry() int32 {
	if o == nil || o.Expiry == nil {
		var ret int32
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginByLoginIdRequest) GetExpiryOk() (*int32, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *UserLoginByLoginIdRequest) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given int32 and assigns it to the Expiry field.
func (o *UserLoginByLoginIdRequest) SetExpiry(v int32) {
	o.Expiry = &v
}

// GetLoginId returns the LoginId field value
func (o *UserLoginByLoginIdRequest) GetLoginId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginId
}

// GetLoginIdOk returns a tuple with the LoginId field value
// and a boolean to check if the value has been set.
func (o *UserLoginByLoginIdRequest) GetLoginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginId, true
}

// SetLoginId sets field value
func (o *UserLoginByLoginIdRequest) SetLoginId(v string) {
	o.LoginId = v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *UserLoginByLoginIdRequest) GetOtp() string {
	if o == nil || o.Otp == nil {
		var ret string
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginByLoginIdRequest) GetOtpOk() (*string, bool) {
	if o == nil || o.Otp == nil {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *UserLoginByLoginIdRequest) HasOtp() bool {
	if o != nil && o.Otp != nil {
		return true
	}

	return false
}

// SetOtp gets a reference to the given string and assigns it to the Otp field.
func (o *UserLoginByLoginIdRequest) SetOtp(v string) {
	o.Otp = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *UserLoginByLoginIdRequest) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginByLoginIdRequest) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *UserLoginByLoginIdRequest) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *UserLoginByLoginIdRequest) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPassword returns the Password field value
func (o *UserLoginByLoginIdRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserLoginByLoginIdRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserLoginByLoginIdRequest) SetPassword(v string) {
	o.Password = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *UserLoginByLoginIdRequest) GetPrivileges() string {
	if o == nil || o.Privileges == nil {
		var ret string
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginByLoginIdRequest) GetPrivilegesOk() (*string, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *UserLoginByLoginIdRequest) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given string and assigns it to the Privileges field.
func (o *UserLoginByLoginIdRequest) SetPrivileges(v string) {
	o.Privileges = &v
}

func (o UserLoginByLoginIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["loginId"] = o.LoginId
	}
	if o.Otp != nil {
		toSerialize["otp"] = o.Otp
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginByLoginIdRequest struct {
	value *UserLoginByLoginIdRequest
	isSet bool
}

func (v NullableUserLoginByLoginIdRequest) Get() *UserLoginByLoginIdRequest {
	return v.value
}

func (v *NullableUserLoginByLoginIdRequest) Set(val *UserLoginByLoginIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginByLoginIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginByLoginIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginByLoginIdRequest(val *UserLoginByLoginIdRequest) *NullableUserLoginByLoginIdRequest {
	return &NullableUserLoginByLoginIdRequest{value: val, isSet: true}
}

func (v NullableUserLoginByLoginIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginByLoginIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


