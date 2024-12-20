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

// Represents a create request for a `TeamMember` object.
type CreateTeamMemberRequest struct {
	// A unique string that identifies this `CreateTeamMember` request. Keys can be any valid string, but must be unique for every request. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).  The minimum length is 1 and the maximum length is 45.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	TeamMember *TeamMember `json:"team_member,omitempty"`
}
