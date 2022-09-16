# EventNotificationTemplateRegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationTemplateSystemName** | **string** |  | 
**PushNotificationParams** | [**KalturaPushNotificationParams**](KalturaPushNotificationParams.md) |  | 

## Methods

### NewEventNotificationTemplateRegisterRequest

`func NewEventNotificationTemplateRegisterRequest(notificationTemplateSystemName string, pushNotificationParams KalturaPushNotificationParams, ) *EventNotificationTemplateRegisterRequest`

NewEventNotificationTemplateRegisterRequest instantiates a new EventNotificationTemplateRegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationTemplateRegisterRequestWithDefaults

`func NewEventNotificationTemplateRegisterRequestWithDefaults() *EventNotificationTemplateRegisterRequest`

NewEventNotificationTemplateRegisterRequestWithDefaults instantiates a new EventNotificationTemplateRegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationTemplateSystemName

`func (o *EventNotificationTemplateRegisterRequest) GetNotificationTemplateSystemName() string`

GetNotificationTemplateSystemName returns the NotificationTemplateSystemName field if non-nil, zero value otherwise.

### GetNotificationTemplateSystemNameOk

`func (o *EventNotificationTemplateRegisterRequest) GetNotificationTemplateSystemNameOk() (*string, bool)`

GetNotificationTemplateSystemNameOk returns a tuple with the NotificationTemplateSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTemplateSystemName

`func (o *EventNotificationTemplateRegisterRequest) SetNotificationTemplateSystemName(v string)`

SetNotificationTemplateSystemName sets NotificationTemplateSystemName field to given value.


### GetPushNotificationParams

`func (o *EventNotificationTemplateRegisterRequest) GetPushNotificationParams() KalturaPushNotificationParams`

GetPushNotificationParams returns the PushNotificationParams field if non-nil, zero value otherwise.

### GetPushNotificationParamsOk

`func (o *EventNotificationTemplateRegisterRequest) GetPushNotificationParamsOk() (*KalturaPushNotificationParams, bool)`

GetPushNotificationParamsOk returns a tuple with the PushNotificationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotificationParams

`func (o *EventNotificationTemplateRegisterRequest) SetPushNotificationParams(v KalturaPushNotificationParams)`

SetPushNotificationParams sets PushNotificationParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


