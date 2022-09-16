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

// KalturaCategoryEntry struct for KalturaCategoryEntry
type KalturaCategoryEntry struct {
	// `readOnly`  The full ids of the Category
	CategoryFullIds *string `json:"categoryFullIds,omitempty"`
	CategoryId *int32 `json:"categoryId,omitempty"`
	// `readOnly`  Creation date as Unix timestamp (In seconds)
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `readOnly`  CategroyEntry creator puser ID
	CreatorUserId *string `json:"creatorUserId,omitempty"`
	// entry id
	EntryId *string `json:"entryId,omitempty"`
	// `readOnly`  Enum Type: `KalturaCategoryEntryStatus`  CategroyEntry status
	Status *int32 `json:"status,omitempty"`
}

// NewKalturaCategoryEntry instantiates a new KalturaCategoryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCategoryEntry() *KalturaCategoryEntry {
	this := KalturaCategoryEntry{}
	return &this
}

// NewKalturaCategoryEntryWithDefaults instantiates a new KalturaCategoryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCategoryEntryWithDefaults() *KalturaCategoryEntry {
	this := KalturaCategoryEntry{}
	return &this
}

// GetCategoryFullIds returns the CategoryFullIds field value if set, zero value otherwise.
func (o *KalturaCategoryEntry) GetCategoryFullIds() string {
	if o == nil || o.CategoryFullIds == nil {
		var ret string
		return ret
	}
	return *o.CategoryFullIds
}

// GetCategoryFullIdsOk returns a tuple with the CategoryFullIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCategoryEntry) GetCategoryFullIdsOk() (*string, bool) {
	if o == nil || o.CategoryFullIds == nil {
		return nil, false
	}
	return o.CategoryFullIds, true
}

// HasCategoryFullIds returns a boolean if a field has been set.
func (o *KalturaCategoryEntry) HasCategoryFullIds() bool {
	if o != nil && o.CategoryFullIds != nil {
		return true
	}

	return false
}

// SetCategoryFullIds gets a reference to the given string and assigns it to the CategoryFullIds field.
func (o *KalturaCategoryEntry) SetCategoryFullIds(v string) {
	o.CategoryFullIds = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *KalturaCategoryEntry) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCategoryEntry) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *KalturaCategoryEntry) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *KalturaCategoryEntry) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaCategoryEntry) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCategoryEntry) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaCategoryEntry) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaCategoryEntry) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise.
func (o *KalturaCategoryEntry) GetCreatorUserId() string {
	if o == nil || o.CreatorUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatorUserId
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCategoryEntry) GetCreatorUserIdOk() (*string, bool) {
	if o == nil || o.CreatorUserId == nil {
		return nil, false
	}
	return o.CreatorUserId, true
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *KalturaCategoryEntry) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId != nil {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given string and assigns it to the CreatorUserId field.
func (o *KalturaCategoryEntry) SetCreatorUserId(v string) {
	o.CreatorUserId = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaCategoryEntry) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCategoryEntry) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaCategoryEntry) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaCategoryEntry) SetEntryId(v string) {
	o.EntryId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaCategoryEntry) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCategoryEntry) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaCategoryEntry) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaCategoryEntry) SetStatus(v int32) {
	o.Status = &v
}

func (o KalturaCategoryEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryFullIds != nil {
		toSerialize["categoryFullIds"] = o.CategoryFullIds
	}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatorUserId != nil {
		toSerialize["creatorUserId"] = o.CreatorUserId
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCategoryEntry struct {
	value *KalturaCategoryEntry
	isSet bool
}

func (v NullableKalturaCategoryEntry) Get() *KalturaCategoryEntry {
	return v.value
}

func (v *NullableKalturaCategoryEntry) Set(val *KalturaCategoryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCategoryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCategoryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCategoryEntry(val *KalturaCategoryEntry) *NullableKalturaCategoryEntry {
	return &NullableKalturaCategoryEntry{value: val, isSet: true}
}

func (v NullableKalturaCategoryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCategoryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

