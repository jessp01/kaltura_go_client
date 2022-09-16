# KalturaSchedulerStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlPanelCommands** | Pointer to [**[]KalturaControlPanelCommand**](KalturaControlPanelCommand.md) |  | [optional] 
**QueuesStatus** | Pointer to [**[]KalturaBatchQueuesStatus**](KalturaBatchQueuesStatus.md) |  | [optional] 
**SchedulerConfigs** | Pointer to [**[]KalturaSchedulerConfig**](KalturaSchedulerConfig.md) |  | [optional] 

## Methods

### NewKalturaSchedulerStatusResponse

`func NewKalturaSchedulerStatusResponse() *KalturaSchedulerStatusResponse`

NewKalturaSchedulerStatusResponse instantiates a new KalturaSchedulerStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSchedulerStatusResponseWithDefaults

`func NewKalturaSchedulerStatusResponseWithDefaults() *KalturaSchedulerStatusResponse`

NewKalturaSchedulerStatusResponseWithDefaults instantiates a new KalturaSchedulerStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlPanelCommands

`func (o *KalturaSchedulerStatusResponse) GetControlPanelCommands() []KalturaControlPanelCommand`

GetControlPanelCommands returns the ControlPanelCommands field if non-nil, zero value otherwise.

### GetControlPanelCommandsOk

`func (o *KalturaSchedulerStatusResponse) GetControlPanelCommandsOk() (*[]KalturaControlPanelCommand, bool)`

GetControlPanelCommandsOk returns a tuple with the ControlPanelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanelCommands

`func (o *KalturaSchedulerStatusResponse) SetControlPanelCommands(v []KalturaControlPanelCommand)`

SetControlPanelCommands sets ControlPanelCommands field to given value.

### HasControlPanelCommands

`func (o *KalturaSchedulerStatusResponse) HasControlPanelCommands() bool`

HasControlPanelCommands returns a boolean if a field has been set.

### GetQueuesStatus

`func (o *KalturaSchedulerStatusResponse) GetQueuesStatus() []KalturaBatchQueuesStatus`

GetQueuesStatus returns the QueuesStatus field if non-nil, zero value otherwise.

### GetQueuesStatusOk

`func (o *KalturaSchedulerStatusResponse) GetQueuesStatusOk() (*[]KalturaBatchQueuesStatus, bool)`

GetQueuesStatusOk returns a tuple with the QueuesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuesStatus

`func (o *KalturaSchedulerStatusResponse) SetQueuesStatus(v []KalturaBatchQueuesStatus)`

SetQueuesStatus sets QueuesStatus field to given value.

### HasQueuesStatus

`func (o *KalturaSchedulerStatusResponse) HasQueuesStatus() bool`

HasQueuesStatus returns a boolean if a field has been set.

### GetSchedulerConfigs

`func (o *KalturaSchedulerStatusResponse) GetSchedulerConfigs() []KalturaSchedulerConfig`

GetSchedulerConfigs returns the SchedulerConfigs field if non-nil, zero value otherwise.

### GetSchedulerConfigsOk

`func (o *KalturaSchedulerStatusResponse) GetSchedulerConfigsOk() (*[]KalturaSchedulerConfig, bool)`

GetSchedulerConfigsOk returns a tuple with the SchedulerConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerConfigs

`func (o *KalturaSchedulerStatusResponse) SetSchedulerConfigs(v []KalturaSchedulerConfig)`

SetSchedulerConfigs sets SchedulerConfigs field to given value.

### HasSchedulerConfigs

`func (o *KalturaSchedulerStatusResponse) HasSchedulerConfigs() bool`

HasSchedulerConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


