# PlayReadyDrmGetLicenseDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**DeviceType** | **int32** |  | 
**EntryId** | Pointer to **string** |  | [optional] 
**KeyId** | **string** |  | 
**Referrer** | Pointer to **string** |  | [optional] 

## Methods

### NewPlayReadyDrmGetLicenseDetailsRequest

`func NewPlayReadyDrmGetLicenseDetailsRequest(deviceId string, deviceType int32, keyId string, ) *PlayReadyDrmGetLicenseDetailsRequest`

NewPlayReadyDrmGetLicenseDetailsRequest instantiates a new PlayReadyDrmGetLicenseDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayReadyDrmGetLicenseDetailsRequestWithDefaults

`func NewPlayReadyDrmGetLicenseDetailsRequestWithDefaults() *PlayReadyDrmGetLicenseDetailsRequest`

NewPlayReadyDrmGetLicenseDetailsRequestWithDefaults instantiates a new PlayReadyDrmGetLicenseDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PlayReadyDrmGetLicenseDetailsRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.


### GetEntryId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetKeyId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PlayReadyDrmGetLicenseDetailsRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetReferrer

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *PlayReadyDrmGetLicenseDetailsRequest) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *PlayReadyDrmGetLicenseDetailsRequest) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *PlayReadyDrmGetLicenseDetailsRequest) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


