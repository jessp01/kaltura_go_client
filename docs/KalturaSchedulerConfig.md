# KalturaSchedulerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandId** | Pointer to **string** | Id of the control panel command that created this config item | [optional] 
**CommandStatus** | Pointer to **string** | The status of the control panel command | [optional] 
**CreatedBy** | Pointer to **string** | Creator name | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Category | [optional] [readonly] 
**SchedulerConfiguredId** | Pointer to **int32** | The configured id of the scheduler | [optional] 
**SchedulerId** | Pointer to **int32** | The id of the scheduler | [optional] 
**SchedulerName** | Pointer to **string** | The name of the scheduler | [optional] 
**UpdatedBy** | Pointer to **string** | Updater name | [optional] 
**Value** | Pointer to **string** | The value of the variable | [optional] 
**Variable** | Pointer to **string** | The name of the variable | [optional] 
**VariablePart** | Pointer to **string** | The part of the variable | [optional] 
**WorkerConfiguredId** | Pointer to **int32** | The configured id of the job worker | [optional] 
**WorkerId** | Pointer to **int32** | The id of the job worker | [optional] 
**WorkerName** | Pointer to **string** | The name of the job worker | [optional] 

## Methods

### NewKalturaSchedulerConfig

`func NewKalturaSchedulerConfig() *KalturaSchedulerConfig`

NewKalturaSchedulerConfig instantiates a new KalturaSchedulerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSchedulerConfigWithDefaults

`func NewKalturaSchedulerConfigWithDefaults() *KalturaSchedulerConfig`

NewKalturaSchedulerConfigWithDefaults instantiates a new KalturaSchedulerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandId

`func (o *KalturaSchedulerConfig) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *KalturaSchedulerConfig) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *KalturaSchedulerConfig) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *KalturaSchedulerConfig) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetCommandStatus

`func (o *KalturaSchedulerConfig) GetCommandStatus() string`

GetCommandStatus returns the CommandStatus field if non-nil, zero value otherwise.

### GetCommandStatusOk

`func (o *KalturaSchedulerConfig) GetCommandStatusOk() (*string, bool)`

GetCommandStatusOk returns a tuple with the CommandStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandStatus

`func (o *KalturaSchedulerConfig) SetCommandStatus(v string)`

SetCommandStatus sets CommandStatus field to given value.

### HasCommandStatus

`func (o *KalturaSchedulerConfig) HasCommandStatus() bool`

HasCommandStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *KalturaSchedulerConfig) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KalturaSchedulerConfig) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KalturaSchedulerConfig) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *KalturaSchedulerConfig) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *KalturaSchedulerConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaSchedulerConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaSchedulerConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaSchedulerConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchedulerConfiguredId

`func (o *KalturaSchedulerConfig) GetSchedulerConfiguredId() int32`

GetSchedulerConfiguredId returns the SchedulerConfiguredId field if non-nil, zero value otherwise.

### GetSchedulerConfiguredIdOk

`func (o *KalturaSchedulerConfig) GetSchedulerConfiguredIdOk() (*int32, bool)`

GetSchedulerConfiguredIdOk returns a tuple with the SchedulerConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerConfiguredId

`func (o *KalturaSchedulerConfig) SetSchedulerConfiguredId(v int32)`

SetSchedulerConfiguredId sets SchedulerConfiguredId field to given value.

### HasSchedulerConfiguredId

`func (o *KalturaSchedulerConfig) HasSchedulerConfiguredId() bool`

HasSchedulerConfiguredId returns a boolean if a field has been set.

### GetSchedulerId

`func (o *KalturaSchedulerConfig) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *KalturaSchedulerConfig) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *KalturaSchedulerConfig) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.

### HasSchedulerId

`func (o *KalturaSchedulerConfig) HasSchedulerId() bool`

HasSchedulerId returns a boolean if a field has been set.

### GetSchedulerName

`func (o *KalturaSchedulerConfig) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *KalturaSchedulerConfig) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *KalturaSchedulerConfig) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *KalturaSchedulerConfig) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *KalturaSchedulerConfig) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *KalturaSchedulerConfig) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *KalturaSchedulerConfig) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *KalturaSchedulerConfig) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValue

`func (o *KalturaSchedulerConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KalturaSchedulerConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KalturaSchedulerConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KalturaSchedulerConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVariable

`func (o *KalturaSchedulerConfig) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *KalturaSchedulerConfig) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *KalturaSchedulerConfig) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *KalturaSchedulerConfig) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### GetVariablePart

`func (o *KalturaSchedulerConfig) GetVariablePart() string`

GetVariablePart returns the VariablePart field if non-nil, zero value otherwise.

### GetVariablePartOk

`func (o *KalturaSchedulerConfig) GetVariablePartOk() (*string, bool)`

GetVariablePartOk returns a tuple with the VariablePart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePart

`func (o *KalturaSchedulerConfig) SetVariablePart(v string)`

SetVariablePart sets VariablePart field to given value.

### HasVariablePart

`func (o *KalturaSchedulerConfig) HasVariablePart() bool`

HasVariablePart returns a boolean if a field has been set.

### GetWorkerConfiguredId

`func (o *KalturaSchedulerConfig) GetWorkerConfiguredId() int32`

GetWorkerConfiguredId returns the WorkerConfiguredId field if non-nil, zero value otherwise.

### GetWorkerConfiguredIdOk

`func (o *KalturaSchedulerConfig) GetWorkerConfiguredIdOk() (*int32, bool)`

GetWorkerConfiguredIdOk returns a tuple with the WorkerConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerConfiguredId

`func (o *KalturaSchedulerConfig) SetWorkerConfiguredId(v int32)`

SetWorkerConfiguredId sets WorkerConfiguredId field to given value.

### HasWorkerConfiguredId

`func (o *KalturaSchedulerConfig) HasWorkerConfiguredId() bool`

HasWorkerConfiguredId returns a boolean if a field has been set.

### GetWorkerId

`func (o *KalturaSchedulerConfig) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *KalturaSchedulerConfig) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *KalturaSchedulerConfig) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *KalturaSchedulerConfig) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkerName

`func (o *KalturaSchedulerConfig) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *KalturaSchedulerConfig) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *KalturaSchedulerConfig) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.

### HasWorkerName

`func (o *KalturaSchedulerConfig) HasWorkerName() bool`

HasWorkerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


