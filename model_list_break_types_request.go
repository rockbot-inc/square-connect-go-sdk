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

// A request for a filtered set of `BreakType` objects.
type ListBreakTypesRequest struct {
	// Filter the returned `BreakType` results to only those that are associated with the specified location.
	LocationId string `json:"location_id,omitempty"`
	// The maximum number of `BreakType` results to return per page. The number can range between 1 and 200. The default is 200.
	Limit int32 `json:"limit,omitempty"`
	// A pointer to the next page of `BreakType` results to fetch.
	Cursor string `json:"cursor,omitempty"`
}
