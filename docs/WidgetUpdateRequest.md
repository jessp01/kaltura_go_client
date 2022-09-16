# WidgetUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Widget** | [**KalturaWidget**](KalturaWidget.md) |  | 

## Methods

### NewWidgetUpdateRequest

`func NewWidgetUpdateRequest(id string, widget KalturaWidget, ) *WidgetUpdateRequest`

NewWidgetUpdateRequest instantiates a new WidgetUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetUpdateRequestWithDefaults

`func NewWidgetUpdateRequestWithDefaults() *WidgetUpdateRequest`

NewWidgetUpdateRequestWithDefaults instantiates a new WidgetUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WidgetUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WidgetUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WidgetUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetWidget

`func (o *WidgetUpdateRequest) GetWidget() KalturaWidget`

GetWidget returns the Widget field if non-nil, zero value otherwise.

### GetWidgetOk

`func (o *WidgetUpdateRequest) GetWidgetOk() (*KalturaWidget, bool)`

GetWidgetOk returns a tuple with the Widget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidget

`func (o *WidgetUpdateRequest) SetWidget(v KalturaWidget)`

SetWidget sets Widget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


