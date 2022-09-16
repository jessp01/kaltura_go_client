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

// SystemPartnerGetAdminSessionRequest struct for SystemPartnerGetAdminSessionRequest
type SystemPartnerGetAdminSessionRequest struct {
	PId int32 `json:"pId"`
	UserId *string `json:"userId,omitempty"`
}

// NewSystemPartnerGetAdminSessionRequest instantiates a new SystemPartnerGetAdminSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPartnerGetAdminSessionRequest(pId int32) *SystemPartnerGetAdminSessionRequest {
	this := SystemPartnerGetAdminSessionRequest{}
	this.PId = pId
	return &this
}

// NewSystemPartnerGetAdminSessionRequestWithDefaults instantiates a new SystemPartnerGetAdminSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPartnerGetAdminSessionRequestWithDefaults() *SystemPartnerGetAdminSessionRequest {
	this := SystemPartnerGetAdminSessionRequest{}
	return &this
}

// GetPId returns the PId field value
func (o *SystemPartnerGetAdminSessionRequest) GetPId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PId
}

// GetPIdOk returns a tuple with the PId field value
// and a boolean to check if the value has been set.
func (o *SystemPartnerGetAdminSessionRequest) GetPIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PId, true
}

// SetPId sets field value
func (o *SystemPartnerGetAdminSessionRequest) SetPId(v int32) {
	o.PId = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SystemPartnerGetAdminSessionRequest) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPartnerGetAdminSessionRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SystemPartnerGetAdminSessionRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SystemPartnerGetAdminSessionRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o SystemPartnerGetAdminSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pId"] = o.PId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPartnerGetAdminSessionRequest struct {
	value *SystemPartnerGetAdminSessionRequest
	isSet bool
}

func (v NullableSystemPartnerGetAdminSessionRequest) Get() *SystemPartnerGetAdminSessionRequest {
	return v.value
}

func (v *NullableSystemPartnerGetAdminSessionRequest) Set(val *SystemPartnerGetAdminSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPartnerGetAdminSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPartnerGetAdminSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPartnerGetAdminSessionRequest(val *SystemPartnerGetAdminSessionRequest) *NullableSystemPartnerGetAdminSessionRequest {
	return &NullableSystemPartnerGetAdminSessionRequest{value: val, isSet: true}
}

func (v NullableSystemPartnerGetAdminSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPartnerGetAdminSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


