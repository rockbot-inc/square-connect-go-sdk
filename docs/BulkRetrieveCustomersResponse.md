# BulkRetrieveCustomersResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | [**map[string]RetrieveCustomerResponse**](RetrieveCustomerResponse.md) | A map of responses that correspond to individual retrieve requests, represented by key-value pairs.  Each key is the customer ID that was specified for a retrieve request and each value is the corresponding response. If the request succeeds, the value is the requested customer profile. If the request fails, the value contains any errors that occurred during the request. | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Any top-level errors that prevented the bulk operation from running. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

