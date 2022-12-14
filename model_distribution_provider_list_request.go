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

// DistributionProviderListRequest struct for DistributionProviderListRequest
type DistributionProviderListRequest struct {
	Filter *KalturaDistributionProviderFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewDistributionProviderListRequest instantiates a new DistributionProviderListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionProviderListRequest() *DistributionProviderListRequest {
	this := DistributionProviderListRequest{}
	return &this
}

// NewDistributionProviderListRequestWithDefaults instantiates a new DistributionProviderListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionProviderListRequestWithDefaults() *DistributionProviderListRequest {
	this := DistributionProviderListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DistributionProviderListRequest) GetFilter() KalturaDistributionProviderFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaDistributionProviderFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionProviderListRequest) GetFilterOk() (*KalturaDistributionProviderFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DistributionProviderListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaDistributionProviderFilter and assigns it to the Filter field.
func (o *DistributionProviderListRequest) SetFilter(v KalturaDistributionProviderFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DistributionProviderListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionProviderListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DistributionProviderListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *DistributionProviderListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o DistributionProviderListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableDistributionProviderListRequest struct {
	value *DistributionProviderListRequest
	isSet bool
}

func (v NullableDistributionProviderListRequest) Get() *DistributionProviderListRequest {
	return v.value
}

func (v *NullableDistributionProviderListRequest) Set(val *DistributionProviderListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionProviderListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionProviderListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionProviderListRequest(val *DistributionProviderListRequest) *NullableDistributionProviderListRequest {
	return &NullableDistributionProviderListRequest{value: val, isSet: true}
}

func (v NullableDistributionProviderListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionProviderListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


