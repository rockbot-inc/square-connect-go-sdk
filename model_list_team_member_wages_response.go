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

// The response to a request for a set of `TeamMemberWage` objects. The response contains a set of `TeamMemberWage` objects.
type ListTeamMemberWagesResponse struct {
	// A page of `TeamMemberWage` results.
	TeamMemberWages []TeamMemberWage `json:"team_member_wages,omitempty"`
	// The value supplied in the subsequent request to fetch the next page of `TeamMemberWage` results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}