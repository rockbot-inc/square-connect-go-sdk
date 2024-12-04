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

// Defines the query parameters that can be included in a request to the [ListTransactions](api-endpoint:Transactions-ListTransactions) endpoint.  Deprecated - recommend using [SearchOrders](api-endpoint:Orders-SearchOrders)
type ListTransactionsRequest struct {
	// The beginning of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`
	// The end of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time.
	EndTime string `json:"end_time,omitempty"`
	SortOrder *SortOrder `json:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`
}
