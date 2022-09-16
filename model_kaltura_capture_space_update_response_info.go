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

// KalturaCaptureSpaceUpdateResponseInfo struct for KalturaCaptureSpaceUpdateResponseInfo
type KalturaCaptureSpaceUpdateResponseInfo struct {
	Hash *KalturaCaptureSpaceUpdateResponseInfoHash `json:"hash,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewKalturaCaptureSpaceUpdateResponseInfo instantiates a new KalturaCaptureSpaceUpdateResponseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCaptureSpaceUpdateResponseInfo() *KalturaCaptureSpaceUpdateResponseInfo {
	this := KalturaCaptureSpaceUpdateResponseInfo{}
	return &this
}

// NewKalturaCaptureSpaceUpdateResponseInfoWithDefaults instantiates a new KalturaCaptureSpaceUpdateResponseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCaptureSpaceUpdateResponseInfoWithDefaults() *KalturaCaptureSpaceUpdateResponseInfo {
	this := KalturaCaptureSpaceUpdateResponseInfo{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *KalturaCaptureSpaceUpdateResponseInfo) GetHash() KalturaCaptureSpaceUpdateResponseInfoHash {
	if o == nil || o.Hash == nil {
		var ret KalturaCaptureSpaceUpdateResponseInfoHash
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCaptureSpaceUpdateResponseInfo) GetHashOk() (*KalturaCaptureSpaceUpdateResponseInfoHash, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *KalturaCaptureSpaceUpdateResponseInfo) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given KalturaCaptureSpaceUpdateResponseInfoHash and assigns it to the Hash field.
func (o *KalturaCaptureSpaceUpdateResponseInfo) SetHash(v KalturaCaptureSpaceUpdateResponseInfoHash) {
	o.Hash = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *KalturaCaptureSpaceUpdateResponseInfo) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCaptureSpaceUpdateResponseInfo) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *KalturaCaptureSpaceUpdateResponseInfo) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *KalturaCaptureSpaceUpdateResponseInfo) SetUrl(v string) {
	o.Url = &v
}

func (o KalturaCaptureSpaceUpdateResponseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCaptureSpaceUpdateResponseInfo struct {
	value *KalturaCaptureSpaceUpdateResponseInfo
	isSet bool
}

func (v NullableKalturaCaptureSpaceUpdateResponseInfo) Get() *KalturaCaptureSpaceUpdateResponseInfo {
	return v.value
}

func (v *NullableKalturaCaptureSpaceUpdateResponseInfo) Set(val *KalturaCaptureSpaceUpdateResponseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCaptureSpaceUpdateResponseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCaptureSpaceUpdateResponseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCaptureSpaceUpdateResponseInfo(val *KalturaCaptureSpaceUpdateResponseInfo) *NullableKalturaCaptureSpaceUpdateResponseInfo {
	return &NullableKalturaCaptureSpaceUpdateResponseInfo{value: val, isSet: true}
}

func (v NullableKalturaCaptureSpaceUpdateResponseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCaptureSpaceUpdateResponseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

