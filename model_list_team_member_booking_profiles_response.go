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

type ListTeamMemberBookingProfilesResponse struct {
	// The list of team member booking profiles. The results are returned in the ascending order of the time when the team member booking profiles were last updated. Multiple booking profiles updated at the same time are further sorted in the ascending order of their IDs.
	TeamMemberBookingProfiles []TeamMemberBookingProfile `json:"team_member_booking_profiles,omitempty"`
	// The pagination cursor to be used in the subsequent request to get the next page of the results. Stop retrieving the next page of the results when the cursor is not set.
	Cursor string `json:"cursor,omitempty"`
	// Errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
