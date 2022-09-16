# AnnotationAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | [**KalturaCuePoint**](KalturaCuePoint.md) |  | 

## Methods

### NewAnnotationAddRequest

`func NewAnnotationAddRequest(annotation KalturaCuePoint, ) *AnnotationAddRequest`

NewAnnotationAddRequest instantiates a new AnnotationAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationAddRequestWithDefaults

`func NewAnnotationAddRequestWithDefaults() *AnnotationAddRequest`

NewAnnotationAddRequestWithDefaults instantiates a new AnnotationAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *AnnotationAddRequest) GetAnnotation() KalturaCuePoint`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *AnnotationAddRequest) GetAnnotationOk() (*KalturaCuePoint, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *AnnotationAddRequest) SetAnnotation(v KalturaCuePoint)`

SetAnnotation sets Annotation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


