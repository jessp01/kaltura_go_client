# BatchResetJobExecutionAttemptsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**JobType** | **string** |  | 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 

## Methods

### NewBatchResetJobExecutionAttemptsRequest

`func NewBatchResetJobExecutionAttemptsRequest(id int32, jobType string, lockKey KalturaExclusiveLockKey, ) *BatchResetJobExecutionAttemptsRequest`

NewBatchResetJobExecutionAttemptsRequest instantiates a new BatchResetJobExecutionAttemptsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResetJobExecutionAttemptsRequestWithDefaults

`func NewBatchResetJobExecutionAttemptsRequestWithDefaults() *BatchResetJobExecutionAttemptsRequest`

NewBatchResetJobExecutionAttemptsRequestWithDefaults instantiates a new BatchResetJobExecutionAttemptsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchResetJobExecutionAttemptsRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchResetJobExecutionAttemptsRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchResetJobExecutionAttemptsRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetJobType

`func (o *BatchResetJobExecutionAttemptsRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BatchResetJobExecutionAttemptsRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BatchResetJobExecutionAttemptsRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetLockKey

`func (o *BatchResetJobExecutionAttemptsRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *BatchResetJobExecutionAttemptsRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *BatchResetJobExecutionAttemptsRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


