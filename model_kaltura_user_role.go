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

// KalturaUserRole struct for KalturaUserRole
type KalturaUserRole struct {
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	Description *string `json:"description,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	PermissionNames *string `json:"permissionNames,omitempty"`
	// Enum Type: `KalturaUserRoleStatus`
	Status *int32 `json:"status,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	Tags *string `json:"tags,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
}

// NewKalturaUserRole instantiates a new KalturaUserRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUserRole() *KalturaUserRole {
	this := KalturaUserRole{}
	return &this
}

// NewKalturaUserRoleWithDefaults instantiates a new KalturaUserRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUserRoleWithDefaults() *KalturaUserRole {
	this := KalturaUserRole{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaUserRole) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaUserRole) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaUserRole) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaUserRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaUserRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaUserRole) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaUserRole) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaUserRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaUserRole) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaUserRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaUserRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaUserRole) SetName(v string) {
	o.Name = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaUserRole) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaUserRole) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaUserRole) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPermissionNames returns the PermissionNames field value if set, zero value otherwise.
func (o *KalturaUserRole) GetPermissionNames() string {
	if o == nil || o.PermissionNames == nil {
		var ret string
		return ret
	}
	return *o.PermissionNames
}

// GetPermissionNamesOk returns a tuple with the PermissionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetPermissionNamesOk() (*string, bool) {
	if o == nil || o.PermissionNames == nil {
		return nil, false
	}
	return o.PermissionNames, true
}

// HasPermissionNames returns a boolean if a field has been set.
func (o *KalturaUserRole) HasPermissionNames() bool {
	if o != nil && o.PermissionNames != nil {
		return true
	}

	return false
}

// SetPermissionNames gets a reference to the given string and assigns it to the PermissionNames field.
func (o *KalturaUserRole) SetPermissionNames(v string) {
	o.PermissionNames = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaUserRole) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaUserRole) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaUserRole) SetStatus(v int32) {
	o.Status = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaUserRole) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaUserRole) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaUserRole) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *KalturaUserRole) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *KalturaUserRole) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *KalturaUserRole) SetTags(v string) {
	o.Tags = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaUserRole) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUserRole) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaUserRole) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaUserRole) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o KalturaUserRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.PermissionNames != nil {
		toSerialize["permissionNames"] = o.PermissionNames
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUserRole struct {
	value *KalturaUserRole
	isSet bool
}

func (v NullableKalturaUserRole) Get() *KalturaUserRole {
	return v.value
}

func (v *NullableKalturaUserRole) Set(val *KalturaUserRole) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUserRole) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUserRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUserRole(val *KalturaUserRole) *NullableKalturaUserRole {
	return &NullableKalturaUserRole{value: val, isSet: true}
}

func (v NullableKalturaUserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUserRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


