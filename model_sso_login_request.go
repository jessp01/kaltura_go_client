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

// SsoLoginRequest struct for SsoLoginRequest
type SsoLoginRequest struct {
	ApplicationType string `json:"applicationType"`
	PartnerId *int32 `json:"partnerId,omitempty"`
	UserId string `json:"userId"`
}

// NewSsoLoginRequest instantiates a new SsoLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoLoginRequest(applicationType string, userId string) *SsoLoginRequest {
	this := SsoLoginRequest{}
	this.ApplicationType = applicationType
	this.UserId = userId
	return &this
}

// NewSsoLoginRequestWithDefaults instantiates a new SsoLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoLoginRequestWithDefaults() *SsoLoginRequest {
	this := SsoLoginRequest{}
	return &this
}

// GetApplicationType returns the ApplicationType field value
func (o *SsoLoginRequest) GetApplicationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value
// and a boolean to check if the value has been set.
func (o *SsoLoginRequest) GetApplicationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationType, true
}

// SetApplicationType sets field value
func (o *SsoLoginRequest) SetApplicationType(v string) {
	o.ApplicationType = v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *SsoLoginRequest) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoLoginRequest) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *SsoLoginRequest) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *SsoLoginRequest) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetUserId returns the UserId field value
func (o *SsoLoginRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SsoLoginRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SsoLoginRequest) SetUserId(v string) {
	o.UserId = v
}

func (o SsoLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applicationType"] = o.ApplicationType
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableSsoLoginRequest struct {
	value *SsoLoginRequest
	isSet bool
}

func (v NullableSsoLoginRequest) Get() *SsoLoginRequest {
	return v.value
}

func (v *NullableSsoLoginRequest) Set(val *SsoLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoLoginRequest(val *SsoLoginRequest) *NullableSsoLoginRequest {
	return &NullableSsoLoginRequest{value: val, isSet: true}
}

func (v NullableSsoLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


