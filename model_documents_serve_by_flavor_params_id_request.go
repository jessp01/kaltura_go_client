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

// DocumentsServeByFlavorParamsIdRequest struct for DocumentsServeByFlavorParamsIdRequest
type DocumentsServeByFlavorParamsIdRequest struct {
	EntryId string `json:"entryId"`
	FlavorParamsId *string `json:"flavorParamsId,omitempty"`
	ForceProxy *bool `json:"forceProxy,omitempty"`
}

// NewDocumentsServeByFlavorParamsIdRequest instantiates a new DocumentsServeByFlavorParamsIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentsServeByFlavorParamsIdRequest(entryId string) *DocumentsServeByFlavorParamsIdRequest {
	this := DocumentsServeByFlavorParamsIdRequest{}
	this.EntryId = entryId
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// NewDocumentsServeByFlavorParamsIdRequestWithDefaults instantiates a new DocumentsServeByFlavorParamsIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentsServeByFlavorParamsIdRequestWithDefaults() *DocumentsServeByFlavorParamsIdRequest {
	this := DocumentsServeByFlavorParamsIdRequest{}
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// GetEntryId returns the EntryId field value
func (o *DocumentsServeByFlavorParamsIdRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *DocumentsServeByFlavorParamsIdRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *DocumentsServeByFlavorParamsIdRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetFlavorParamsId returns the FlavorParamsId field value if set, zero value otherwise.
func (o *DocumentsServeByFlavorParamsIdRequest) GetFlavorParamsId() string {
	if o == nil || o.FlavorParamsId == nil {
		var ret string
		return ret
	}
	return *o.FlavorParamsId
}

// GetFlavorParamsIdOk returns a tuple with the FlavorParamsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsServeByFlavorParamsIdRequest) GetFlavorParamsIdOk() (*string, bool) {
	if o == nil || o.FlavorParamsId == nil {
		return nil, false
	}
	return o.FlavorParamsId, true
}

// HasFlavorParamsId returns a boolean if a field has been set.
func (o *DocumentsServeByFlavorParamsIdRequest) HasFlavorParamsId() bool {
	if o != nil && o.FlavorParamsId != nil {
		return true
	}

	return false
}

// SetFlavorParamsId gets a reference to the given string and assigns it to the FlavorParamsId field.
func (o *DocumentsServeByFlavorParamsIdRequest) SetFlavorParamsId(v string) {
	o.FlavorParamsId = &v
}

// GetForceProxy returns the ForceProxy field value if set, zero value otherwise.
func (o *DocumentsServeByFlavorParamsIdRequest) GetForceProxy() bool {
	if o == nil || o.ForceProxy == nil {
		var ret bool
		return ret
	}
	return *o.ForceProxy
}

// GetForceProxyOk returns a tuple with the ForceProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsServeByFlavorParamsIdRequest) GetForceProxyOk() (*bool, bool) {
	if o == nil || o.ForceProxy == nil {
		return nil, false
	}
	return o.ForceProxy, true
}

// HasForceProxy returns a boolean if a field has been set.
func (o *DocumentsServeByFlavorParamsIdRequest) HasForceProxy() bool {
	if o != nil && o.ForceProxy != nil {
		return true
	}

	return false
}

// SetForceProxy gets a reference to the given bool and assigns it to the ForceProxy field.
func (o *DocumentsServeByFlavorParamsIdRequest) SetForceProxy(v bool) {
	o.ForceProxy = &v
}

func (o DocumentsServeByFlavorParamsIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if o.FlavorParamsId != nil {
		toSerialize["flavorParamsId"] = o.FlavorParamsId
	}
	if o.ForceProxy != nil {
		toSerialize["forceProxy"] = o.ForceProxy
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentsServeByFlavorParamsIdRequest struct {
	value *DocumentsServeByFlavorParamsIdRequest
	isSet bool
}

func (v NullableDocumentsServeByFlavorParamsIdRequest) Get() *DocumentsServeByFlavorParamsIdRequest {
	return v.value
}

func (v *NullableDocumentsServeByFlavorParamsIdRequest) Set(val *DocumentsServeByFlavorParamsIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentsServeByFlavorParamsIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentsServeByFlavorParamsIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentsServeByFlavorParamsIdRequest(val *DocumentsServeByFlavorParamsIdRequest) *NullableDocumentsServeByFlavorParamsIdRequest {
	return &NullableDocumentsServeByFlavorParamsIdRequest{value: val, isSet: true}
}

func (v NullableDocumentsServeByFlavorParamsIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentsServeByFlavorParamsIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


