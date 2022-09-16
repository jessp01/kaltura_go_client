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

// KalturaSipResponse struct for KalturaSipResponse
type KalturaSipResponse struct {
	Action *string `json:"action,omitempty"`
	HostName *string `json:"hostName,omitempty"`
	Msg *string `json:"msg,omitempty"`
	SessionId *string `json:"sessionId,omitempty"`
}

// NewKalturaSipResponse instantiates a new KalturaSipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSipResponse() *KalturaSipResponse {
	this := KalturaSipResponse{}
	return &this
}

// NewKalturaSipResponseWithDefaults instantiates a new KalturaSipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSipResponseWithDefaults() *KalturaSipResponse {
	this := KalturaSipResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *KalturaSipResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSipResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *KalturaSipResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *KalturaSipResponse) SetAction(v string) {
	o.Action = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *KalturaSipResponse) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSipResponse) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *KalturaSipResponse) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *KalturaSipResponse) SetHostName(v string) {
	o.HostName = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *KalturaSipResponse) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSipResponse) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *KalturaSipResponse) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *KalturaSipResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *KalturaSipResponse) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSipResponse) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *KalturaSipResponse) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *KalturaSipResponse) SetSessionId(v string) {
	o.SessionId = &v
}

func (o KalturaSipResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSipResponse struct {
	value *KalturaSipResponse
	isSet bool
}

func (v NullableKalturaSipResponse) Get() *KalturaSipResponse {
	return v.value
}

func (v *NullableKalturaSipResponse) Set(val *KalturaSipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSipResponse(val *KalturaSipResponse) *NullableKalturaSipResponse {
	return &NullableKalturaSipResponse{value: val, isSet: true}
}

func (v NullableKalturaSipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


