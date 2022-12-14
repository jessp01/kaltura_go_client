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

// SyndicationFeedUpdateRequest struct for SyndicationFeedUpdateRequest
type SyndicationFeedUpdateRequest struct {
	Id string `json:"id"`
	SyndicationFeed KalturaBaseSyndicationFeed `json:"syndicationFeed"`
}

// NewSyndicationFeedUpdateRequest instantiates a new SyndicationFeedUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyndicationFeedUpdateRequest(id string, syndicationFeed KalturaBaseSyndicationFeed) *SyndicationFeedUpdateRequest {
	this := SyndicationFeedUpdateRequest{}
	this.Id = id
	this.SyndicationFeed = syndicationFeed
	return &this
}

// NewSyndicationFeedUpdateRequestWithDefaults instantiates a new SyndicationFeedUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyndicationFeedUpdateRequestWithDefaults() *SyndicationFeedUpdateRequest {
	this := SyndicationFeedUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *SyndicationFeedUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyndicationFeedUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyndicationFeedUpdateRequest) SetId(v string) {
	o.Id = v
}

// GetSyndicationFeed returns the SyndicationFeed field value
func (o *SyndicationFeedUpdateRequest) GetSyndicationFeed() KalturaBaseSyndicationFeed {
	if o == nil {
		var ret KalturaBaseSyndicationFeed
		return ret
	}

	return o.SyndicationFeed
}

// GetSyndicationFeedOk returns a tuple with the SyndicationFeed field value
// and a boolean to check if the value has been set.
func (o *SyndicationFeedUpdateRequest) GetSyndicationFeedOk() (*KalturaBaseSyndicationFeed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyndicationFeed, true
}

// SetSyndicationFeed sets field value
func (o *SyndicationFeedUpdateRequest) SetSyndicationFeed(v KalturaBaseSyndicationFeed) {
	o.SyndicationFeed = v
}

func (o SyndicationFeedUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["syndicationFeed"] = o.SyndicationFeed
	}
	return json.Marshal(toSerialize)
}

type NullableSyndicationFeedUpdateRequest struct {
	value *SyndicationFeedUpdateRequest
	isSet bool
}

func (v NullableSyndicationFeedUpdateRequest) Get() *SyndicationFeedUpdateRequest {
	return v.value
}

func (v *NullableSyndicationFeedUpdateRequest) Set(val *SyndicationFeedUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyndicationFeedUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyndicationFeedUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyndicationFeedUpdateRequest(val *SyndicationFeedUpdateRequest) *NullableSyndicationFeedUpdateRequest {
	return &NullableSyndicationFeedUpdateRequest{value: val, isSet: true}
}

func (v NullableSyndicationFeedUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyndicationFeedUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


