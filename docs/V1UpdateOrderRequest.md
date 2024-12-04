# V1UpdateOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [***V1UpdateOrderRequestAction**](V1UpdateOrderRequestAction.md) |  | [default to null]
**ShippedTrackingNumber** | **string** | The tracking number of the shipment associated with the order. Only valid if action is COMPLETE. | [optional] [default to null]
**CompletedNote** | **string** | A merchant-specified note about the completion of the order. Only valid if action is COMPLETE. | [optional] [default to null]
**RefundedNote** | **string** | A merchant-specified note about the refunding of the order. Only valid if action is REFUND. | [optional] [default to null]
**CanceledNote** | **string** | A merchant-specified note about the canceling of the order. Only valid if action is CANCEL. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

