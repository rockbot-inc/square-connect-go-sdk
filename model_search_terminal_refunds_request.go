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

type SearchTerminalRefundsRequest struct {
	Query *TerminalRefundQuery `json:"query,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.
	Cursor string `json:"cursor,omitempty"`
	// Limits the number of results returned for a single request.
	Limit int32 `json:"limit,omitempty"`
}
