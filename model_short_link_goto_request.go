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

// ShortLinkGotoRequest struct for ShortLinkGotoRequest
type ShortLinkGotoRequest struct {
	Id string `json:"id"`
	Proxy *bool `json:"proxy,omitempty"`
}

// NewShortLinkGotoRequest instantiates a new ShortLinkGotoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShortLinkGotoRequest(id string) *ShortLinkGotoRequest {
	this := ShortLinkGotoRequest{}
	this.Id = id
	var proxy bool = false
	this.Proxy = &proxy
	return &this
}

// NewShortLinkGotoRequestWithDefaults instantiates a new ShortLinkGotoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortLinkGotoRequestWithDefaults() *ShortLinkGotoRequest {
	this := ShortLinkGotoRequest{}
	var proxy bool = false
	this.Proxy = &proxy
	return &this
}

// GetId returns the Id field value
func (o *ShortLinkGotoRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShortLinkGotoRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShortLinkGotoRequest) SetId(v string) {
	o.Id = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *ShortLinkGotoRequest) GetProxy() bool {
	if o == nil || o.Proxy == nil {
		var ret bool
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortLinkGotoRequest) GetProxyOk() (*bool, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *ShortLinkGotoRequest) HasProxy() bool {
	if o != nil && o.Proxy != nil {
		return true
	}

	return false
}

// SetProxy gets a reference to the given bool and assigns it to the Proxy field.
func (o *ShortLinkGotoRequest) SetProxy(v bool) {
	o.Proxy = &v
}

func (o ShortLinkGotoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	return json.Marshal(toSerialize)
}

type NullableShortLinkGotoRequest struct {
	value *ShortLinkGotoRequest
	isSet bool
}

func (v NullableShortLinkGotoRequest) Get() *ShortLinkGotoRequest {
	return v.value
}

func (v *NullableShortLinkGotoRequest) Set(val *ShortLinkGotoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableShortLinkGotoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableShortLinkGotoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShortLinkGotoRequest(val *ShortLinkGotoRequest) *NullableShortLinkGotoRequest {
	return &NullableShortLinkGotoRequest{value: val, isSet: true}
}

func (v NullableShortLinkGotoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShortLinkGotoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


