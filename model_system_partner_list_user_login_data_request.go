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

// SystemPartnerListUserLoginDataRequest struct for SystemPartnerListUserLoginDataRequest
type SystemPartnerListUserLoginDataRequest struct {
	Filter *KalturaUserLoginDataFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewSystemPartnerListUserLoginDataRequest instantiates a new SystemPartnerListUserLoginDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPartnerListUserLoginDataRequest() *SystemPartnerListUserLoginDataRequest {
	this := SystemPartnerListUserLoginDataRequest{}
	return &this
}

// NewSystemPartnerListUserLoginDataRequestWithDefaults instantiates a new SystemPartnerListUserLoginDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPartnerListUserLoginDataRequestWithDefaults() *SystemPartnerListUserLoginDataRequest {
	this := SystemPartnerListUserLoginDataRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SystemPartnerListUserLoginDataRequest) GetFilter() KalturaUserLoginDataFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaUserLoginDataFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPartnerListUserLoginDataRequest) GetFilterOk() (*KalturaUserLoginDataFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SystemPartnerListUserLoginDataRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaUserLoginDataFilter and assigns it to the Filter field.
func (o *SystemPartnerListUserLoginDataRequest) SetFilter(v KalturaUserLoginDataFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *SystemPartnerListUserLoginDataRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPartnerListUserLoginDataRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *SystemPartnerListUserLoginDataRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *SystemPartnerListUserLoginDataRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o SystemPartnerListUserLoginDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPartnerListUserLoginDataRequest struct {
	value *SystemPartnerListUserLoginDataRequest
	isSet bool
}

func (v NullableSystemPartnerListUserLoginDataRequest) Get() *SystemPartnerListUserLoginDataRequest {
	return v.value
}

func (v *NullableSystemPartnerListUserLoginDataRequest) Set(val *SystemPartnerListUserLoginDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPartnerListUserLoginDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPartnerListUserLoginDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPartnerListUserLoginDataRequest(val *SystemPartnerListUserLoginDataRequest) *NullableSystemPartnerListUserLoginDataRequest {
	return &NullableSystemPartnerListUserLoginDataRequest{value: val, isSet: true}
}

func (v NullableSystemPartnerListUserLoginDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPartnerListUserLoginDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


