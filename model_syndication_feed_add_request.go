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

// SyndicationFeedAddRequest struct for SyndicationFeedAddRequest
type SyndicationFeedAddRequest struct {
	SyndicationFeed KalturaBaseSyndicationFeed `json:"syndicationFeed"`
}

// NewSyndicationFeedAddRequest instantiates a new SyndicationFeedAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyndicationFeedAddRequest(syndicationFeed KalturaBaseSyndicationFeed) *SyndicationFeedAddRequest {
	this := SyndicationFeedAddRequest{}
	this.SyndicationFeed = syndicationFeed
	return &this
}

// NewSyndicationFeedAddRequestWithDefaults instantiates a new SyndicationFeedAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyndicationFeedAddRequestWithDefaults() *SyndicationFeedAddRequest {
	this := SyndicationFeedAddRequest{}
	return &this
}

// GetSyndicationFeed returns the SyndicationFeed field value
func (o *SyndicationFeedAddRequest) GetSyndicationFeed() KalturaBaseSyndicationFeed {
	if o == nil {
		var ret KalturaBaseSyndicationFeed
		return ret
	}

	return o.SyndicationFeed
}

// GetSyndicationFeedOk returns a tuple with the SyndicationFeed field value
// and a boolean to check if the value has been set.
func (o *SyndicationFeedAddRequest) GetSyndicationFeedOk() (*KalturaBaseSyndicationFeed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyndicationFeed, true
}

// SetSyndicationFeed sets field value
func (o *SyndicationFeedAddRequest) SetSyndicationFeed(v KalturaBaseSyndicationFeed) {
	o.SyndicationFeed = v
}

func (o SyndicationFeedAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["syndicationFeed"] = o.SyndicationFeed
	}
	return json.Marshal(toSerialize)
}

type NullableSyndicationFeedAddRequest struct {
	value *SyndicationFeedAddRequest
	isSet bool
}

func (v NullableSyndicationFeedAddRequest) Get() *SyndicationFeedAddRequest {
	return v.value
}

func (v *NullableSyndicationFeedAddRequest) Set(val *SyndicationFeedAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyndicationFeedAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyndicationFeedAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyndicationFeedAddRequest(val *SyndicationFeedAddRequest) *NullableSyndicationFeedAddRequest {
	return &NullableSyndicationFeedAddRequest{value: val, isSet: true}
}

func (v NullableSyndicationFeedAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyndicationFeedAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

