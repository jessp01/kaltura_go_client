# KalturaVarPartnerUsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioEntriesCount** | Pointer to **int32** | Number of new audio entries created during specific date range | [optional] 
**AvgStorage** | Pointer to **float32** | The average amount of storage consumption during the given date range for the specific publisher | [optional] 
**Bandwidth** | Pointer to **float32** | The total bandwidth usage during the given date range (in MB) | [optional] 
**CombinedStorageBandwidth** | Pointer to **float32** | The combined amount of bandwidth and storage consumed during the given date range for the specific publisher | [optional] 
**DateId** | Pointer to **string** | TGhe date at which the report was taken - Unix Timestamp | [optional] 
**DeletedStorage** | Pointer to **float32** | The deleted storage consumption (new uploads) during the given date range (in MB) | [optional] 
**EntriesCount** | Pointer to **int32** | Number of new entries created during specific date range | [optional] 
**ImageEntriesCount** | Pointer to **int32** | Number of new image entries created during specific date range | [optional] 
**MixEntriesCount** | Pointer to **int32** | Number of new mix entries created during specific date range | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerCreatedAt** | Pointer to **int32** | Partner creation date (Unix timestamp) | [optional] 
**PartnerId** | Pointer to **int32** | Partner ID | [optional] 
**PartnerName** | Pointer to **string** | Partner name | [optional] 
**PartnerPackage** | Pointer to **int32** | Partner package | [optional] 
**PartnerStatus** | Pointer to **int32** | Enum Type: &#x60;KalturaPartnerStatus&#x60;  Partner status | [optional] 
**PeakStorage** | Pointer to **float32** | The peak amount of storage consumption during the given date range for the specific publisher | [optional] 
**Plays** | Pointer to **int32** | Number of plays in the specific date range | [optional] 
**Storage** | Pointer to **float32** | The added storage consumption (new uploads) during the given date range (in MB) | [optional] 
**TotalEntriesCount** | Pointer to **int32** | Total number of entries | [optional] 
**TotalStorage** | Pointer to **float32** | The total storage consumption (in MB) | [optional] 
**TranscodingUsage** | Pointer to **float32** | Amount of transcoding usage in MB | [optional] 
**VideoEntriesCount** | Pointer to **int32** | Number of new video entries created during specific date range | [optional] 
**Views** | Pointer to **int32** | Number of player loads in the specific date range | [optional] 

## Methods

### NewKalturaVarPartnerUsageItem

`func NewKalturaVarPartnerUsageItem() *KalturaVarPartnerUsageItem`

NewKalturaVarPartnerUsageItem instantiates a new KalturaVarPartnerUsageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaVarPartnerUsageItemWithDefaults

`func NewKalturaVarPartnerUsageItemWithDefaults() *KalturaVarPartnerUsageItem`

NewKalturaVarPartnerUsageItemWithDefaults instantiates a new KalturaVarPartnerUsageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioEntriesCount

`func (o *KalturaVarPartnerUsageItem) GetAudioEntriesCount() int32`

GetAudioEntriesCount returns the AudioEntriesCount field if non-nil, zero value otherwise.

### GetAudioEntriesCountOk

`func (o *KalturaVarPartnerUsageItem) GetAudioEntriesCountOk() (*int32, bool)`

GetAudioEntriesCountOk returns a tuple with the AudioEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEntriesCount

`func (o *KalturaVarPartnerUsageItem) SetAudioEntriesCount(v int32)`

SetAudioEntriesCount sets AudioEntriesCount field to given value.

### HasAudioEntriesCount

`func (o *KalturaVarPartnerUsageItem) HasAudioEntriesCount() bool`

HasAudioEntriesCount returns a boolean if a field has been set.

### GetAvgStorage

`func (o *KalturaVarPartnerUsageItem) GetAvgStorage() float32`

GetAvgStorage returns the AvgStorage field if non-nil, zero value otherwise.

### GetAvgStorageOk

`func (o *KalturaVarPartnerUsageItem) GetAvgStorageOk() (*float32, bool)`

GetAvgStorageOk returns a tuple with the AvgStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgStorage

`func (o *KalturaVarPartnerUsageItem) SetAvgStorage(v float32)`

SetAvgStorage sets AvgStorage field to given value.

### HasAvgStorage

`func (o *KalturaVarPartnerUsageItem) HasAvgStorage() bool`

HasAvgStorage returns a boolean if a field has been set.

### GetBandwidth

`func (o *KalturaVarPartnerUsageItem) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *KalturaVarPartnerUsageItem) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *KalturaVarPartnerUsageItem) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *KalturaVarPartnerUsageItem) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetCombinedStorageBandwidth

`func (o *KalturaVarPartnerUsageItem) GetCombinedStorageBandwidth() float32`

GetCombinedStorageBandwidth returns the CombinedStorageBandwidth field if non-nil, zero value otherwise.

### GetCombinedStorageBandwidthOk

`func (o *KalturaVarPartnerUsageItem) GetCombinedStorageBandwidthOk() (*float32, bool)`

GetCombinedStorageBandwidthOk returns a tuple with the CombinedStorageBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedStorageBandwidth

`func (o *KalturaVarPartnerUsageItem) SetCombinedStorageBandwidth(v float32)`

SetCombinedStorageBandwidth sets CombinedStorageBandwidth field to given value.

### HasCombinedStorageBandwidth

`func (o *KalturaVarPartnerUsageItem) HasCombinedStorageBandwidth() bool`

HasCombinedStorageBandwidth returns a boolean if a field has been set.

### GetDateId

`func (o *KalturaVarPartnerUsageItem) GetDateId() string`

GetDateId returns the DateId field if non-nil, zero value otherwise.

### GetDateIdOk

`func (o *KalturaVarPartnerUsageItem) GetDateIdOk() (*string, bool)`

GetDateIdOk returns a tuple with the DateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateId

`func (o *KalturaVarPartnerUsageItem) SetDateId(v string)`

SetDateId sets DateId field to given value.

### HasDateId

`func (o *KalturaVarPartnerUsageItem) HasDateId() bool`

HasDateId returns a boolean if a field has been set.

### GetDeletedStorage

`func (o *KalturaVarPartnerUsageItem) GetDeletedStorage() float32`

GetDeletedStorage returns the DeletedStorage field if non-nil, zero value otherwise.

### GetDeletedStorageOk

`func (o *KalturaVarPartnerUsageItem) GetDeletedStorageOk() (*float32, bool)`

GetDeletedStorageOk returns a tuple with the DeletedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedStorage

`func (o *KalturaVarPartnerUsageItem) SetDeletedStorage(v float32)`

SetDeletedStorage sets DeletedStorage field to given value.

### HasDeletedStorage

`func (o *KalturaVarPartnerUsageItem) HasDeletedStorage() bool`

HasDeletedStorage returns a boolean if a field has been set.

### GetEntriesCount

`func (o *KalturaVarPartnerUsageItem) GetEntriesCount() int32`

GetEntriesCount returns the EntriesCount field if non-nil, zero value otherwise.

### GetEntriesCountOk

`func (o *KalturaVarPartnerUsageItem) GetEntriesCountOk() (*int32, bool)`

GetEntriesCountOk returns a tuple with the EntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesCount

`func (o *KalturaVarPartnerUsageItem) SetEntriesCount(v int32)`

SetEntriesCount sets EntriesCount field to given value.

### HasEntriesCount

`func (o *KalturaVarPartnerUsageItem) HasEntriesCount() bool`

HasEntriesCount returns a boolean if a field has been set.

### GetImageEntriesCount

`func (o *KalturaVarPartnerUsageItem) GetImageEntriesCount() int32`

GetImageEntriesCount returns the ImageEntriesCount field if non-nil, zero value otherwise.

### GetImageEntriesCountOk

`func (o *KalturaVarPartnerUsageItem) GetImageEntriesCountOk() (*int32, bool)`

GetImageEntriesCountOk returns a tuple with the ImageEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageEntriesCount

`func (o *KalturaVarPartnerUsageItem) SetImageEntriesCount(v int32)`

SetImageEntriesCount sets ImageEntriesCount field to given value.

### HasImageEntriesCount

`func (o *KalturaVarPartnerUsageItem) HasImageEntriesCount() bool`

HasImageEntriesCount returns a boolean if a field has been set.

### GetMixEntriesCount

`func (o *KalturaVarPartnerUsageItem) GetMixEntriesCount() int32`

GetMixEntriesCount returns the MixEntriesCount field if non-nil, zero value otherwise.

### GetMixEntriesCountOk

`func (o *KalturaVarPartnerUsageItem) GetMixEntriesCountOk() (*int32, bool)`

GetMixEntriesCountOk returns a tuple with the MixEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixEntriesCount

`func (o *KalturaVarPartnerUsageItem) SetMixEntriesCount(v int32)`

SetMixEntriesCount sets MixEntriesCount field to given value.

### HasMixEntriesCount

`func (o *KalturaVarPartnerUsageItem) HasMixEntriesCount() bool`

HasMixEntriesCount returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaVarPartnerUsageItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaVarPartnerUsageItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaVarPartnerUsageItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaVarPartnerUsageItem) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerCreatedAt

`func (o *KalturaVarPartnerUsageItem) GetPartnerCreatedAt() int32`

GetPartnerCreatedAt returns the PartnerCreatedAt field if non-nil, zero value otherwise.

### GetPartnerCreatedAtOk

`func (o *KalturaVarPartnerUsageItem) GetPartnerCreatedAtOk() (*int32, bool)`

GetPartnerCreatedAtOk returns a tuple with the PartnerCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCreatedAt

`func (o *KalturaVarPartnerUsageItem) SetPartnerCreatedAt(v int32)`

SetPartnerCreatedAt sets PartnerCreatedAt field to given value.

### HasPartnerCreatedAt

`func (o *KalturaVarPartnerUsageItem) HasPartnerCreatedAt() bool`

HasPartnerCreatedAt returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaVarPartnerUsageItem) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaVarPartnerUsageItem) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaVarPartnerUsageItem) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaVarPartnerUsageItem) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerName

`func (o *KalturaVarPartnerUsageItem) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *KalturaVarPartnerUsageItem) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *KalturaVarPartnerUsageItem) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *KalturaVarPartnerUsageItem) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### GetPartnerPackage

`func (o *KalturaVarPartnerUsageItem) GetPartnerPackage() int32`

GetPartnerPackage returns the PartnerPackage field if non-nil, zero value otherwise.

### GetPartnerPackageOk

`func (o *KalturaVarPartnerUsageItem) GetPartnerPackageOk() (*int32, bool)`

GetPartnerPackageOk returns a tuple with the PartnerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPackage

`func (o *KalturaVarPartnerUsageItem) SetPartnerPackage(v int32)`

SetPartnerPackage sets PartnerPackage field to given value.

### HasPartnerPackage

`func (o *KalturaVarPartnerUsageItem) HasPartnerPackage() bool`

HasPartnerPackage returns a boolean if a field has been set.

### GetPartnerStatus

`func (o *KalturaVarPartnerUsageItem) GetPartnerStatus() int32`

GetPartnerStatus returns the PartnerStatus field if non-nil, zero value otherwise.

### GetPartnerStatusOk

`func (o *KalturaVarPartnerUsageItem) GetPartnerStatusOk() (*int32, bool)`

GetPartnerStatusOk returns a tuple with the PartnerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerStatus

`func (o *KalturaVarPartnerUsageItem) SetPartnerStatus(v int32)`

SetPartnerStatus sets PartnerStatus field to given value.

### HasPartnerStatus

`func (o *KalturaVarPartnerUsageItem) HasPartnerStatus() bool`

HasPartnerStatus returns a boolean if a field has been set.

### GetPeakStorage

`func (o *KalturaVarPartnerUsageItem) GetPeakStorage() float32`

GetPeakStorage returns the PeakStorage field if non-nil, zero value otherwise.

### GetPeakStorageOk

`func (o *KalturaVarPartnerUsageItem) GetPeakStorageOk() (*float32, bool)`

GetPeakStorageOk returns a tuple with the PeakStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakStorage

`func (o *KalturaVarPartnerUsageItem) SetPeakStorage(v float32)`

SetPeakStorage sets PeakStorage field to given value.

### HasPeakStorage

`func (o *KalturaVarPartnerUsageItem) HasPeakStorage() bool`

HasPeakStorage returns a boolean if a field has been set.

### GetPlays

`func (o *KalturaVarPartnerUsageItem) GetPlays() int32`

GetPlays returns the Plays field if non-nil, zero value otherwise.

### GetPlaysOk

`func (o *KalturaVarPartnerUsageItem) GetPlaysOk() (*int32, bool)`

GetPlaysOk returns a tuple with the Plays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlays

`func (o *KalturaVarPartnerUsageItem) SetPlays(v int32)`

SetPlays sets Plays field to given value.

### HasPlays

`func (o *KalturaVarPartnerUsageItem) HasPlays() bool`

HasPlays returns a boolean if a field has been set.

### GetStorage

`func (o *KalturaVarPartnerUsageItem) GetStorage() float32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *KalturaVarPartnerUsageItem) GetStorageOk() (*float32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *KalturaVarPartnerUsageItem) SetStorage(v float32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *KalturaVarPartnerUsageItem) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTotalEntriesCount

`func (o *KalturaVarPartnerUsageItem) GetTotalEntriesCount() int32`

GetTotalEntriesCount returns the TotalEntriesCount field if non-nil, zero value otherwise.

### GetTotalEntriesCountOk

`func (o *KalturaVarPartnerUsageItem) GetTotalEntriesCountOk() (*int32, bool)`

GetTotalEntriesCountOk returns a tuple with the TotalEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntriesCount

`func (o *KalturaVarPartnerUsageItem) SetTotalEntriesCount(v int32)`

SetTotalEntriesCount sets TotalEntriesCount field to given value.

### HasTotalEntriesCount

`func (o *KalturaVarPartnerUsageItem) HasTotalEntriesCount() bool`

HasTotalEntriesCount returns a boolean if a field has been set.

### GetTotalStorage

`func (o *KalturaVarPartnerUsageItem) GetTotalStorage() float32`

GetTotalStorage returns the TotalStorage field if non-nil, zero value otherwise.

### GetTotalStorageOk

`func (o *KalturaVarPartnerUsageItem) GetTotalStorageOk() (*float32, bool)`

GetTotalStorageOk returns a tuple with the TotalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorage

`func (o *KalturaVarPartnerUsageItem) SetTotalStorage(v float32)`

SetTotalStorage sets TotalStorage field to given value.

### HasTotalStorage

`func (o *KalturaVarPartnerUsageItem) HasTotalStorage() bool`

HasTotalStorage returns a boolean if a field has been set.

### GetTranscodingUsage

`func (o *KalturaVarPartnerUsageItem) GetTranscodingUsage() float32`

GetTranscodingUsage returns the TranscodingUsage field if non-nil, zero value otherwise.

### GetTranscodingUsageOk

`func (o *KalturaVarPartnerUsageItem) GetTranscodingUsageOk() (*float32, bool)`

GetTranscodingUsageOk returns a tuple with the TranscodingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingUsage

`func (o *KalturaVarPartnerUsageItem) SetTranscodingUsage(v float32)`

SetTranscodingUsage sets TranscodingUsage field to given value.

### HasTranscodingUsage

`func (o *KalturaVarPartnerUsageItem) HasTranscodingUsage() bool`

HasTranscodingUsage returns a boolean if a field has been set.

### GetVideoEntriesCount

`func (o *KalturaVarPartnerUsageItem) GetVideoEntriesCount() int32`

GetVideoEntriesCount returns the VideoEntriesCount field if non-nil, zero value otherwise.

### GetVideoEntriesCountOk

`func (o *KalturaVarPartnerUsageItem) GetVideoEntriesCountOk() (*int32, bool)`

GetVideoEntriesCountOk returns a tuple with the VideoEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEntriesCount

`func (o *KalturaVarPartnerUsageItem) SetVideoEntriesCount(v int32)`

SetVideoEntriesCount sets VideoEntriesCount field to given value.

### HasVideoEntriesCount

`func (o *KalturaVarPartnerUsageItem) HasVideoEntriesCount() bool`

HasVideoEntriesCount returns a boolean if a field has been set.

### GetViews

`func (o *KalturaVarPartnerUsageItem) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *KalturaVarPartnerUsageItem) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *KalturaVarPartnerUsageItem) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *KalturaVarPartnerUsageItem) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


