# KalturaWorkerQueueFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaBatchJobFilter**](KalturaBatchJobFilter.md) |  | [optional] 
**JobType** | Pointer to **string** | Enum Type: &#x60;KalturaBatchJobType&#x60; | [optional] 
**SchedulerId** | Pointer to **int32** |  | [optional] 
**WorkerId** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaWorkerQueueFilter

`func NewKalturaWorkerQueueFilter() *KalturaWorkerQueueFilter`

NewKalturaWorkerQueueFilter instantiates a new KalturaWorkerQueueFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaWorkerQueueFilterWithDefaults

`func NewKalturaWorkerQueueFilterWithDefaults() *KalturaWorkerQueueFilter`

NewKalturaWorkerQueueFilterWithDefaults instantiates a new KalturaWorkerQueueFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *KalturaWorkerQueueFilter) GetFilter() KalturaBatchJobFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *KalturaWorkerQueueFilter) GetFilterOk() (*KalturaBatchJobFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *KalturaWorkerQueueFilter) SetFilter(v KalturaBatchJobFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *KalturaWorkerQueueFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetJobType

`func (o *KalturaWorkerQueueFilter) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KalturaWorkerQueueFilter) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KalturaWorkerQueueFilter) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KalturaWorkerQueueFilter) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetSchedulerId

`func (o *KalturaWorkerQueueFilter) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *KalturaWorkerQueueFilter) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *KalturaWorkerQueueFilter) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.

### HasSchedulerId

`func (o *KalturaWorkerQueueFilter) HasSchedulerId() bool`

HasSchedulerId returns a boolean if a field has been set.

### GetWorkerId

`func (o *KalturaWorkerQueueFilter) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *KalturaWorkerQueueFilter) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *KalturaWorkerQueueFilter) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *KalturaWorkerQueueFilter) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


