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

// Represents a bulk create request for `TeamMember` objects.
type BulkCreateTeamMembersRequest struct {
	// The data used to create the `TeamMember` objects. Each key is the `idempotency_key` that maps to the `CreateTeamMemberRequest`. The maximum number of create objects is 25.
	TeamMembers map[string]CreateTeamMemberRequest `json:"team_members"`
}
