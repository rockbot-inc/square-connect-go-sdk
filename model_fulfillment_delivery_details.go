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

// Describes delivery details of an order fulfillment.
type FulfillmentDeliveryDetails struct {
	Recipient *FulfillmentRecipient `json:"recipient,omitempty"`
	ScheduleType *FulfillmentDeliveryDetailsOrderFulfillmentDeliveryDetailsScheduleType `json:"schedule_type,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the fulfillment was placed. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").  Must be in RFC 3339 timestamp format, e.g., \"2016-09-04T23:59:33.123Z\".
	PlacedAt string `json:"placed_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) that represents the start of the delivery period. When the fulfillment `schedule_type` is `ASAP`, the field is automatically set to the current time plus the `prep_time_duration`. Otherwise, the application can set this field while the fulfillment `state` is `PROPOSED`, `RESERVED`, or `PREPARED` (any time before the terminal state such as `COMPLETED`, `CANCELED`, and `FAILED`).  The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	DeliverAt string `json:"deliver_at,omitempty"`
	// The duration of time it takes to prepare and deliver this fulfillment. The duration must be in RFC 3339 format (for example, \"P1W3D\").
	PrepTimeDuration string `json:"prep_time_duration,omitempty"`
	// The time period after `deliver_at` in which to deliver the order. Applications can set this field when the fulfillment `state` is `PROPOSED`, `RESERVED`, or `PREPARED` (any time before the terminal state such as `COMPLETED`, `CANCELED`, and `FAILED`).  The duration must be in RFC 3339 format (for example, \"P1W3D\").
	DeliveryWindowDuration string `json:"delivery_window_duration,omitempty"`
	// Provides additional instructions about the delivery fulfillment. It is displayed in the Square Point of Sale application and set by the API.
	Note string `json:"note,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicates when the seller completed the fulfillment. This field is automatically set when  fulfillment `state` changes to `COMPLETED`. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	CompletedAt string `json:"completed_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicates when the seller started processing the fulfillment. This field is automatically set when the fulfillment `state` changes to `RESERVED`. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	InProgressAt string `json:"in_progress_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the fulfillment was rejected. This field is automatically set when the fulfillment `state` changes to `FAILED`. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	RejectedAt string `json:"rejected_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the seller marked the fulfillment as ready for courier pickup. This field is automatically set when the fulfillment `state` changes to PREPARED. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	ReadyAt string `json:"ready_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the fulfillment was delivered to the recipient. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	DeliveredAt string `json:"delivered_at,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the fulfillment was canceled. This field is automatically set when the fulfillment `state` changes to `CANCELED`.  The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	CanceledAt string `json:"canceled_at,omitempty"`
	// The delivery cancellation reason. Max length: 100 characters.
	CancelReason string `json:"cancel_reason,omitempty"`
	// The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when an order can be picked up by the courier for delivery. The timestamp must be in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	CourierPickupAt string `json:"courier_pickup_at,omitempty"`
	// The time period after `courier_pickup_at` in which the courier should pick up the order. The duration must be in RFC 3339 format (for example, \"P1W3D\").
	CourierPickupWindowDuration string `json:"courier_pickup_window_duration,omitempty"`
	// Whether the delivery is preferred to be no contact.
	IsNoContactDelivery bool `json:"is_no_contact_delivery,omitempty"`
	// A note to provide additional instructions about how to deliver the order.
	DropoffNotes string `json:"dropoff_notes,omitempty"`
	// The name of the courier provider.
	CourierProviderName string `json:"courier_provider_name,omitempty"`
	// The support phone number of the courier.
	CourierSupportPhoneNumber string `json:"courier_support_phone_number,omitempty"`
	// The identifier for the delivery created by Square.
	SquareDeliveryId string `json:"square_delivery_id,omitempty"`
	// The identifier for the delivery created by the third-party courier service.
	ExternalDeliveryId string `json:"external_delivery_id,omitempty"`
	// The flag to indicate the delivery is managed by a third party (ie DoorDash), which means we may not receive all recipient information for PII purposes.
	ManagedDelivery bool `json:"managed_delivery,omitempty"`
}
