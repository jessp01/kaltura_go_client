# KalturaScheduler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**[]KalturaSchedulerConfig**](KalturaSchedulerConfig.md) |  | [optional] 
**ConfiguredId** | Pointer to **int32** | The id as configured in the batch config | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  creation time | [optional] [readonly] 
**Host** | Pointer to **string** | The host name | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Scheduler | [optional] [readonly] 
**LastStatus** | Pointer to **int32** | &#x60;readOnly&#x60;  last status time | [optional] [readonly] 
**LastStatusStr** | Pointer to **string** | &#x60;readOnly&#x60;  last status formated | [optional] [readonly] 
**Name** | Pointer to **string** | The scheduler name | [optional] 
**Statuses** | Pointer to [**[]KalturaSchedulerStatus**](KalturaSchedulerStatus.md) |  | [optional] 
**Workers** | Pointer to [**[]KalturaSchedulerWorker**](KalturaSchedulerWorker.md) |  | [optional] 

## Methods

### NewKalturaScheduler

`func NewKalturaScheduler() *KalturaScheduler`

NewKalturaScheduler instantiates a new KalturaScheduler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSchedulerWithDefaults

`func NewKalturaSchedulerWithDefaults() *KalturaScheduler`

NewKalturaSchedulerWithDefaults instantiates a new KalturaScheduler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *KalturaScheduler) GetConfigs() []KalturaSchedulerConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *KalturaScheduler) GetConfigsOk() (*[]KalturaSchedulerConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *KalturaScheduler) SetConfigs(v []KalturaSchedulerConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *KalturaScheduler) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetConfiguredId

`func (o *KalturaScheduler) GetConfiguredId() int32`

GetConfiguredId returns the ConfiguredId field if non-nil, zero value otherwise.

### GetConfiguredIdOk

`func (o *KalturaScheduler) GetConfiguredIdOk() (*int32, bool)`

GetConfiguredIdOk returns a tuple with the ConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredId

`func (o *KalturaScheduler) SetConfiguredId(v int32)`

SetConfiguredId sets ConfiguredId field to given value.

### HasConfiguredId

`func (o *KalturaScheduler) HasConfiguredId() bool`

HasConfiguredId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaScheduler) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaScheduler) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaScheduler) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaScheduler) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *KalturaScheduler) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KalturaScheduler) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KalturaScheduler) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KalturaScheduler) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *KalturaScheduler) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaScheduler) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaScheduler) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaScheduler) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastStatus

`func (o *KalturaScheduler) GetLastStatus() int32`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *KalturaScheduler) GetLastStatusOk() (*int32, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *KalturaScheduler) SetLastStatus(v int32)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *KalturaScheduler) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastStatusStr

`func (o *KalturaScheduler) GetLastStatusStr() string`

GetLastStatusStr returns the LastStatusStr field if non-nil, zero value otherwise.

### GetLastStatusStrOk

`func (o *KalturaScheduler) GetLastStatusStrOk() (*string, bool)`

GetLastStatusStrOk returns a tuple with the LastStatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusStr

`func (o *KalturaScheduler) SetLastStatusStr(v string)`

SetLastStatusStr sets LastStatusStr field to given value.

### HasLastStatusStr

`func (o *KalturaScheduler) HasLastStatusStr() bool`

HasLastStatusStr returns a boolean if a field has been set.

### GetName

`func (o *KalturaScheduler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaScheduler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaScheduler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaScheduler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatuses

`func (o *KalturaScheduler) GetStatuses() []KalturaSchedulerStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *KalturaScheduler) GetStatusesOk() (*[]KalturaSchedulerStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *KalturaScheduler) SetStatuses(v []KalturaSchedulerStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *KalturaScheduler) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkers

`func (o *KalturaScheduler) GetWorkers() []KalturaSchedulerWorker`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *KalturaScheduler) GetWorkersOk() (*[]KalturaSchedulerWorker, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *KalturaScheduler) SetWorkers(v []KalturaSchedulerWorker)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *KalturaScheduler) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


