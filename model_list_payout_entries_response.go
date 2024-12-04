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

// The response to retrieve payout records entries.
type ListPayoutEntriesResponse struct {
	// The requested list of payout entries, ordered with the given or default sort order.
	PayoutEntries []PayoutEntry `json:"payout_entries,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty, this is the final response. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
}