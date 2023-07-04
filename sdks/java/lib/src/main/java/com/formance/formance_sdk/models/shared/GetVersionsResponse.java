/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * GetVersionsResponse - OK
 */
public class GetVersionsResponse {
    @JsonProperty("env")
    public String env;

    public GetVersionsResponse withEnv(String env) {
        this.env = env;
        return this;
    }
    
    @JsonProperty("region")
    public String region;

    public GetVersionsResponse withRegion(String region) {
        this.region = region;
        return this;
    }
    
    @JsonProperty("versions")
    public Version[] versions;

    public GetVersionsResponse withVersions(Version[] versions) {
        this.versions = versions;
        return this;
    }
    
    public GetVersionsResponse(@JsonProperty("env") String env, @JsonProperty("region") String region, @JsonProperty("versions") Version[] versions) {
        this.env = env;
        this.region = region;
        this.versions = versions;
  }
}