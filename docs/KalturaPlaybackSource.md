# KalturaPlaybackSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryProfileId** | Pointer to **string** |  | [optional] 
**Drm** | Pointer to [**[]KalturaDrmPlaybackPluginData**](KalturaDrmPlaybackPluginData.md) |  | [optional] 
**FlavorIds** | Pointer to **string** | comma separated string of flavor ids | [optional] 
**Format** | Pointer to **string** | source format according to delivery profile streamer type (applehttp, mpegdash etc.) | [optional] 
**Protocols** | Pointer to **string** | comma separated string according to deliveryProfile media protocols (&#39;http,https&#39; etc.) | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaPlaybackSource

`func NewKalturaPlaybackSource() *KalturaPlaybackSource`

NewKalturaPlaybackSource instantiates a new KalturaPlaybackSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaPlaybackSourceWithDefaults

`func NewKalturaPlaybackSourceWithDefaults() *KalturaPlaybackSource`

NewKalturaPlaybackSourceWithDefaults instantiates a new KalturaPlaybackSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryProfileId

`func (o *KalturaPlaybackSource) GetDeliveryProfileId() string`

GetDeliveryProfileId returns the DeliveryProfileId field if non-nil, zero value otherwise.

### GetDeliveryProfileIdOk

`func (o *KalturaPlaybackSource) GetDeliveryProfileIdOk() (*string, bool)`

GetDeliveryProfileIdOk returns a tuple with the DeliveryProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProfileId

`func (o *KalturaPlaybackSource) SetDeliveryProfileId(v string)`

SetDeliveryProfileId sets DeliveryProfileId field to given value.

### HasDeliveryProfileId

`func (o *KalturaPlaybackSource) HasDeliveryProfileId() bool`

HasDeliveryProfileId returns a boolean if a field has been set.

### GetDrm

`func (o *KalturaPlaybackSource) GetDrm() []KalturaDrmPlaybackPluginData`

GetDrm returns the Drm field if non-nil, zero value otherwise.

### GetDrmOk

`func (o *KalturaPlaybackSource) GetDrmOk() (*[]KalturaDrmPlaybackPluginData, bool)`

GetDrmOk returns a tuple with the Drm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrm

`func (o *KalturaPlaybackSource) SetDrm(v []KalturaDrmPlaybackPluginData)`

SetDrm sets Drm field to given value.

### HasDrm

`func (o *KalturaPlaybackSource) HasDrm() bool`

HasDrm returns a boolean if a field has been set.

### GetFlavorIds

`func (o *KalturaPlaybackSource) GetFlavorIds() string`

GetFlavorIds returns the FlavorIds field if non-nil, zero value otherwise.

### GetFlavorIdsOk

`func (o *KalturaPlaybackSource) GetFlavorIdsOk() (*string, bool)`

GetFlavorIdsOk returns a tuple with the FlavorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorIds

`func (o *KalturaPlaybackSource) SetFlavorIds(v string)`

SetFlavorIds sets FlavorIds field to given value.

### HasFlavorIds

`func (o *KalturaPlaybackSource) HasFlavorIds() bool`

HasFlavorIds returns a boolean if a field has been set.

### GetFormat

`func (o *KalturaPlaybackSource) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *KalturaPlaybackSource) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *KalturaPlaybackSource) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *KalturaPlaybackSource) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetProtocols

`func (o *KalturaPlaybackSource) GetProtocols() string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *KalturaPlaybackSource) GetProtocolsOk() (*string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *KalturaPlaybackSource) SetProtocols(v string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *KalturaPlaybackSource) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetUrl

`func (o *KalturaPlaybackSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KalturaPlaybackSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KalturaPlaybackSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KalturaPlaybackSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


