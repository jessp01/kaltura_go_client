# BatchGetQueueSizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerQueueFilter** | [**KalturaWorkerQueueFilter**](KalturaWorkerQueueFilter.md) |  | 

## Methods

### NewBatchGetQueueSizeRequest

`func NewBatchGetQueueSizeRequest(workerQueueFilter KalturaWorkerQueueFilter, ) *BatchGetQueueSizeRequest`

NewBatchGetQueueSizeRequest instantiates a new BatchGetQueueSizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGetQueueSizeRequestWithDefaults

`func NewBatchGetQueueSizeRequestWithDefaults() *BatchGetQueueSizeRequest`

NewBatchGetQueueSizeRequestWithDefaults instantiates a new BatchGetQueueSizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerQueueFilter

`func (o *BatchGetQueueSizeRequest) GetWorkerQueueFilter() KalturaWorkerQueueFilter`

GetWorkerQueueFilter returns the WorkerQueueFilter field if non-nil, zero value otherwise.

### GetWorkerQueueFilterOk

`func (o *BatchGetQueueSizeRequest) GetWorkerQueueFilterOk() (*KalturaWorkerQueueFilter, bool)`

GetWorkerQueueFilterOk returns a tuple with the WorkerQueueFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerQueueFilter

`func (o *BatchGetQueueSizeRequest) SetWorkerQueueFilter(v KalturaWorkerQueueFilter)`

SetWorkerQueueFilter sets WorkerQueueFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


