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

// LiveConversionProfileServeRequest struct for LiveConversionProfileServeRequest
type LiveConversionProfileServeRequest struct {
	ExtraParams *string `json:"extraParams,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	StreamName string `json:"streamName"`
}

// NewLiveConversionProfileServeRequest instantiates a new LiveConversionProfileServeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveConversionProfileServeRequest(streamName string) *LiveConversionProfileServeRequest {
	this := LiveConversionProfileServeRequest{}
	this.StreamName = streamName
	return &this
}

// NewLiveConversionProfileServeRequestWithDefaults instantiates a new LiveConversionProfileServeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveConversionProfileServeRequestWithDefaults() *LiveConversionProfileServeRequest {
	this := LiveConversionProfileServeRequest{}
	return &this
}

// GetExtraParams returns the ExtraParams field value if set, zero value otherwise.
func (o *LiveConversionProfileServeRequest) GetExtraParams() string {
	if o == nil || o.ExtraParams == nil {
		var ret string
		return ret
	}
	return *o.ExtraParams
}

// GetExtraParamsOk returns a tuple with the ExtraParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveConversionProfileServeRequest) GetExtraParamsOk() (*string, bool) {
	if o == nil || o.ExtraParams == nil {
		return nil, false
	}
	return o.ExtraParams, true
}

// HasExtraParams returns a boolean if a field has been set.
func (o *LiveConversionProfileServeRequest) HasExtraParams() bool {
	if o != nil && o.ExtraParams != nil {
		return true
	}

	return false
}

// SetExtraParams gets a reference to the given string and assigns it to the ExtraParams field.
func (o *LiveConversionProfileServeRequest) SetExtraParams(v string) {
	o.ExtraParams = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *LiveConversionProfileServeRequest) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveConversionProfileServeRequest) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *LiveConversionProfileServeRequest) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *LiveConversionProfileServeRequest) SetHostname(v string) {
	o.Hostname = &v
}

// GetStreamName returns the StreamName field value
func (o *LiveConversionProfileServeRequest) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *LiveConversionProfileServeRequest) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *LiveConversionProfileServeRequest) SetStreamName(v string) {
	o.StreamName = v
}

func (o LiveConversionProfileServeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraParams != nil {
		toSerialize["extraParams"] = o.ExtraParams
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["streamName"] = o.StreamName
	}
	return json.Marshal(toSerialize)
}

type NullableLiveConversionProfileServeRequest struct {
	value *LiveConversionProfileServeRequest
	isSet bool
}

func (v NullableLiveConversionProfileServeRequest) Get() *LiveConversionProfileServeRequest {
	return v.value
}

func (v *NullableLiveConversionProfileServeRequest) Set(val *LiveConversionProfileServeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveConversionProfileServeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveConversionProfileServeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveConversionProfileServeRequest(val *LiveConversionProfileServeRequest) *NullableLiveConversionProfileServeRequest {
	return &NullableLiveConversionProfileServeRequest{value: val, isSet: true}
}

func (v NullableLiveConversionProfileServeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveConversionProfileServeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


