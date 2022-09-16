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

// KalturaPushNotificationParams Object which contains contextual entry-related data.
type KalturaPushNotificationParams struct {
	UserParams []KalturaPushEventNotificationParameter `json:"userParams,omitempty"`
}

// NewKalturaPushNotificationParams instantiates a new KalturaPushNotificationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPushNotificationParams() *KalturaPushNotificationParams {
	this := KalturaPushNotificationParams{}
	return &this
}

// NewKalturaPushNotificationParamsWithDefaults instantiates a new KalturaPushNotificationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPushNotificationParamsWithDefaults() *KalturaPushNotificationParams {
	this := KalturaPushNotificationParams{}
	return &this
}

// GetUserParams returns the UserParams field value if set, zero value otherwise.
func (o *KalturaPushNotificationParams) GetUserParams() []KalturaPushEventNotificationParameter {
	if o == nil || o.UserParams == nil {
		var ret []KalturaPushEventNotificationParameter
		return ret
	}
	return o.UserParams
}

// GetUserParamsOk returns a tuple with the UserParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPushNotificationParams) GetUserParamsOk() ([]KalturaPushEventNotificationParameter, bool) {
	if o == nil || o.UserParams == nil {
		return nil, false
	}
	return o.UserParams, true
}

// HasUserParams returns a boolean if a field has been set.
func (o *KalturaPushNotificationParams) HasUserParams() bool {
	if o != nil && o.UserParams != nil {
		return true
	}

	return false
}

// SetUserParams gets a reference to the given []KalturaPushEventNotificationParameter and assigns it to the UserParams field.
func (o *KalturaPushNotificationParams) SetUserParams(v []KalturaPushEventNotificationParameter) {
	o.UserParams = v
}

func (o KalturaPushNotificationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserParams != nil {
		toSerialize["userParams"] = o.UserParams
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPushNotificationParams struct {
	value *KalturaPushNotificationParams
	isSet bool
}

func (v NullableKalturaPushNotificationParams) Get() *KalturaPushNotificationParams {
	return v.value
}

func (v *NullableKalturaPushNotificationParams) Set(val *KalturaPushNotificationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPushNotificationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPushNotificationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPushNotificationParams(val *KalturaPushNotificationParams) *NullableKalturaPushNotificationParams {
	return &NullableKalturaPushNotificationParams{value: val, isSet: true}
}

func (v NullableKalturaPushNotificationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPushNotificationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

