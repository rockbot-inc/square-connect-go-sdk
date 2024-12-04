# RetrieveOrderCustomAttributeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control, include this optional field and specify the current version of the custom attribute. | [optional] [default to null]
**WithDefinition** | **bool** | Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the &#x60;definition&#x60; field of each  custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute,  information about the data type, or other definition details. The default value is &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

