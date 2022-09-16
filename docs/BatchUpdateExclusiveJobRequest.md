# BatchUpdateExclusiveJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Job** | [**KalturaBatchJob**](KalturaBatchJob.md) |  | 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 

## Methods

### NewBatchUpdateExclusiveJobRequest

`func NewBatchUpdateExclusiveJobRequest(id int32, job KalturaBatchJob, lockKey KalturaExclusiveLockKey, ) *BatchUpdateExclusiveJobRequest`

NewBatchUpdateExclusiveJobRequest instantiates a new BatchUpdateExclusiveJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateExclusiveJobRequestWithDefaults

`func NewBatchUpdateExclusiveJobRequestWithDefaults() *BatchUpdateExclusiveJobRequest`

NewBatchUpdateExclusiveJobRequestWithDefaults instantiates a new BatchUpdateExclusiveJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchUpdateExclusiveJobRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchUpdateExclusiveJobRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchUpdateExclusiveJobRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetJob

`func (o *BatchUpdateExclusiveJobRequest) GetJob() KalturaBatchJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *BatchUpdateExclusiveJobRequest) GetJobOk() (*KalturaBatchJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *BatchUpdateExclusiveJobRequest) SetJob(v KalturaBatchJob)`

SetJob sets Job field to given value.


### GetLockKey

`func (o *BatchUpdateExclusiveJobRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *BatchUpdateExclusiveJobRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *BatchUpdateExclusiveJobRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


