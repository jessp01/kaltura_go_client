# SsoLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | **string** |  | 
**PartnerId** | Pointer to **int32** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewSsoLoginRequest

`func NewSsoLoginRequest(applicationType string, userId string, ) *SsoLoginRequest`

NewSsoLoginRequest instantiates a new SsoLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoLoginRequestWithDefaults

`func NewSsoLoginRequestWithDefaults() *SsoLoginRequest`

NewSsoLoginRequestWithDefaults instantiates a new SsoLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *SsoLoginRequest) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *SsoLoginRequest) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *SsoLoginRequest) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.


### GetPartnerId

`func (o *SsoLoginRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *SsoLoginRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *SsoLoginRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *SsoLoginRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetUserId

`func (o *SsoLoginRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SsoLoginRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SsoLoginRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


