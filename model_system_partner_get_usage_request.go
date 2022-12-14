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

// SystemPartnerGetUsageRequest struct for SystemPartnerGetUsageRequest
type SystemPartnerGetUsageRequest struct {
	Pager *KalturaFilterPager `json:"pager,omitempty"`
	PartnerFilter *KalturaPartnerFilter `json:"partnerFilter,omitempty"`
	UsageFilter *KalturaSystemPartnerUsageFilter `json:"usageFilter,omitempty"`
}

// NewSystemPartnerGetUsageRequest instantiates a new SystemPartnerGetUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPartnerGetUsageRequest() *SystemPartnerGetUsageRequest {
	this := SystemPartnerGetUsageRequest{}
	return &this
}

// NewSystemPartnerGetUsageRequestWithDefaults instantiates a new SystemPartnerGetUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPartnerGetUsageRequestWithDefaults() *SystemPartnerGetUsageRequest {
	this := SystemPartnerGetUsageRequest{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *SystemPartnerGetUsageRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPartnerGetUsageRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *SystemPartnerGetUsageRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *SystemPartnerGetUsageRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

// GetPartnerFilter returns the PartnerFilter field value if set, zero value otherwise.
func (o *SystemPartnerGetUsageRequest) GetPartnerFilter() KalturaPartnerFilter {
	if o == nil || o.PartnerFilter == nil {
		var ret KalturaPartnerFilter
		return ret
	}
	return *o.PartnerFilter
}

// GetPartnerFilterOk returns a tuple with the PartnerFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPartnerGetUsageRequest) GetPartnerFilterOk() (*KalturaPartnerFilter, bool) {
	if o == nil || o.PartnerFilter == nil {
		return nil, false
	}
	return o.PartnerFilter, true
}

// HasPartnerFilter returns a boolean if a field has been set.
func (o *SystemPartnerGetUsageRequest) HasPartnerFilter() bool {
	if o != nil && o.PartnerFilter != nil {
		return true
	}

	return false
}

// SetPartnerFilter gets a reference to the given KalturaPartnerFilter and assigns it to the PartnerFilter field.
func (o *SystemPartnerGetUsageRequest) SetPartnerFilter(v KalturaPartnerFilter) {
	o.PartnerFilter = &v
}

// GetUsageFilter returns the UsageFilter field value if set, zero value otherwise.
func (o *SystemPartnerGetUsageRequest) GetUsageFilter() KalturaSystemPartnerUsageFilter {
	if o == nil || o.UsageFilter == nil {
		var ret KalturaSystemPartnerUsageFilter
		return ret
	}
	return *o.UsageFilter
}

// GetUsageFilterOk returns a tuple with the UsageFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPartnerGetUsageRequest) GetUsageFilterOk() (*KalturaSystemPartnerUsageFilter, bool) {
	if o == nil || o.UsageFilter == nil {
		return nil, false
	}
	return o.UsageFilter, true
}

// HasUsageFilter returns a boolean if a field has been set.
func (o *SystemPartnerGetUsageRequest) HasUsageFilter() bool {
	if o != nil && o.UsageFilter != nil {
		return true
	}

	return false
}

// SetUsageFilter gets a reference to the given KalturaSystemPartnerUsageFilter and assigns it to the UsageFilter field.
func (o *SystemPartnerGetUsageRequest) SetUsageFilter(v KalturaSystemPartnerUsageFilter) {
	o.UsageFilter = &v
}

func (o SystemPartnerGetUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	if o.PartnerFilter != nil {
		toSerialize["partnerFilter"] = o.PartnerFilter
	}
	if o.UsageFilter != nil {
		toSerialize["usageFilter"] = o.UsageFilter
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPartnerGetUsageRequest struct {
	value *SystemPartnerGetUsageRequest
	isSet bool
}

func (v NullableSystemPartnerGetUsageRequest) Get() *SystemPartnerGetUsageRequest {
	return v.value
}

func (v *NullableSystemPartnerGetUsageRequest) Set(val *SystemPartnerGetUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPartnerGetUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPartnerGetUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPartnerGetUsageRequest(val *SystemPartnerGetUsageRequest) *NullableSystemPartnerGetUsageRequest {
	return &NullableSystemPartnerGetUsageRequest{value: val, isSet: true}
}

func (v NullableSystemPartnerGetUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPartnerGetUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


