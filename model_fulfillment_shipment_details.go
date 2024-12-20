/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Contains the details necessary to fulfill a shipment order.
type FulfillmentShipmentDetails struct {
	Recipient *FulfillmentRecipient `json:"recipient,omitempty"`
	// The shipping carrier being used to ship this fulfillment (such as UPS, FedEx, or USPS).
	Carrier string `json:"carrier,omitempty"`
	// A note with additional information for the shipping carrier.
	ShippingNote string `json:"shipping_note,omitempty"`
	// A description of the type of shipping product purchased from the carrier (such as First Class, Priority, or Express).
	ShippingType string `json:"shipping_type,omitempty"`
	// The reference number provided by the carrier to track the shipment's progress.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// A link to the tracking webpage on the carrier's website.
	TrackingUrl string `json:"tracking_url,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the shipment was requested. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	PlacedAt string `json:"placed_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when this fulfillment was moved to the `RESERVED` state, which  indicates that preparation of this shipment has begun. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	InProgressAt string `json:"in_progress_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when this fulfillment was moved to the `PREPARED` state, which indicates that the fulfillment is packaged. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	PackagedAt string `json:"packaged_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the shipment is expected to be delivered to the shipping carrier. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	ExpectedShippedAt string `json:"expected_shipped_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when this fulfillment was moved to the `COMPLETED` state, which indicates that the fulfillment has been given to the shipping carrier. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	ShippedAt string `json:"shipped_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating the shipment was canceled. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	CanceledAt string `json:"canceled_at,omitempty"`
	// A description of why the shipment was canceled.
	CancelReason string `json:"cancel_reason,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the shipment failed to be completed. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	FailedAt string `json:"failed_at,omitempty"`
	// A description of why the shipment failed to be completed.
	FailureReason string `json:"failure_reason,omitempty"`
}
