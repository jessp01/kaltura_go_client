# BatchUpdateExclusiveConvertCollectionJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlavorsData** | Pointer to [**[]KalturaConvertCollectionFlavorData**](KalturaConvertCollectionFlavorData.md) |  | [optional] 
**Id** | **int32** |  | 
**Job** | [**KalturaBatchJob**](KalturaBatchJob.md) |  | 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 

## Methods

### NewBatchUpdateExclusiveConvertCollectionJobRequest

`func NewBatchUpdateExclusiveConvertCollectionJobRequest(id int32, job KalturaBatchJob, lockKey KalturaExclusiveLockKey, ) *BatchUpdateExclusiveConvertCollectionJobRequest`

NewBatchUpdateExclusiveConvertCollectionJobRequest instantiates a new BatchUpdateExclusiveConvertCollectionJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateExclusiveConvertCollectionJobRequestWithDefaults

`func NewBatchUpdateExclusiveConvertCollectionJobRequestWithDefaults() *BatchUpdateExclusiveConvertCollectionJobRequest`

NewBatchUpdateExclusiveConvertCollectionJobRequestWithDefaults instantiates a new BatchUpdateExclusiveConvertCollectionJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavorsData

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetFlavorsData() []KalturaConvertCollectionFlavorData`

GetFlavorsData returns the FlavorsData field if non-nil, zero value otherwise.

### GetFlavorsDataOk

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetFlavorsDataOk() (*[]KalturaConvertCollectionFlavorData, bool)`

GetFlavorsDataOk returns a tuple with the FlavorsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorsData

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) SetFlavorsData(v []KalturaConvertCollectionFlavorData)`

SetFlavorsData sets FlavorsData field to given value.

### HasFlavorsData

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) HasFlavorsData() bool`

HasFlavorsData returns a boolean if a field has been set.

### GetId

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetJob

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetJob() KalturaBatchJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetJobOk() (*KalturaBatchJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) SetJob(v KalturaBatchJob)`

SetJob sets Job field to given value.


### GetLockKey

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *BatchUpdateExclusiveConvertCollectionJobRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


