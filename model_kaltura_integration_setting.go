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

// KalturaIntegrationSetting `abstract`
type KalturaIntegrationSetting struct {
	// `readOnly`
	AccountId *string `json:"accountId,omitempty"`
	ConversionProfileId *int32 `json:"conversionProfileId,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	CreateUserIfNotExist *int32 `json:"createUserIfNotExist,omitempty"`
	// `readOnly`
	CreatedAt *string `json:"createdAt,omitempty"`
	DefaultUserId *string `json:"defaultUserId,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	DeletionPolicy *int32 `json:"deletionPolicy,omitempty"`
	// Enum Type: `KalturaHandleParticipantsMode`
	HandleParticipantsMode *int32 `json:"handleParticipantsMode,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  Enum Type: `KalturaVendorIntegrationStatus`
	Status *int32 `json:"status,omitempty"`
	// `readOnly`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewKalturaIntegrationSetting instantiates a new KalturaIntegrationSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaIntegrationSetting() *KalturaIntegrationSetting {
	this := KalturaIntegrationSetting{}
	return &this
}

// NewKalturaIntegrationSettingWithDefaults instantiates a new KalturaIntegrationSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaIntegrationSettingWithDefaults() *KalturaIntegrationSetting {
	this := KalturaIntegrationSetting{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *KalturaIntegrationSetting) SetAccountId(v string) {
	o.AccountId = &v
}

// GetConversionProfileId returns the ConversionProfileId field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetConversionProfileId() int32 {
	if o == nil || o.ConversionProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ConversionProfileId
}

// GetConversionProfileIdOk returns a tuple with the ConversionProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetConversionProfileIdOk() (*int32, bool) {
	if o == nil || o.ConversionProfileId == nil {
		return nil, false
	}
	return o.ConversionProfileId, true
}

// HasConversionProfileId returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasConversionProfileId() bool {
	if o != nil && o.ConversionProfileId != nil {
		return true
	}

	return false
}

// SetConversionProfileId gets a reference to the given int32 and assigns it to the ConversionProfileId field.
func (o *KalturaIntegrationSetting) SetConversionProfileId(v int32) {
	o.ConversionProfileId = &v
}

// GetCreateUserIfNotExist returns the CreateUserIfNotExist field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetCreateUserIfNotExist() int32 {
	if o == nil || o.CreateUserIfNotExist == nil {
		var ret int32
		return ret
	}
	return *o.CreateUserIfNotExist
}

// GetCreateUserIfNotExistOk returns a tuple with the CreateUserIfNotExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetCreateUserIfNotExistOk() (*int32, bool) {
	if o == nil || o.CreateUserIfNotExist == nil {
		return nil, false
	}
	return o.CreateUserIfNotExist, true
}

// HasCreateUserIfNotExist returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasCreateUserIfNotExist() bool {
	if o != nil && o.CreateUserIfNotExist != nil {
		return true
	}

	return false
}

// SetCreateUserIfNotExist gets a reference to the given int32 and assigns it to the CreateUserIfNotExist field.
func (o *KalturaIntegrationSetting) SetCreateUserIfNotExist(v int32) {
	o.CreateUserIfNotExist = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *KalturaIntegrationSetting) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDefaultUserId returns the DefaultUserId field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetDefaultUserId() string {
	if o == nil || o.DefaultUserId == nil {
		var ret string
		return ret
	}
	return *o.DefaultUserId
}

// GetDefaultUserIdOk returns a tuple with the DefaultUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetDefaultUserIdOk() (*string, bool) {
	if o == nil || o.DefaultUserId == nil {
		return nil, false
	}
	return o.DefaultUserId, true
}

// HasDefaultUserId returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasDefaultUserId() bool {
	if o != nil && o.DefaultUserId != nil {
		return true
	}

	return false
}

// SetDefaultUserId gets a reference to the given string and assigns it to the DefaultUserId field.
func (o *KalturaIntegrationSetting) SetDefaultUserId(v string) {
	o.DefaultUserId = &v
}

// GetDeletionPolicy returns the DeletionPolicy field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetDeletionPolicy() int32 {
	if o == nil || o.DeletionPolicy == nil {
		var ret int32
		return ret
	}
	return *o.DeletionPolicy
}

// GetDeletionPolicyOk returns a tuple with the DeletionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetDeletionPolicyOk() (*int32, bool) {
	if o == nil || o.DeletionPolicy == nil {
		return nil, false
	}
	return o.DeletionPolicy, true
}

// HasDeletionPolicy returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasDeletionPolicy() bool {
	if o != nil && o.DeletionPolicy != nil {
		return true
	}

	return false
}

// SetDeletionPolicy gets a reference to the given int32 and assigns it to the DeletionPolicy field.
func (o *KalturaIntegrationSetting) SetDeletionPolicy(v int32) {
	o.DeletionPolicy = &v
}

// GetHandleParticipantsMode returns the HandleParticipantsMode field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetHandleParticipantsMode() int32 {
	if o == nil || o.HandleParticipantsMode == nil {
		var ret int32
		return ret
	}
	return *o.HandleParticipantsMode
}

// GetHandleParticipantsModeOk returns a tuple with the HandleParticipantsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetHandleParticipantsModeOk() (*int32, bool) {
	if o == nil || o.HandleParticipantsMode == nil {
		return nil, false
	}
	return o.HandleParticipantsMode, true
}

// HasHandleParticipantsMode returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasHandleParticipantsMode() bool {
	if o != nil && o.HandleParticipantsMode != nil {
		return true
	}

	return false
}

// SetHandleParticipantsMode gets a reference to the given int32 and assigns it to the HandleParticipantsMode field.
func (o *KalturaIntegrationSetting) SetHandleParticipantsMode(v int32) {
	o.HandleParticipantsMode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaIntegrationSetting) SetId(v int32) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaIntegrationSetting) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaIntegrationSetting) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaIntegrationSetting) SetStatus(v int32) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaIntegrationSetting) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaIntegrationSetting) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaIntegrationSetting) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *KalturaIntegrationSetting) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o KalturaIntegrationSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.ConversionProfileId != nil {
		toSerialize["conversionProfileId"] = o.ConversionProfileId
	}
	if o.CreateUserIfNotExist != nil {
		toSerialize["createUserIfNotExist"] = o.CreateUserIfNotExist
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DefaultUserId != nil {
		toSerialize["defaultUserId"] = o.DefaultUserId
	}
	if o.DeletionPolicy != nil {
		toSerialize["deletionPolicy"] = o.DeletionPolicy
	}
	if o.HandleParticipantsMode != nil {
		toSerialize["handleParticipantsMode"] = o.HandleParticipantsMode
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaIntegrationSetting struct {
	value *KalturaIntegrationSetting
	isSet bool
}

func (v NullableKalturaIntegrationSetting) Get() *KalturaIntegrationSetting {
	return v.value
}

func (v *NullableKalturaIntegrationSetting) Set(val *KalturaIntegrationSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaIntegrationSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaIntegrationSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaIntegrationSetting(val *KalturaIntegrationSetting) *NullableKalturaIntegrationSetting {
	return &NullableKalturaIntegrationSetting{value: val, isSet: true}
}

func (v NullableKalturaIntegrationSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaIntegrationSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

