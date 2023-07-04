/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class UpdateClientResponse {
    
    public String contentType;

    public UpdateClientResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    
    public Integer statusCode;

    public UpdateClientResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public UpdateClientResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    /**
     * Updated client
     */
    
    public org.openapis.openapi.models.shared.UpdateClientResponse updateClientResponse;

    public UpdateClientResponse withUpdateClientResponse(org.openapis.openapi.models.shared.UpdateClientResponse updateClientResponse) {
        this.updateClientResponse = updateClientResponse;
        return this;
    }
    
    public UpdateClientResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}