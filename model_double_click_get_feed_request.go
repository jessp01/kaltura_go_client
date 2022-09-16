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

// DoubleClickGetFeedRequest struct for DoubleClickGetFeedRequest
type DoubleClickGetFeedRequest struct {
	DistributionProfileId int32 `json:"distributionProfileId"`
	Hash string `json:"hash"`
	IgnoreScheduling *bool `json:"ignoreScheduling,omitempty"`
	Page *int32 `json:"page,omitempty"`
	Period *int32 `json:"period,omitempty"`
	State *string `json:"state,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewDoubleClickGetFeedRequest instantiates a new DoubleClickGetFeedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDoubleClickGetFeedRequest(distributionProfileId int32, hash string) *DoubleClickGetFeedRequest {
	this := DoubleClickGetFeedRequest{}
	this.DistributionProfileId = distributionProfileId
	this.Hash = hash
	var ignoreScheduling bool = false
	this.IgnoreScheduling = &ignoreScheduling
	return &this
}

// NewDoubleClickGetFeedRequestWithDefaults instantiates a new DoubleClickGetFeedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDoubleClickGetFeedRequestWithDefaults() *DoubleClickGetFeedRequest {
	this := DoubleClickGetFeedRequest{}
	var ignoreScheduling bool = false
	this.IgnoreScheduling = &ignoreScheduling
	return &this
}

// GetDistributionProfileId returns the DistributionProfileId field value
func (o *DoubleClickGetFeedRequest) GetDistributionProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistributionProfileId
}

// GetDistributionProfileIdOk returns a tuple with the DistributionProfileId field value
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetDistributionProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistributionProfileId, true
}

// SetDistributionProfileId sets field value
func (o *DoubleClickGetFeedRequest) SetDistributionProfileId(v int32) {
	o.DistributionProfileId = v
}

// GetHash returns the Hash field value
func (o *DoubleClickGetFeedRequest) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *DoubleClickGetFeedRequest) SetHash(v string) {
	o.Hash = v
}

// GetIgnoreScheduling returns the IgnoreScheduling field value if set, zero value otherwise.
func (o *DoubleClickGetFeedRequest) GetIgnoreScheduling() bool {
	if o == nil || o.IgnoreScheduling == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreScheduling
}

// GetIgnoreSchedulingOk returns a tuple with the IgnoreScheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetIgnoreSchedulingOk() (*bool, bool) {
	if o == nil || o.IgnoreScheduling == nil {
		return nil, false
	}
	return o.IgnoreScheduling, true
}

// HasIgnoreScheduling returns a boolean if a field has been set.
func (o *DoubleClickGetFeedRequest) HasIgnoreScheduling() bool {
	if o != nil && o.IgnoreScheduling != nil {
		return true
	}

	return false
}

// SetIgnoreScheduling gets a reference to the given bool and assigns it to the IgnoreScheduling field.
func (o *DoubleClickGetFeedRequest) SetIgnoreScheduling(v bool) {
	o.IgnoreScheduling = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DoubleClickGetFeedRequest) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DoubleClickGetFeedRequest) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *DoubleClickGetFeedRequest) SetPage(v int32) {
	o.Page = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DoubleClickGetFeedRequest) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DoubleClickGetFeedRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *DoubleClickGetFeedRequest) SetPeriod(v int32) {
	o.Period = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DoubleClickGetFeedRequest) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DoubleClickGetFeedRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DoubleClickGetFeedRequest) SetState(v string) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DoubleClickGetFeedRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleClickGetFeedRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DoubleClickGetFeedRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DoubleClickGetFeedRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o DoubleClickGetFeedRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distributionProfileId"] = o.DistributionProfileId
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if o.IgnoreScheduling != nil {
		toSerialize["ignoreScheduling"] = o.IgnoreScheduling
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDoubleClickGetFeedRequest struct {
	value *DoubleClickGetFeedRequest
	isSet bool
}

func (v NullableDoubleClickGetFeedRequest) Get() *DoubleClickGetFeedRequest {
	return v.value
}

func (v *NullableDoubleClickGetFeedRequest) Set(val *DoubleClickGetFeedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDoubleClickGetFeedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDoubleClickGetFeedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDoubleClickGetFeedRequest(val *DoubleClickGetFeedRequest) *NullableDoubleClickGetFeedRequest {
	return &NullableDoubleClickGetFeedRequest{value: val, isSet: true}
}

func (v NullableDoubleClickGetFeedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDoubleClickGetFeedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


