# BatchcontrolSetWorkerConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **int32** |  | 
**Cause** | Pointer to **string** |  | [optional] 
**ConfigParam** | **string** |  | 
**ConfigParamPart** | Pointer to **string** |  | [optional] 
**ConfigValue** | **string** |  | 
**WorkerId** | **int32** |  | 

## Methods

### NewBatchcontrolSetWorkerConfigRequest

`func NewBatchcontrolSetWorkerConfigRequest(adminId int32, configParam string, configValue string, workerId int32, ) *BatchcontrolSetWorkerConfigRequest`

NewBatchcontrolSetWorkerConfigRequest instantiates a new BatchcontrolSetWorkerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolSetWorkerConfigRequestWithDefaults

`func NewBatchcontrolSetWorkerConfigRequestWithDefaults() *BatchcontrolSetWorkerConfigRequest`

NewBatchcontrolSetWorkerConfigRequestWithDefaults instantiates a new BatchcontrolSetWorkerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *BatchcontrolSetWorkerConfigRequest) GetAdminId() int32`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *BatchcontrolSetWorkerConfigRequest) GetAdminIdOk() (*int32, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *BatchcontrolSetWorkerConfigRequest) SetAdminId(v int32)`

SetAdminId sets AdminId field to given value.


### GetCause

`func (o *BatchcontrolSetWorkerConfigRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BatchcontrolSetWorkerConfigRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BatchcontrolSetWorkerConfigRequest) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *BatchcontrolSetWorkerConfigRequest) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetConfigParam

`func (o *BatchcontrolSetWorkerConfigRequest) GetConfigParam() string`

GetConfigParam returns the ConfigParam field if non-nil, zero value otherwise.

### GetConfigParamOk

`func (o *BatchcontrolSetWorkerConfigRequest) GetConfigParamOk() (*string, bool)`

GetConfigParamOk returns a tuple with the ConfigParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParam

`func (o *BatchcontrolSetWorkerConfigRequest) SetConfigParam(v string)`

SetConfigParam sets ConfigParam field to given value.


### GetConfigParamPart

`func (o *BatchcontrolSetWorkerConfigRequest) GetConfigParamPart() string`

GetConfigParamPart returns the ConfigParamPart field if non-nil, zero value otherwise.

### GetConfigParamPartOk

`func (o *BatchcontrolSetWorkerConfigRequest) GetConfigParamPartOk() (*string, bool)`

GetConfigParamPartOk returns a tuple with the ConfigParamPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParamPart

`func (o *BatchcontrolSetWorkerConfigRequest) SetConfigParamPart(v string)`

SetConfigParamPart sets ConfigParamPart field to given value.

### HasConfigParamPart

`func (o *BatchcontrolSetWorkerConfigRequest) HasConfigParamPart() bool`

HasConfigParamPart returns a boolean if a field has been set.

### GetConfigValue

`func (o *BatchcontrolSetWorkerConfigRequest) GetConfigValue() string`

GetConfigValue returns the ConfigValue field if non-nil, zero value otherwise.

### GetConfigValueOk

`func (o *BatchcontrolSetWorkerConfigRequest) GetConfigValueOk() (*string, bool)`

GetConfigValueOk returns a tuple with the ConfigValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigValue

`func (o *BatchcontrolSetWorkerConfigRequest) SetConfigValue(v string)`

SetConfigValue sets ConfigValue field to given value.


### GetWorkerId

`func (o *BatchcontrolSetWorkerConfigRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *BatchcontrolSetWorkerConfigRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *BatchcontrolSetWorkerConfigRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


