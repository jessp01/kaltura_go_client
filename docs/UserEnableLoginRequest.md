# UserEnableLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginId** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewUserEnableLoginRequest

`func NewUserEnableLoginRequest(loginId string, userId string, ) *UserEnableLoginRequest`

NewUserEnableLoginRequest instantiates a new UserEnableLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEnableLoginRequestWithDefaults

`func NewUserEnableLoginRequestWithDefaults() *UserEnableLoginRequest`

NewUserEnableLoginRequestWithDefaults instantiates a new UserEnableLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginId

`func (o *UserEnableLoginRequest) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *UserEnableLoginRequest) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *UserEnableLoginRequest) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.


### GetPassword

`func (o *UserEnableLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserEnableLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserEnableLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserEnableLoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserId

`func (o *UserEnableLoginRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserEnableLoginRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserEnableLoginRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


