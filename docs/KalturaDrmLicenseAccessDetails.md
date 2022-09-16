# KalturaDrmLicenseAccessDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteDuration** | Pointer to **int32** | playback window in seconds | [optional] 
**Duration** | Pointer to **int32** | movie duration in seconds | [optional] 
**LicenseParams** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**Policy** | Pointer to **string** | Drm policy name | [optional] 

## Methods

### NewKalturaDrmLicenseAccessDetails

`func NewKalturaDrmLicenseAccessDetails() *KalturaDrmLicenseAccessDetails`

NewKalturaDrmLicenseAccessDetails instantiates a new KalturaDrmLicenseAccessDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDrmLicenseAccessDetailsWithDefaults

`func NewKalturaDrmLicenseAccessDetailsWithDefaults() *KalturaDrmLicenseAccessDetails`

NewKalturaDrmLicenseAccessDetailsWithDefaults instantiates a new KalturaDrmLicenseAccessDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteDuration

`func (o *KalturaDrmLicenseAccessDetails) GetAbsoluteDuration() int32`

GetAbsoluteDuration returns the AbsoluteDuration field if non-nil, zero value otherwise.

### GetAbsoluteDurationOk

`func (o *KalturaDrmLicenseAccessDetails) GetAbsoluteDurationOk() (*int32, bool)`

GetAbsoluteDurationOk returns a tuple with the AbsoluteDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteDuration

`func (o *KalturaDrmLicenseAccessDetails) SetAbsoluteDuration(v int32)`

SetAbsoluteDuration sets AbsoluteDuration field to given value.

### HasAbsoluteDuration

`func (o *KalturaDrmLicenseAccessDetails) HasAbsoluteDuration() bool`

HasAbsoluteDuration returns a boolean if a field has been set.

### GetDuration

`func (o *KalturaDrmLicenseAccessDetails) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KalturaDrmLicenseAccessDetails) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KalturaDrmLicenseAccessDetails) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *KalturaDrmLicenseAccessDetails) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLicenseParams

`func (o *KalturaDrmLicenseAccessDetails) GetLicenseParams() []KalturaKeyValue`

GetLicenseParams returns the LicenseParams field if non-nil, zero value otherwise.

### GetLicenseParamsOk

`func (o *KalturaDrmLicenseAccessDetails) GetLicenseParamsOk() (*[]KalturaKeyValue, bool)`

GetLicenseParamsOk returns a tuple with the LicenseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseParams

`func (o *KalturaDrmLicenseAccessDetails) SetLicenseParams(v []KalturaKeyValue)`

SetLicenseParams sets LicenseParams field to given value.

### HasLicenseParams

`func (o *KalturaDrmLicenseAccessDetails) HasLicenseParams() bool`

HasLicenseParams returns a boolean if a field has been set.

### GetPolicy

`func (o *KalturaDrmLicenseAccessDetails) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *KalturaDrmLicenseAccessDetails) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *KalturaDrmLicenseAccessDetails) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *KalturaDrmLicenseAccessDetails) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


