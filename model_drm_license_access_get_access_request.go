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

// DrmLicenseAccessGetAccessRequest struct for DrmLicenseAccessGetAccessRequest
type DrmLicenseAccessGetAccessRequest struct {
	EntryId string `json:"entryId"`
	FlavorIds string `json:"flavorIds"`
	Referrer string `json:"referrer"`
}

// NewDrmLicenseAccessGetAccessRequest instantiates a new DrmLicenseAccessGetAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrmLicenseAccessGetAccessRequest(entryId string, flavorIds string, referrer string) *DrmLicenseAccessGetAccessRequest {
	this := DrmLicenseAccessGetAccessRequest{}
	this.EntryId = entryId
	this.FlavorIds = flavorIds
	this.Referrer = referrer
	return &this
}

// NewDrmLicenseAccessGetAccessRequestWithDefaults instantiates a new DrmLicenseAccessGetAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrmLicenseAccessGetAccessRequestWithDefaults() *DrmLicenseAccessGetAccessRequest {
	this := DrmLicenseAccessGetAccessRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *DrmLicenseAccessGetAccessRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *DrmLicenseAccessGetAccessRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *DrmLicenseAccessGetAccessRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetFlavorIds returns the FlavorIds field value
func (o *DrmLicenseAccessGetAccessRequest) GetFlavorIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlavorIds
}

// GetFlavorIdsOk returns a tuple with the FlavorIds field value
// and a boolean to check if the value has been set.
func (o *DrmLicenseAccessGetAccessRequest) GetFlavorIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlavorIds, true
}

// SetFlavorIds sets field value
func (o *DrmLicenseAccessGetAccessRequest) SetFlavorIds(v string) {
	o.FlavorIds = v
}

// GetReferrer returns the Referrer field value
func (o *DrmLicenseAccessGetAccessRequest) GetReferrer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value
// and a boolean to check if the value has been set.
func (o *DrmLicenseAccessGetAccessRequest) GetReferrerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referrer, true
}

// SetReferrer sets field value
func (o *DrmLicenseAccessGetAccessRequest) SetReferrer(v string) {
	o.Referrer = v
}

func (o DrmLicenseAccessGetAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["flavorIds"] = o.FlavorIds
	}
	if true {
		toSerialize["referrer"] = o.Referrer
	}
	return json.Marshal(toSerialize)
}

type NullableDrmLicenseAccessGetAccessRequest struct {
	value *DrmLicenseAccessGetAccessRequest
	isSet bool
}

func (v NullableDrmLicenseAccessGetAccessRequest) Get() *DrmLicenseAccessGetAccessRequest {
	return v.value
}

func (v *NullableDrmLicenseAccessGetAccessRequest) Set(val *DrmLicenseAccessGetAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDrmLicenseAccessGetAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDrmLicenseAccessGetAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrmLicenseAccessGetAccessRequest(val *DrmLicenseAccessGetAccessRequest) *NullableDrmLicenseAccessGetAccessRequest {
	return &NullableDrmLicenseAccessGetAccessRequest{value: val, isSet: true}
}

func (v NullableDrmLicenseAccessGetAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrmLicenseAccessGetAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

