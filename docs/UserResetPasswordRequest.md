# UserResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**LinkType** | Pointer to **string** |  | [optional] 

## Methods

### NewUserResetPasswordRequest

`func NewUserResetPasswordRequest(email string, ) *UserResetPasswordRequest`

NewUserResetPasswordRequest instantiates a new UserResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResetPasswordRequestWithDefaults

`func NewUserResetPasswordRequestWithDefaults() *UserResetPasswordRequest`

NewUserResetPasswordRequestWithDefaults instantiates a new UserResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserResetPasswordRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResetPasswordRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResetPasswordRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLinkType

`func (o *UserResetPasswordRequest) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *UserResetPasswordRequest) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *UserResetPasswordRequest) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *UserResetPasswordRequest) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


