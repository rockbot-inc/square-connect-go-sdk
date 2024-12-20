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

// Represents a response from a bulk create request containing the created `TeamMember` objects or error messages.
type BulkCreateTeamMembersResponse struct {
	// The successfully created `TeamMember` objects. Each key is the `idempotency_key` that maps to the `CreateTeamMemberRequest`.
	TeamMembers map[string]CreateTeamMemberResponse `json:"team_members,omitempty"`
	// The errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
