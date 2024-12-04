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

// Defines the fields that are included in the response body of a request to the [SearchEvents]($e/Events/SearchEvents) endpoint.  Note: if there are errors processing the request, the events field will not be present.
type SearchEventsResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The list of [Event](entity:Event)s returned by the search.
	Events []Event `json:"events,omitempty"`
	// Contains the metadata of an event. For more information, see [Event](entity:Event).
	Metadata []EventMetadata `json:"metadata,omitempty"`
	// When a response is truncated, it includes a cursor that you can use in a subsequent request to fetch the next set of events. If empty, this is the final response.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
