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

// KalturaScheduledTaskProfile struct for KalturaScheduledTaskProfile
type KalturaScheduledTaskProfile struct {
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	Description *string `json:"description,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	LastExecutionStartedAt *int32 `json:"lastExecutionStartedAt,omitempty"`
	// The maximum number of result count allowed to be processed by this profile per execution
	MaxTotalCountAllowed *int32 `json:"maxTotalCountAllowed,omitempty"`
	Name *string `json:"name,omitempty"`
	ObjectFilter *KalturaFilter `json:"objectFilter,omitempty"`
	// Enum Type: `KalturaObjectFilterEngineType`  The type of engine to use to list objects using the given \"objectFilter\"
	ObjectFilterEngineType *string `json:"objectFilterEngineType,omitempty"`
	ObjectTasks []KalturaObjectTask `json:"objectTasks,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// Enum Type: `KalturaScheduledTaskProfileStatus`
	Status *int32 `json:"status,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
}

// NewKalturaScheduledTaskProfile instantiates a new KalturaScheduledTaskProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduledTaskProfile() *KalturaScheduledTaskProfile {
	this := KalturaScheduledTaskProfile{}
	return &this
}

// NewKalturaScheduledTaskProfileWithDefaults instantiates a new KalturaScheduledTaskProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduledTaskProfileWithDefaults() *KalturaScheduledTaskProfile {
	this := KalturaScheduledTaskProfile{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaScheduledTaskProfile) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaScheduledTaskProfile) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaScheduledTaskProfile) SetId(v int32) {
	o.Id = &v
}

// GetLastExecutionStartedAt returns the LastExecutionStartedAt field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetLastExecutionStartedAt() int32 {
	if o == nil || o.LastExecutionStartedAt == nil {
		var ret int32
		return ret
	}
	return *o.LastExecutionStartedAt
}

// GetLastExecutionStartedAtOk returns a tuple with the LastExecutionStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetLastExecutionStartedAtOk() (*int32, bool) {
	if o == nil || o.LastExecutionStartedAt == nil {
		return nil, false
	}
	return o.LastExecutionStartedAt, true
}

// HasLastExecutionStartedAt returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasLastExecutionStartedAt() bool {
	if o != nil && o.LastExecutionStartedAt != nil {
		return true
	}

	return false
}

// SetLastExecutionStartedAt gets a reference to the given int32 and assigns it to the LastExecutionStartedAt field.
func (o *KalturaScheduledTaskProfile) SetLastExecutionStartedAt(v int32) {
	o.LastExecutionStartedAt = &v
}

// GetMaxTotalCountAllowed returns the MaxTotalCountAllowed field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetMaxTotalCountAllowed() int32 {
	if o == nil || o.MaxTotalCountAllowed == nil {
		var ret int32
		return ret
	}
	return *o.MaxTotalCountAllowed
}

// GetMaxTotalCountAllowedOk returns a tuple with the MaxTotalCountAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetMaxTotalCountAllowedOk() (*int32, bool) {
	if o == nil || o.MaxTotalCountAllowed == nil {
		return nil, false
	}
	return o.MaxTotalCountAllowed, true
}

// HasMaxTotalCountAllowed returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasMaxTotalCountAllowed() bool {
	if o != nil && o.MaxTotalCountAllowed != nil {
		return true
	}

	return false
}

// SetMaxTotalCountAllowed gets a reference to the given int32 and assigns it to the MaxTotalCountAllowed field.
func (o *KalturaScheduledTaskProfile) SetMaxTotalCountAllowed(v int32) {
	o.MaxTotalCountAllowed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaScheduledTaskProfile) SetName(v string) {
	o.Name = &v
}

// GetObjectFilter returns the ObjectFilter field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetObjectFilter() KalturaFilter {
	if o == nil || o.ObjectFilter == nil {
		var ret KalturaFilter
		return ret
	}
	return *o.ObjectFilter
}

// GetObjectFilterOk returns a tuple with the ObjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetObjectFilterOk() (*KalturaFilter, bool) {
	if o == nil || o.ObjectFilter == nil {
		return nil, false
	}
	return o.ObjectFilter, true
}

// HasObjectFilter returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasObjectFilter() bool {
	if o != nil && o.ObjectFilter != nil {
		return true
	}

	return false
}

// SetObjectFilter gets a reference to the given KalturaFilter and assigns it to the ObjectFilter field.
func (o *KalturaScheduledTaskProfile) SetObjectFilter(v KalturaFilter) {
	o.ObjectFilter = &v
}

// GetObjectFilterEngineType returns the ObjectFilterEngineType field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetObjectFilterEngineType() string {
	if o == nil || o.ObjectFilterEngineType == nil {
		var ret string
		return ret
	}
	return *o.ObjectFilterEngineType
}

// GetObjectFilterEngineTypeOk returns a tuple with the ObjectFilterEngineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetObjectFilterEngineTypeOk() (*string, bool) {
	if o == nil || o.ObjectFilterEngineType == nil {
		return nil, false
	}
	return o.ObjectFilterEngineType, true
}

// HasObjectFilterEngineType returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasObjectFilterEngineType() bool {
	if o != nil && o.ObjectFilterEngineType != nil {
		return true
	}

	return false
}

// SetObjectFilterEngineType gets a reference to the given string and assigns it to the ObjectFilterEngineType field.
func (o *KalturaScheduledTaskProfile) SetObjectFilterEngineType(v string) {
	o.ObjectFilterEngineType = &v
}

// GetObjectTasks returns the ObjectTasks field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetObjectTasks() []KalturaObjectTask {
	if o == nil || o.ObjectTasks == nil {
		var ret []KalturaObjectTask
		return ret
	}
	return o.ObjectTasks
}

// GetObjectTasksOk returns a tuple with the ObjectTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetObjectTasksOk() ([]KalturaObjectTask, bool) {
	if o == nil || o.ObjectTasks == nil {
		return nil, false
	}
	return o.ObjectTasks, true
}

// HasObjectTasks returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasObjectTasks() bool {
	if o != nil && o.ObjectTasks != nil {
		return true
	}

	return false
}

// SetObjectTasks gets a reference to the given []KalturaObjectTask and assigns it to the ObjectTasks field.
func (o *KalturaScheduledTaskProfile) SetObjectTasks(v []KalturaObjectTask) {
	o.ObjectTasks = v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaScheduledTaskProfile) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaScheduledTaskProfile) SetStatus(v int32) {
	o.Status = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaScheduledTaskProfile) SetSystemName(v string) {
	o.SystemName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaScheduledTaskProfile) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduledTaskProfile) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaScheduledTaskProfile) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaScheduledTaskProfile) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o KalturaScheduledTaskProfile) MarshalJSON() ([]byte, error) {
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
	if o.LastExecutionStartedAt != nil {
		toSerialize["lastExecutionStartedAt"] = o.LastExecutionStartedAt
	}
	if o.MaxTotalCountAllowed != nil {
		toSerialize["maxTotalCountAllowed"] = o.MaxTotalCountAllowed
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectFilter != nil {
		toSerialize["objectFilter"] = o.ObjectFilter
	}
	if o.ObjectFilterEngineType != nil {
		toSerialize["objectFilterEngineType"] = o.ObjectFilterEngineType
	}
	if o.ObjectTasks != nil {
		toSerialize["objectTasks"] = o.ObjectTasks
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScheduledTaskProfile struct {
	value *KalturaScheduledTaskProfile
	isSet bool
}

func (v NullableKalturaScheduledTaskProfile) Get() *KalturaScheduledTaskProfile {
	return v.value
}

func (v *NullableKalturaScheduledTaskProfile) Set(val *KalturaScheduledTaskProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduledTaskProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduledTaskProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduledTaskProfile(val *KalturaScheduledTaskProfile) *NullableKalturaScheduledTaskProfile {
	return &NullableKalturaScheduledTaskProfile{value: val, isSet: true}
}

func (v NullableKalturaScheduledTaskProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduledTaskProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

