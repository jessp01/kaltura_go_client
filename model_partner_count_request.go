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

// PartnerCountRequest struct for PartnerCountRequest
type PartnerCountRequest struct {
	Filter *KalturaPartnerFilter `json:"filter,omitempty"`
}

// NewPartnerCountRequest instantiates a new PartnerCountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCountRequest() *PartnerCountRequest {
	this := PartnerCountRequest{}
	return &this
}

// NewPartnerCountRequestWithDefaults instantiates a new PartnerCountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCountRequestWithDefaults() *PartnerCountRequest {
	this := PartnerCountRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PartnerCountRequest) GetFilter() KalturaPartnerFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaPartnerFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCountRequest) GetFilterOk() (*KalturaPartnerFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PartnerCountRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaPartnerFilter and assigns it to the Filter field.
func (o *PartnerCountRequest) SetFilter(v KalturaPartnerFilter) {
	o.Filter = &v
}

func (o PartnerCountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerCountRequest struct {
	value *PartnerCountRequest
	isSet bool
}

func (v NullablePartnerCountRequest) Get() *PartnerCountRequest {
	return v.value
}

func (v *NullablePartnerCountRequest) Set(val *PartnerCountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCountRequest(val *PartnerCountRequest) *NullablePartnerCountRequest {
	return &NullablePartnerCountRequest{value: val, isSet: true}
}

func (v NullablePartnerCountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


