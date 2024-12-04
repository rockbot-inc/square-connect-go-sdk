# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelInvoice**](InvoicesApi.md#CancelInvoice) | **Post** /v2/invoices/{invoice_id}/cancel | CancelInvoice
[**CreateInvoice**](InvoicesApi.md#CreateInvoice) | **Post** /v2/invoices | CreateInvoice
[**CreateInvoiceAttachment**](InvoicesApi.md#CreateInvoiceAttachment) | **Post** /v2/invoices/{invoice_id}/attachments | CreateInvoiceAttachment
[**DeleteInvoice**](InvoicesApi.md#DeleteInvoice) | **Delete** /v2/invoices/{invoice_id} | DeleteInvoice
[**DeleteInvoiceAttachment**](InvoicesApi.md#DeleteInvoiceAttachment) | **Delete** /v2/invoices/{invoice_id}/attachments/{attachment_id} | DeleteInvoiceAttachment
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /v2/invoices/{invoice_id} | GetInvoice
[**ListInvoices**](InvoicesApi.md#ListInvoices) | **Get** /v2/invoices | ListInvoices
[**PublishInvoice**](InvoicesApi.md#PublishInvoice) | **Post** /v2/invoices/{invoice_id}/publish | PublishInvoice
[**SearchInvoices**](InvoicesApi.md#SearchInvoices) | **Post** /v2/invoices/search | SearchInvoices
[**UpdateInvoice**](InvoicesApi.md#UpdateInvoice) | **Put** /v2/invoices/{invoice_id} | UpdateInvoice

# **CancelInvoice**
> CancelInvoiceResponse CancelInvoice(ctx, body, invoiceId)
CancelInvoice

Cancels an invoice. The seller cannot collect payments for  the canceled invoice.  You cannot cancel an invoice in the `DRAFT` state or in a terminal state: `PAID`, `REFUNDED`, `CANCELED`, or `FAILED`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelInvoiceRequest**](CancelInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **invoiceId** | **string**| The ID of the [invoice](entity:Invoice) to cancel. | 

### Return type

[**CancelInvoiceResponse**](CancelInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInvoice**
> CreateInvoiceResponse CreateInvoice(ctx, body)
CreateInvoice

Creates a draft [invoice]($m/Invoice)  for an order created using the Orders API.  A draft invoice remains in your account and no action is taken.  You must publish the invoice before Square can process it (send it to the customer's email address or charge the customer’s card on file).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInvoiceRequest**](CreateInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateInvoiceResponse**](CreateInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInvoiceAttachment**
> CreateInvoiceAttachmentResponse CreateInvoiceAttachment(ctx, request, imageFile, invoiceId)
CreateInvoiceAttachment

Uploads a file and attaches it to an invoice. This endpoint accepts HTTP multipart/form-data file uploads with a JSON `request` part and a `file` part. The `file` part must be a `readable stream` that contains a file in a supported format: GIF, JPEG, PNG, TIFF, BMP, or PDF.  Invoices can have up to 10 attachments with a total file size of 25 MB. Attachments can be added only to invoices in the `DRAFT`, `SCHEDULED`, `UNPAID`, or `PARTIALLY_PAID` state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**CreateInvoiceAttachmentRequest**](.md)|  | 
  **imageFile** | ***os.File*****os.File**|  | 
  **invoiceId** | **string**| The ID of the [invoice](entity:Invoice) to attach the file to. | 

### Return type

[**CreateInvoiceAttachmentResponse**](CreateInvoiceAttachmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInvoice**
> DeleteInvoiceResponse DeleteInvoice(ctx, invoiceId, optional)
DeleteInvoice

Deletes the specified invoice. When an invoice is deleted, the  associated order status changes to CANCELED. You can only delete a draft  invoice (you cannot delete a published invoice, including one that is scheduled for processing).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The ID of the invoice to delete. | 
 **optional** | ***InvoicesApiDeleteInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicesApiDeleteInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| The version of the [invoice](entity:Invoice) to delete. If you do not know the version, you can call [GetInvoice](api-endpoint:Invoices-GetInvoice) or  [ListInvoices](api-endpoint:Invoices-ListInvoices). | 

### Return type

[**DeleteInvoiceResponse**](DeleteInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInvoiceAttachment**
> DeleteInvoiceAttachmentResponse DeleteInvoiceAttachment(ctx, invoiceId, attachmentId)
DeleteInvoiceAttachment

Removes an attachment from an invoice and permanently deletes the file. Attachments can be removed only from invoices in the `DRAFT`, `SCHEDULED`, `UNPAID`, or `PARTIALLY_PAID` state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The ID of the [invoice](entity:Invoice) to delete the attachment from. | 
  **attachmentId** | **string**| The ID of the [attachment](entity:InvoiceAttachment) to delete. | 

### Return type

[**DeleteInvoiceAttachmentResponse**](DeleteInvoiceAttachmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoice**
> GetInvoiceResponse GetInvoice(ctx, invoiceId)
GetInvoice

Retrieves an invoice by invoice ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The ID of the invoice to retrieve. | 

### Return type

[**GetInvoiceResponse**](GetInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInvoices**
> ListInvoicesResponse ListInvoices(ctx, locationId, optional)
ListInvoices

Returns a list of invoices for a given location. The response  is paginated. If truncated, the response includes a `cursor` that you     use in a subsequent request to retrieve the next set of invoices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location for which to list invoices. | 
 **optional** | ***InvoicesApiListInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicesApiListInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint.  Provide this cursor to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **limit** | **optional.Int32**| The maximum number of invoices to return (200 is the maximum &#x60;limit&#x60;).  If not provided, the server uses a default limit of 100 invoices. | 

### Return type

[**ListInvoicesResponse**](ListInvoicesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishInvoice**
> PublishInvoiceResponse PublishInvoice(ctx, body, invoiceId)
PublishInvoice

Publishes the specified draft invoice.   After an invoice is published, Square  follows up based on the invoice configuration. For example, Square  sends the invoice to the customer's email address, charges the customer's card on file, or does  nothing. Square also makes the invoice available on a Square-hosted invoice page.   The invoice `status` also changes from `DRAFT` to a status  based on the invoice configuration. For example, the status changes to `UNPAID` if  Square emails the invoice or `PARTIALLY_PAID` if Square charges a card on file for a portion of the  invoice amount.  In addition to the required `ORDERS_WRITE` and `INVOICES_WRITE` permissions, `CUSTOMERS_READ` and `PAYMENTS_WRITE` are required when publishing invoices configured for card-on-file payments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublishInvoiceRequest**](PublishInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **invoiceId** | **string**| The ID of the invoice to publish. | 

### Return type

[**PublishInvoiceResponse**](PublishInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchInvoices**
> SearchInvoicesResponse SearchInvoices(ctx, body)
SearchInvoices

Searches for invoices from a location specified in  the filter. You can optionally specify customers in the filter for whom to  retrieve invoices. In the current implementation, you can only specify one location and  optionally one customer.  The response is paginated. If truncated, the response includes a `cursor`  that you use in a subsequent request to retrieve the next set of invoices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchInvoicesRequest**](SearchInvoicesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchInvoicesResponse**](SearchInvoicesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInvoice**
> UpdateInvoiceResponse UpdateInvoice(ctx, body, invoiceId)
UpdateInvoice

Updates an invoice. This endpoint supports sparse updates, so you only need to specify the fields you want to change along with the required `version` field. Some restrictions apply to updating invoices. For example, you cannot change the `order_id` or `location_id` field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateInvoiceRequest**](UpdateInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **invoiceId** | **string**| The ID of the invoice to update. | 

### Return type

[**UpdateInvoiceResponse**](UpdateInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

