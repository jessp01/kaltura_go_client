# SystemPartnerGetAdminSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PId** | **int32** |  | 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemPartnerGetAdminSessionRequest

`func NewSystemPartnerGetAdminSessionRequest(pId int32, ) *SystemPartnerGetAdminSessionRequest`

NewSystemPartnerGetAdminSessionRequest instantiates a new SystemPartnerGetAdminSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPartnerGetAdminSessionRequestWithDefaults

`func NewSystemPartnerGetAdminSessionRequestWithDefaults() *SystemPartnerGetAdminSessionRequest`

NewSystemPartnerGetAdminSessionRequestWithDefaults instantiates a new SystemPartnerGetAdminSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPId

`func (o *SystemPartnerGetAdminSessionRequest) GetPId() int32`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *SystemPartnerGetAdminSessionRequest) GetPIdOk() (*int32, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *SystemPartnerGetAdminSessionRequest) SetPId(v int32)`

SetPId sets PId field to given value.


### GetUserId

`func (o *SystemPartnerGetAdminSessionRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SystemPartnerGetAdminSessionRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SystemPartnerGetAdminSessionRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SystemPartnerGetAdminSessionRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


