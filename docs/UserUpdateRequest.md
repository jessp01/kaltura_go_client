# UserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**KalturaUser**](KalturaUser.md) |  | 
**UserId** | **string** |  | 

## Methods

### NewUserUpdateRequest

`func NewUserUpdateRequest(user KalturaUser, userId string, ) *UserUpdateRequest`

NewUserUpdateRequest instantiates a new UserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateRequestWithDefaults

`func NewUserUpdateRequestWithDefaults() *UserUpdateRequest`

NewUserUpdateRequestWithDefaults instantiates a new UserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserUpdateRequest) GetUser() KalturaUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserUpdateRequest) GetUserOk() (*KalturaUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserUpdateRequest) SetUser(v KalturaUser)`

SetUser sets User field to given value.


### GetUserId

`func (o *UserUpdateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


