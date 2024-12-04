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

// Request payload for bulk retrieval of bookings.
type BulkRetrieveBookingsRequest struct {
	// A non-empty list of [Booking](entity:Booking) IDs specifying bookings to retrieve.
	BookingIds []string `json:"booking_ids"`
}
