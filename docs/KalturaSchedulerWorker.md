# KalturaSchedulerWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgWait** | Pointer to **int32** | Avarage time between creation and queue time | [optional] 
**AvgWork** | Pointer to **int32** | Avarage time between queue time end finish time | [optional] 
**Configs** | Pointer to [**[]KalturaSchedulerConfig**](KalturaSchedulerConfig.md) |  | [optional] 
**ConfiguredId** | Pointer to **int32** | The id as configured in the batch config | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Worker | [optional] [readonly] 
**LastStatus** | Pointer to **int32** | last status time | [optional] 
**LastStatusStr** | Pointer to **string** | last status formated | [optional] 
**LockedJobs** | Pointer to [**[]KalturaBatchJob**](KalturaBatchJob.md) |  | [optional] 
**Name** | Pointer to **string** | The scheduler name | [optional] 
**SchedulerConfiguredId** | Pointer to **int32** | The id of the scheduler as configured in the batch config | [optional] 
**SchedulerId** | Pointer to **int32** | The id of the Scheduler | [optional] 
**Statuses** | Pointer to [**[]KalturaSchedulerStatus**](KalturaSchedulerStatus.md) |  | [optional] 
**Type** | Pointer to **string** | Enum Type: &#x60;KalturaBatchJobType&#x60;  The worker type | [optional] 
**TypeName** | Pointer to **string** | The friendly name of the type | [optional] 

## Methods

### NewKalturaSchedulerWorker

`func NewKalturaSchedulerWorker() *KalturaSchedulerWorker`

NewKalturaSchedulerWorker instantiates a new KalturaSchedulerWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSchedulerWorkerWithDefaults

`func NewKalturaSchedulerWorkerWithDefaults() *KalturaSchedulerWorker`

NewKalturaSchedulerWorkerWithDefaults instantiates a new KalturaSchedulerWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgWait

`func (o *KalturaSchedulerWorker) GetAvgWait() int32`

GetAvgWait returns the AvgWait field if non-nil, zero value otherwise.

### GetAvgWaitOk

`func (o *KalturaSchedulerWorker) GetAvgWaitOk() (*int32, bool)`

GetAvgWaitOk returns a tuple with the AvgWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgWait

`func (o *KalturaSchedulerWorker) SetAvgWait(v int32)`

SetAvgWait sets AvgWait field to given value.

### HasAvgWait

`func (o *KalturaSchedulerWorker) HasAvgWait() bool`

HasAvgWait returns a boolean if a field has been set.

### GetAvgWork

`func (o *KalturaSchedulerWorker) GetAvgWork() int32`

GetAvgWork returns the AvgWork field if non-nil, zero value otherwise.

### GetAvgWorkOk

`func (o *KalturaSchedulerWorker) GetAvgWorkOk() (*int32, bool)`

GetAvgWorkOk returns a tuple with the AvgWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgWork

`func (o *KalturaSchedulerWorker) SetAvgWork(v int32)`

SetAvgWork sets AvgWork field to given value.

### HasAvgWork

`func (o *KalturaSchedulerWorker) HasAvgWork() bool`

HasAvgWork returns a boolean if a field has been set.

### GetConfigs

`func (o *KalturaSchedulerWorker) GetConfigs() []KalturaSchedulerConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *KalturaSchedulerWorker) GetConfigsOk() (*[]KalturaSchedulerConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *KalturaSchedulerWorker) SetConfigs(v []KalturaSchedulerConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *KalturaSchedulerWorker) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetConfiguredId

`func (o *KalturaSchedulerWorker) GetConfiguredId() int32`

GetConfiguredId returns the ConfiguredId field if non-nil, zero value otherwise.

### GetConfiguredIdOk

`func (o *KalturaSchedulerWorker) GetConfiguredIdOk() (*int32, bool)`

GetConfiguredIdOk returns a tuple with the ConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredId

`func (o *KalturaSchedulerWorker) SetConfiguredId(v int32)`

SetConfiguredId sets ConfiguredId field to given value.

### HasConfiguredId

`func (o *KalturaSchedulerWorker) HasConfiguredId() bool`

HasConfiguredId returns a boolean if a field has been set.

### GetId

`func (o *KalturaSchedulerWorker) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaSchedulerWorker) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaSchedulerWorker) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaSchedulerWorker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastStatus

`func (o *KalturaSchedulerWorker) GetLastStatus() int32`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *KalturaSchedulerWorker) GetLastStatusOk() (*int32, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *KalturaSchedulerWorker) SetLastStatus(v int32)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *KalturaSchedulerWorker) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastStatusStr

`func (o *KalturaSchedulerWorker) GetLastStatusStr() string`

GetLastStatusStr returns the LastStatusStr field if non-nil, zero value otherwise.

### GetLastStatusStrOk

`func (o *KalturaSchedulerWorker) GetLastStatusStrOk() (*string, bool)`

GetLastStatusStrOk returns a tuple with the LastStatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusStr

`func (o *KalturaSchedulerWorker) SetLastStatusStr(v string)`

SetLastStatusStr sets LastStatusStr field to given value.

### HasLastStatusStr

`func (o *KalturaSchedulerWorker) HasLastStatusStr() bool`

HasLastStatusStr returns a boolean if a field has been set.

### GetLockedJobs

`func (o *KalturaSchedulerWorker) GetLockedJobs() []KalturaBatchJob`

GetLockedJobs returns the LockedJobs field if non-nil, zero value otherwise.

### GetLockedJobsOk

`func (o *KalturaSchedulerWorker) GetLockedJobsOk() (*[]KalturaBatchJob, bool)`

GetLockedJobsOk returns a tuple with the LockedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedJobs

`func (o *KalturaSchedulerWorker) SetLockedJobs(v []KalturaBatchJob)`

SetLockedJobs sets LockedJobs field to given value.

### HasLockedJobs

`func (o *KalturaSchedulerWorker) HasLockedJobs() bool`

HasLockedJobs returns a boolean if a field has been set.

### GetName

`func (o *KalturaSchedulerWorker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaSchedulerWorker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaSchedulerWorker) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaSchedulerWorker) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedulerConfiguredId

`func (o *KalturaSchedulerWorker) GetSchedulerConfiguredId() int32`

GetSchedulerConfiguredId returns the SchedulerConfiguredId field if non-nil, zero value otherwise.

### GetSchedulerConfiguredIdOk

`func (o *KalturaSchedulerWorker) GetSchedulerConfiguredIdOk() (*int32, bool)`

GetSchedulerConfiguredIdOk returns a tuple with the SchedulerConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerConfiguredId

`func (o *KalturaSchedulerWorker) SetSchedulerConfiguredId(v int32)`

SetSchedulerConfiguredId sets SchedulerConfiguredId field to given value.

### HasSchedulerConfiguredId

`func (o *KalturaSchedulerWorker) HasSchedulerConfiguredId() bool`

HasSchedulerConfiguredId returns a boolean if a field has been set.

### GetSchedulerId

`func (o *KalturaSchedulerWorker) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *KalturaSchedulerWorker) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *KalturaSchedulerWorker) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.

### HasSchedulerId

`func (o *KalturaSchedulerWorker) HasSchedulerId() bool`

HasSchedulerId returns a boolean if a field has been set.

### GetStatuses

`func (o *KalturaSchedulerWorker) GetStatuses() []KalturaSchedulerStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *KalturaSchedulerWorker) GetStatusesOk() (*[]KalturaSchedulerStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *KalturaSchedulerWorker) SetStatuses(v []KalturaSchedulerStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *KalturaSchedulerWorker) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetType

`func (o *KalturaSchedulerWorker) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaSchedulerWorker) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaSchedulerWorker) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaSchedulerWorker) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *KalturaSchedulerWorker) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *KalturaSchedulerWorker) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *KalturaSchedulerWorker) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *KalturaSchedulerWorker) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


