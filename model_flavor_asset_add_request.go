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

// FlavorAssetAddRequest struct for FlavorAssetAddRequest
type FlavorAssetAddRequest struct {
	EntryId string `json:"entryId"`
	FlavorAsset KalturaFlavorAsset `json:"flavorAsset"`
}

// NewFlavorAssetAddRequest instantiates a new FlavorAssetAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorAssetAddRequest(entryId string, flavorAsset KalturaFlavorAsset) *FlavorAssetAddRequest {
	this := FlavorAssetAddRequest{}
	this.EntryId = entryId
	this.FlavorAsset = flavorAsset
	return &this
}

// NewFlavorAssetAddRequestWithDefaults instantiates a new FlavorAssetAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorAssetAddRequestWithDefaults() *FlavorAssetAddRequest {
	this := FlavorAssetAddRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *FlavorAssetAddRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *FlavorAssetAddRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *FlavorAssetAddRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetFlavorAsset returns the FlavorAsset field value
func (o *FlavorAssetAddRequest) GetFlavorAsset() KalturaFlavorAsset {
	if o == nil {
		var ret KalturaFlavorAsset
		return ret
	}

	return o.FlavorAsset
}

// GetFlavorAssetOk returns a tuple with the FlavorAsset field value
// and a boolean to check if the value has been set.
func (o *FlavorAssetAddRequest) GetFlavorAssetOk() (*KalturaFlavorAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlavorAsset, true
}

// SetFlavorAsset sets field value
func (o *FlavorAssetAddRequest) SetFlavorAsset(v KalturaFlavorAsset) {
	o.FlavorAsset = v
}

func (o FlavorAssetAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["flavorAsset"] = o.FlavorAsset
	}
	return json.Marshal(toSerialize)
}

type NullableFlavorAssetAddRequest struct {
	value *FlavorAssetAddRequest
	isSet bool
}

func (v NullableFlavorAssetAddRequest) Get() *FlavorAssetAddRequest {
	return v.value
}

func (v *NullableFlavorAssetAddRequest) Set(val *FlavorAssetAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorAssetAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorAssetAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorAssetAddRequest(val *FlavorAssetAddRequest) *NullableFlavorAssetAddRequest {
	return &NullableFlavorAssetAddRequest{value: val, isSet: true}
}

func (v NullableFlavorAssetAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorAssetAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


