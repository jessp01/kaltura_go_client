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

// KalturaEmailNotificationRecipientProvider `abstract`  Abstract core class  which provides the recipients (to, CC, BCC) for an email notification
type KalturaEmailNotificationRecipientProvider struct {
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaEmailNotificationRecipientProvider instantiates a new KalturaEmailNotificationRecipientProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEmailNotificationRecipientProvider() *KalturaEmailNotificationRecipientProvider {
	this := KalturaEmailNotificationRecipientProvider{}
	return &this
}

// NewKalturaEmailNotificationRecipientProviderWithDefaults instantiates a new KalturaEmailNotificationRecipientProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEmailNotificationRecipientProviderWithDefaults() *KalturaEmailNotificationRecipientProvider {
	this := KalturaEmailNotificationRecipientProvider{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaEmailNotificationRecipientProvider) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEmailNotificationRecipientProvider) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaEmailNotificationRecipientProvider) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaEmailNotificationRecipientProvider) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaEmailNotificationRecipientProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEmailNotificationRecipientProvider struct {
	value *KalturaEmailNotificationRecipientProvider
	isSet bool
}

func (v NullableKalturaEmailNotificationRecipientProvider) Get() *KalturaEmailNotificationRecipientProvider {
	return v.value
}

func (v *NullableKalturaEmailNotificationRecipientProvider) Set(val *KalturaEmailNotificationRecipientProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEmailNotificationRecipientProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEmailNotificationRecipientProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEmailNotificationRecipientProvider(val *KalturaEmailNotificationRecipientProvider) *NullableKalturaEmailNotificationRecipientProvider {
	return &NullableKalturaEmailNotificationRecipientProvider{value: val, isSet: true}
}

func (v NullableKalturaEmailNotificationRecipientProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEmailNotificationRecipientProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


