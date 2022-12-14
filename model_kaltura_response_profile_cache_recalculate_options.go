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

// KalturaResponseProfileCacheRecalculateOptions struct for KalturaResponseProfileCacheRecalculateOptions
type KalturaResponseProfileCacheRecalculateOptions struct {
	// Class name
	CachedObjectType *string `json:"cachedObjectType,omitempty"`
	EndObjectKey *string `json:"endObjectKey,omitempty"`
	IsFirstLoop *bool `json:"isFirstLoop,omitempty"`
	JobCreatedAt *int32 `json:"jobCreatedAt,omitempty"`
	// Maximum number of keys to recalculate
	Limit *int32 `json:"limit,omitempty"`
	ObjectId *string `json:"objectId,omitempty"`
	StartObjectKey *string `json:"startObjectKey,omitempty"`
}

// NewKalturaResponseProfileCacheRecalculateOptions instantiates a new KalturaResponseProfileCacheRecalculateOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaResponseProfileCacheRecalculateOptions() *KalturaResponseProfileCacheRecalculateOptions {
	this := KalturaResponseProfileCacheRecalculateOptions{}
	return &this
}

// NewKalturaResponseProfileCacheRecalculateOptionsWithDefaults instantiates a new KalturaResponseProfileCacheRecalculateOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaResponseProfileCacheRecalculateOptionsWithDefaults() *KalturaResponseProfileCacheRecalculateOptions {
	this := KalturaResponseProfileCacheRecalculateOptions{}
	return &this
}

// GetCachedObjectType returns the CachedObjectType field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetCachedObjectType() string {
	if o == nil || o.CachedObjectType == nil {
		var ret string
		return ret
	}
	return *o.CachedObjectType
}

// GetCachedObjectTypeOk returns a tuple with the CachedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetCachedObjectTypeOk() (*string, bool) {
	if o == nil || o.CachedObjectType == nil {
		return nil, false
	}
	return o.CachedObjectType, true
}

// HasCachedObjectType returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasCachedObjectType() bool {
	if o != nil && o.CachedObjectType != nil {
		return true
	}

	return false
}

// SetCachedObjectType gets a reference to the given string and assigns it to the CachedObjectType field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetCachedObjectType(v string) {
	o.CachedObjectType = &v
}

// GetEndObjectKey returns the EndObjectKey field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetEndObjectKey() string {
	if o == nil || o.EndObjectKey == nil {
		var ret string
		return ret
	}
	return *o.EndObjectKey
}

// GetEndObjectKeyOk returns a tuple with the EndObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetEndObjectKeyOk() (*string, bool) {
	if o == nil || o.EndObjectKey == nil {
		return nil, false
	}
	return o.EndObjectKey, true
}

// HasEndObjectKey returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasEndObjectKey() bool {
	if o != nil && o.EndObjectKey != nil {
		return true
	}

	return false
}

// SetEndObjectKey gets a reference to the given string and assigns it to the EndObjectKey field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetEndObjectKey(v string) {
	o.EndObjectKey = &v
}

// GetIsFirstLoop returns the IsFirstLoop field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetIsFirstLoop() bool {
	if o == nil || o.IsFirstLoop == nil {
		var ret bool
		return ret
	}
	return *o.IsFirstLoop
}

// GetIsFirstLoopOk returns a tuple with the IsFirstLoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetIsFirstLoopOk() (*bool, bool) {
	if o == nil || o.IsFirstLoop == nil {
		return nil, false
	}
	return o.IsFirstLoop, true
}

// HasIsFirstLoop returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasIsFirstLoop() bool {
	if o != nil && o.IsFirstLoop != nil {
		return true
	}

	return false
}

// SetIsFirstLoop gets a reference to the given bool and assigns it to the IsFirstLoop field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetIsFirstLoop(v bool) {
	o.IsFirstLoop = &v
}

// GetJobCreatedAt returns the JobCreatedAt field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetJobCreatedAt() int32 {
	if o == nil || o.JobCreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.JobCreatedAt
}

// GetJobCreatedAtOk returns a tuple with the JobCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetJobCreatedAtOk() (*int32, bool) {
	if o == nil || o.JobCreatedAt == nil {
		return nil, false
	}
	return o.JobCreatedAt, true
}

// HasJobCreatedAt returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasJobCreatedAt() bool {
	if o != nil && o.JobCreatedAt != nil {
		return true
	}

	return false
}

// SetJobCreatedAt gets a reference to the given int32 and assigns it to the JobCreatedAt field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetJobCreatedAt(v int32) {
	o.JobCreatedAt = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetLimit(v int32) {
	o.Limit = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetStartObjectKey returns the StartObjectKey field value if set, zero value otherwise.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetStartObjectKey() string {
	if o == nil || o.StartObjectKey == nil {
		var ret string
		return ret
	}
	return *o.StartObjectKey
}

// GetStartObjectKeyOk returns a tuple with the StartObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) GetStartObjectKeyOk() (*string, bool) {
	if o == nil || o.StartObjectKey == nil {
		return nil, false
	}
	return o.StartObjectKey, true
}

// HasStartObjectKey returns a boolean if a field has been set.
func (o *KalturaResponseProfileCacheRecalculateOptions) HasStartObjectKey() bool {
	if o != nil && o.StartObjectKey != nil {
		return true
	}

	return false
}

// SetStartObjectKey gets a reference to the given string and assigns it to the StartObjectKey field.
func (o *KalturaResponseProfileCacheRecalculateOptions) SetStartObjectKey(v string) {
	o.StartObjectKey = &v
}

func (o KalturaResponseProfileCacheRecalculateOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CachedObjectType != nil {
		toSerialize["cachedObjectType"] = o.CachedObjectType
	}
	if o.EndObjectKey != nil {
		toSerialize["endObjectKey"] = o.EndObjectKey
	}
	if o.IsFirstLoop != nil {
		toSerialize["isFirstLoop"] = o.IsFirstLoop
	}
	if o.JobCreatedAt != nil {
		toSerialize["jobCreatedAt"] = o.JobCreatedAt
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.StartObjectKey != nil {
		toSerialize["startObjectKey"] = o.StartObjectKey
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaResponseProfileCacheRecalculateOptions struct {
	value *KalturaResponseProfileCacheRecalculateOptions
	isSet bool
}

func (v NullableKalturaResponseProfileCacheRecalculateOptions) Get() *KalturaResponseProfileCacheRecalculateOptions {
	return v.value
}

func (v *NullableKalturaResponseProfileCacheRecalculateOptions) Set(val *KalturaResponseProfileCacheRecalculateOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaResponseProfileCacheRecalculateOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaResponseProfileCacheRecalculateOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaResponseProfileCacheRecalculateOptions(val *KalturaResponseProfileCacheRecalculateOptions) *NullableKalturaResponseProfileCacheRecalculateOptions {
	return &NullableKalturaResponseProfileCacheRecalculateOptions{value: val, isSet: true}
}

func (v NullableKalturaResponseProfileCacheRecalculateOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaResponseProfileCacheRecalculateOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


