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

// AccessControlProfileListRequest struct for AccessControlProfileListRequest
type AccessControlProfileListRequest struct {
	Filter *KalturaAccessControlProfileFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewAccessControlProfileListRequest instantiates a new AccessControlProfileListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessControlProfileListRequest() *AccessControlProfileListRequest {
	this := AccessControlProfileListRequest{}
	return &this
}

// NewAccessControlProfileListRequestWithDefaults instantiates a new AccessControlProfileListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessControlProfileListRequestWithDefaults() *AccessControlProfileListRequest {
	this := AccessControlProfileListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AccessControlProfileListRequest) GetFilter() KalturaAccessControlProfileFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaAccessControlProfileFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessControlProfileListRequest) GetFilterOk() (*KalturaAccessControlProfileFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AccessControlProfileListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaAccessControlProfileFilter and assigns it to the Filter field.
func (o *AccessControlProfileListRequest) SetFilter(v KalturaAccessControlProfileFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *AccessControlProfileListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessControlProfileListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *AccessControlProfileListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *AccessControlProfileListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o AccessControlProfileListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableAccessControlProfileListRequest struct {
	value *AccessControlProfileListRequest
	isSet bool
}

func (v NullableAccessControlProfileListRequest) Get() *AccessControlProfileListRequest {
	return v.value
}

func (v *NullableAccessControlProfileListRequest) Set(val *AccessControlProfileListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessControlProfileListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessControlProfileListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessControlProfileListRequest(val *AccessControlProfileListRequest) *NullableAccessControlProfileListRequest {
	return &NullableAccessControlProfileListRequest{value: val, isSet: true}
}

func (v NullableAccessControlProfileListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessControlProfileListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


