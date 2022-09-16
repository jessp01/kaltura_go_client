# KalturaLiveStreamParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bitrate** | Pointer to **int32** | Bit rate of the stream. (i.e. 900) | [optional] 
**Codec** | Pointer to **string** | Live stream&#39;s codec | [optional] 
**FlavorId** | Pointer to **string** | flavor asset id | [optional] 
**FrameRate** | Pointer to **int32** | Live stream&#39;s farme rate | [optional] 
**Height** | Pointer to **int32** | Stream&#39;s height | [optional] 
**KeyFrameInterval** | Pointer to **float32** | Live stream&#39;s key frame interval | [optional] 
**Language** | Pointer to **string** | Live stream&#39;s language | [optional] 
**Width** | Pointer to **int32** | Stream&#39;s width | [optional] 

## Methods

### NewKalturaLiveStreamParams

`func NewKalturaLiveStreamParams() *KalturaLiveStreamParams`

NewKalturaLiveStreamParams instantiates a new KalturaLiveStreamParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLiveStreamParamsWithDefaults

`func NewKalturaLiveStreamParamsWithDefaults() *KalturaLiveStreamParams`

NewKalturaLiveStreamParamsWithDefaults instantiates a new KalturaLiveStreamParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitrate

`func (o *KalturaLiveStreamParams) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *KalturaLiveStreamParams) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *KalturaLiveStreamParams) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *KalturaLiveStreamParams) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetCodec

`func (o *KalturaLiveStreamParams) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *KalturaLiveStreamParams) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *KalturaLiveStreamParams) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *KalturaLiveStreamParams) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetFlavorId

`func (o *KalturaLiveStreamParams) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *KalturaLiveStreamParams) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *KalturaLiveStreamParams) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *KalturaLiveStreamParams) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### GetFrameRate

`func (o *KalturaLiveStreamParams) GetFrameRate() int32`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *KalturaLiveStreamParams) GetFrameRateOk() (*int32, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameRate

`func (o *KalturaLiveStreamParams) SetFrameRate(v int32)`

SetFrameRate sets FrameRate field to given value.

### HasFrameRate

`func (o *KalturaLiveStreamParams) HasFrameRate() bool`

HasFrameRate returns a boolean if a field has been set.

### GetHeight

`func (o *KalturaLiveStreamParams) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *KalturaLiveStreamParams) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *KalturaLiveStreamParams) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *KalturaLiveStreamParams) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetKeyFrameInterval

`func (o *KalturaLiveStreamParams) GetKeyFrameInterval() float32`

GetKeyFrameInterval returns the KeyFrameInterval field if non-nil, zero value otherwise.

### GetKeyFrameIntervalOk

`func (o *KalturaLiveStreamParams) GetKeyFrameIntervalOk() (*float32, bool)`

GetKeyFrameIntervalOk returns a tuple with the KeyFrameInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFrameInterval

`func (o *KalturaLiveStreamParams) SetKeyFrameInterval(v float32)`

SetKeyFrameInterval sets KeyFrameInterval field to given value.

### HasKeyFrameInterval

`func (o *KalturaLiveStreamParams) HasKeyFrameInterval() bool`

HasKeyFrameInterval returns a boolean if a field has been set.

### GetLanguage

`func (o *KalturaLiveStreamParams) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *KalturaLiveStreamParams) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *KalturaLiveStreamParams) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *KalturaLiveStreamParams) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetWidth

`func (o *KalturaLiveStreamParams) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *KalturaLiveStreamParams) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *KalturaLiveStreamParams) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *KalturaLiveStreamParams) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


