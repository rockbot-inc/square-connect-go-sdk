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

// Represents a response from a create request containing the created `TeamMember` object or error messages.
type CreateTeamMemberResponse struct {
	TeamMember *TeamMember `json:"team_member,omitempty"`
	// The errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
