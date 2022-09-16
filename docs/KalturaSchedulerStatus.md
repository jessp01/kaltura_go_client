# KalturaSchedulerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Category | [optional] [readonly] 
**SchedulerConfiguredId** | Pointer to **int32** | The configured id of the scheduler | [optional] 
**SchedulerId** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the scheduler | [optional] [readonly] 
**Type** | Pointer to **int32** | Enum Type: &#x60;KalturaSchedulerStatusType&#x60;  The status type | [optional] 
**Value** | Pointer to **int32** | The status value | [optional] 
**WorkerConfiguredId** | Pointer to **int32** | The configured id of the job worker | [optional] 
**WorkerId** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the worker | [optional] [readonly] 
**WorkerType** | Pointer to **string** | Enum Type: &#x60;KalturaBatchJobType&#x60;  The type of the job worker. | [optional] 

## Methods

### NewKalturaSchedulerStatus

`func NewKalturaSchedulerStatus() *KalturaSchedulerStatus`

NewKalturaSchedulerStatus instantiates a new KalturaSchedulerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSchedulerStatusWithDefaults

`func NewKalturaSchedulerStatusWithDefaults() *KalturaSchedulerStatus`

NewKalturaSchedulerStatusWithDefaults instantiates a new KalturaSchedulerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KalturaSchedulerStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaSchedulerStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaSchedulerStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaSchedulerStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchedulerConfiguredId

`func (o *KalturaSchedulerStatus) GetSchedulerConfiguredId() int32`

GetSchedulerConfiguredId returns the SchedulerConfiguredId field if non-nil, zero value otherwise.

### GetSchedulerConfiguredIdOk

`func (o *KalturaSchedulerStatus) GetSchedulerConfiguredIdOk() (*int32, bool)`

GetSchedulerConfiguredIdOk returns a tuple with the SchedulerConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerConfiguredId

`func (o *KalturaSchedulerStatus) SetSchedulerConfiguredId(v int32)`

SetSchedulerConfiguredId sets SchedulerConfiguredId field to given value.

### HasSchedulerConfiguredId

`func (o *KalturaSchedulerStatus) HasSchedulerConfiguredId() bool`

HasSchedulerConfiguredId returns a boolean if a field has been set.

### GetSchedulerId

`func (o *KalturaSchedulerStatus) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *KalturaSchedulerStatus) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *KalturaSchedulerStatus) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.

### HasSchedulerId

`func (o *KalturaSchedulerStatus) HasSchedulerId() bool`

HasSchedulerId returns a boolean if a field has been set.

### GetType

`func (o *KalturaSchedulerStatus) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaSchedulerStatus) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaSchedulerStatus) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaSchedulerStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *KalturaSchedulerStatus) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KalturaSchedulerStatus) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KalturaSchedulerStatus) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *KalturaSchedulerStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWorkerConfiguredId

`func (o *KalturaSchedulerStatus) GetWorkerConfiguredId() int32`

GetWorkerConfiguredId returns the WorkerConfiguredId field if non-nil, zero value otherwise.

### GetWorkerConfiguredIdOk

`func (o *KalturaSchedulerStatus) GetWorkerConfiguredIdOk() (*int32, bool)`

GetWorkerConfiguredIdOk returns a tuple with the WorkerConfiguredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerConfiguredId

`func (o *KalturaSchedulerStatus) SetWorkerConfiguredId(v int32)`

SetWorkerConfiguredId sets WorkerConfiguredId field to given value.

### HasWorkerConfiguredId

`func (o *KalturaSchedulerStatus) HasWorkerConfiguredId() bool`

HasWorkerConfiguredId returns a boolean if a field has been set.

### GetWorkerId

`func (o *KalturaSchedulerStatus) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *KalturaSchedulerStatus) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *KalturaSchedulerStatus) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *KalturaSchedulerStatus) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkerType

`func (o *KalturaSchedulerStatus) GetWorkerType() string`

GetWorkerType returns the WorkerType field if non-nil, zero value otherwise.

### GetWorkerTypeOk

`func (o *KalturaSchedulerStatus) GetWorkerTypeOk() (*string, bool)`

GetWorkerTypeOk returns a tuple with the WorkerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerType

`func (o *KalturaSchedulerStatus) SetWorkerType(v string)`

SetWorkerType sets WorkerType field to given value.

### HasWorkerType

`func (o *KalturaSchedulerStatus) HasWorkerType() bool`

HasWorkerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


