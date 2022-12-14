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

// DocumentsServeRequest struct for DocumentsServeRequest
type DocumentsServeRequest struct {
	EntryId string `json:"entryId"`
	FlavorAssetId *string `json:"flavorAssetId,omitempty"`
	ForceProxy *bool `json:"forceProxy,omitempty"`
}

// NewDocumentsServeRequest instantiates a new DocumentsServeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentsServeRequest(entryId string) *DocumentsServeRequest {
	this := DocumentsServeRequest{}
	this.EntryId = entryId
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// NewDocumentsServeRequestWithDefaults instantiates a new DocumentsServeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentsServeRequestWithDefaults() *DocumentsServeRequest {
	this := DocumentsServeRequest{}
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// GetEntryId returns the EntryId field value
func (o *DocumentsServeRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *DocumentsServeRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *DocumentsServeRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetFlavorAssetId returns the FlavorAssetId field value if set, zero value otherwise.
func (o *DocumentsServeRequest) GetFlavorAssetId() string {
	if o == nil || o.FlavorAssetId == nil {
		var ret string
		return ret
	}
	return *o.FlavorAssetId
}

// GetFlavorAssetIdOk returns a tuple with the FlavorAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsServeRequest) GetFlavorAssetIdOk() (*string, bool) {
	if o == nil || o.FlavorAssetId == nil {
		return nil, false
	}
	return o.FlavorAssetId, true
}

// HasFlavorAssetId returns a boolean if a field has been set.
func (o *DocumentsServeRequest) HasFlavorAssetId() bool {
	if o != nil && o.FlavorAssetId != nil {
		return true
	}

	return false
}

// SetFlavorAssetId gets a reference to the given string and assigns it to the FlavorAssetId field.
func (o *DocumentsServeRequest) SetFlavorAssetId(v string) {
	o.FlavorAssetId = &v
}

// GetForceProxy returns the ForceProxy field value if set, zero value otherwise.
func (o *DocumentsServeRequest) GetForceProxy() bool {
	if o == nil || o.ForceProxy == nil {
		var ret bool
		return ret
	}
	return *o.ForceProxy
}

// GetForceProxyOk returns a tuple with the ForceProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsServeRequest) GetForceProxyOk() (*bool, bool) {
	if o == nil || o.ForceProxy == nil {
		return nil, false
	}
	return o.ForceProxy, true
}

// HasForceProxy returns a boolean if a field has been set.
func (o *DocumentsServeRequest) HasForceProxy() bool {
	if o != nil && o.ForceProxy != nil {
		return true
	}

	return false
}

// SetForceProxy gets a reference to the given bool and assigns it to the ForceProxy field.
func (o *DocumentsServeRequest) SetForceProxy(v bool) {
	o.ForceProxy = &v
}

func (o DocumentsServeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if o.FlavorAssetId != nil {
		toSerialize["flavorAssetId"] = o.FlavorAssetId
	}
	if o.ForceProxy != nil {
		toSerialize["forceProxy"] = o.ForceProxy
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentsServeRequest struct {
	value *DocumentsServeRequest
	isSet bool
}

func (v NullableDocumentsServeRequest) Get() *DocumentsServeRequest {
	return v.value
}

func (v *NullableDocumentsServeRequest) Set(val *DocumentsServeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentsServeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentsServeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentsServeRequest(val *DocumentsServeRequest) *NullableDocumentsServeRequest {
	return &NullableDocumentsServeRequest{value: val, isSet: true}
}

func (v NullableDocumentsServeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentsServeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


