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

// SessionGetRequest struct for SessionGetRequest
type SessionGetRequest struct {
	Session *string `json:"session,omitempty"`
}

// NewSessionGetRequest instantiates a new SessionGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionGetRequest() *SessionGetRequest {
	this := SessionGetRequest{}
	return &this
}

// NewSessionGetRequestWithDefaults instantiates a new SessionGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionGetRequestWithDefaults() *SessionGetRequest {
	this := SessionGetRequest{}
	return &this
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *SessionGetRequest) GetSession() string {
	if o == nil || o.Session == nil {
		var ret string
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionGetRequest) GetSessionOk() (*string, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *SessionGetRequest) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given string and assigns it to the Session field.
func (o *SessionGetRequest) SetSession(v string) {
	o.Session = &v
}

func (o SessionGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	return json.Marshal(toSerialize)
}

type NullableSessionGetRequest struct {
	value *SessionGetRequest
	isSet bool
}

func (v NullableSessionGetRequest) Get() *SessionGetRequest {
	return v.value
}

func (v *NullableSessionGetRequest) Set(val *SessionGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionGetRequest(val *SessionGetRequest) *NullableSessionGetRequest {
	return &NullableSessionGetRequest{value: val, isSet: true}
}

func (v NullableSessionGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


