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

// DocumentsAddFromFlavorAssetRequest struct for DocumentsAddFromFlavorAssetRequest
type DocumentsAddFromFlavorAssetRequest struct {
	DocumentEntry *KalturaDocumentEntry `json:"documentEntry,omitempty"`
	SourceFlavorAssetId string `json:"sourceFlavorAssetId"`
}

// NewDocumentsAddFromFlavorAssetRequest instantiates a new DocumentsAddFromFlavorAssetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentsAddFromFlavorAssetRequest(sourceFlavorAssetId string) *DocumentsAddFromFlavorAssetRequest {
	this := DocumentsAddFromFlavorAssetRequest{}
	this.SourceFlavorAssetId = sourceFlavorAssetId
	return &this
}

// NewDocumentsAddFromFlavorAssetRequestWithDefaults instantiates a new DocumentsAddFromFlavorAssetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentsAddFromFlavorAssetRequestWithDefaults() *DocumentsAddFromFlavorAssetRequest {
	this := DocumentsAddFromFlavorAssetRequest{}
	return &this
}

// GetDocumentEntry returns the DocumentEntry field value if set, zero value otherwise.
func (o *DocumentsAddFromFlavorAssetRequest) GetDocumentEntry() KalturaDocumentEntry {
	if o == nil || o.DocumentEntry == nil {
		var ret KalturaDocumentEntry
		return ret
	}
	return *o.DocumentEntry
}

// GetDocumentEntryOk returns a tuple with the DocumentEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsAddFromFlavorAssetRequest) GetDocumentEntryOk() (*KalturaDocumentEntry, bool) {
	if o == nil || o.DocumentEntry == nil {
		return nil, false
	}
	return o.DocumentEntry, true
}

// HasDocumentEntry returns a boolean if a field has been set.
func (o *DocumentsAddFromFlavorAssetRequest) HasDocumentEntry() bool {
	if o != nil && o.DocumentEntry != nil {
		return true
	}

	return false
}

// SetDocumentEntry gets a reference to the given KalturaDocumentEntry and assigns it to the DocumentEntry field.
func (o *DocumentsAddFromFlavorAssetRequest) SetDocumentEntry(v KalturaDocumentEntry) {
	o.DocumentEntry = &v
}

// GetSourceFlavorAssetId returns the SourceFlavorAssetId field value
func (o *DocumentsAddFromFlavorAssetRequest) GetSourceFlavorAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceFlavorAssetId
}

// GetSourceFlavorAssetIdOk returns a tuple with the SourceFlavorAssetId field value
// and a boolean to check if the value has been set.
func (o *DocumentsAddFromFlavorAssetRequest) GetSourceFlavorAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceFlavorAssetId, true
}

// SetSourceFlavorAssetId sets field value
func (o *DocumentsAddFromFlavorAssetRequest) SetSourceFlavorAssetId(v string) {
	o.SourceFlavorAssetId = v
}

func (o DocumentsAddFromFlavorAssetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentEntry != nil {
		toSerialize["documentEntry"] = o.DocumentEntry
	}
	if true {
		toSerialize["sourceFlavorAssetId"] = o.SourceFlavorAssetId
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentsAddFromFlavorAssetRequest struct {
	value *DocumentsAddFromFlavorAssetRequest
	isSet bool
}

func (v NullableDocumentsAddFromFlavorAssetRequest) Get() *DocumentsAddFromFlavorAssetRequest {
	return v.value
}

func (v *NullableDocumentsAddFromFlavorAssetRequest) Set(val *DocumentsAddFromFlavorAssetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentsAddFromFlavorAssetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentsAddFromFlavorAssetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentsAddFromFlavorAssetRequest(val *DocumentsAddFromFlavorAssetRequest) *NullableDocumentsAddFromFlavorAssetRequest {
	return &NullableDocumentsAddFromFlavorAssetRequest{value: val, isSet: true}
}

func (v NullableDocumentsAddFromFlavorAssetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentsAddFromFlavorAssetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


