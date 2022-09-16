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

// DistributionProfileListByPartnerRequest struct for DistributionProfileListByPartnerRequest
type DistributionProfileListByPartnerRequest struct {
	Filter *KalturaPartnerFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewDistributionProfileListByPartnerRequest instantiates a new DistributionProfileListByPartnerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionProfileListByPartnerRequest() *DistributionProfileListByPartnerRequest {
	this := DistributionProfileListByPartnerRequest{}
	return &this
}

// NewDistributionProfileListByPartnerRequestWithDefaults instantiates a new DistributionProfileListByPartnerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionProfileListByPartnerRequestWithDefaults() *DistributionProfileListByPartnerRequest {
	this := DistributionProfileListByPartnerRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DistributionProfileListByPartnerRequest) GetFilter() KalturaPartnerFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaPartnerFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionProfileListByPartnerRequest) GetFilterOk() (*KalturaPartnerFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DistributionProfileListByPartnerRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaPartnerFilter and assigns it to the Filter field.
func (o *DistributionProfileListByPartnerRequest) SetFilter(v KalturaPartnerFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DistributionProfileListByPartnerRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionProfileListByPartnerRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DistributionProfileListByPartnerRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *DistributionProfileListByPartnerRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o DistributionProfileListByPartnerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableDistributionProfileListByPartnerRequest struct {
	value *DistributionProfileListByPartnerRequest
	isSet bool
}

func (v NullableDistributionProfileListByPartnerRequest) Get() *DistributionProfileListByPartnerRequest {
	return v.value
}

func (v *NullableDistributionProfileListByPartnerRequest) Set(val *DistributionProfileListByPartnerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionProfileListByPartnerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionProfileListByPartnerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionProfileListByPartnerRequest(val *DistributionProfileListByPartnerRequest) *NullableDistributionProfileListByPartnerRequest {
	return &NullableDistributionProfileListByPartnerRequest{value: val, isSet: true}
}

func (v NullableDistributionProfileListByPartnerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionProfileListByPartnerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


