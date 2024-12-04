/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// V1UpdateOrderRequest
type V1UpdateOrderRequest struct {
	Action *V1UpdateOrderRequestAction `json:"action"`
	// The tracking number of the shipment associated with the order. Only valid if action is COMPLETE.
	ShippedTrackingNumber string `json:"shipped_tracking_number,omitempty"`
	// A merchant-specified note about the completion of the order. Only valid if action is COMPLETE.
	CompletedNote string `json:"completed_note,omitempty"`
	// A merchant-specified note about the refunding of the order. Only valid if action is REFUND.
	RefundedNote string `json:"refunded_note,omitempty"`
	// A merchant-specified note about the canceling of the order. Only valid if action is CANCEL.
	CanceledNote string `json:"canceled_note,omitempty"`
}