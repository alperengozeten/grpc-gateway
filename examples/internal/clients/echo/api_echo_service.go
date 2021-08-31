/*
 * examples/internal/proto/examplepb/echo_service.proto
 *
 * Echo Service    Echo Service API consists of a single service which returns  a message.
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package echo

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type EchoServiceApiService service

/* 
EchoServiceApiService Echo method receives a simple message and returns it.    The message posted as the id parameter will also be  returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id represents the message identifier.

@return ExamplepbSimpleMessage
*/
func (a *EchoServiceApiService) EchoServiceEcho(ctx context.Context, id string) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService Echo method receives a simple message and returns it.    The message posted as the id parameter will also be  returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id represents the message identifier.
 * @param num
 * @param optional nil or *EchoServiceEcho2Opts - Optional Parameters:
     * @param "LineNum" (optional.String) - 
     * @param "Lang" (optional.String) - 
     * @param "StatusProgress" (optional.String) - 
     * @param "StatusNote" (optional.String) - 
     * @param "En" (optional.String) - 
     * @param "NoProgress" (optional.String) - 
     * @param "NoNote" (optional.String) - 

@return ExamplepbSimpleMessage
*/

type EchoServiceEcho2Opts struct { 
	LineNum optional.String
	Lang optional.String
	StatusProgress optional.String
	StatusNote optional.String
	En optional.String
	NoProgress optional.String
	NoNote optional.String
}

func (a *EchoServiceApiService) EchoServiceEcho2(ctx context.Context, id string, num string, localVarOptionals *EchoServiceEcho2Opts) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo/{id}/{num}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"num"+"}", fmt.Sprintf("%v", num), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.LineNum.IsSet() {
		localVarQueryParams.Add("lineNum", parameterToString(localVarOptionals.LineNum.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lang.IsSet() {
		localVarQueryParams.Add("lang", parameterToString(localVarOptionals.Lang.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusProgress.IsSet() {
		localVarQueryParams.Add("status.progress", parameterToString(localVarOptionals.StatusProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusNote.IsSet() {
		localVarQueryParams.Add("status.note", parameterToString(localVarOptionals.StatusNote.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.En.IsSet() {
		localVarQueryParams.Add("en", parameterToString(localVarOptionals.En.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoProgress.IsSet() {
		localVarQueryParams.Add("no.progress", parameterToString(localVarOptionals.NoProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoNote.IsSet() {
		localVarQueryParams.Add("no.note", parameterToString(localVarOptionals.NoNote.Value(), ""))
	}
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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService Echo method receives a simple message and returns it.    The message posted as the id parameter will also be  returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id represents the message identifier.
 * @param num
 * @param lang
 * @param optional nil or *EchoServiceEcho3Opts - Optional Parameters:
     * @param "LineNum" (optional.String) - 
     * @param "StatusProgress" (optional.String) - 
     * @param "StatusNote" (optional.String) - 
     * @param "En" (optional.String) - 
     * @param "NoProgress" (optional.String) - 
     * @param "NoNote" (optional.String) - 

@return ExamplepbSimpleMessage
*/

type EchoServiceEcho3Opts struct { 
	LineNum optional.String
	StatusProgress optional.String
	StatusNote optional.String
	En optional.String
	NoProgress optional.String
	NoNote optional.String
}

func (a *EchoServiceApiService) EchoServiceEcho3(ctx context.Context, id string, num string, lang string, localVarOptionals *EchoServiceEcho3Opts) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo/{id}/{num}/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"num"+"}", fmt.Sprintf("%v", num), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", fmt.Sprintf("%v", lang), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.LineNum.IsSet() {
		localVarQueryParams.Add("lineNum", parameterToString(localVarOptionals.LineNum.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusProgress.IsSet() {
		localVarQueryParams.Add("status.progress", parameterToString(localVarOptionals.StatusProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusNote.IsSet() {
		localVarQueryParams.Add("status.note", parameterToString(localVarOptionals.StatusNote.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.En.IsSet() {
		localVarQueryParams.Add("en", parameterToString(localVarOptionals.En.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoProgress.IsSet() {
		localVarQueryParams.Add("no.progress", parameterToString(localVarOptionals.NoProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoNote.IsSet() {
		localVarQueryParams.Add("no.note", parameterToString(localVarOptionals.NoNote.Value(), ""))
	}
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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService Echo method receives a simple message and returns it.    The message posted as the id parameter will also be  returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Id represents the message identifier.
 * @param lineNum
 * @param statusNote
 * @param optional nil or *EchoServiceEcho4Opts - Optional Parameters:
     * @param "Num" (optional.String) - 
     * @param "Lang" (optional.String) - 
     * @param "StatusProgress" (optional.String) - 
     * @param "En" (optional.String) - 
     * @param "NoProgress" (optional.String) - 

@return ExamplepbSimpleMessage
*/

type EchoServiceEcho4Opts struct { 
	Num optional.String
	Lang optional.String
	StatusProgress optional.String
	En optional.String
	NoProgress optional.String
}

func (a *EchoServiceApiService) EchoServiceEcho4(ctx context.Context, id string, lineNum string, statusNote string, localVarOptionals *EchoServiceEcho4Opts) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo1/{id}/{lineNum}/{status.note}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lineNum"+"}", fmt.Sprintf("%v", lineNum), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"status.note"+"}", fmt.Sprintf("%v", statusNote), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Num.IsSet() {
		localVarQueryParams.Add("num", parameterToString(localVarOptionals.Num.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lang.IsSet() {
		localVarQueryParams.Add("lang", parameterToString(localVarOptionals.Lang.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusProgress.IsSet() {
		localVarQueryParams.Add("status.progress", parameterToString(localVarOptionals.StatusProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.En.IsSet() {
		localVarQueryParams.Add("en", parameterToString(localVarOptionals.En.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoProgress.IsSet() {
		localVarQueryParams.Add("no.progress", parameterToString(localVarOptionals.NoProgress.Value(), ""))
	}
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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService Echo method receives a simple message and returns it.    The message posted as the id parameter will also be  returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param noNote
 * @param optional nil or *EchoServiceEcho5Opts - Optional Parameters:
     * @param "Id" (optional.String) -  Id represents the message identifier.
     * @param "Num" (optional.String) - 
     * @param "LineNum" (optional.String) - 
     * @param "Lang" (optional.String) - 
     * @param "StatusProgress" (optional.String) - 
     * @param "En" (optional.String) - 
     * @param "NoProgress" (optional.String) - 

@return ExamplepbSimpleMessage
*/

type EchoServiceEcho5Opts struct { 
	Id optional.String
	Num optional.String
	LineNum optional.String
	Lang optional.String
	StatusProgress optional.String
	En optional.String
	NoProgress optional.String
}

func (a *EchoServiceApiService) EchoServiceEcho5(ctx context.Context, noNote string, localVarOptionals *EchoServiceEcho5Opts) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo2/{no.note}"
	localVarPath = strings.Replace(localVarPath, "{"+"no.note"+"}", fmt.Sprintf("%v", noNote), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Num.IsSet() {
		localVarQueryParams.Add("num", parameterToString(localVarOptionals.Num.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LineNum.IsSet() {
		localVarQueryParams.Add("lineNum", parameterToString(localVarOptionals.LineNum.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lang.IsSet() {
		localVarQueryParams.Add("lang", parameterToString(localVarOptionals.Lang.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusProgress.IsSet() {
		localVarQueryParams.Add("status.progress", parameterToString(localVarOptionals.StatusProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.En.IsSet() {
		localVarQueryParams.Add("en", parameterToString(localVarOptionals.En.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoProgress.IsSet() {
		localVarQueryParams.Add("no.progress", parameterToString(localVarOptionals.NoProgress.Value(), ""))
	}
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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService EchoBody method receives a simple message and returns it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body

@return ExamplepbSimpleMessage
*/
func (a *EchoServiceApiService) EchoServiceEchoBody(ctx context.Context, body ExamplepbSimpleMessage) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo_body"

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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService EchoDelete method receives a simple message and returns it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *EchoServiceEchoDeleteOpts - Optional Parameters:
     * @param "Id" (optional.String) -  Id represents the message identifier.
     * @param "Num" (optional.String) - 
     * @param "LineNum" (optional.String) - 
     * @param "Lang" (optional.String) - 
     * @param "StatusProgress" (optional.String) - 
     * @param "StatusNote" (optional.String) - 
     * @param "En" (optional.String) - 
     * @param "NoProgress" (optional.String) - 
     * @param "NoNote" (optional.String) - 

@return ExamplepbSimpleMessage
*/

type EchoServiceEchoDeleteOpts struct { 
	Id optional.String
	Num optional.String
	LineNum optional.String
	Lang optional.String
	StatusProgress optional.String
	StatusNote optional.String
	En optional.String
	NoProgress optional.String
	NoNote optional.String
}

func (a *EchoServiceApiService) EchoServiceEchoDelete(ctx context.Context, localVarOptionals *EchoServiceEchoDeleteOpts) (ExamplepbSimpleMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbSimpleMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo_delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Num.IsSet() {
		localVarQueryParams.Add("num", parameterToString(localVarOptionals.Num.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LineNum.IsSet() {
		localVarQueryParams.Add("lineNum", parameterToString(localVarOptionals.LineNum.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lang.IsSet() {
		localVarQueryParams.Add("lang", parameterToString(localVarOptionals.Lang.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusProgress.IsSet() {
		localVarQueryParams.Add("status.progress", parameterToString(localVarOptionals.StatusProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusNote.IsSet() {
		localVarQueryParams.Add("status.note", parameterToString(localVarOptionals.StatusNote.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.En.IsSet() {
		localVarQueryParams.Add("en", parameterToString(localVarOptionals.En.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoProgress.IsSet() {
		localVarQueryParams.Add("no.progress", parameterToString(localVarOptionals.NoProgress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NoNote.IsSet() {
		localVarQueryParams.Add("no.note", parameterToString(localVarOptionals.NoNote.Value(), ""))
	}
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
			var v ExamplepbSimpleMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
EchoServiceApiService EchoPatch method receives a NonStandardUpdateRequest and returns it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param optional nil or *EchoServiceEchoPatchOpts - Optional Parameters:
     * @param "UpdateMask" (optional.String) - 

@return ExamplepbDynamicMessageUpdate
*/

type EchoServiceEchoPatchOpts struct { 
	UpdateMask optional.String
}

func (a *EchoServiceApiService) EchoServiceEchoPatch(ctx context.Context, body ExamplepbDynamicMessage, localVarOptionals *EchoServiceEchoPatchOpts) (ExamplepbDynamicMessageUpdate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExamplepbDynamicMessageUpdate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/example/echo_patch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UpdateMask.IsSet() {
		localVarQueryParams.Add("updateMask", parameterToString(localVarOptionals.UpdateMask.Value(), ""))
	}
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
			var v ExamplepbDynamicMessageUpdate
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v RpcStatus
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
