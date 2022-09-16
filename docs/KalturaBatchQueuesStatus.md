# KalturaBatchQueuesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | Pointer to **string** | Enum Type: &#x60;KalturaBatchJobType&#x60; | [optional] 
**Size** | Pointer to **int32** | The size of the queue | [optional] 
**TypeName** | Pointer to **string** | The friendly name of the type | [optional] 
**WaitTime** | Pointer to **int32** | The avarage wait time | [optional] 
**WorkerId** | Pointer to **int32** | The worker configured id | [optional] 

## Methods

### NewKalturaBatchQueuesStatus

`func NewKalturaBatchQueuesStatus() *KalturaBatchQueuesStatus`

NewKalturaBatchQueuesStatus instantiates a new KalturaBatchQueuesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBatchQueuesStatusWithDefaults

`func NewKalturaBatchQueuesStatusWithDefaults() *KalturaBatchQueuesStatus`

NewKalturaBatchQueuesStatusWithDefaults instantiates a new KalturaBatchQueuesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *KalturaBatchQueuesStatus) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KalturaBatchQueuesStatus) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KalturaBatchQueuesStatus) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KalturaBatchQueuesStatus) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetSize

`func (o *KalturaBatchQueuesStatus) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KalturaBatchQueuesStatus) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KalturaBatchQueuesStatus) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *KalturaBatchQueuesStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTypeName

`func (o *KalturaBatchQueuesStatus) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *KalturaBatchQueuesStatus) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *KalturaBatchQueuesStatus) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *KalturaBatchQueuesStatus) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetWaitTime

`func (o *KalturaBatchQueuesStatus) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *KalturaBatchQueuesStatus) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *KalturaBatchQueuesStatus) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *KalturaBatchQueuesStatus) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.

### GetWorkerId

`func (o *KalturaBatchQueuesStatus) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *KalturaBatchQueuesStatus) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *KalturaBatchQueuesStatus) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *KalturaBatchQueuesStatus) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


