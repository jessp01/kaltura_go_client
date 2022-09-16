# SystemPartnerResetUserPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | **string** |  | 
**PId** | **int32** |  | 
**UserId** | **string** |  | 

## Methods

### NewSystemPartnerResetUserPasswordRequest

`func NewSystemPartnerResetUserPasswordRequest(newPassword string, pId int32, userId string, ) *SystemPartnerResetUserPasswordRequest`

NewSystemPartnerResetUserPasswordRequest instantiates a new SystemPartnerResetUserPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPartnerResetUserPasswordRequestWithDefaults

`func NewSystemPartnerResetUserPasswordRequestWithDefaults() *SystemPartnerResetUserPasswordRequest`

NewSystemPartnerResetUserPasswordRequestWithDefaults instantiates a new SystemPartnerResetUserPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *SystemPartnerResetUserPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *SystemPartnerResetUserPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *SystemPartnerResetUserPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetPId

`func (o *SystemPartnerResetUserPasswordRequest) GetPId() int32`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *SystemPartnerResetUserPasswordRequest) GetPIdOk() (*int32, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *SystemPartnerResetUserPasswordRequest) SetPId(v int32)`

SetPId sets PId field to given value.


### GetUserId

`func (o *SystemPartnerResetUserPasswordRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SystemPartnerResetUserPasswordRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SystemPartnerResetUserPasswordRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


