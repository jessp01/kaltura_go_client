# KalturaLockFileSyncsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** |  | [optional] 
**DcSecret** | Pointer to **string** |  | [optional] 
**FileSyncs** | Pointer to [**[]KalturaFileSync**](KalturaFileSync.md) |  | [optional] 
**LimitReached** | Pointer to **bool** |  | [optional] 

## Methods

### NewKalturaLockFileSyncsResponse

`func NewKalturaLockFileSyncsResponse() *KalturaLockFileSyncsResponse`

NewKalturaLockFileSyncsResponse instantiates a new KalturaLockFileSyncsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLockFileSyncsResponseWithDefaults

`func NewKalturaLockFileSyncsResponseWithDefaults() *KalturaLockFileSyncsResponse`

NewKalturaLockFileSyncsResponseWithDefaults instantiates a new KalturaLockFileSyncsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *KalturaLockFileSyncsResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *KalturaLockFileSyncsResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *KalturaLockFileSyncsResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *KalturaLockFileSyncsResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetDcSecret

`func (o *KalturaLockFileSyncsResponse) GetDcSecret() string`

GetDcSecret returns the DcSecret field if non-nil, zero value otherwise.

### GetDcSecretOk

`func (o *KalturaLockFileSyncsResponse) GetDcSecretOk() (*string, bool)`

GetDcSecretOk returns a tuple with the DcSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSecret

`func (o *KalturaLockFileSyncsResponse) SetDcSecret(v string)`

SetDcSecret sets DcSecret field to given value.

### HasDcSecret

`func (o *KalturaLockFileSyncsResponse) HasDcSecret() bool`

HasDcSecret returns a boolean if a field has been set.

### GetFileSyncs

`func (o *KalturaLockFileSyncsResponse) GetFileSyncs() []KalturaFileSync`

GetFileSyncs returns the FileSyncs field if non-nil, zero value otherwise.

### GetFileSyncsOk

`func (o *KalturaLockFileSyncsResponse) GetFileSyncsOk() (*[]KalturaFileSync, bool)`

GetFileSyncsOk returns a tuple with the FileSyncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSyncs

`func (o *KalturaLockFileSyncsResponse) SetFileSyncs(v []KalturaFileSync)`

SetFileSyncs sets FileSyncs field to given value.

### HasFileSyncs

`func (o *KalturaLockFileSyncsResponse) HasFileSyncs() bool`

HasFileSyncs returns a boolean if a field has been set.

### GetLimitReached

`func (o *KalturaLockFileSyncsResponse) GetLimitReached() bool`

GetLimitReached returns the LimitReached field if non-nil, zero value otherwise.

### GetLimitReachedOk

`func (o *KalturaLockFileSyncsResponse) GetLimitReachedOk() (*bool, bool)`

GetLimitReachedOk returns a tuple with the LimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitReached

`func (o *KalturaLockFileSyncsResponse) SetLimitReached(v bool)`

SetLimitReached sets LimitReached field to given value.

### HasLimitReached

`func (o *KalturaLockFileSyncsResponse) HasLimitReached() bool`

HasLimitReached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


