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

// Represents an output from a call to [SearchVendors]($e/Vendors/SearchVendors).
type SearchVendorsResponse struct {
	// Errors encountered when the request fails.
	Errors []ModelError `json:"errors,omitempty"`
	// The [Vendor](entity:Vendor) objects matching the specified search filter.
	Vendors []Vendor `json:"vendors,omitempty"`
	// The pagination cursor to be used in a subsequent request. If unset, this is the final response.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`
}
