# BatchcontrolSetSchedulerConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **int32** |  | 
**Cause** | Pointer to **string** |  | [optional] 
**ConfigParam** | **string** |  | 
**ConfigParamPart** | Pointer to **string** |  | [optional] 
**ConfigValue** | **string** |  | 
**SchedulerId** | **int32** |  | 

## Methods

### NewBatchcontrolSetSchedulerConfigRequest

`func NewBatchcontrolSetSchedulerConfigRequest(adminId int32, configParam string, configValue string, schedulerId int32, ) *BatchcontrolSetSchedulerConfigRequest`

NewBatchcontrolSetSchedulerConfigRequest instantiates a new BatchcontrolSetSchedulerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolSetSchedulerConfigRequestWithDefaults

`func NewBatchcontrolSetSchedulerConfigRequestWithDefaults() *BatchcontrolSetSchedulerConfigRequest`

NewBatchcontrolSetSchedulerConfigRequestWithDefaults instantiates a new BatchcontrolSetSchedulerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *BatchcontrolSetSchedulerConfigRequest) GetAdminId() int32`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *BatchcontrolSetSchedulerConfigRequest) GetAdminIdOk() (*int32, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *BatchcontrolSetSchedulerConfigRequest) SetAdminId(v int32)`

SetAdminId sets AdminId field to given value.


### GetCause

`func (o *BatchcontrolSetSchedulerConfigRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BatchcontrolSetSchedulerConfigRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BatchcontrolSetSchedulerConfigRequest) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *BatchcontrolSetSchedulerConfigRequest) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetConfigParam

`func (o *BatchcontrolSetSchedulerConfigRequest) GetConfigParam() string`

GetConfigParam returns the ConfigParam field if non-nil, zero value otherwise.

### GetConfigParamOk

`func (o *BatchcontrolSetSchedulerConfigRequest) GetConfigParamOk() (*string, bool)`

GetConfigParamOk returns a tuple with the ConfigParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParam

`func (o *BatchcontrolSetSchedulerConfigRequest) SetConfigParam(v string)`

SetConfigParam sets ConfigParam field to given value.


### GetConfigParamPart

`func (o *BatchcontrolSetSchedulerConfigRequest) GetConfigParamPart() string`

GetConfigParamPart returns the ConfigParamPart field if non-nil, zero value otherwise.

### GetConfigParamPartOk

`func (o *BatchcontrolSetSchedulerConfigRequest) GetConfigParamPartOk() (*string, bool)`

GetConfigParamPartOk returns a tuple with the ConfigParamPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParamPart

`func (o *BatchcontrolSetSchedulerConfigRequest) SetConfigParamPart(v string)`

SetConfigParamPart sets ConfigParamPart field to given value.

### HasConfigParamPart

`func (o *BatchcontrolSetSchedulerConfigRequest) HasConfigParamPart() bool`

HasConfigParamPart returns a boolean if a field has been set.

### GetConfigValue

`func (o *BatchcontrolSetSchedulerConfigRequest) GetConfigValue() string`

GetConfigValue returns the ConfigValue field if non-nil, zero value otherwise.

### GetConfigValueOk

`func (o *BatchcontrolSetSchedulerConfigRequest) GetConfigValueOk() (*string, bool)`

GetConfigValueOk returns a tuple with the ConfigValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigValue

`func (o *BatchcontrolSetSchedulerConfigRequest) SetConfigValue(v string)`

SetConfigValue sets ConfigValue field to given value.


### GetSchedulerId

`func (o *BatchcontrolSetSchedulerConfigRequest) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *BatchcontrolSetSchedulerConfigRequest) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *BatchcontrolSetSchedulerConfigRequest) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


