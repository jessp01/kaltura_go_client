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

// KalturaBulkUploadResult struct for KalturaBulkUploadResult
type KalturaBulkUploadResult struct {
	// Enum Type: `KalturaBulkUploadAction`
	Action *string `json:"action,omitempty"`
	// The id of the parent job
	BulkUploadJobId *int32 `json:"bulkUploadJobId,omitempty"`
	// Enum Type: `KalturaBulkUploadObjectType`
	BulkUploadResultObjectType *string `json:"bulkUploadResultObjectType,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	ErrorType *int32 `json:"errorType,omitempty"`
	// `readOnly`  The id of the result
	Id *int32 `json:"id,omitempty"`
	// The index of the line in the CSV
	LineIndex *int32 `json:"lineIndex,omitempty"`
	ObjectErrorDescription *string `json:"objectErrorDescription,omitempty"`
	ObjectId *string `json:"objectId,omitempty"`
	ObjectStatus *int32 `json:"objectStatus,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	PartnerData *string `json:"partnerData,omitempty"`
	PartnerId *int32 `json:"partnerId,omitempty"`
	PluginsData []KalturaBulkUploadPluginData `json:"pluginsData,omitempty"`
	// The data as recieved in the csv
	RowData *string `json:"rowData,omitempty"`
	// Enum Type: `KalturaBulkUploadResultStatus`
	Status *string `json:"status,omitempty"`
}

// NewKalturaBulkUploadResult instantiates a new KalturaBulkUploadResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadResult() *KalturaBulkUploadResult {
	this := KalturaBulkUploadResult{}
	return &this
}

// NewKalturaBulkUploadResultWithDefaults instantiates a new KalturaBulkUploadResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadResultWithDefaults() *KalturaBulkUploadResult {
	this := KalturaBulkUploadResult{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *KalturaBulkUploadResult) SetAction(v string) {
	o.Action = &v
}

// GetBulkUploadJobId returns the BulkUploadJobId field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetBulkUploadJobId() int32 {
	if o == nil || o.BulkUploadJobId == nil {
		var ret int32
		return ret
	}
	return *o.BulkUploadJobId
}

// GetBulkUploadJobIdOk returns a tuple with the BulkUploadJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetBulkUploadJobIdOk() (*int32, bool) {
	if o == nil || o.BulkUploadJobId == nil {
		return nil, false
	}
	return o.BulkUploadJobId, true
}

// HasBulkUploadJobId returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasBulkUploadJobId() bool {
	if o != nil && o.BulkUploadJobId != nil {
		return true
	}

	return false
}

// SetBulkUploadJobId gets a reference to the given int32 and assigns it to the BulkUploadJobId field.
func (o *KalturaBulkUploadResult) SetBulkUploadJobId(v int32) {
	o.BulkUploadJobId = &v
}

// GetBulkUploadResultObjectType returns the BulkUploadResultObjectType field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetBulkUploadResultObjectType() string {
	if o == nil || o.BulkUploadResultObjectType == nil {
		var ret string
		return ret
	}
	return *o.BulkUploadResultObjectType
}

// GetBulkUploadResultObjectTypeOk returns a tuple with the BulkUploadResultObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetBulkUploadResultObjectTypeOk() (*string, bool) {
	if o == nil || o.BulkUploadResultObjectType == nil {
		return nil, false
	}
	return o.BulkUploadResultObjectType, true
}

// HasBulkUploadResultObjectType returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasBulkUploadResultObjectType() bool {
	if o != nil && o.BulkUploadResultObjectType != nil {
		return true
	}

	return false
}

// SetBulkUploadResultObjectType gets a reference to the given string and assigns it to the BulkUploadResultObjectType field.
func (o *KalturaBulkUploadResult) SetBulkUploadResultObjectType(v string) {
	o.BulkUploadResultObjectType = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *KalturaBulkUploadResult) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *KalturaBulkUploadResult) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetErrorType() int32 {
	if o == nil || o.ErrorType == nil {
		var ret int32
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetErrorTypeOk() (*int32, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given int32 and assigns it to the ErrorType field.
func (o *KalturaBulkUploadResult) SetErrorType(v int32) {
	o.ErrorType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaBulkUploadResult) SetId(v int32) {
	o.Id = &v
}

// GetLineIndex returns the LineIndex field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetLineIndex() int32 {
	if o == nil || o.LineIndex == nil {
		var ret int32
		return ret
	}
	return *o.LineIndex
}

// GetLineIndexOk returns a tuple with the LineIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetLineIndexOk() (*int32, bool) {
	if o == nil || o.LineIndex == nil {
		return nil, false
	}
	return o.LineIndex, true
}

// HasLineIndex returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasLineIndex() bool {
	if o != nil && o.LineIndex != nil {
		return true
	}

	return false
}

// SetLineIndex gets a reference to the given int32 and assigns it to the LineIndex field.
func (o *KalturaBulkUploadResult) SetLineIndex(v int32) {
	o.LineIndex = &v
}

// GetObjectErrorDescription returns the ObjectErrorDescription field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetObjectErrorDescription() string {
	if o == nil || o.ObjectErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ObjectErrorDescription
}

// GetObjectErrorDescriptionOk returns a tuple with the ObjectErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetObjectErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ObjectErrorDescription == nil {
		return nil, false
	}
	return o.ObjectErrorDescription, true
}

// HasObjectErrorDescription returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasObjectErrorDescription() bool {
	if o != nil && o.ObjectErrorDescription != nil {
		return true
	}

	return false
}

// SetObjectErrorDescription gets a reference to the given string and assigns it to the ObjectErrorDescription field.
func (o *KalturaBulkUploadResult) SetObjectErrorDescription(v string) {
	o.ObjectErrorDescription = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *KalturaBulkUploadResult) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObjectStatus returns the ObjectStatus field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetObjectStatus() int32 {
	if o == nil || o.ObjectStatus == nil {
		var ret int32
		return ret
	}
	return *o.ObjectStatus
}

// GetObjectStatusOk returns a tuple with the ObjectStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetObjectStatusOk() (*int32, bool) {
	if o == nil || o.ObjectStatus == nil {
		return nil, false
	}
	return o.ObjectStatus, true
}

// HasObjectStatus returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasObjectStatus() bool {
	if o != nil && o.ObjectStatus != nil {
		return true
	}

	return false
}

// SetObjectStatus gets a reference to the given int32 and assigns it to the ObjectStatus field.
func (o *KalturaBulkUploadResult) SetObjectStatus(v int32) {
	o.ObjectStatus = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaBulkUploadResult) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerData returns the PartnerData field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetPartnerData() string {
	if o == nil || o.PartnerData == nil {
		var ret string
		return ret
	}
	return *o.PartnerData
}

// GetPartnerDataOk returns a tuple with the PartnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetPartnerDataOk() (*string, bool) {
	if o == nil || o.PartnerData == nil {
		return nil, false
	}
	return o.PartnerData, true
}

// HasPartnerData returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasPartnerData() bool {
	if o != nil && o.PartnerData != nil {
		return true
	}

	return false
}

// SetPartnerData gets a reference to the given string and assigns it to the PartnerData field.
func (o *KalturaBulkUploadResult) SetPartnerData(v string) {
	o.PartnerData = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaBulkUploadResult) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPluginsData returns the PluginsData field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetPluginsData() []KalturaBulkUploadPluginData {
	if o == nil || o.PluginsData == nil {
		var ret []KalturaBulkUploadPluginData
		return ret
	}
	return o.PluginsData
}

// GetPluginsDataOk returns a tuple with the PluginsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetPluginsDataOk() ([]KalturaBulkUploadPluginData, bool) {
	if o == nil || o.PluginsData == nil {
		return nil, false
	}
	return o.PluginsData, true
}

// HasPluginsData returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasPluginsData() bool {
	if o != nil && o.PluginsData != nil {
		return true
	}

	return false
}

// SetPluginsData gets a reference to the given []KalturaBulkUploadPluginData and assigns it to the PluginsData field.
func (o *KalturaBulkUploadResult) SetPluginsData(v []KalturaBulkUploadPluginData) {
	o.PluginsData = v
}

// GetRowData returns the RowData field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetRowData() string {
	if o == nil || o.RowData == nil {
		var ret string
		return ret
	}
	return *o.RowData
}

// GetRowDataOk returns a tuple with the RowData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetRowDataOk() (*string, bool) {
	if o == nil || o.RowData == nil {
		return nil, false
	}
	return o.RowData, true
}

// HasRowData returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasRowData() bool {
	if o != nil && o.RowData != nil {
		return true
	}

	return false
}

// SetRowData gets a reference to the given string and assigns it to the RowData field.
func (o *KalturaBulkUploadResult) SetRowData(v string) {
	o.RowData = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaBulkUploadResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBulkUploadResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaBulkUploadResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KalturaBulkUploadResult) SetStatus(v string) {
	o.Status = &v
}

func (o KalturaBulkUploadResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.BulkUploadJobId != nil {
		toSerialize["bulkUploadJobId"] = o.BulkUploadJobId
	}
	if o.BulkUploadResultObjectType != nil {
		toSerialize["bulkUploadResultObjectType"] = o.BulkUploadResultObjectType
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorType != nil {
		toSerialize["errorType"] = o.ErrorType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LineIndex != nil {
		toSerialize["lineIndex"] = o.LineIndex
	}
	if o.ObjectErrorDescription != nil {
		toSerialize["objectErrorDescription"] = o.ObjectErrorDescription
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectStatus != nil {
		toSerialize["objectStatus"] = o.ObjectStatus
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerData != nil {
		toSerialize["partnerData"] = o.PartnerData
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.PluginsData != nil {
		toSerialize["pluginsData"] = o.PluginsData
	}
	if o.RowData != nil {
		toSerialize["rowData"] = o.RowData
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBulkUploadResult struct {
	value *KalturaBulkUploadResult
	isSet bool
}

func (v NullableKalturaBulkUploadResult) Get() *KalturaBulkUploadResult {
	return v.value
}

func (v *NullableKalturaBulkUploadResult) Set(val *KalturaBulkUploadResult) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadResult) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadResult(val *KalturaBulkUploadResult) *NullableKalturaBulkUploadResult {
	return &NullableKalturaBulkUploadResult{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


