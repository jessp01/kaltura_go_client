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

// ZoomVendorLocalRegistrationPageRequest struct for ZoomVendorLocalRegistrationPageRequest
type ZoomVendorLocalRegistrationPageRequest struct {
	Jwt string `json:"jwt"`
}

// NewZoomVendorLocalRegistrationPageRequest instantiates a new ZoomVendorLocalRegistrationPageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoomVendorLocalRegistrationPageRequest(jwt string) *ZoomVendorLocalRegistrationPageRequest {
	this := ZoomVendorLocalRegistrationPageRequest{}
	this.Jwt = jwt
	return &this
}

// NewZoomVendorLocalRegistrationPageRequestWithDefaults instantiates a new ZoomVendorLocalRegistrationPageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoomVendorLocalRegistrationPageRequestWithDefaults() *ZoomVendorLocalRegistrationPageRequest {
	this := ZoomVendorLocalRegistrationPageRequest{}
	return &this
}

// GetJwt returns the Jwt field value
func (o *ZoomVendorLocalRegistrationPageRequest) GetJwt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value
// and a boolean to check if the value has been set.
func (o *ZoomVendorLocalRegistrationPageRequest) GetJwtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwt, true
}

// SetJwt sets field value
func (o *ZoomVendorLocalRegistrationPageRequest) SetJwt(v string) {
	o.Jwt = v
}

func (o ZoomVendorLocalRegistrationPageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jwt"] = o.Jwt
	}
	return json.Marshal(toSerialize)
}

type NullableZoomVendorLocalRegistrationPageRequest struct {
	value *ZoomVendorLocalRegistrationPageRequest
	isSet bool
}

func (v NullableZoomVendorLocalRegistrationPageRequest) Get() *ZoomVendorLocalRegistrationPageRequest {
	return v.value
}

func (v *NullableZoomVendorLocalRegistrationPageRequest) Set(val *ZoomVendorLocalRegistrationPageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableZoomVendorLocalRegistrationPageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableZoomVendorLocalRegistrationPageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoomVendorLocalRegistrationPageRequest(val *ZoomVendorLocalRegistrationPageRequest) *NullableZoomVendorLocalRegistrationPageRequest {
	return &NullableZoomVendorLocalRegistrationPageRequest{value: val, isSet: true}
}

func (v NullableZoomVendorLocalRegistrationPageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoomVendorLocalRegistrationPageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


