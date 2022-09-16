# KalturaEntryDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIds** | Pointer to **string** | Comma separated asset ids | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry distribution creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**DirtyStatus** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryDistributionFlag&#x60; | [optional] [readonly] 
**DistributionProfileId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**EntryId** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ErrorNumber** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ErrorType** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaBatchJobErrorTypes&#x60; | [optional] [readonly] 
**FlavorAssetIds** | Pointer to **string** | Comma separated flavor asset ids | [optional] 
**HasDeleteResultsLog** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] [readonly] 
**HasDeleteSentDataLog** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] [readonly] 
**HasSubmitResultsLog** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] [readonly] 
**HasSubmitSentDataLog** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] [readonly] 
**HasUpdateResultsLog** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] [readonly] 
**HasUpdateSentDataLog** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Auto generated unique id | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Plays** | Pointer to **int32** | &#x60;readOnly&#x60;  The plays as retrieved from the remote destination reports | [optional] [readonly] 
**RemoteId** | Pointer to **string** | &#x60;readOnly&#x60;  The id as returned from the distributed destination | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryDistributionStatus&#x60; | [optional] [readonly] 
**SubmittedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry distribution submission date as Unix timestamp (In seconds) | [optional] [readonly] 
**SunStatus** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryDistributionSunStatus&#x60; | [optional] [readonly] 
**Sunrise** | Pointer to **int32** | Entry distribution publish time as Unix timestamp (In seconds) | [optional] 
**Sunset** | Pointer to **int32** | Entry distribution un-publish time as Unix timestamp (In seconds) | [optional] 
**ThumbAssetIds** | Pointer to **string** | Comma separated thumbnail asset ids | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry distribution last update date as Unix timestamp (In seconds) | [optional] [readonly] 
**ValidationErrors** | Pointer to [**[]KalturaDistributionValidationError**](KalturaDistributionValidationError.md) |  | [optional] 
**Views** | Pointer to **int32** | &#x60;readOnly&#x60;  The views as retrieved from the remote destination reports | [optional] [readonly] 

## Methods

### NewKalturaEntryDistribution

`func NewKalturaEntryDistribution() *KalturaEntryDistribution`

NewKalturaEntryDistribution instantiates a new KalturaEntryDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEntryDistributionWithDefaults

`func NewKalturaEntryDistributionWithDefaults() *KalturaEntryDistribution`

NewKalturaEntryDistributionWithDefaults instantiates a new KalturaEntryDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIds

`func (o *KalturaEntryDistribution) GetAssetIds() string`

GetAssetIds returns the AssetIds field if non-nil, zero value otherwise.

### GetAssetIdsOk

`func (o *KalturaEntryDistribution) GetAssetIdsOk() (*string, bool)`

GetAssetIdsOk returns a tuple with the AssetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIds

`func (o *KalturaEntryDistribution) SetAssetIds(v string)`

SetAssetIds sets AssetIds field to given value.

### HasAssetIds

`func (o *KalturaEntryDistribution) HasAssetIds() bool`

HasAssetIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaEntryDistribution) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaEntryDistribution) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaEntryDistribution) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaEntryDistribution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirtyStatus

`func (o *KalturaEntryDistribution) GetDirtyStatus() int32`

GetDirtyStatus returns the DirtyStatus field if non-nil, zero value otherwise.

### GetDirtyStatusOk

`func (o *KalturaEntryDistribution) GetDirtyStatusOk() (*int32, bool)`

GetDirtyStatusOk returns a tuple with the DirtyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirtyStatus

`func (o *KalturaEntryDistribution) SetDirtyStatus(v int32)`

SetDirtyStatus sets DirtyStatus field to given value.

### HasDirtyStatus

`func (o *KalturaEntryDistribution) HasDirtyStatus() bool`

HasDirtyStatus returns a boolean if a field has been set.

### GetDistributionProfileId

`func (o *KalturaEntryDistribution) GetDistributionProfileId() int32`

GetDistributionProfileId returns the DistributionProfileId field if non-nil, zero value otherwise.

### GetDistributionProfileIdOk

`func (o *KalturaEntryDistribution) GetDistributionProfileIdOk() (*int32, bool)`

GetDistributionProfileIdOk returns a tuple with the DistributionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionProfileId

`func (o *KalturaEntryDistribution) SetDistributionProfileId(v int32)`

SetDistributionProfileId sets DistributionProfileId field to given value.

### HasDistributionProfileId

`func (o *KalturaEntryDistribution) HasDistributionProfileId() bool`

HasDistributionProfileId returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaEntryDistribution) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaEntryDistribution) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaEntryDistribution) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaEntryDistribution) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaEntryDistribution) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaEntryDistribution) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaEntryDistribution) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaEntryDistribution) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorNumber

`func (o *KalturaEntryDistribution) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *KalturaEntryDistribution) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *KalturaEntryDistribution) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumber

`func (o *KalturaEntryDistribution) HasErrorNumber() bool`

HasErrorNumber returns a boolean if a field has been set.

### GetErrorType

`func (o *KalturaEntryDistribution) GetErrorType() int32`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *KalturaEntryDistribution) GetErrorTypeOk() (*int32, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *KalturaEntryDistribution) SetErrorType(v int32)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *KalturaEntryDistribution) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetFlavorAssetIds

`func (o *KalturaEntryDistribution) GetFlavorAssetIds() string`

GetFlavorAssetIds returns the FlavorAssetIds field if non-nil, zero value otherwise.

### GetFlavorAssetIdsOk

`func (o *KalturaEntryDistribution) GetFlavorAssetIdsOk() (*string, bool)`

GetFlavorAssetIdsOk returns a tuple with the FlavorAssetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorAssetIds

`func (o *KalturaEntryDistribution) SetFlavorAssetIds(v string)`

SetFlavorAssetIds sets FlavorAssetIds field to given value.

### HasFlavorAssetIds

`func (o *KalturaEntryDistribution) HasFlavorAssetIds() bool`

HasFlavorAssetIds returns a boolean if a field has been set.

### GetHasDeleteResultsLog

`func (o *KalturaEntryDistribution) GetHasDeleteResultsLog() int32`

GetHasDeleteResultsLog returns the HasDeleteResultsLog field if non-nil, zero value otherwise.

### GetHasDeleteResultsLogOk

`func (o *KalturaEntryDistribution) GetHasDeleteResultsLogOk() (*int32, bool)`

GetHasDeleteResultsLogOk returns a tuple with the HasDeleteResultsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeleteResultsLog

`func (o *KalturaEntryDistribution) SetHasDeleteResultsLog(v int32)`

SetHasDeleteResultsLog sets HasDeleteResultsLog field to given value.

### HasHasDeleteResultsLog

`func (o *KalturaEntryDistribution) HasHasDeleteResultsLog() bool`

HasHasDeleteResultsLog returns a boolean if a field has been set.

### GetHasDeleteSentDataLog

`func (o *KalturaEntryDistribution) GetHasDeleteSentDataLog() int32`

GetHasDeleteSentDataLog returns the HasDeleteSentDataLog field if non-nil, zero value otherwise.

### GetHasDeleteSentDataLogOk

`func (o *KalturaEntryDistribution) GetHasDeleteSentDataLogOk() (*int32, bool)`

GetHasDeleteSentDataLogOk returns a tuple with the HasDeleteSentDataLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeleteSentDataLog

`func (o *KalturaEntryDistribution) SetHasDeleteSentDataLog(v int32)`

SetHasDeleteSentDataLog sets HasDeleteSentDataLog field to given value.

### HasHasDeleteSentDataLog

`func (o *KalturaEntryDistribution) HasHasDeleteSentDataLog() bool`

HasHasDeleteSentDataLog returns a boolean if a field has been set.

### GetHasSubmitResultsLog

`func (o *KalturaEntryDistribution) GetHasSubmitResultsLog() int32`

GetHasSubmitResultsLog returns the HasSubmitResultsLog field if non-nil, zero value otherwise.

### GetHasSubmitResultsLogOk

`func (o *KalturaEntryDistribution) GetHasSubmitResultsLogOk() (*int32, bool)`

GetHasSubmitResultsLogOk returns a tuple with the HasSubmitResultsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubmitResultsLog

`func (o *KalturaEntryDistribution) SetHasSubmitResultsLog(v int32)`

SetHasSubmitResultsLog sets HasSubmitResultsLog field to given value.

### HasHasSubmitResultsLog

`func (o *KalturaEntryDistribution) HasHasSubmitResultsLog() bool`

HasHasSubmitResultsLog returns a boolean if a field has been set.

### GetHasSubmitSentDataLog

`func (o *KalturaEntryDistribution) GetHasSubmitSentDataLog() int32`

GetHasSubmitSentDataLog returns the HasSubmitSentDataLog field if non-nil, zero value otherwise.

### GetHasSubmitSentDataLogOk

`func (o *KalturaEntryDistribution) GetHasSubmitSentDataLogOk() (*int32, bool)`

GetHasSubmitSentDataLogOk returns a tuple with the HasSubmitSentDataLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubmitSentDataLog

`func (o *KalturaEntryDistribution) SetHasSubmitSentDataLog(v int32)`

SetHasSubmitSentDataLog sets HasSubmitSentDataLog field to given value.

### HasHasSubmitSentDataLog

`func (o *KalturaEntryDistribution) HasHasSubmitSentDataLog() bool`

HasHasSubmitSentDataLog returns a boolean if a field has been set.

### GetHasUpdateResultsLog

`func (o *KalturaEntryDistribution) GetHasUpdateResultsLog() int32`

GetHasUpdateResultsLog returns the HasUpdateResultsLog field if non-nil, zero value otherwise.

### GetHasUpdateResultsLogOk

`func (o *KalturaEntryDistribution) GetHasUpdateResultsLogOk() (*int32, bool)`

GetHasUpdateResultsLogOk returns a tuple with the HasUpdateResultsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateResultsLog

`func (o *KalturaEntryDistribution) SetHasUpdateResultsLog(v int32)`

SetHasUpdateResultsLog sets HasUpdateResultsLog field to given value.

### HasHasUpdateResultsLog

`func (o *KalturaEntryDistribution) HasHasUpdateResultsLog() bool`

HasHasUpdateResultsLog returns a boolean if a field has been set.

### GetHasUpdateSentDataLog

`func (o *KalturaEntryDistribution) GetHasUpdateSentDataLog() int32`

GetHasUpdateSentDataLog returns the HasUpdateSentDataLog field if non-nil, zero value otherwise.

### GetHasUpdateSentDataLogOk

`func (o *KalturaEntryDistribution) GetHasUpdateSentDataLogOk() (*int32, bool)`

GetHasUpdateSentDataLogOk returns a tuple with the HasUpdateSentDataLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateSentDataLog

`func (o *KalturaEntryDistribution) SetHasUpdateSentDataLog(v int32)`

SetHasUpdateSentDataLog sets HasUpdateSentDataLog field to given value.

### HasHasUpdateSentDataLog

`func (o *KalturaEntryDistribution) HasHasUpdateSentDataLog() bool`

HasHasUpdateSentDataLog returns a boolean if a field has been set.

### GetId

`func (o *KalturaEntryDistribution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaEntryDistribution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaEntryDistribution) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaEntryDistribution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaEntryDistribution) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaEntryDistribution) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaEntryDistribution) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaEntryDistribution) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPlays

`func (o *KalturaEntryDistribution) GetPlays() int32`

GetPlays returns the Plays field if non-nil, zero value otherwise.

### GetPlaysOk

`func (o *KalturaEntryDistribution) GetPlaysOk() (*int32, bool)`

GetPlaysOk returns a tuple with the Plays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlays

`func (o *KalturaEntryDistribution) SetPlays(v int32)`

SetPlays sets Plays field to given value.

### HasPlays

`func (o *KalturaEntryDistribution) HasPlays() bool`

HasPlays returns a boolean if a field has been set.

### GetRemoteId

`func (o *KalturaEntryDistribution) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *KalturaEntryDistribution) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *KalturaEntryDistribution) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *KalturaEntryDistribution) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaEntryDistribution) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaEntryDistribution) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaEntryDistribution) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaEntryDistribution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *KalturaEntryDistribution) GetSubmittedAt() int32`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *KalturaEntryDistribution) GetSubmittedAtOk() (*int32, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *KalturaEntryDistribution) SetSubmittedAt(v int32)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *KalturaEntryDistribution) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetSunStatus

`func (o *KalturaEntryDistribution) GetSunStatus() int32`

GetSunStatus returns the SunStatus field if non-nil, zero value otherwise.

### GetSunStatusOk

`func (o *KalturaEntryDistribution) GetSunStatusOk() (*int32, bool)`

GetSunStatusOk returns a tuple with the SunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunStatus

`func (o *KalturaEntryDistribution) SetSunStatus(v int32)`

SetSunStatus sets SunStatus field to given value.

### HasSunStatus

`func (o *KalturaEntryDistribution) HasSunStatus() bool`

HasSunStatus returns a boolean if a field has been set.

### GetSunrise

`func (o *KalturaEntryDistribution) GetSunrise() int32`

GetSunrise returns the Sunrise field if non-nil, zero value otherwise.

### GetSunriseOk

`func (o *KalturaEntryDistribution) GetSunriseOk() (*int32, bool)`

GetSunriseOk returns a tuple with the Sunrise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunrise

`func (o *KalturaEntryDistribution) SetSunrise(v int32)`

SetSunrise sets Sunrise field to given value.

### HasSunrise

`func (o *KalturaEntryDistribution) HasSunrise() bool`

HasSunrise returns a boolean if a field has been set.

### GetSunset

`func (o *KalturaEntryDistribution) GetSunset() int32`

GetSunset returns the Sunset field if non-nil, zero value otherwise.

### GetSunsetOk

`func (o *KalturaEntryDistribution) GetSunsetOk() (*int32, bool)`

GetSunsetOk returns a tuple with the Sunset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunset

`func (o *KalturaEntryDistribution) SetSunset(v int32)`

SetSunset sets Sunset field to given value.

### HasSunset

`func (o *KalturaEntryDistribution) HasSunset() bool`

HasSunset returns a boolean if a field has been set.

### GetThumbAssetIds

`func (o *KalturaEntryDistribution) GetThumbAssetIds() string`

GetThumbAssetIds returns the ThumbAssetIds field if non-nil, zero value otherwise.

### GetThumbAssetIdsOk

`func (o *KalturaEntryDistribution) GetThumbAssetIdsOk() (*string, bool)`

GetThumbAssetIdsOk returns a tuple with the ThumbAssetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbAssetIds

`func (o *KalturaEntryDistribution) SetThumbAssetIds(v string)`

SetThumbAssetIds sets ThumbAssetIds field to given value.

### HasThumbAssetIds

`func (o *KalturaEntryDistribution) HasThumbAssetIds() bool`

HasThumbAssetIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaEntryDistribution) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaEntryDistribution) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaEntryDistribution) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaEntryDistribution) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValidationErrors

`func (o *KalturaEntryDistribution) GetValidationErrors() []KalturaDistributionValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *KalturaEntryDistribution) GetValidationErrorsOk() (*[]KalturaDistributionValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *KalturaEntryDistribution) SetValidationErrors(v []KalturaDistributionValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *KalturaEntryDistribution) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetViews

`func (o *KalturaEntryDistribution) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *KalturaEntryDistribution) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *KalturaEntryDistribution) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *KalturaEntryDistribution) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


