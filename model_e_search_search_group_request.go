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

// ESearchSearchGroupRequest struct for ESearchSearchGroupRequest
type ESearchSearchGroupRequest struct {
	Pager *KalturaPager `json:"pager,omitempty"`
	SearchParams KalturaESearchGroupParams `json:"searchParams"`
}

// NewESearchSearchGroupRequest instantiates a new ESearchSearchGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewESearchSearchGroupRequest(searchParams KalturaESearchGroupParams) *ESearchSearchGroupRequest {
	this := ESearchSearchGroupRequest{}
	this.SearchParams = searchParams
	return &this
}

// NewESearchSearchGroupRequestWithDefaults instantiates a new ESearchSearchGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewESearchSearchGroupRequestWithDefaults() *ESearchSearchGroupRequest {
	this := ESearchSearchGroupRequest{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *ESearchSearchGroupRequest) GetPager() KalturaPager {
	if o == nil || o.Pager == nil {
		var ret KalturaPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESearchSearchGroupRequest) GetPagerOk() (*KalturaPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *ESearchSearchGroupRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaPager and assigns it to the Pager field.
func (o *ESearchSearchGroupRequest) SetPager(v KalturaPager) {
	o.Pager = &v
}

// GetSearchParams returns the SearchParams field value
func (o *ESearchSearchGroupRequest) GetSearchParams() KalturaESearchGroupParams {
	if o == nil {
		var ret KalturaESearchGroupParams
		return ret
	}

	return o.SearchParams
}

// GetSearchParamsOk returns a tuple with the SearchParams field value
// and a boolean to check if the value has been set.
func (o *ESearchSearchGroupRequest) GetSearchParamsOk() (*KalturaESearchGroupParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchParams, true
}

// SetSearchParams sets field value
func (o *ESearchSearchGroupRequest) SetSearchParams(v KalturaESearchGroupParams) {
	o.SearchParams = v
}

func (o ESearchSearchGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	if true {
		toSerialize["searchParams"] = o.SearchParams
	}
	return json.Marshal(toSerialize)
}

type NullableESearchSearchGroupRequest struct {
	value *ESearchSearchGroupRequest
	isSet bool
}

func (v NullableESearchSearchGroupRequest) Get() *ESearchSearchGroupRequest {
	return v.value
}

func (v *NullableESearchSearchGroupRequest) Set(val *ESearchSearchGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableESearchSearchGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableESearchSearchGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESearchSearchGroupRequest(val *ESearchSearchGroupRequest) *NullableESearchSearchGroupRequest {
	return &NullableESearchSearchGroupRequest{value: val, isSet: true}
}

func (v NullableESearchSearchGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESearchSearchGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

