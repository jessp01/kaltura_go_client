# BatchcontrolConfigLoadedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigParam** | **string** |  | 
**ConfigParamPart** | Pointer to **string** |  | [optional] 
**ConfigValue** | **string** |  | 
**Scheduler** | [**KalturaScheduler**](KalturaScheduler.md) |  | 
**WorkerConfigId** | Pointer to **int32** |  | [optional] 
**WorkerName** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchcontrolConfigLoadedRequest

`func NewBatchcontrolConfigLoadedRequest(configParam string, configValue string, scheduler KalturaScheduler, ) *BatchcontrolConfigLoadedRequest`

NewBatchcontrolConfigLoadedRequest instantiates a new BatchcontrolConfigLoadedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolConfigLoadedRequestWithDefaults

`func NewBatchcontrolConfigLoadedRequestWithDefaults() *BatchcontrolConfigLoadedRequest`

NewBatchcontrolConfigLoadedRequestWithDefaults instantiates a new BatchcontrolConfigLoadedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigParam

`func (o *BatchcontrolConfigLoadedRequest) GetConfigParam() string`

GetConfigParam returns the ConfigParam field if non-nil, zero value otherwise.

### GetConfigParamOk

`func (o *BatchcontrolConfigLoadedRequest) GetConfigParamOk() (*string, bool)`

GetConfigParamOk returns a tuple with the ConfigParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParam

`func (o *BatchcontrolConfigLoadedRequest) SetConfigParam(v string)`

SetConfigParam sets ConfigParam field to given value.


### GetConfigParamPart

`func (o *BatchcontrolConfigLoadedRequest) GetConfigParamPart() string`

GetConfigParamPart returns the ConfigParamPart field if non-nil, zero value otherwise.

### GetConfigParamPartOk

`func (o *BatchcontrolConfigLoadedRequest) GetConfigParamPartOk() (*string, bool)`

GetConfigParamPartOk returns a tuple with the ConfigParamPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParamPart

`func (o *BatchcontrolConfigLoadedRequest) SetConfigParamPart(v string)`

SetConfigParamPart sets ConfigParamPart field to given value.

### HasConfigParamPart

`func (o *BatchcontrolConfigLoadedRequest) HasConfigParamPart() bool`

HasConfigParamPart returns a boolean if a field has been set.

### GetConfigValue

`func (o *BatchcontrolConfigLoadedRequest) GetConfigValue() string`

GetConfigValue returns the ConfigValue field if non-nil, zero value otherwise.

### GetConfigValueOk

`func (o *BatchcontrolConfigLoadedRequest) GetConfigValueOk() (*string, bool)`

GetConfigValueOk returns a tuple with the ConfigValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigValue

`func (o *BatchcontrolConfigLoadedRequest) SetConfigValue(v string)`

SetConfigValue sets ConfigValue field to given value.


### GetScheduler

`func (o *BatchcontrolConfigLoadedRequest) GetScheduler() KalturaScheduler`

GetScheduler returns the Scheduler field if non-nil, zero value otherwise.

### GetSchedulerOk

`func (o *BatchcontrolConfigLoadedRequest) GetSchedulerOk() (*KalturaScheduler, bool)`

GetSchedulerOk returns a tuple with the Scheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduler

`func (o *BatchcontrolConfigLoadedRequest) SetScheduler(v KalturaScheduler)`

SetScheduler sets Scheduler field to given value.


### GetWorkerConfigId

`func (o *BatchcontrolConfigLoadedRequest) GetWorkerConfigId() int32`

GetWorkerConfigId returns the WorkerConfigId field if non-nil, zero value otherwise.

### GetWorkerConfigIdOk

`func (o *BatchcontrolConfigLoadedRequest) GetWorkerConfigIdOk() (*int32, bool)`

GetWorkerConfigIdOk returns a tuple with the WorkerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerConfigId

`func (o *BatchcontrolConfigLoadedRequest) SetWorkerConfigId(v int32)`

SetWorkerConfigId sets WorkerConfigId field to given value.

### HasWorkerConfigId

`func (o *BatchcontrolConfigLoadedRequest) HasWorkerConfigId() bool`

HasWorkerConfigId returns a boolean if a field has been set.

### GetWorkerName

`func (o *BatchcontrolConfigLoadedRequest) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *BatchcontrolConfigLoadedRequest) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *BatchcontrolConfigLoadedRequest) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.

### HasWorkerName

`func (o *BatchcontrolConfigLoadedRequest) HasWorkerName() bool`

HasWorkerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


