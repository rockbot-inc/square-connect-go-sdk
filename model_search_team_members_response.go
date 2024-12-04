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

// Represents a response from a search request containing a filtered list of `TeamMember` objects.
type SearchTeamMembersResponse struct {
	// The filtered list of `TeamMember` objects.
	TeamMembers []TeamMember `json:"team_members,omitempty"`
	// The opaque cursor for fetching the next page. For more information, see [pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
