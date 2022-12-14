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

// EmailIngestionProfileGetByEmailAddressRequest struct for EmailIngestionProfileGetByEmailAddressRequest
type EmailIngestionProfileGetByEmailAddressRequest struct {
	EmailAddress string `json:"emailAddress"`
}

// NewEmailIngestionProfileGetByEmailAddressRequest instantiates a new EmailIngestionProfileGetByEmailAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailIngestionProfileGetByEmailAddressRequest(emailAddress string) *EmailIngestionProfileGetByEmailAddressRequest {
	this := EmailIngestionProfileGetByEmailAddressRequest{}
	this.EmailAddress = emailAddress
	return &this
}

// NewEmailIngestionProfileGetByEmailAddressRequestWithDefaults instantiates a new EmailIngestionProfileGetByEmailAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailIngestionProfileGetByEmailAddressRequestWithDefaults() *EmailIngestionProfileGetByEmailAddressRequest {
	this := EmailIngestionProfileGetByEmailAddressRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *EmailIngestionProfileGetByEmailAddressRequest) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *EmailIngestionProfileGetByEmailAddressRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *EmailIngestionProfileGetByEmailAddressRequest) SetEmailAddress(v string) {
	o.EmailAddress = v
}

func (o EmailIngestionProfileGetByEmailAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableEmailIngestionProfileGetByEmailAddressRequest struct {
	value *EmailIngestionProfileGetByEmailAddressRequest
	isSet bool
}

func (v NullableEmailIngestionProfileGetByEmailAddressRequest) Get() *EmailIngestionProfileGetByEmailAddressRequest {
	return v.value
}

func (v *NullableEmailIngestionProfileGetByEmailAddressRequest) Set(val *EmailIngestionProfileGetByEmailAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailIngestionProfileGetByEmailAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailIngestionProfileGetByEmailAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailIngestionProfileGetByEmailAddressRequest(val *EmailIngestionProfileGetByEmailAddressRequest) *NullableEmailIngestionProfileGetByEmailAddressRequest {
	return &NullableEmailIngestionProfileGetByEmailAddressRequest{value: val, isSet: true}
}

func (v NullableEmailIngestionProfileGetByEmailAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailIngestionProfileGetByEmailAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


