# KalturaStatsKmcEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientVer** | Pointer to **string** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**EventTimestamp** | Pointer to **float32** | the client&#39;s timestamp of this event | [optional] 
**KmcEventActionPath** | Pointer to **string** |  | [optional] 
**KmcEventType** | Pointer to **int32** | Enum Type: &#x60;KalturaStatsKmcEventType&#x60; | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**SessionId** | Pointer to **string** | a unique string generated by the client that will represent the client-side session: the primary component will pass it on to other components that sprout from it | [optional] 
**UiconfId** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** | the partner&#39;s user id | [optional] 
**UserIp** | Pointer to **string** | &#x60;readOnly&#x60;  will be retrieved from the request of the user | [optional] [readonly] 
**WidgetId** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaStatsKmcEvent

`func NewKalturaStatsKmcEvent() *KalturaStatsKmcEvent`

NewKalturaStatsKmcEvent instantiates a new KalturaStatsKmcEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaStatsKmcEventWithDefaults

`func NewKalturaStatsKmcEventWithDefaults() *KalturaStatsKmcEvent`

NewKalturaStatsKmcEventWithDefaults instantiates a new KalturaStatsKmcEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientVer

`func (o *KalturaStatsKmcEvent) GetClientVer() string`

GetClientVer returns the ClientVer field if non-nil, zero value otherwise.

### GetClientVerOk

`func (o *KalturaStatsKmcEvent) GetClientVerOk() (*string, bool)`

GetClientVerOk returns a tuple with the ClientVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVer

`func (o *KalturaStatsKmcEvent) SetClientVer(v string)`

SetClientVer sets ClientVer field to given value.

### HasClientVer

`func (o *KalturaStatsKmcEvent) HasClientVer() bool`

HasClientVer returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaStatsKmcEvent) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaStatsKmcEvent) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaStatsKmcEvent) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaStatsKmcEvent) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEventTimestamp

`func (o *KalturaStatsKmcEvent) GetEventTimestamp() float32`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *KalturaStatsKmcEvent) GetEventTimestampOk() (*float32, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *KalturaStatsKmcEvent) SetEventTimestamp(v float32)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *KalturaStatsKmcEvent) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### GetKmcEventActionPath

`func (o *KalturaStatsKmcEvent) GetKmcEventActionPath() string`

GetKmcEventActionPath returns the KmcEventActionPath field if non-nil, zero value otherwise.

### GetKmcEventActionPathOk

`func (o *KalturaStatsKmcEvent) GetKmcEventActionPathOk() (*string, bool)`

GetKmcEventActionPathOk returns a tuple with the KmcEventActionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmcEventActionPath

`func (o *KalturaStatsKmcEvent) SetKmcEventActionPath(v string)`

SetKmcEventActionPath sets KmcEventActionPath field to given value.

### HasKmcEventActionPath

`func (o *KalturaStatsKmcEvent) HasKmcEventActionPath() bool`

HasKmcEventActionPath returns a boolean if a field has been set.

### GetKmcEventType

`func (o *KalturaStatsKmcEvent) GetKmcEventType() int32`

GetKmcEventType returns the KmcEventType field if non-nil, zero value otherwise.

### GetKmcEventTypeOk

`func (o *KalturaStatsKmcEvent) GetKmcEventTypeOk() (*int32, bool)`

GetKmcEventTypeOk returns a tuple with the KmcEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmcEventType

`func (o *KalturaStatsKmcEvent) SetKmcEventType(v int32)`

SetKmcEventType sets KmcEventType field to given value.

### HasKmcEventType

`func (o *KalturaStatsKmcEvent) HasKmcEventType() bool`

HasKmcEventType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaStatsKmcEvent) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaStatsKmcEvent) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaStatsKmcEvent) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaStatsKmcEvent) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSessionId

`func (o *KalturaStatsKmcEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *KalturaStatsKmcEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *KalturaStatsKmcEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *KalturaStatsKmcEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUiconfId

`func (o *KalturaStatsKmcEvent) GetUiconfId() int32`

GetUiconfId returns the UiconfId field if non-nil, zero value otherwise.

### GetUiconfIdOk

`func (o *KalturaStatsKmcEvent) GetUiconfIdOk() (*int32, bool)`

GetUiconfIdOk returns a tuple with the UiconfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiconfId

`func (o *KalturaStatsKmcEvent) SetUiconfId(v int32)`

SetUiconfId sets UiconfId field to given value.

### HasUiconfId

`func (o *KalturaStatsKmcEvent) HasUiconfId() bool`

HasUiconfId returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaStatsKmcEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaStatsKmcEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaStatsKmcEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaStatsKmcEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserIp

`func (o *KalturaStatsKmcEvent) GetUserIp() string`

GetUserIp returns the UserIp field if non-nil, zero value otherwise.

### GetUserIpOk

`func (o *KalturaStatsKmcEvent) GetUserIpOk() (*string, bool)`

GetUserIpOk returns a tuple with the UserIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIp

`func (o *KalturaStatsKmcEvent) SetUserIp(v string)`

SetUserIp sets UserIp field to given value.

### HasUserIp

`func (o *KalturaStatsKmcEvent) HasUserIp() bool`

HasUserIp returns a boolean if a field has been set.

### GetWidgetId

`func (o *KalturaStatsKmcEvent) GetWidgetId() string`

GetWidgetId returns the WidgetId field if non-nil, zero value otherwise.

### GetWidgetIdOk

`func (o *KalturaStatsKmcEvent) GetWidgetIdOk() (*string, bool)`

GetWidgetIdOk returns a tuple with the WidgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetId

`func (o *KalturaStatsKmcEvent) SetWidgetId(v string)`

SetWidgetId sets WidgetId field to given value.

### HasWidgetId

`func (o *KalturaStatsKmcEvent) HasWidgetId() bool`

HasWidgetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

