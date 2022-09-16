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

// KalturaAssetServeOptions struct for KalturaAssetServeOptions
type KalturaAssetServeOptions struct {
	Download *bool `json:"download,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	Referrer *string `json:"referrer,omitempty"`
}

// NewKalturaAssetServeOptions instantiates a new KalturaAssetServeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAssetServeOptions() *KalturaAssetServeOptions {
	this := KalturaAssetServeOptions{}
	return &this
}

// NewKalturaAssetServeOptionsWithDefaults instantiates a new KalturaAssetServeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetServeOptionsWithDefaults() *KalturaAssetServeOptions {
	this := KalturaAssetServeOptions{}
	return &this
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *KalturaAssetServeOptions) GetDownload() bool {
	if o == nil || o.Download == nil {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetServeOptions) GetDownloadOk() (*bool, bool) {
	if o == nil || o.Download == nil {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *KalturaAssetServeOptions) HasDownload() bool {
	if o != nil && o.Download != nil {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *KalturaAssetServeOptions) SetDownload(v bool) {
	o.Download = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaAssetServeOptions) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetServeOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaAssetServeOptions) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaAssetServeOptions) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetReferrer returns the Referrer field value if set, zero value otherwise.
func (o *KalturaAssetServeOptions) GetReferrer() string {
	if o == nil || o.Referrer == nil {
		var ret string
		return ret
	}
	return *o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetServeOptions) GetReferrerOk() (*string, bool) {
	if o == nil || o.Referrer == nil {
		return nil, false
	}
	return o.Referrer, true
}

// HasReferrer returns a boolean if a field has been set.
func (o *KalturaAssetServeOptions) HasReferrer() bool {
	if o != nil && o.Referrer != nil {
		return true
	}

	return false
}

// SetReferrer gets a reference to the given string and assigns it to the Referrer field.
func (o *KalturaAssetServeOptions) SetReferrer(v string) {
	o.Referrer = &v
}

func (o KalturaAssetServeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Download != nil {
		toSerialize["download"] = o.Download
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.Referrer != nil {
		toSerialize["referrer"] = o.Referrer
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAssetServeOptions struct {
	value *KalturaAssetServeOptions
	isSet bool
}

func (v NullableKalturaAssetServeOptions) Get() *KalturaAssetServeOptions {
	return v.value
}

func (v *NullableKalturaAssetServeOptions) Set(val *KalturaAssetServeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAssetServeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAssetServeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAssetServeOptions(val *KalturaAssetServeOptions) *NullableKalturaAssetServeOptions {
	return &NullableKalturaAssetServeOptions{value: val, isSet: true}
}

func (v NullableKalturaAssetServeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAssetServeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

