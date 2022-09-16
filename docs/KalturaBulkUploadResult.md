# KalturaBulkUploadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Enum Type: &#x60;KalturaBulkUploadAction&#x60; | [optional] 
**BulkUploadJobId** | Pointer to **int32** | The id of the parent job | [optional] 
**BulkUploadResultObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaBulkUploadObjectType&#x60; | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**ErrorType** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the result | [optional] [readonly] 
**LineIndex** | Pointer to **int32** | The index of the line in the CSV | [optional] 
**ObjectErrorDescription** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectStatus** | Pointer to **int32** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerData** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**PluginsData** | Pointer to [**[]KalturaBulkUploadPluginData**](KalturaBulkUploadPluginData.md) |  | [optional] 
**RowData** | Pointer to **string** | The data as recieved in the csv | [optional] 
**Status** | Pointer to **string** | Enum Type: &#x60;KalturaBulkUploadResultStatus&#x60; | [optional] 

## Methods

### NewKalturaBulkUploadResult

`func NewKalturaBulkUploadResult() *KalturaBulkUploadResult`

NewKalturaBulkUploadResult instantiates a new KalturaBulkUploadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBulkUploadResultWithDefaults

`func NewKalturaBulkUploadResultWithDefaults() *KalturaBulkUploadResult`

NewKalturaBulkUploadResultWithDefaults instantiates a new KalturaBulkUploadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *KalturaBulkUploadResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *KalturaBulkUploadResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *KalturaBulkUploadResult) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *KalturaBulkUploadResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBulkUploadJobId

`func (o *KalturaBulkUploadResult) GetBulkUploadJobId() int32`

GetBulkUploadJobId returns the BulkUploadJobId field if non-nil, zero value otherwise.

### GetBulkUploadJobIdOk

`func (o *KalturaBulkUploadResult) GetBulkUploadJobIdOk() (*int32, bool)`

GetBulkUploadJobIdOk returns a tuple with the BulkUploadJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadJobId

`func (o *KalturaBulkUploadResult) SetBulkUploadJobId(v int32)`

SetBulkUploadJobId sets BulkUploadJobId field to given value.

### HasBulkUploadJobId

`func (o *KalturaBulkUploadResult) HasBulkUploadJobId() bool`

HasBulkUploadJobId returns a boolean if a field has been set.

### GetBulkUploadResultObjectType

`func (o *KalturaBulkUploadResult) GetBulkUploadResultObjectType() string`

GetBulkUploadResultObjectType returns the BulkUploadResultObjectType field if non-nil, zero value otherwise.

### GetBulkUploadResultObjectTypeOk

`func (o *KalturaBulkUploadResult) GetBulkUploadResultObjectTypeOk() (*string, bool)`

GetBulkUploadResultObjectTypeOk returns a tuple with the BulkUploadResultObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadResultObjectType

`func (o *KalturaBulkUploadResult) SetBulkUploadResultObjectType(v string)`

SetBulkUploadResultObjectType sets BulkUploadResultObjectType field to given value.

### HasBulkUploadResultObjectType

`func (o *KalturaBulkUploadResult) HasBulkUploadResultObjectType() bool`

HasBulkUploadResultObjectType returns a boolean if a field has been set.

### GetErrorCode

`func (o *KalturaBulkUploadResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *KalturaBulkUploadResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *KalturaBulkUploadResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *KalturaBulkUploadResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaBulkUploadResult) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaBulkUploadResult) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaBulkUploadResult) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaBulkUploadResult) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorType

`func (o *KalturaBulkUploadResult) GetErrorType() int32`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *KalturaBulkUploadResult) GetErrorTypeOk() (*int32, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *KalturaBulkUploadResult) SetErrorType(v int32)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *KalturaBulkUploadResult) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetId

`func (o *KalturaBulkUploadResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBulkUploadResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBulkUploadResult) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBulkUploadResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineIndex

`func (o *KalturaBulkUploadResult) GetLineIndex() int32`

GetLineIndex returns the LineIndex field if non-nil, zero value otherwise.

### GetLineIndexOk

`func (o *KalturaBulkUploadResult) GetLineIndexOk() (*int32, bool)`

GetLineIndexOk returns a tuple with the LineIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIndex

`func (o *KalturaBulkUploadResult) SetLineIndex(v int32)`

SetLineIndex sets LineIndex field to given value.

### HasLineIndex

`func (o *KalturaBulkUploadResult) HasLineIndex() bool`

HasLineIndex returns a boolean if a field has been set.

### GetObjectErrorDescription

`func (o *KalturaBulkUploadResult) GetObjectErrorDescription() string`

GetObjectErrorDescription returns the ObjectErrorDescription field if non-nil, zero value otherwise.

### GetObjectErrorDescriptionOk

`func (o *KalturaBulkUploadResult) GetObjectErrorDescriptionOk() (*string, bool)`

GetObjectErrorDescriptionOk returns a tuple with the ObjectErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectErrorDescription

`func (o *KalturaBulkUploadResult) SetObjectErrorDescription(v string)`

SetObjectErrorDescription sets ObjectErrorDescription field to given value.

### HasObjectErrorDescription

`func (o *KalturaBulkUploadResult) HasObjectErrorDescription() bool`

HasObjectErrorDescription returns a boolean if a field has been set.

### GetObjectId

`func (o *KalturaBulkUploadResult) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *KalturaBulkUploadResult) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *KalturaBulkUploadResult) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *KalturaBulkUploadResult) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectStatus

`func (o *KalturaBulkUploadResult) GetObjectStatus() int32`

GetObjectStatus returns the ObjectStatus field if non-nil, zero value otherwise.

### GetObjectStatusOk

`func (o *KalturaBulkUploadResult) GetObjectStatusOk() (*int32, bool)`

GetObjectStatusOk returns a tuple with the ObjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStatus

`func (o *KalturaBulkUploadResult) SetObjectStatus(v int32)`

SetObjectStatus sets ObjectStatus field to given value.

### HasObjectStatus

`func (o *KalturaBulkUploadResult) HasObjectStatus() bool`

HasObjectStatus returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaBulkUploadResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaBulkUploadResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaBulkUploadResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaBulkUploadResult) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaBulkUploadResult) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaBulkUploadResult) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaBulkUploadResult) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaBulkUploadResult) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaBulkUploadResult) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaBulkUploadResult) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaBulkUploadResult) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaBulkUploadResult) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPluginsData

`func (o *KalturaBulkUploadResult) GetPluginsData() []KalturaBulkUploadPluginData`

GetPluginsData returns the PluginsData field if non-nil, zero value otherwise.

### GetPluginsDataOk

`func (o *KalturaBulkUploadResult) GetPluginsDataOk() (*[]KalturaBulkUploadPluginData, bool)`

GetPluginsDataOk returns a tuple with the PluginsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginsData

`func (o *KalturaBulkUploadResult) SetPluginsData(v []KalturaBulkUploadPluginData)`

SetPluginsData sets PluginsData field to given value.

### HasPluginsData

`func (o *KalturaBulkUploadResult) HasPluginsData() bool`

HasPluginsData returns a boolean if a field has been set.

### GetRowData

`func (o *KalturaBulkUploadResult) GetRowData() string`

GetRowData returns the RowData field if non-nil, zero value otherwise.

### GetRowDataOk

`func (o *KalturaBulkUploadResult) GetRowDataOk() (*string, bool)`

GetRowDataOk returns a tuple with the RowData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowData

`func (o *KalturaBulkUploadResult) SetRowData(v string)`

SetRowData sets RowData field to given value.

### HasRowData

`func (o *KalturaBulkUploadResult) HasRowData() bool`

HasRowData returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBulkUploadResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBulkUploadResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBulkUploadResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBulkUploadResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


