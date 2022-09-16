# EventNotificationTemplateSendCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**NotificationTemplateSystemName** | **string** |  | 
**PushNotificationParams** | [**KalturaPushNotificationParams**](KalturaPushNotificationParams.md) |  | 

## Methods

### NewEventNotificationTemplateSendCommandRequest

`func NewEventNotificationTemplateSendCommandRequest(command string, notificationTemplateSystemName string, pushNotificationParams KalturaPushNotificationParams, ) *EventNotificationTemplateSendCommandRequest`

NewEventNotificationTemplateSendCommandRequest instantiates a new EventNotificationTemplateSendCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationTemplateSendCommandRequestWithDefaults

`func NewEventNotificationTemplateSendCommandRequestWithDefaults() *EventNotificationTemplateSendCommandRequest`

NewEventNotificationTemplateSendCommandRequestWithDefaults instantiates a new EventNotificationTemplateSendCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *EventNotificationTemplateSendCommandRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *EventNotificationTemplateSendCommandRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *EventNotificationTemplateSendCommandRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetNotificationTemplateSystemName

`func (o *EventNotificationTemplateSendCommandRequest) GetNotificationTemplateSystemName() string`

GetNotificationTemplateSystemName returns the NotificationTemplateSystemName field if non-nil, zero value otherwise.

### GetNotificationTemplateSystemNameOk

`func (o *EventNotificationTemplateSendCommandRequest) GetNotificationTemplateSystemNameOk() (*string, bool)`

GetNotificationTemplateSystemNameOk returns a tuple with the NotificationTemplateSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTemplateSystemName

`func (o *EventNotificationTemplateSendCommandRequest) SetNotificationTemplateSystemName(v string)`

SetNotificationTemplateSystemName sets NotificationTemplateSystemName field to given value.


### GetPushNotificationParams

`func (o *EventNotificationTemplateSendCommandRequest) GetPushNotificationParams() KalturaPushNotificationParams`

GetPushNotificationParams returns the PushNotificationParams field if non-nil, zero value otherwise.

### GetPushNotificationParamsOk

`func (o *EventNotificationTemplateSendCommandRequest) GetPushNotificationParamsOk() (*KalturaPushNotificationParams, bool)`

GetPushNotificationParamsOk returns a tuple with the PushNotificationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotificationParams

`func (o *EventNotificationTemplateSendCommandRequest) SetPushNotificationParams(v KalturaPushNotificationParams)`

SetPushNotificationParams sets PushNotificationParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


