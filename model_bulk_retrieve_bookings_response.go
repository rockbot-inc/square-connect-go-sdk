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

// Response payload for bulk retrieval of bookings.
type BulkRetrieveBookingsResponse struct {
	// Requested bookings returned as a map containing `booking_id` as the key and `RetrieveBookingResponse` as the value.
	Bookings map[string]RetrieveBookingResponse `json:"bookings,omitempty"`
	// Errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
