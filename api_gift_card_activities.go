
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

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type GiftCardActivitiesApiService service
/*
GiftCardActivitiesApiService CreateGiftCardActivity
Creates a gift card activity to manage the balance or state of a [gift card]($m/GiftCard). For example, create an &#x60;ACTIVATE&#x60; activity to activate a gift card with an initial balance before first use.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body An object containing the fields to POST for the request.

See the corresponding object definition for field details.
@return CreateGiftCardActivityResponse
*/
func (a *GiftCardActivitiesApiService) CreateGiftCardActivity(ctx context.Context, body CreateGiftCardActivityRequest) (CreateGiftCardActivityResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CreateGiftCardActivityResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/gift-cards/activities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CreateGiftCardActivityResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
GiftCardActivitiesApiService ListGiftCardActivities
Lists gift card activities. By default, you get gift card activities for all gift cards in the seller&#x27;s account. You can optionally specify query parameters to filter the list. For example, you can get a list of gift card activities for a gift card, for all gift cards in a specific region, or for activities within a time window.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GiftCardActivitiesApiListGiftCardActivitiesOpts - Optional Parameters:
     * @param "GiftCardId" (optional.String) -  If a gift card ID is provided, the endpoint returns activities related  to the specified gift card. Otherwise, the endpoint returns all gift card activities for  the seller.
     * @param "Type_" (optional.String) -  If a [type](entity:GiftCardActivityType) is provided, the endpoint returns gift card activities of the specified type.  Otherwise, the endpoint returns all types of gift card activities.
     * @param "LocationId" (optional.String) -  If a location ID is provided, the endpoint returns gift card activities for the specified location.  Otherwise, the endpoint returns gift card activities for all locations.
     * @param "BeginTime" (optional.String) -  The timestamp for the beginning of the reporting period, in RFC 3339 format. This start time is inclusive. The default value is the current time minus one year.
     * @param "EndTime" (optional.String) -  The timestamp for the end of the reporting period, in RFC 3339 format. This end time is inclusive. The default value is the current time.
     * @param "Limit" (optional.Int32) -  If a limit is provided, the endpoint returns the specified number  of results (or fewer) per page. The maximum value is 100. The default value is 50. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
     * @param "Cursor" (optional.String) -  A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If a cursor is not provided, the endpoint returns the first page of the results. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
     * @param "SortOrder" (optional.String) -  The order in which the endpoint returns the activities, based on &#x60;created_at&#x60;. - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default).
@return ListGiftCardActivitiesResponse
*/

type GiftCardActivitiesApiListGiftCardActivitiesOpts struct {
    GiftCardId optional.String
    Type_ optional.String
    LocationId optional.String
    BeginTime optional.String
    EndTime optional.String
    Limit optional.Int32
    Cursor optional.String
    SortOrder optional.String
}

func (a *GiftCardActivitiesApiService) ListGiftCardActivities(ctx context.Context, localVarOptionals *GiftCardActivitiesApiListGiftCardActivitiesOpts) (ListGiftCardActivitiesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ListGiftCardActivitiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/gift-cards/activities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.GiftCardId.IsSet() {
		localVarQueryParams.Add("gift_card_id", parameterToString(localVarOptionals.GiftCardId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LocationId.IsSet() {
		localVarQueryParams.Add("location_id", parameterToString(localVarOptionals.LocationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeginTime.IsSet() {
		localVarQueryParams.Add("begin_time", parameterToString(localVarOptionals.BeginTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndTime.IsSet() {
		localVarQueryParams.Add("end_time", parameterToString(localVarOptionals.EndTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortOrder.IsSet() {
		localVarQueryParams.Add("sort_order", parameterToString(localVarOptionals.SortOrder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ListGiftCardActivitiesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
