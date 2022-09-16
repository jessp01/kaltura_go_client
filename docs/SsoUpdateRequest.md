# SsoUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sso** | [**KalturaSso**](KalturaSso.md) |  | 
**SsoId** | **int32** |  | 

## Methods

### NewSsoUpdateRequest

`func NewSsoUpdateRequest(sso KalturaSso, ssoId int32, ) *SsoUpdateRequest`

NewSsoUpdateRequest instantiates a new SsoUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoUpdateRequestWithDefaults

`func NewSsoUpdateRequestWithDefaults() *SsoUpdateRequest`

NewSsoUpdateRequestWithDefaults instantiates a new SsoUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSso

`func (o *SsoUpdateRequest) GetSso() KalturaSso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SsoUpdateRequest) GetSsoOk() (*KalturaSso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SsoUpdateRequest) SetSso(v KalturaSso)`

SetSso sets Sso field to given value.


### GetSsoId

`func (o *SsoUpdateRequest) GetSsoId() int32`

GetSsoId returns the SsoId field if non-nil, zero value otherwise.

### GetSsoIdOk

`func (o *SsoUpdateRequest) GetSsoIdOk() (*int32, bool)`

GetSsoIdOk returns a tuple with the SsoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoId

`func (o *SsoUpdateRequest) SetSsoId(v int32)`

SetSsoId sets SsoId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


