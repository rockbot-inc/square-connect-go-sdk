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

type ListDevicesRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`
	SortOrder *SortOrder `json:"sort_order,omitempty"`
	// The number of results to return in a single page.
	Limit int32 `json:"limit,omitempty"`
	// If present, only returns devices at the target location.
	LocationId string `json:"location_id,omitempty"`
}
