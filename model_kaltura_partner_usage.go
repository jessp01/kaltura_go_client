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

// KalturaPartnerUsage struct for KalturaPartnerUsage
type KalturaPartnerUsage struct {
	// `readOnly`  percent of usage out of partner's package. if usageGB is 5 and package is 10GB, this value will be 50
	Percent *float32 `json:"Percent,omitempty"`
	// `readOnly`  Partner total hosting in GB on the disk
	HostingGB *float32 `json:"hostingGB,omitempty"`
	// `readOnly`  package total BW - actually this is usage, which represents BW+storage
	PackageBW *int32 `json:"packageBW,omitempty"`
	// `readOnly`  date when partner reached the limit of his package (timestamp)
	ReachedLimitDate *int32 `json:"reachedLimitDate,omitempty"`
	// `readOnly`  total usage in GB - including bandwidth and storage
	UsageGB *float32 `json:"usageGB,omitempty"`
	// `readOnly`  a semi-colon separated list of comma-separated key-values to represent a usage graph.  keys could be 1-12 for a year view (1,1.2;2,1.1;3,0.9;...;12,1.4;)  keys could be 1-[28,29,30,31] depending on the requested month, for a daily view in a given month (1,0.4;2,0.2;...;31,0.1;)
	UsageGraph *string `json:"usageGraph,omitempty"`
}

// NewKalturaPartnerUsage instantiates a new KalturaPartnerUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPartnerUsage() *KalturaPartnerUsage {
	this := KalturaPartnerUsage{}
	return &this
}

// NewKalturaPartnerUsageWithDefaults instantiates a new KalturaPartnerUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPartnerUsageWithDefaults() *KalturaPartnerUsage {
	this := KalturaPartnerUsage{}
	return &this
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *KalturaPartnerUsage) GetPercent() float32 {
	if o == nil || o.Percent == nil {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerUsage) GetPercentOk() (*float32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *KalturaPartnerUsage) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *KalturaPartnerUsage) SetPercent(v float32) {
	o.Percent = &v
}

// GetHostingGB returns the HostingGB field value if set, zero value otherwise.
func (o *KalturaPartnerUsage) GetHostingGB() float32 {
	if o == nil || o.HostingGB == nil {
		var ret float32
		return ret
	}
	return *o.HostingGB
}

// GetHostingGBOk returns a tuple with the HostingGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerUsage) GetHostingGBOk() (*float32, bool) {
	if o == nil || o.HostingGB == nil {
		return nil, false
	}
	return o.HostingGB, true
}

// HasHostingGB returns a boolean if a field has been set.
func (o *KalturaPartnerUsage) HasHostingGB() bool {
	if o != nil && o.HostingGB != nil {
		return true
	}

	return false
}

// SetHostingGB gets a reference to the given float32 and assigns it to the HostingGB field.
func (o *KalturaPartnerUsage) SetHostingGB(v float32) {
	o.HostingGB = &v
}

// GetPackageBW returns the PackageBW field value if set, zero value otherwise.
func (o *KalturaPartnerUsage) GetPackageBW() int32 {
	if o == nil || o.PackageBW == nil {
		var ret int32
		return ret
	}
	return *o.PackageBW
}

// GetPackageBWOk returns a tuple with the PackageBW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerUsage) GetPackageBWOk() (*int32, bool) {
	if o == nil || o.PackageBW == nil {
		return nil, false
	}
	return o.PackageBW, true
}

// HasPackageBW returns a boolean if a field has been set.
func (o *KalturaPartnerUsage) HasPackageBW() bool {
	if o != nil && o.PackageBW != nil {
		return true
	}

	return false
}

// SetPackageBW gets a reference to the given int32 and assigns it to the PackageBW field.
func (o *KalturaPartnerUsage) SetPackageBW(v int32) {
	o.PackageBW = &v
}

// GetReachedLimitDate returns the ReachedLimitDate field value if set, zero value otherwise.
func (o *KalturaPartnerUsage) GetReachedLimitDate() int32 {
	if o == nil || o.ReachedLimitDate == nil {
		var ret int32
		return ret
	}
	return *o.ReachedLimitDate
}

// GetReachedLimitDateOk returns a tuple with the ReachedLimitDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerUsage) GetReachedLimitDateOk() (*int32, bool) {
	if o == nil || o.ReachedLimitDate == nil {
		return nil, false
	}
	return o.ReachedLimitDate, true
}

// HasReachedLimitDate returns a boolean if a field has been set.
func (o *KalturaPartnerUsage) HasReachedLimitDate() bool {
	if o != nil && o.ReachedLimitDate != nil {
		return true
	}

	return false
}

// SetReachedLimitDate gets a reference to the given int32 and assigns it to the ReachedLimitDate field.
func (o *KalturaPartnerUsage) SetReachedLimitDate(v int32) {
	o.ReachedLimitDate = &v
}

// GetUsageGB returns the UsageGB field value if set, zero value otherwise.
func (o *KalturaPartnerUsage) GetUsageGB() float32 {
	if o == nil || o.UsageGB == nil {
		var ret float32
		return ret
	}
	return *o.UsageGB
}

// GetUsageGBOk returns a tuple with the UsageGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerUsage) GetUsageGBOk() (*float32, bool) {
	if o == nil || o.UsageGB == nil {
		return nil, false
	}
	return o.UsageGB, true
}

// HasUsageGB returns a boolean if a field has been set.
func (o *KalturaPartnerUsage) HasUsageGB() bool {
	if o != nil && o.UsageGB != nil {
		return true
	}

	return false
}

// SetUsageGB gets a reference to the given float32 and assigns it to the UsageGB field.
func (o *KalturaPartnerUsage) SetUsageGB(v float32) {
	o.UsageGB = &v
}

// GetUsageGraph returns the UsageGraph field value if set, zero value otherwise.
func (o *KalturaPartnerUsage) GetUsageGraph() string {
	if o == nil || o.UsageGraph == nil {
		var ret string
		return ret
	}
	return *o.UsageGraph
}

// GetUsageGraphOk returns a tuple with the UsageGraph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPartnerUsage) GetUsageGraphOk() (*string, bool) {
	if o == nil || o.UsageGraph == nil {
		return nil, false
	}
	return o.UsageGraph, true
}

// HasUsageGraph returns a boolean if a field has been set.
func (o *KalturaPartnerUsage) HasUsageGraph() bool {
	if o != nil && o.UsageGraph != nil {
		return true
	}

	return false
}

// SetUsageGraph gets a reference to the given string and assigns it to the UsageGraph field.
func (o *KalturaPartnerUsage) SetUsageGraph(v string) {
	o.UsageGraph = &v
}

func (o KalturaPartnerUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Percent != nil {
		toSerialize["Percent"] = o.Percent
	}
	if o.HostingGB != nil {
		toSerialize["hostingGB"] = o.HostingGB
	}
	if o.PackageBW != nil {
		toSerialize["packageBW"] = o.PackageBW
	}
	if o.ReachedLimitDate != nil {
		toSerialize["reachedLimitDate"] = o.ReachedLimitDate
	}
	if o.UsageGB != nil {
		toSerialize["usageGB"] = o.UsageGB
	}
	if o.UsageGraph != nil {
		toSerialize["usageGraph"] = o.UsageGraph
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPartnerUsage struct {
	value *KalturaPartnerUsage
	isSet bool
}

func (v NullableKalturaPartnerUsage) Get() *KalturaPartnerUsage {
	return v.value
}

func (v *NullableKalturaPartnerUsage) Set(val *KalturaPartnerUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPartnerUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPartnerUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPartnerUsage(val *KalturaPartnerUsage) *NullableKalturaPartnerUsage {
	return &NullableKalturaPartnerUsage{value: val, isSet: true}
}

func (v NullableKalturaPartnerUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPartnerUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


