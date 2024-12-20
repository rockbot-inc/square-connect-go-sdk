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

// Defines the fields that are included in the response body of a request to the [TestWebhookSubscription]($e/WebhookSubscriptions/TestWebhookSubscription) endpoint.  Note: If there are errors processing the request, the [SubscriptionTestResult]($m/SubscriptionTestResult) field is not present.
type TestWebhookSubscriptionResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	SubscriptionTestResult *SubscriptionTestResult `json:"subscription_test_result,omitempty"`
}
