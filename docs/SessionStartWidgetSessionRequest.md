# SessionStartWidgetSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**WidgetId** | **string** |  | 

## Methods

### NewSessionStartWidgetSessionRequest

`func NewSessionStartWidgetSessionRequest(widgetId string, ) *SessionStartWidgetSessionRequest`

NewSessionStartWidgetSessionRequest instantiates a new SessionStartWidgetSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStartWidgetSessionRequestWithDefaults

`func NewSessionStartWidgetSessionRequestWithDefaults() *SessionStartWidgetSessionRequest`

NewSessionStartWidgetSessionRequestWithDefaults instantiates a new SessionStartWidgetSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *SessionStartWidgetSessionRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SessionStartWidgetSessionRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SessionStartWidgetSessionRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SessionStartWidgetSessionRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetWidgetId

`func (o *SessionStartWidgetSessionRequest) GetWidgetId() string`

GetWidgetId returns the WidgetId field if non-nil, zero value otherwise.

### GetWidgetIdOk

`func (o *SessionStartWidgetSessionRequest) GetWidgetIdOk() (*string, bool)`

GetWidgetIdOk returns a tuple with the WidgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetId

`func (o *SessionStartWidgetSessionRequest) SetWidgetId(v string)`

SetWidgetId sets WidgetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


