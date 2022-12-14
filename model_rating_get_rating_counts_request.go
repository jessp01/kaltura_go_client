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

// RatingGetRatingCountsRequest struct for RatingGetRatingCountsRequest
type RatingGetRatingCountsRequest struct {
	Filter KalturaRatingCountFilter `json:"filter"`
}

// NewRatingGetRatingCountsRequest instantiates a new RatingGetRatingCountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatingGetRatingCountsRequest(filter KalturaRatingCountFilter) *RatingGetRatingCountsRequest {
	this := RatingGetRatingCountsRequest{}
	this.Filter = filter
	return &this
}

// NewRatingGetRatingCountsRequestWithDefaults instantiates a new RatingGetRatingCountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingGetRatingCountsRequestWithDefaults() *RatingGetRatingCountsRequest {
	this := RatingGetRatingCountsRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *RatingGetRatingCountsRequest) GetFilter() KalturaRatingCountFilter {
	if o == nil {
		var ret KalturaRatingCountFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *RatingGetRatingCountsRequest) GetFilterOk() (*KalturaRatingCountFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *RatingGetRatingCountsRequest) SetFilter(v KalturaRatingCountFilter) {
	o.Filter = v
}

func (o RatingGetRatingCountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableRatingGetRatingCountsRequest struct {
	value *RatingGetRatingCountsRequest
	isSet bool
}

func (v NullableRatingGetRatingCountsRequest) Get() *RatingGetRatingCountsRequest {
	return v.value
}

func (v *NullableRatingGetRatingCountsRequest) Set(val *RatingGetRatingCountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingGetRatingCountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingGetRatingCountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingGetRatingCountsRequest(val *RatingGetRatingCountsRequest) *NullableRatingGetRatingCountsRequest {
	return &NullableRatingGetRatingCountsRequest{value: val, isSet: true}
}

func (v NullableRatingGetRatingCountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingGetRatingCountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


