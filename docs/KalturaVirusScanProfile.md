# KalturaVirusScanProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionIfInfected** | Pointer to **int32** | Enum Type: &#x60;KalturaVirusFoundAction&#x60; | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**EngineType** | Pointer to **string** | Enum Type: &#x60;KalturaVirusScanEngineType&#x60; | [optional] 
**EntryFilter** | Pointer to [**KalturaBaseEntryFilter**](KalturaBaseEntryFilter.md) |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaVirusScanProfileStatus&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaVirusScanProfile

`func NewKalturaVirusScanProfile() *KalturaVirusScanProfile`

NewKalturaVirusScanProfile instantiates a new KalturaVirusScanProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaVirusScanProfileWithDefaults

`func NewKalturaVirusScanProfileWithDefaults() *KalturaVirusScanProfile`

NewKalturaVirusScanProfileWithDefaults instantiates a new KalturaVirusScanProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionIfInfected

`func (o *KalturaVirusScanProfile) GetActionIfInfected() int32`

GetActionIfInfected returns the ActionIfInfected field if non-nil, zero value otherwise.

### GetActionIfInfectedOk

`func (o *KalturaVirusScanProfile) GetActionIfInfectedOk() (*int32, bool)`

GetActionIfInfectedOk returns a tuple with the ActionIfInfected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionIfInfected

`func (o *KalturaVirusScanProfile) SetActionIfInfected(v int32)`

SetActionIfInfected sets ActionIfInfected field to given value.

### HasActionIfInfected

`func (o *KalturaVirusScanProfile) HasActionIfInfected() bool`

HasActionIfInfected returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaVirusScanProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaVirusScanProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaVirusScanProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaVirusScanProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEngineType

`func (o *KalturaVirusScanProfile) GetEngineType() string`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *KalturaVirusScanProfile) GetEngineTypeOk() (*string, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *KalturaVirusScanProfile) SetEngineType(v string)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *KalturaVirusScanProfile) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### GetEntryFilter

`func (o *KalturaVirusScanProfile) GetEntryFilter() KalturaBaseEntryFilter`

GetEntryFilter returns the EntryFilter field if non-nil, zero value otherwise.

### GetEntryFilterOk

`func (o *KalturaVirusScanProfile) GetEntryFilterOk() (*KalturaBaseEntryFilter, bool)`

GetEntryFilterOk returns a tuple with the EntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryFilter

`func (o *KalturaVirusScanProfile) SetEntryFilter(v KalturaBaseEntryFilter)`

SetEntryFilter sets EntryFilter field to given value.

### HasEntryFilter

`func (o *KalturaVirusScanProfile) HasEntryFilter() bool`

HasEntryFilter returns a boolean if a field has been set.

### GetId

`func (o *KalturaVirusScanProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaVirusScanProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaVirusScanProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaVirusScanProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaVirusScanProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaVirusScanProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaVirusScanProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaVirusScanProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaVirusScanProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaVirusScanProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaVirusScanProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaVirusScanProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaVirusScanProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaVirusScanProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaVirusScanProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaVirusScanProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaVirusScanProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaVirusScanProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaVirusScanProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaVirusScanProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


