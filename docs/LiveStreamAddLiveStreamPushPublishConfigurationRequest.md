# LiveStreamAddLiveStreamPushPublishConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**LiveStreamConfiguration** | Pointer to [**KalturaLiveStreamConfiguration**](KalturaLiveStreamConfiguration.md) |  | [optional] 
**Protocol** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewLiveStreamAddLiveStreamPushPublishConfigurationRequest

`func NewLiveStreamAddLiveStreamPushPublishConfigurationRequest(entryId string, protocol string, ) *LiveStreamAddLiveStreamPushPublishConfigurationRequest`

NewLiveStreamAddLiveStreamPushPublishConfigurationRequest instantiates a new LiveStreamAddLiveStreamPushPublishConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamAddLiveStreamPushPublishConfigurationRequestWithDefaults

`func NewLiveStreamAddLiveStreamPushPublishConfigurationRequestWithDefaults() *LiveStreamAddLiveStreamPushPublishConfigurationRequest`

NewLiveStreamAddLiveStreamPushPublishConfigurationRequestWithDefaults instantiates a new LiveStreamAddLiveStreamPushPublishConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetLiveStreamConfiguration

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetLiveStreamConfiguration() KalturaLiveStreamConfiguration`

GetLiveStreamConfiguration returns the LiveStreamConfiguration field if non-nil, zero value otherwise.

### GetLiveStreamConfigurationOk

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetLiveStreamConfigurationOk() (*KalturaLiveStreamConfiguration, bool)`

GetLiveStreamConfigurationOk returns a tuple with the LiveStreamConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamConfiguration

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) SetLiveStreamConfiguration(v KalturaLiveStreamConfiguration)`

SetLiveStreamConfiguration sets LiveStreamConfiguration field to given value.

### HasLiveStreamConfiguration

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) HasLiveStreamConfiguration() bool`

HasLiveStreamConfiguration returns a boolean if a field has been set.

### GetProtocol

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetUrl

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LiveStreamAddLiveStreamPushPublishConfigurationRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


