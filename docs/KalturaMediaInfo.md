# KalturaMediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioBitRate** | Pointer to **int32** | The audio bit rate | [optional] 
**AudioBitRateMode** | Pointer to **int32** | Enum Type: &#x60;KalturaBitRateMode&#x60;  The audio bit rate mode | [optional] 
**AudioChannels** | Pointer to **int32** | The number of audio channels | [optional] 
**AudioCodecId** | Pointer to **string** | The audio codec id | [optional] 
**AudioDuration** | Pointer to **int32** | The audio duration | [optional] 
**AudioFormat** | Pointer to **string** | The audio format | [optional] 
**AudioResolution** | Pointer to **int32** | The audio resolution | [optional] 
**AudioSamplingRate** | Pointer to **int32** | The audio sampling rate | [optional] 
**BitsDepth** | Pointer to **int32** |  | [optional] 
**ChromaSubsampling** | Pointer to **string** |  | [optional] 
**ColorPrimaries** | Pointer to **string** |  | [optional] 
**ColorSpace** | Pointer to **string** |  | [optional] 
**ColorTransfer** | Pointer to **string** |  | [optional] 
**ComplexityValue** | Pointer to **int32** |  | [optional] 
**ContainerBitRate** | Pointer to **int32** | The container bit rate | [optional] 
**ContainerDuration** | Pointer to **int32** | The container duration | [optional] 
**ContainerFormat** | Pointer to **string** | The container format | [optional] 
**ContainerId** | Pointer to **string** | The container id | [optional] 
**ContainerProfile** | Pointer to **string** | The container profile | [optional] 
**ContentStreams** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** | The file size | [optional] 
**FlavorAssetId** | Pointer to **string** | The id of the related flavor asset | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the media info | [optional] [readonly] 
**IsFastStart** | Pointer to **int32** |  | [optional] 
**MatrixCoefficients** | Pointer to **string** |  | [optional] 
**MaxGOP** | Pointer to **float32** |  | [optional] 
**MultiStream** | Pointer to **string** |  | [optional] 
**MultiStreamInfo** | Pointer to **string** |  | [optional] 
**PixelFormat** | Pointer to **string** |  | [optional] 
**RawData** | Pointer to **string** | The data as returned by the mediainfo command line | [optional] 
**ScanType** | Pointer to **int32** |  | [optional] 
**VideoBitRate** | Pointer to **int32** | The video bit rate | [optional] 
**VideoBitRateMode** | Pointer to **int32** | Enum Type: &#x60;KalturaBitRateMode&#x60;  The video bit rate mode | [optional] 
**VideoCodecId** | Pointer to **string** | The video codec id | [optional] 
**VideoDar** | Pointer to **float32** | The video display aspect ratio (dar) | [optional] 
**VideoDuration** | Pointer to **int32** | The video duration | [optional] 
**VideoFormat** | Pointer to **string** | The video format | [optional] 
**VideoFrameRate** | Pointer to **float32** | The video frame rate | [optional] 
**VideoHeight** | Pointer to **int32** | The video height | [optional] 
**VideoRotation** | Pointer to **int32** |  | [optional] 
**VideoWidth** | Pointer to **int32** | The video width | [optional] 
**WritingLib** | Pointer to **string** | The writing library | [optional] 

## Methods

### NewKalturaMediaInfo

`func NewKalturaMediaInfo() *KalturaMediaInfo`

NewKalturaMediaInfo instantiates a new KalturaMediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaMediaInfoWithDefaults

`func NewKalturaMediaInfoWithDefaults() *KalturaMediaInfo`

NewKalturaMediaInfoWithDefaults instantiates a new KalturaMediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioBitRate

`func (o *KalturaMediaInfo) GetAudioBitRate() int32`

GetAudioBitRate returns the AudioBitRate field if non-nil, zero value otherwise.

### GetAudioBitRateOk

`func (o *KalturaMediaInfo) GetAudioBitRateOk() (*int32, bool)`

GetAudioBitRateOk returns a tuple with the AudioBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitRate

`func (o *KalturaMediaInfo) SetAudioBitRate(v int32)`

SetAudioBitRate sets AudioBitRate field to given value.

### HasAudioBitRate

`func (o *KalturaMediaInfo) HasAudioBitRate() bool`

HasAudioBitRate returns a boolean if a field has been set.

### GetAudioBitRateMode

`func (o *KalturaMediaInfo) GetAudioBitRateMode() int32`

GetAudioBitRateMode returns the AudioBitRateMode field if non-nil, zero value otherwise.

### GetAudioBitRateModeOk

`func (o *KalturaMediaInfo) GetAudioBitRateModeOk() (*int32, bool)`

GetAudioBitRateModeOk returns a tuple with the AudioBitRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitRateMode

`func (o *KalturaMediaInfo) SetAudioBitRateMode(v int32)`

SetAudioBitRateMode sets AudioBitRateMode field to given value.

### HasAudioBitRateMode

`func (o *KalturaMediaInfo) HasAudioBitRateMode() bool`

HasAudioBitRateMode returns a boolean if a field has been set.

### GetAudioChannels

`func (o *KalturaMediaInfo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *KalturaMediaInfo) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *KalturaMediaInfo) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *KalturaMediaInfo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### GetAudioCodecId

`func (o *KalturaMediaInfo) GetAudioCodecId() string`

GetAudioCodecId returns the AudioCodecId field if non-nil, zero value otherwise.

### GetAudioCodecIdOk

`func (o *KalturaMediaInfo) GetAudioCodecIdOk() (*string, bool)`

GetAudioCodecIdOk returns a tuple with the AudioCodecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodecId

`func (o *KalturaMediaInfo) SetAudioCodecId(v string)`

SetAudioCodecId sets AudioCodecId field to given value.

### HasAudioCodecId

`func (o *KalturaMediaInfo) HasAudioCodecId() bool`

HasAudioCodecId returns a boolean if a field has been set.

### GetAudioDuration

`func (o *KalturaMediaInfo) GetAudioDuration() int32`

GetAudioDuration returns the AudioDuration field if non-nil, zero value otherwise.

### GetAudioDurationOk

`func (o *KalturaMediaInfo) GetAudioDurationOk() (*int32, bool)`

GetAudioDurationOk returns a tuple with the AudioDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioDuration

`func (o *KalturaMediaInfo) SetAudioDuration(v int32)`

SetAudioDuration sets AudioDuration field to given value.

### HasAudioDuration

`func (o *KalturaMediaInfo) HasAudioDuration() bool`

HasAudioDuration returns a boolean if a field has been set.

### GetAudioFormat

`func (o *KalturaMediaInfo) GetAudioFormat() string`

GetAudioFormat returns the AudioFormat field if non-nil, zero value otherwise.

### GetAudioFormatOk

`func (o *KalturaMediaInfo) GetAudioFormatOk() (*string, bool)`

GetAudioFormatOk returns a tuple with the AudioFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioFormat

`func (o *KalturaMediaInfo) SetAudioFormat(v string)`

SetAudioFormat sets AudioFormat field to given value.

### HasAudioFormat

`func (o *KalturaMediaInfo) HasAudioFormat() bool`

HasAudioFormat returns a boolean if a field has been set.

### GetAudioResolution

`func (o *KalturaMediaInfo) GetAudioResolution() int32`

GetAudioResolution returns the AudioResolution field if non-nil, zero value otherwise.

### GetAudioResolutionOk

`func (o *KalturaMediaInfo) GetAudioResolutionOk() (*int32, bool)`

GetAudioResolutionOk returns a tuple with the AudioResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioResolution

`func (o *KalturaMediaInfo) SetAudioResolution(v int32)`

SetAudioResolution sets AudioResolution field to given value.

### HasAudioResolution

`func (o *KalturaMediaInfo) HasAudioResolution() bool`

HasAudioResolution returns a boolean if a field has been set.

### GetAudioSamplingRate

`func (o *KalturaMediaInfo) GetAudioSamplingRate() int32`

GetAudioSamplingRate returns the AudioSamplingRate field if non-nil, zero value otherwise.

### GetAudioSamplingRateOk

`func (o *KalturaMediaInfo) GetAudioSamplingRateOk() (*int32, bool)`

GetAudioSamplingRateOk returns a tuple with the AudioSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSamplingRate

`func (o *KalturaMediaInfo) SetAudioSamplingRate(v int32)`

SetAudioSamplingRate sets AudioSamplingRate field to given value.

### HasAudioSamplingRate

`func (o *KalturaMediaInfo) HasAudioSamplingRate() bool`

HasAudioSamplingRate returns a boolean if a field has been set.

### GetBitsDepth

`func (o *KalturaMediaInfo) GetBitsDepth() int32`

GetBitsDepth returns the BitsDepth field if non-nil, zero value otherwise.

### GetBitsDepthOk

`func (o *KalturaMediaInfo) GetBitsDepthOk() (*int32, bool)`

GetBitsDepthOk returns a tuple with the BitsDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitsDepth

`func (o *KalturaMediaInfo) SetBitsDepth(v int32)`

SetBitsDepth sets BitsDepth field to given value.

### HasBitsDepth

`func (o *KalturaMediaInfo) HasBitsDepth() bool`

HasBitsDepth returns a boolean if a field has been set.

### GetChromaSubsampling

`func (o *KalturaMediaInfo) GetChromaSubsampling() string`

GetChromaSubsampling returns the ChromaSubsampling field if non-nil, zero value otherwise.

### GetChromaSubsamplingOk

`func (o *KalturaMediaInfo) GetChromaSubsamplingOk() (*string, bool)`

GetChromaSubsamplingOk returns a tuple with the ChromaSubsampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromaSubsampling

`func (o *KalturaMediaInfo) SetChromaSubsampling(v string)`

SetChromaSubsampling sets ChromaSubsampling field to given value.

### HasChromaSubsampling

`func (o *KalturaMediaInfo) HasChromaSubsampling() bool`

HasChromaSubsampling returns a boolean if a field has been set.

### GetColorPrimaries

`func (o *KalturaMediaInfo) GetColorPrimaries() string`

GetColorPrimaries returns the ColorPrimaries field if non-nil, zero value otherwise.

### GetColorPrimariesOk

`func (o *KalturaMediaInfo) GetColorPrimariesOk() (*string, bool)`

GetColorPrimariesOk returns a tuple with the ColorPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorPrimaries

`func (o *KalturaMediaInfo) SetColorPrimaries(v string)`

SetColorPrimaries sets ColorPrimaries field to given value.

### HasColorPrimaries

`func (o *KalturaMediaInfo) HasColorPrimaries() bool`

HasColorPrimaries returns a boolean if a field has been set.

### GetColorSpace

`func (o *KalturaMediaInfo) GetColorSpace() string`

GetColorSpace returns the ColorSpace field if non-nil, zero value otherwise.

### GetColorSpaceOk

`func (o *KalturaMediaInfo) GetColorSpaceOk() (*string, bool)`

GetColorSpaceOk returns a tuple with the ColorSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorSpace

`func (o *KalturaMediaInfo) SetColorSpace(v string)`

SetColorSpace sets ColorSpace field to given value.

### HasColorSpace

`func (o *KalturaMediaInfo) HasColorSpace() bool`

HasColorSpace returns a boolean if a field has been set.

### GetColorTransfer

`func (o *KalturaMediaInfo) GetColorTransfer() string`

GetColorTransfer returns the ColorTransfer field if non-nil, zero value otherwise.

### GetColorTransferOk

`func (o *KalturaMediaInfo) GetColorTransferOk() (*string, bool)`

GetColorTransferOk returns a tuple with the ColorTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorTransfer

`func (o *KalturaMediaInfo) SetColorTransfer(v string)`

SetColorTransfer sets ColorTransfer field to given value.

### HasColorTransfer

`func (o *KalturaMediaInfo) HasColorTransfer() bool`

HasColorTransfer returns a boolean if a field has been set.

### GetComplexityValue

`func (o *KalturaMediaInfo) GetComplexityValue() int32`

GetComplexityValue returns the ComplexityValue field if non-nil, zero value otherwise.

### GetComplexityValueOk

`func (o *KalturaMediaInfo) GetComplexityValueOk() (*int32, bool)`

GetComplexityValueOk returns a tuple with the ComplexityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexityValue

`func (o *KalturaMediaInfo) SetComplexityValue(v int32)`

SetComplexityValue sets ComplexityValue field to given value.

### HasComplexityValue

`func (o *KalturaMediaInfo) HasComplexityValue() bool`

HasComplexityValue returns a boolean if a field has been set.

### GetContainerBitRate

`func (o *KalturaMediaInfo) GetContainerBitRate() int32`

GetContainerBitRate returns the ContainerBitRate field if non-nil, zero value otherwise.

### GetContainerBitRateOk

`func (o *KalturaMediaInfo) GetContainerBitRateOk() (*int32, bool)`

GetContainerBitRateOk returns a tuple with the ContainerBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerBitRate

`func (o *KalturaMediaInfo) SetContainerBitRate(v int32)`

SetContainerBitRate sets ContainerBitRate field to given value.

### HasContainerBitRate

`func (o *KalturaMediaInfo) HasContainerBitRate() bool`

HasContainerBitRate returns a boolean if a field has been set.

### GetContainerDuration

`func (o *KalturaMediaInfo) GetContainerDuration() int32`

GetContainerDuration returns the ContainerDuration field if non-nil, zero value otherwise.

### GetContainerDurationOk

`func (o *KalturaMediaInfo) GetContainerDurationOk() (*int32, bool)`

GetContainerDurationOk returns a tuple with the ContainerDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDuration

`func (o *KalturaMediaInfo) SetContainerDuration(v int32)`

SetContainerDuration sets ContainerDuration field to given value.

### HasContainerDuration

`func (o *KalturaMediaInfo) HasContainerDuration() bool`

HasContainerDuration returns a boolean if a field has been set.

### GetContainerFormat

`func (o *KalturaMediaInfo) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *KalturaMediaInfo) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *KalturaMediaInfo) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *KalturaMediaInfo) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### GetContainerId

`func (o *KalturaMediaInfo) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *KalturaMediaInfo) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *KalturaMediaInfo) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *KalturaMediaInfo) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetContainerProfile

`func (o *KalturaMediaInfo) GetContainerProfile() string`

GetContainerProfile returns the ContainerProfile field if non-nil, zero value otherwise.

### GetContainerProfileOk

`func (o *KalturaMediaInfo) GetContainerProfileOk() (*string, bool)`

GetContainerProfileOk returns a tuple with the ContainerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfile

`func (o *KalturaMediaInfo) SetContainerProfile(v string)`

SetContainerProfile sets ContainerProfile field to given value.

### HasContainerProfile

`func (o *KalturaMediaInfo) HasContainerProfile() bool`

HasContainerProfile returns a boolean if a field has been set.

### GetContentStreams

`func (o *KalturaMediaInfo) GetContentStreams() string`

GetContentStreams returns the ContentStreams field if non-nil, zero value otherwise.

### GetContentStreamsOk

`func (o *KalturaMediaInfo) GetContentStreamsOk() (*string, bool)`

GetContentStreamsOk returns a tuple with the ContentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentStreams

`func (o *KalturaMediaInfo) SetContentStreams(v string)`

SetContentStreams sets ContentStreams field to given value.

### HasContentStreams

`func (o *KalturaMediaInfo) HasContentStreams() bool`

HasContentStreams returns a boolean if a field has been set.

### GetFileSize

`func (o *KalturaMediaInfo) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *KalturaMediaInfo) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *KalturaMediaInfo) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *KalturaMediaInfo) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFlavorAssetId

`func (o *KalturaMediaInfo) GetFlavorAssetId() string`

GetFlavorAssetId returns the FlavorAssetId field if non-nil, zero value otherwise.

### GetFlavorAssetIdOk

`func (o *KalturaMediaInfo) GetFlavorAssetIdOk() (*string, bool)`

GetFlavorAssetIdOk returns a tuple with the FlavorAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorAssetId

`func (o *KalturaMediaInfo) SetFlavorAssetId(v string)`

SetFlavorAssetId sets FlavorAssetId field to given value.

### HasFlavorAssetId

`func (o *KalturaMediaInfo) HasFlavorAssetId() bool`

HasFlavorAssetId returns a boolean if a field has been set.

### GetId

`func (o *KalturaMediaInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaMediaInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaMediaInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaMediaInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFastStart

`func (o *KalturaMediaInfo) GetIsFastStart() int32`

GetIsFastStart returns the IsFastStart field if non-nil, zero value otherwise.

### GetIsFastStartOk

`func (o *KalturaMediaInfo) GetIsFastStartOk() (*int32, bool)`

GetIsFastStartOk returns a tuple with the IsFastStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFastStart

`func (o *KalturaMediaInfo) SetIsFastStart(v int32)`

SetIsFastStart sets IsFastStart field to given value.

### HasIsFastStart

`func (o *KalturaMediaInfo) HasIsFastStart() bool`

HasIsFastStart returns a boolean if a field has been set.

### GetMatrixCoefficients

`func (o *KalturaMediaInfo) GetMatrixCoefficients() string`

GetMatrixCoefficients returns the MatrixCoefficients field if non-nil, zero value otherwise.

### GetMatrixCoefficientsOk

`func (o *KalturaMediaInfo) GetMatrixCoefficientsOk() (*string, bool)`

GetMatrixCoefficientsOk returns a tuple with the MatrixCoefficients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrixCoefficients

`func (o *KalturaMediaInfo) SetMatrixCoefficients(v string)`

SetMatrixCoefficients sets MatrixCoefficients field to given value.

### HasMatrixCoefficients

`func (o *KalturaMediaInfo) HasMatrixCoefficients() bool`

HasMatrixCoefficients returns a boolean if a field has been set.

### GetMaxGOP

`func (o *KalturaMediaInfo) GetMaxGOP() float32`

GetMaxGOP returns the MaxGOP field if non-nil, zero value otherwise.

### GetMaxGOPOk

`func (o *KalturaMediaInfo) GetMaxGOPOk() (*float32, bool)`

GetMaxGOPOk returns a tuple with the MaxGOP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGOP

`func (o *KalturaMediaInfo) SetMaxGOP(v float32)`

SetMaxGOP sets MaxGOP field to given value.

### HasMaxGOP

`func (o *KalturaMediaInfo) HasMaxGOP() bool`

HasMaxGOP returns a boolean if a field has been set.

### GetMultiStream

`func (o *KalturaMediaInfo) GetMultiStream() string`

GetMultiStream returns the MultiStream field if non-nil, zero value otherwise.

### GetMultiStreamOk

`func (o *KalturaMediaInfo) GetMultiStreamOk() (*string, bool)`

GetMultiStreamOk returns a tuple with the MultiStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStream

`func (o *KalturaMediaInfo) SetMultiStream(v string)`

SetMultiStream sets MultiStream field to given value.

### HasMultiStream

`func (o *KalturaMediaInfo) HasMultiStream() bool`

HasMultiStream returns a boolean if a field has been set.

### GetMultiStreamInfo

`func (o *KalturaMediaInfo) GetMultiStreamInfo() string`

GetMultiStreamInfo returns the MultiStreamInfo field if non-nil, zero value otherwise.

### GetMultiStreamInfoOk

`func (o *KalturaMediaInfo) GetMultiStreamInfoOk() (*string, bool)`

GetMultiStreamInfoOk returns a tuple with the MultiStreamInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStreamInfo

`func (o *KalturaMediaInfo) SetMultiStreamInfo(v string)`

SetMultiStreamInfo sets MultiStreamInfo field to given value.

### HasMultiStreamInfo

`func (o *KalturaMediaInfo) HasMultiStreamInfo() bool`

HasMultiStreamInfo returns a boolean if a field has been set.

### GetPixelFormat

`func (o *KalturaMediaInfo) GetPixelFormat() string`

GetPixelFormat returns the PixelFormat field if non-nil, zero value otherwise.

### GetPixelFormatOk

`func (o *KalturaMediaInfo) GetPixelFormatOk() (*string, bool)`

GetPixelFormatOk returns a tuple with the PixelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelFormat

`func (o *KalturaMediaInfo) SetPixelFormat(v string)`

SetPixelFormat sets PixelFormat field to given value.

### HasPixelFormat

`func (o *KalturaMediaInfo) HasPixelFormat() bool`

HasPixelFormat returns a boolean if a field has been set.

### GetRawData

`func (o *KalturaMediaInfo) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *KalturaMediaInfo) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *KalturaMediaInfo) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *KalturaMediaInfo) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetScanType

`func (o *KalturaMediaInfo) GetScanType() int32`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *KalturaMediaInfo) GetScanTypeOk() (*int32, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *KalturaMediaInfo) SetScanType(v int32)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *KalturaMediaInfo) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetVideoBitRate

`func (o *KalturaMediaInfo) GetVideoBitRate() int32`

GetVideoBitRate returns the VideoBitRate field if non-nil, zero value otherwise.

### GetVideoBitRateOk

`func (o *KalturaMediaInfo) GetVideoBitRateOk() (*int32, bool)`

GetVideoBitRateOk returns a tuple with the VideoBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBitRate

`func (o *KalturaMediaInfo) SetVideoBitRate(v int32)`

SetVideoBitRate sets VideoBitRate field to given value.

### HasVideoBitRate

`func (o *KalturaMediaInfo) HasVideoBitRate() bool`

HasVideoBitRate returns a boolean if a field has been set.

### GetVideoBitRateMode

`func (o *KalturaMediaInfo) GetVideoBitRateMode() int32`

GetVideoBitRateMode returns the VideoBitRateMode field if non-nil, zero value otherwise.

### GetVideoBitRateModeOk

`func (o *KalturaMediaInfo) GetVideoBitRateModeOk() (*int32, bool)`

GetVideoBitRateModeOk returns a tuple with the VideoBitRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBitRateMode

`func (o *KalturaMediaInfo) SetVideoBitRateMode(v int32)`

SetVideoBitRateMode sets VideoBitRateMode field to given value.

### HasVideoBitRateMode

`func (o *KalturaMediaInfo) HasVideoBitRateMode() bool`

HasVideoBitRateMode returns a boolean if a field has been set.

### GetVideoCodecId

`func (o *KalturaMediaInfo) GetVideoCodecId() string`

GetVideoCodecId returns the VideoCodecId field if non-nil, zero value otherwise.

### GetVideoCodecIdOk

`func (o *KalturaMediaInfo) GetVideoCodecIdOk() (*string, bool)`

GetVideoCodecIdOk returns a tuple with the VideoCodecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodecId

`func (o *KalturaMediaInfo) SetVideoCodecId(v string)`

SetVideoCodecId sets VideoCodecId field to given value.

### HasVideoCodecId

`func (o *KalturaMediaInfo) HasVideoCodecId() bool`

HasVideoCodecId returns a boolean if a field has been set.

### GetVideoDar

`func (o *KalturaMediaInfo) GetVideoDar() float32`

GetVideoDar returns the VideoDar field if non-nil, zero value otherwise.

### GetVideoDarOk

`func (o *KalturaMediaInfo) GetVideoDarOk() (*float32, bool)`

GetVideoDarOk returns a tuple with the VideoDar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDar

`func (o *KalturaMediaInfo) SetVideoDar(v float32)`

SetVideoDar sets VideoDar field to given value.

### HasVideoDar

`func (o *KalturaMediaInfo) HasVideoDar() bool`

HasVideoDar returns a boolean if a field has been set.

### GetVideoDuration

`func (o *KalturaMediaInfo) GetVideoDuration() int32`

GetVideoDuration returns the VideoDuration field if non-nil, zero value otherwise.

### GetVideoDurationOk

`func (o *KalturaMediaInfo) GetVideoDurationOk() (*int32, bool)`

GetVideoDurationOk returns a tuple with the VideoDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDuration

`func (o *KalturaMediaInfo) SetVideoDuration(v int32)`

SetVideoDuration sets VideoDuration field to given value.

### HasVideoDuration

`func (o *KalturaMediaInfo) HasVideoDuration() bool`

HasVideoDuration returns a boolean if a field has been set.

### GetVideoFormat

`func (o *KalturaMediaInfo) GetVideoFormat() string`

GetVideoFormat returns the VideoFormat field if non-nil, zero value otherwise.

### GetVideoFormatOk

`func (o *KalturaMediaInfo) GetVideoFormatOk() (*string, bool)`

GetVideoFormatOk returns a tuple with the VideoFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFormat

`func (o *KalturaMediaInfo) SetVideoFormat(v string)`

SetVideoFormat sets VideoFormat field to given value.

### HasVideoFormat

`func (o *KalturaMediaInfo) HasVideoFormat() bool`

HasVideoFormat returns a boolean if a field has been set.

### GetVideoFrameRate

`func (o *KalturaMediaInfo) GetVideoFrameRate() float32`

GetVideoFrameRate returns the VideoFrameRate field if non-nil, zero value otherwise.

### GetVideoFrameRateOk

`func (o *KalturaMediaInfo) GetVideoFrameRateOk() (*float32, bool)`

GetVideoFrameRateOk returns a tuple with the VideoFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFrameRate

`func (o *KalturaMediaInfo) SetVideoFrameRate(v float32)`

SetVideoFrameRate sets VideoFrameRate field to given value.

### HasVideoFrameRate

`func (o *KalturaMediaInfo) HasVideoFrameRate() bool`

HasVideoFrameRate returns a boolean if a field has been set.

### GetVideoHeight

`func (o *KalturaMediaInfo) GetVideoHeight() int32`

GetVideoHeight returns the VideoHeight field if non-nil, zero value otherwise.

### GetVideoHeightOk

`func (o *KalturaMediaInfo) GetVideoHeightOk() (*int32, bool)`

GetVideoHeightOk returns a tuple with the VideoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoHeight

`func (o *KalturaMediaInfo) SetVideoHeight(v int32)`

SetVideoHeight sets VideoHeight field to given value.

### HasVideoHeight

`func (o *KalturaMediaInfo) HasVideoHeight() bool`

HasVideoHeight returns a boolean if a field has been set.

### GetVideoRotation

`func (o *KalturaMediaInfo) GetVideoRotation() int32`

GetVideoRotation returns the VideoRotation field if non-nil, zero value otherwise.

### GetVideoRotationOk

`func (o *KalturaMediaInfo) GetVideoRotationOk() (*int32, bool)`

GetVideoRotationOk returns a tuple with the VideoRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRotation

`func (o *KalturaMediaInfo) SetVideoRotation(v int32)`

SetVideoRotation sets VideoRotation field to given value.

### HasVideoRotation

`func (o *KalturaMediaInfo) HasVideoRotation() bool`

HasVideoRotation returns a boolean if a field has been set.

### GetVideoWidth

`func (o *KalturaMediaInfo) GetVideoWidth() int32`

GetVideoWidth returns the VideoWidth field if non-nil, zero value otherwise.

### GetVideoWidthOk

`func (o *KalturaMediaInfo) GetVideoWidthOk() (*int32, bool)`

GetVideoWidthOk returns a tuple with the VideoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoWidth

`func (o *KalturaMediaInfo) SetVideoWidth(v int32)`

SetVideoWidth sets VideoWidth field to given value.

### HasVideoWidth

`func (o *KalturaMediaInfo) HasVideoWidth() bool`

HasVideoWidth returns a boolean if a field has been set.

### GetWritingLib

`func (o *KalturaMediaInfo) GetWritingLib() string`

GetWritingLib returns the WritingLib field if non-nil, zero value otherwise.

### GetWritingLibOk

`func (o *KalturaMediaInfo) GetWritingLibOk() (*string, bool)`

GetWritingLibOk returns a tuple with the WritingLib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritingLib

`func (o *KalturaMediaInfo) SetWritingLib(v string)`

SetWritingLib sets WritingLib field to given value.

### HasWritingLib

`func (o *KalturaMediaInfo) HasWritingLib() bool`

HasWritingLib returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


