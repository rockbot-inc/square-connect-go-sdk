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

// Defines input parameters in a request to the [ResumeSubscription]($e/Subscriptions/ResumeSubscription) endpoint.
type ResumeSubscriptionRequest struct {
	// The `YYYY-MM-DD`-formatted date when the subscription reactivated.
	ResumeEffectiveDate string `json:"resume_effective_date,omitempty"`
	ResumeChangeTiming *ChangeTiming `json:"resume_change_timing,omitempty"`
}
