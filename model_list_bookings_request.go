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

type ListBookingsRequest struct {
	// The maximum number of results per page to return in a paged response.
	Limit int32 `json:"limit,omitempty"`
	// The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results.
	Cursor string `json:"cursor,omitempty"`
	// The [customer](entity:Customer) for whom to retrieve bookings. If this is not set, bookings for all customers are retrieved.
	CustomerId string `json:"customer_id,omitempty"`
	// The team member for whom to retrieve bookings. If this is not set, bookings of all members are retrieved.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// The location for which to retrieve bookings. If this is not set, all locations' bookings are retrieved.
	LocationId string `json:"location_id,omitempty"`
	// The RFC 3339 timestamp specifying the earliest of the start time. If this is not set, the current time is used.
	StartAtMin string `json:"start_at_min,omitempty"`
	// The RFC 3339 timestamp specifying the latest of the start time. If this is not set, the time of 31 days after `start_at_min` is used.
	StartAtMax string `json:"start_at_max,omitempty"`
}
