# DeliveryProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivery** | [**KalturaDeliveryProfile**](KalturaDeliveryProfile.md) |  | 
**Id** | **string** |  | 

## Methods

### NewDeliveryProfileUpdateRequest

`func NewDeliveryProfileUpdateRequest(delivery KalturaDeliveryProfile, id string, ) *DeliveryProfileUpdateRequest`

NewDeliveryProfileUpdateRequest instantiates a new DeliveryProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryProfileUpdateRequestWithDefaults

`func NewDeliveryProfileUpdateRequestWithDefaults() *DeliveryProfileUpdateRequest`

NewDeliveryProfileUpdateRequestWithDefaults instantiates a new DeliveryProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivery

`func (o *DeliveryProfileUpdateRequest) GetDelivery() KalturaDeliveryProfile`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *DeliveryProfileUpdateRequest) GetDeliveryOk() (*KalturaDeliveryProfile, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *DeliveryProfileUpdateRequest) SetDelivery(v KalturaDeliveryProfile)`

SetDelivery sets Delivery field to given value.


### GetId

`func (o *DeliveryProfileUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryProfileUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryProfileUpdateRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


