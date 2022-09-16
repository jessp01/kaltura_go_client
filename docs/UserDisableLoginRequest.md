# UserDisableLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDisableLoginRequest

`func NewUserDisableLoginRequest() *UserDisableLoginRequest`

NewUserDisableLoginRequest instantiates a new UserDisableLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDisableLoginRequestWithDefaults

`func NewUserDisableLoginRequestWithDefaults() *UserDisableLoginRequest`

NewUserDisableLoginRequestWithDefaults instantiates a new UserDisableLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginId

`func (o *UserDisableLoginRequest) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *UserDisableLoginRequest) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *UserDisableLoginRequest) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.

### HasLoginId

`func (o *UserDisableLoginRequest) HasLoginId() bool`

HasLoginId returns a boolean if a field has been set.

### GetUserId

`func (o *UserDisableLoginRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserDisableLoginRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserDisableLoginRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserDisableLoginRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


