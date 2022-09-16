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

// KalturaContextDataResult struct for KalturaContextDataResult
type KalturaContextDataResult struct {
	Actions []KalturaRuleAction `json:"actions,omitempty"`
	Messages []KalturaString `json:"messages,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaContextDataResult instantiates a new KalturaContextDataResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaContextDataResult() *KalturaContextDataResult {
	this := KalturaContextDataResult{}
	return &this
}

// NewKalturaContextDataResultWithDefaults instantiates a new KalturaContextDataResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaContextDataResultWithDefaults() *KalturaContextDataResult {
	this := KalturaContextDataResult{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *KalturaContextDataResult) GetActions() []KalturaRuleAction {
	if o == nil || o.Actions == nil {
		var ret []KalturaRuleAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaContextDataResult) GetActionsOk() ([]KalturaRuleAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *KalturaContextDataResult) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []KalturaRuleAction and assigns it to the Actions field.
func (o *KalturaContextDataResult) SetActions(v []KalturaRuleAction) {
	o.Actions = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *KalturaContextDataResult) GetMessages() []KalturaString {
	if o == nil || o.Messages == nil {
		var ret []KalturaString
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaContextDataResult) GetMessagesOk() ([]KalturaString, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *KalturaContextDataResult) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []KalturaString and assigns it to the Messages field.
func (o *KalturaContextDataResult) SetMessages(v []KalturaString) {
	o.Messages = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaContextDataResult) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaContextDataResult) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaContextDataResult) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaContextDataResult) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaContextDataResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaContextDataResult struct {
	value *KalturaContextDataResult
	isSet bool
}

func (v NullableKalturaContextDataResult) Get() *KalturaContextDataResult {
	return v.value
}

func (v *NullableKalturaContextDataResult) Set(val *KalturaContextDataResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaContextDataResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaContextDataResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaContextDataResult(val *KalturaContextDataResult) *NullableKalturaContextDataResult {
	return &NullableKalturaContextDataResult{value: val, isSet: true}
}

func (v NullableKalturaContextDataResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaContextDataResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

