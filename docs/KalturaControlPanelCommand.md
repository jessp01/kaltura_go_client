# KalturaControlPanelCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchIndex** | Pointer to **int32** | The index of the batch process that the command refers to | [optional] 
**Cause** | Pointer to **string** | The reason for the command | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | Creator name | [optional] 
**CreatedById** | Pointer to **int32** | Creator id | [optional] 
**Description** | Pointer to **string** | Command description | [optional] 
**ErrorDescription** | Pointer to **string** | Error description | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Category | [optional] [readonly] 
**SchedulerId** | Pointer to **int32** | The id of the scheduler that the command refers to | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaControlPanelCommandStatus&#x60;  The command status | [optional] 
**TargetType** | Pointer to **int32** | Enum Type: &#x60;KalturaControlPanelCommandTargetType&#x60;  The command target type - data center / scheduler / job / job type | [optional] 
**Type** | Pointer to **int32** | Enum Type: &#x60;KalturaControlPanelCommandType&#x60;  The command type - stop / start / config | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UpdatedBy** | Pointer to **string** | Updater name | [optional] 
**WorkerConfiguredId** | Pointer to **int32** | The id of the scheduler worker as configured in the ini file | [optional] 
**WorkerId** | Pointer to **int32** | The id of the scheduler worker that the command refers to | [optional] 
**WorkerName** | Pointer to **int32** | The name of the scheduler worker that the command refers to | [optional] 

## Methods

### NewKalturaControlPanelCommand

`func NewKalturaControlPanelCommand() *KalturaControlPanelCommand`

NewKalturaControlPanelCommand instantiates a new KalturaControlPanelCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaControlPanelCommandWithDefaults

`func NewKalturaControlPanelCommandWithDefaults() *KalturaControlPanelCommand`

NewKalturaControlPanelCommandWithDefaults instantiates a new KalturaControlPanelCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchIndex

`func (o *KalturaControlPanelCommand) GetBatchIndex() int32`

GetBatchIndex returns the BatchIndex field if non-nil, zero value otherwise.

### GetBatchIndexOk

`func (o *KalturaControlPanelCommand) GetBatchIndexOk() (*int32, bool)`

GetBatchIndexOk returns a tuple with the BatchIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIndex

`func (o *KalturaControlPanelCommand) SetBatchIndex(v int32)`

SetBatchIndex sets BatchIndex field to given value.

### HasBatchIndex

`func (o *KalturaControlPanelCommand) HasBatchIndex() bool`

HasBatchIndex returns a boolean if a field has been set.

### GetCause

`func (o *KalturaControlPanelCommand) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *KalturaControlPanelCommand) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *KalturaControlPanelCommand) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *KalturaControlPanelCommand) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaControlPanelCommand) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaControlPanelCommand) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaControlPanelCommand) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaControlPanelCommand) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *KalturaControlPanelCommand) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KalturaControlPanelCommand) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KalturaControlPanelCommand) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *KalturaControlPanelCommand) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *KalturaControlPanelCommand) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *KalturaControlPanelCommand) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *KalturaControlPanelCommand) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *KalturaControlPanelCommand) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaControlPanelCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaControlPanelCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaControlPanelCommand) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaControlPanelCommand) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaControlPanelCommand) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaControlPanelCommand) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaControlPanelCommand) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaControlPanelCommand) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaControlPanelCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaControlPanelCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaControlPanelCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaControlPanelCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchedulerId

`func (o *KalturaControlPanelCommand) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *KalturaControlPanelCommand) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *KalturaControlPanelCommand) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.

### HasSchedulerId

`func (o *KalturaControlPanelCommand) HasSchedulerId() bool`

HasSchedulerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaControlPanelCommand) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaControlPanelCommand) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaControlPanelCommand) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaControlPanelCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetType

`func (o *KalturaControlPanelCommand) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *KalturaControlPanelCommand) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *KalturaControlPanelCommand) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *KalturaControlPanelCommand) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetType

`func (o *KalturaControlPanelCommand) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaControlPanelCommand) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaControlPanelCommand) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaControlPanelCommand) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaControlPanelCommand) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaControlPanelCommand) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaControlPanelCommand) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaControlPanelCommand) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *KalturaControlPanelCommand) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *KalturaControlPanelCommand) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *KalturaControlPanelCommand) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *KalturaControlPanelCommand) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWorkerConfiguredId

`func (o *KalturaControlPanelCommand) GetWorkerConfiguredId() int32`

GetWorkerConfiguredId returns the WorkerConfiguredId field if non-nil, zero value otherwise.

### GetWorkerConfiguredIdOk

`func (o *KalturaControlPanelCommand) GetWorkerConfiguredIdOk() (*int32, bool)`

GetWorkerConfiguredIdOk returns a tuple with the WorkerConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerConfiguredId

`func (o *KalturaControlPanelCommand) SetWorkerConfiguredId(v int32)`

SetWorkerConfiguredId sets WorkerConfiguredId field to given value.

### HasWorkerConfiguredId

`func (o *KalturaControlPanelCommand) HasWorkerConfiguredId() bool`

HasWorkerConfiguredId returns a boolean if a field has been set.

### GetWorkerId

`func (o *KalturaControlPanelCommand) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *KalturaControlPanelCommand) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *KalturaControlPanelCommand) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *KalturaControlPanelCommand) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkerName

`func (o *KalturaControlPanelCommand) GetWorkerName() int32`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *KalturaControlPanelCommand) GetWorkerNameOk() (*int32, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *KalturaControlPanelCommand) SetWorkerName(v int32)`

SetWorkerName sets WorkerName field to given value.

### HasWorkerName

`func (o *KalturaControlPanelCommand) HasWorkerName() bool`

HasWorkerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


