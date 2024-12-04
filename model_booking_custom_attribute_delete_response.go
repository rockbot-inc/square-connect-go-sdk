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

// Represents a response for an individual upsert request in a [BulkDeleteBookingCustomAttributes]($e/BookingCustomAttributes/BulkDeleteBookingCustomAttributes) operation.
type BookingCustomAttributeDeleteResponse struct {
	// The ID of the [booking](entity:Booking) associated with the custom attribute.
	BookingId string `json:"booking_id,omitempty"`
	// Any errors that occurred while processing the individual request.
	Errors []ModelError `json:"errors,omitempty"`
}