# KalturaClientNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | The serialized notification data to send | [optional] 
**Url** | Pointer to **string** | The URL where the notification should be sent to | [optional] 

## Methods

### NewKalturaClientNotification

`func NewKalturaClientNotification() *KalturaClientNotification`

NewKalturaClientNotification instantiates a new KalturaClientNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaClientNotificationWithDefaults

`func NewKalturaClientNotificationWithDefaults() *KalturaClientNotification`

NewKalturaClientNotificationWithDefaults instantiates a new KalturaClientNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *KalturaClientNotification) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KalturaClientNotification) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KalturaClientNotification) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *KalturaClientNotification) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUrl

`func (o *KalturaClientNotification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KalturaClientNotification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KalturaClientNotification) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KalturaClientNotification) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


