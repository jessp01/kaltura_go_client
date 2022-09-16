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

// KalturaUrlRecognizer struct for KalturaUrlRecognizer
type KalturaUrlRecognizer struct {
	// The hosts that are recognized
	Hosts *string `json:"hosts,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// The URI prefix we use for security
	UriPrefix *string `json:"uriPrefix,omitempty"`
}

// NewKalturaUrlRecognizer instantiates a new KalturaUrlRecognizer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUrlRecognizer() *KalturaUrlRecognizer {
	this := KalturaUrlRecognizer{}
	return &this
}

// NewKalturaUrlRecognizerWithDefaults instantiates a new KalturaUrlRecognizer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUrlRecognizerWithDefaults() *KalturaUrlRecognizer {
	this := KalturaUrlRecognizer{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *KalturaUrlRecognizer) GetHosts() string {
	if o == nil || o.Hosts == nil {
		var ret string
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUrlRecognizer) GetHostsOk() (*string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *KalturaUrlRecognizer) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given string and assigns it to the Hosts field.
func (o *KalturaUrlRecognizer) SetHosts(v string) {
	o.Hosts = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaUrlRecognizer) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUrlRecognizer) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaUrlRecognizer) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaUrlRecognizer) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetUriPrefix returns the UriPrefix field value if set, zero value otherwise.
func (o *KalturaUrlRecognizer) GetUriPrefix() string {
	if o == nil || o.UriPrefix == nil {
		var ret string
		return ret
	}
	return *o.UriPrefix
}

// GetUriPrefixOk returns a tuple with the UriPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUrlRecognizer) GetUriPrefixOk() (*string, bool) {
	if o == nil || o.UriPrefix == nil {
		return nil, false
	}
	return o.UriPrefix, true
}

// HasUriPrefix returns a boolean if a field has been set.
func (o *KalturaUrlRecognizer) HasUriPrefix() bool {
	if o != nil && o.UriPrefix != nil {
		return true
	}

	return false
}

// SetUriPrefix gets a reference to the given string and assigns it to the UriPrefix field.
func (o *KalturaUrlRecognizer) SetUriPrefix(v string) {
	o.UriPrefix = &v
}

func (o KalturaUrlRecognizer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.UriPrefix != nil {
		toSerialize["uriPrefix"] = o.UriPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUrlRecognizer struct {
	value *KalturaUrlRecognizer
	isSet bool
}

func (v NullableKalturaUrlRecognizer) Get() *KalturaUrlRecognizer {
	return v.value
}

func (v *NullableKalturaUrlRecognizer) Set(val *KalturaUrlRecognizer) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUrlRecognizer) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUrlRecognizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUrlRecognizer(val *KalturaUrlRecognizer) *NullableKalturaUrlRecognizer {
	return &NullableKalturaUrlRecognizer{value: val, isSet: true}
}

func (v NullableKalturaUrlRecognizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUrlRecognizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


