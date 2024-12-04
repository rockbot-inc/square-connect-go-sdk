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

// Defines the fields that are included in the response body of a request to the [ListEventTypes]($e/Events/ListEventTypes) endpoint.  Note: if there are errors processing the request, the event types field will not be present.
type ListEventTypesResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The list of event types.
	EventTypes []string `json:"event_types,omitempty"`
	// Contains the metadata of an event type. For more information, see [EventTypeMetadata](entity:EventTypeMetadata).
	Metadata []EventTypeMetadata `json:"metadata,omitempty"`
}