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

type ListLocationBookingProfilesRequest struct {
	// The maximum number of results to return in a paged response.
	Limit int32 `json:"limit,omitempty"`
	// The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results.
	Cursor string `json:"cursor,omitempty"`
}