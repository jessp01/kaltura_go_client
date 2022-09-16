# AuthGetAuthInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddPermissions** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthGetAuthInfoRequest

`func NewAuthGetAuthInfoRequest() *AuthGetAuthInfoRequest`

NewAuthGetAuthInfoRequest instantiates a new AuthGetAuthInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGetAuthInfoRequestWithDefaults

`func NewAuthGetAuthInfoRequestWithDefaults() *AuthGetAuthInfoRequest`

NewAuthGetAuthInfoRequestWithDefaults instantiates a new AuthGetAuthInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddPermissions

`func (o *AuthGetAuthInfoRequest) GetAddPermissions() bool`

GetAddPermissions returns the AddPermissions field if non-nil, zero value otherwise.

### GetAddPermissionsOk

`func (o *AuthGetAuthInfoRequest) GetAddPermissionsOk() (*bool, bool)`

GetAddPermissionsOk returns a tuple with the AddPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPermissions

`func (o *AuthGetAuthInfoRequest) SetAddPermissions(v bool)`

SetAddPermissions sets AddPermissions field to given value.

### HasAddPermissions

`func (o *AuthGetAuthInfoRequest) HasAddPermissions() bool`

HasAddPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


