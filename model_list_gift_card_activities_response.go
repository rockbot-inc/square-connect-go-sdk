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

// A response that contains a list of `GiftCardActivity` objects. If the request resulted in errors,  the response contains a set of `Error` objects.
type ListGiftCardActivitiesResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The requested gift card activities or an empty object if none are found.
	GiftCardActivities []GiftCardActivity `json:"gift_card_activities,omitempty"`
	// When a response is truncated, it includes a cursor that you can use in a subsequent request to retrieve the next set of activities. If a cursor is not present, this is the final response. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
}
