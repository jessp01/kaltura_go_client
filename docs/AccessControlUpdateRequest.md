# AccessControlUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControl** | [**KalturaAccessControl**](KalturaAccessControl.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewAccessControlUpdateRequest

`func NewAccessControlUpdateRequest(accessControl KalturaAccessControl, id int32, ) *AccessControlUpdateRequest`

NewAccessControlUpdateRequest instantiates a new AccessControlUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlUpdateRequestWithDefaults

`func NewAccessControlUpdateRequestWithDefaults() *AccessControlUpdateRequest`

NewAccessControlUpdateRequestWithDefaults instantiates a new AccessControlUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControl

`func (o *AccessControlUpdateRequest) GetAccessControl() KalturaAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *AccessControlUpdateRequest) GetAccessControlOk() (*KalturaAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *AccessControlUpdateRequest) SetAccessControl(v KalturaAccessControl)`

SetAccessControl sets AccessControl field to given value.


### GetId

`func (o *AccessControlUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessControlUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessControlUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


