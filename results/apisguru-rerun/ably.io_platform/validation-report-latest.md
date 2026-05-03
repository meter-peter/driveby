# API Validation Report

Generated: 2026-05-03T21:05:11+03:00
Environment: production
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 1
- Failed Checks: 5
- Critical Issues: 4
- Warnings: 1
- Info: 0

### Categories
- Error Handling
- Schema
- Versioning
- Specification
- Documentation


### Failed Tags
- standards
- compatibility
- errors
- responses
- request
- compliance
- documentation
- quality
- versioning
- openapi
- specification
- usability
- schema
- validation
- lifecycle


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: parameter "filterLimit": parameter "limit" schema is invalid: invalid default: value must be an integer
Schema:
  {
    "default": "100",
    "type": "integer"
  }

Value:
  "100"

- **Tags:** openapi, specification, compliance

**Checks Performed:**
- OpenAPI version is 3.0.x or 3.1.0
- Required info fields (title, version) are present
- Paths are properly defined
- Components are valid
- References are resolvable
- No duplicate operationIds
- Valid HTTP methods used

**Details:**
```json
{
    "checks": {
      "Components are valid": true,
      "No duplicate operationIds": true,
      "Paths are properly defined": true,
      "Required info fields (title, version) are present": true,
      "Specification structure is valid": false,
      "Specification version is present": true,
      "Valid HTTP methods used": true
    },
    "messages": {
      "Specification structure is valid": "invalid components: parameter \"filterLimit\": parameter \"limit\" schema is invalid: invalid default: value must be an integer\nSchema:\n  {\n    \"default\": \"100\",\n    \"type\": \"integer\"\n  }\n\nValue:\n  \"100\"\n"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All parameters have descriptions: GET /channels: parameter limit, GET /channels/{channel_id}/presence: parameter clientId, GET /channels/{channel_id}/presence: parameter connectionId, GET /channels/{channel_id}/presence: parameter limit, GET /channels/{channel_id}/presence/history: parameter start, GET /channels/{channel_id}/presence/history: parameter limit, GET /channels/{channel_id}/presence/history: parameter end, GET /channels/{channel_id}/presence/history: parameter direction, GET /channels/{channel_id}/messages: parameter start, GET /channels/{channel_id}/messages: parameter limit, GET /channels/{channel_id}/messages: parameter end, GET /channels/{channel_id}/messages: parameter direction, GET /stats: parameter start, GET /stats: parameter limit, GET /stats: parameter end, GET /stats: parameter direction; All request/response bodies have examples: GET /channels: 2XX application/json response, GET /channels: 2XX application/x-msgpack response, GET /channels: 2XX text/html response, GET /channels: default application/json response, GET /channels: default application/x-msgpack response, GET /channels: default text/html response, GET /channels/{channel_id}/presence: 200 application/x-msgpack response, GET /channels/{channel_id}/presence: 200 text/html response, GET /channels/{channel_id}/presence: 200 application/json response, GET /channels/{channel_id}/presence: default application/json response, GET /channels/{channel_id}/presence: default application/x-msgpack response, GET /channels/{channel_id}/presence: default text/html response, POST /keys/{keyName}/requestToken: request body, POST /keys/{keyName}/requestToken: default text/html response, POST /keys/{keyName}/requestToken: default application/json response, POST /keys/{keyName}/requestToken: default application/x-msgpack response, POST /keys/{keyName}/requestToken: 2XX application/json response, POST /keys/{keyName}/requestToken: 2XX application/x-msgpack response, GET /channels/{channel_id}/presence/history: 2XX application/json response, GET /channels/{channel_id}/presence/history: 2XX application/x-msgpack response, GET /channels/{channel_id}/presence/history: 2XX text/html response, GET /channels/{channel_id}/presence/history: default application/json response, GET /channels/{channel_id}/presence/history: default application/x-msgpack response, GET /channels/{channel_id}/presence/history: default text/html response, DELETE /push/deviceRegistrations: default application/x-msgpack response, DELETE /push/deviceRegistrations: default text/html response, DELETE /push/deviceRegistrations: default application/json response, GET /push/deviceRegistrations: 2XX application/x-msgpack response, GET /push/deviceRegistrations: 2XX text/html response, GET /push/deviceRegistrations: 2XX application/json response, GET /push/deviceRegistrations: default application/json response, GET /push/deviceRegistrations: default application/x-msgpack response, GET /push/deviceRegistrations: default text/html response, POST /push/deviceRegistrations: request body, POST /push/deviceRegistrations: application/x-msgpack request body, POST /push/deviceRegistrations: application/json request body, POST /push/deviceRegistrations: 2XX text/html response, POST /push/deviceRegistrations: 2XX application/json response, POST /push/deviceRegistrations: 2XX application/x-msgpack response, POST /push/deviceRegistrations: default application/json response, POST /push/deviceRegistrations: default application/x-msgpack response, POST /push/deviceRegistrations: default text/html response, DELETE /push/deviceRegistrations/{device_id}: default application/json response, DELETE /push/deviceRegistrations/{device_id}: default application/x-msgpack response, DELETE /push/deviceRegistrations/{device_id}: default text/html response, GET /push/deviceRegistrations/{device_id}: 2XX application/json response, GET /push/deviceRegistrations/{device_id}: 2XX application/x-msgpack response, GET /push/deviceRegistrations/{device_id}: 2XX text/html response, GET /push/deviceRegistrations/{device_id}: default application/x-msgpack response, GET /push/deviceRegistrations/{device_id}: default text/html response, GET /push/deviceRegistrations/{device_id}: default application/json response, PATCH /push/deviceRegistrations/{device_id}: request body, PATCH /push/deviceRegistrations/{device_id}: application/x-www-form-urlencoded request body, PATCH /push/deviceRegistrations/{device_id}: application/json request body, PATCH /push/deviceRegistrations/{device_id}: application/x-msgpack request body, PATCH /push/deviceRegistrations/{device_id}: 2XX application/json response, PATCH /push/deviceRegistrations/{device_id}: 2XX application/x-msgpack response, PATCH /push/deviceRegistrations/{device_id}: 2XX text/html response, PATCH /push/deviceRegistrations/{device_id}: default application/json response, PATCH /push/deviceRegistrations/{device_id}: default application/x-msgpack response, PATCH /push/deviceRegistrations/{device_id}: default text/html response, PUT /push/deviceRegistrations/{device_id}: request body, PUT /push/deviceRegistrations/{device_id}: application/x-msgpack request body, PUT /push/deviceRegistrations/{device_id}: application/x-www-form-urlencoded request body, PUT /push/deviceRegistrations/{device_id}: application/json request body, PUT /push/deviceRegistrations/{device_id}: 2XX application/json response, PUT /push/deviceRegistrations/{device_id}: 2XX application/x-msgpack response, PUT /push/deviceRegistrations/{device_id}: 2XX text/html response, PUT /push/deviceRegistrations/{device_id}: default application/json response, PUT /push/deviceRegistrations/{device_id}: default application/x-msgpack response, PUT /push/deviceRegistrations/{device_id}: default text/html response, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: 2XX text/html response, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: 2XX application/json response, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: 2XX application/x-msgpack response, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: default application/json response, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: default application/x-msgpack response, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: default text/html response, GET /channels/{channel_id}: 200 application/json response, GET /channels/{channel_id}: default application/json response, GET /channels/{channel_id}: default application/x-msgpack response, GET /channels/{channel_id}: default text/html response, GET /channels/{channel_id}/messages: 2XX application/json response, GET /channels/{channel_id}/messages: 2XX application/x-msgpack response, GET /channels/{channel_id}/messages: 2XX text/html response, POST /channels/{channel_id}/messages: request body, POST /channels/{channel_id}/messages: application/json request body, POST /channels/{channel_id}/messages: application/x-msgpack request body, POST /channels/{channel_id}/messages: application/x-www-form-urlencoded request body, POST /channels/{channel_id}/messages: 2XX application/json response, POST /channels/{channel_id}/messages: 2XX application/x-msgpack response, POST /channels/{channel_id}/messages: 2XX text/html response, POST /channels/{channel_id}/messages: default application/json response, POST /channels/{channel_id}/messages: default application/x-msgpack response, POST /channels/{channel_id}/messages: default text/html response, DELETE /push/channelSubscriptions: default application/json response, DELETE /push/channelSubscriptions: default application/x-msgpack response, DELETE /push/channelSubscriptions: default text/html response, GET /push/channelSubscriptions: default application/x-msgpack response, GET /push/channelSubscriptions: default text/html response, GET /push/channelSubscriptions: default application/json response, GET /push/channelSubscriptions: 2XX application/json response, POST /push/channelSubscriptions: request body, POST /push/channelSubscriptions: default application/json response, POST /push/channelSubscriptions: default application/x-msgpack response, POST /push/channelSubscriptions: default text/html response, GET /push/channels: 2XX application/json response, GET /push/channels: 2XX application/x-msgpack response, GET /push/channels: 2XX text/html response, GET /push/channels: default application/json response, GET /push/channels: default application/x-msgpack response, GET /push/channels: default text/html response, POST /push/publish: request body, POST /push/publish: application/x-msgpack request body, POST /push/publish: application/x-www-form-urlencoded request body, POST /push/publish: application/json request body, POST /push/publish: default application/x-msgpack response, POST /push/publish: default text/html response, POST /push/publish: default application/json response, GET /stats: 2XX application/json response, GET /stats: default application/x-msgpack response, GET /stats: default text/html response, GET /stats: default application/json response, GET /time: 2XX application/json response, GET /time: 2XX application/x-msgpack response, GET /time: 2XX text/html response, GET /time: default application/json response, GET /time: default application/x-msgpack response, GET /time: default text/html response; All schemas have descriptions: Push, TokenDetails, ChannelDetails, DeviceDetails, PresenceMessage, SignedTokenRequest, TokenRequest, Notification
- **Tags:** documentation, quality, usability

**Checks Performed:**
- All operations have clear summaries
- All operations have detailed descriptions
- All operations have unique operationIds
- All parameters have descriptions
- All request/response bodies have examples
- All schemas have descriptions
- All enums have descriptions
- API has a general description
- Contact information is provided
- License information is provided

**Details:**
```json
{
    "checks": {
      "API has a general description": true,
      "All parameters have descriptions": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": false
    },
    "messages": {
      "License information is provided": "License information is missing"
    },
    "missing_docs": {
      "All parameters have descriptions": [
        "GET /channels: parameter limit",
        "GET /channels/{channel_id}/presence: parameter clientId",
        "GET /channels/{channel_id}/presence: parameter connectionId",
        "GET /channels/{channel_id}/presence: parameter limit",
        "GET /channels/{channel_id}/presence/history: parameter start",
        "GET /channels/{channel_id}/presence/history: parameter limit",
        "GET /channels/{channel_id}/presence/history: parameter end",
        "GET /channels/{channel_id}/presence/history: parameter direction",
        "GET /channels/{channel_id}/messages: parameter start",
        "GET /channels/{channel_id}/messages: parameter limit",
        "GET /channels/{channel_id}/messages: parameter end",
        "GET /channels/{channel_id}/messages: parameter direction",
        "GET /stats: parameter start",
        "GET /stats: parameter limit",
        "GET /stats: parameter end",
        "GET /stats: parameter direction"
      ],
      "All request/response bodies have examples": [
        "GET /channels: 2XX application/json response",
        "GET /channels: 2XX application/x-msgpack response",
        "GET /channels: 2XX text/html response",
        "GET /channels: default application/json response",
        "GET /channels: default application/x-msgpack response",
        "GET /channels: default text/html response",
        "GET /channels/{channel_id}/presence: 200 application/x-msgpack response",
        "GET /channels/{channel_id}/presence: 200 text/html response",
        "GET /channels/{channel_id}/presence: 200 application/json response",
        "GET /channels/{channel_id}/presence: default application/json response",
        "GET /channels/{channel_id}/presence: default application/x-msgpack response",
        "GET /channels/{channel_id}/presence: default text/html response",
        "POST /keys/{keyName}/requestToken: request body",
        "POST /keys/{keyName}/requestToken: default text/html response",
        "POST /keys/{keyName}/requestToken: default application/json response",
        "POST /keys/{keyName}/requestToken: default application/x-msgpack response",
        "POST /keys/{keyName}/requestToken: 2XX application/json response",
        "POST /keys/{keyName}/requestToken: 2XX application/x-msgpack response",
        "GET /channels/{channel_id}/presence/history: 2XX application/json response",
        "GET /channels/{channel_id}/presence/history: 2XX application/x-msgpack response",
        "GET /channels/{channel_id}/presence/history: 2XX text/html response",
        "GET /channels/{channel_id}/presence/history: default application/json response",
        "GET /channels/{channel_id}/presence/history: default application/x-msgpack response",
        "GET /channels/{channel_id}/presence/history: default text/html response",
        "DELETE /push/deviceRegistrations: default application/x-msgpack response",
        "DELETE /push/deviceRegistrations: default text/html response",
        "DELETE /push/deviceRegistrations: default application/json response",
        "GET /push/deviceRegistrations: 2XX application/x-msgpack response",
        "GET /push/deviceRegistrations: 2XX text/html response",
        "GET /push/deviceRegistrations: 2XX application/json response",
        "GET /push/deviceRegistrations: default application/json response",
        "GET /push/deviceRegistrations: default application/x-msgpack response",
        "GET /push/deviceRegistrations: default text/html response",
        "POST /push/deviceRegistrations: request body",
        "POST /push/deviceRegistrations: application/x-msgpack request body",
        "POST /push/deviceRegistrations: application/json request body",
        "POST /push/deviceRegistrations: 2XX text/html response",
        "POST /push/deviceRegistrations: 2XX application/json response",
        "POST /push/deviceRegistrations: 2XX application/x-msgpack response",
        "POST /push/deviceRegistrations: default application/json response",
        "POST /push/deviceRegistrations: default application/x-msgpack response",
        "POST /push/deviceRegistrations: default text/html response",
        "DELETE /push/deviceRegistrations/{device_id}: default application/json response",
        "DELETE /push/deviceRegistrations/{device_id}: default application/x-msgpack response",
        "DELETE /push/deviceRegistrations/{device_id}: default text/html response",
        "GET /push/deviceRegistrations/{device_id}: 2XX application/json response",
        "GET /push/deviceRegistrations/{device_id}: 2XX application/x-msgpack response",
        "GET /push/deviceRegistrations/{device_id}: 2XX text/html response",
        "GET /push/deviceRegistrations/{device_id}: default application/x-msgpack response",
        "GET /push/deviceRegistrations/{device_id}: default text/html response",
        "GET /push/deviceRegistrations/{device_id}: default application/json response",
        "PATCH /push/deviceRegistrations/{device_id}: request body",
        "PATCH /push/deviceRegistrations/{device_id}: application/x-www-form-urlencoded request body",
        "PATCH /push/deviceRegistrations/{device_id}: application/json request body",
        "PATCH /push/deviceRegistrations/{device_id}: application/x-msgpack request body",
        "PATCH /push/deviceRegistrations/{device_id}: 2XX application/json response",
        "PATCH /push/deviceRegistrations/{device_id}: 2XX application/x-msgpack response",
        "PATCH /push/deviceRegistrations/{device_id}: 2XX text/html response",
        "PATCH /push/deviceRegistrations/{device_id}: default application/json response",
        "PATCH /push/deviceRegistrations/{device_id}: default application/x-msgpack response",
        "PATCH /push/deviceRegistrations/{device_id}: default text/html response",
        "PUT /push/deviceRegistrations/{device_id}: request body",
        "PUT /push/deviceRegistrations/{device_id}: application/x-msgpack request body",
        "PUT /push/deviceRegistrations/{device_id}: application/x-www-form-urlencoded request body",
        "PUT /push/deviceRegistrations/{device_id}: application/json request body",
        "PUT /push/deviceRegistrations/{device_id}: 2XX application/json response",
        "PUT /push/deviceRegistrations/{device_id}: 2XX application/x-msgpack response",
        "PUT /push/deviceRegistrations/{device_id}: 2XX text/html response",
        "PUT /push/deviceRegistrations/{device_id}: default application/json response",
        "PUT /push/deviceRegistrations/{device_id}: default application/x-msgpack response",
        "PUT /push/deviceRegistrations/{device_id}: default text/html response",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: 2XX text/html response",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: 2XX application/json response",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: 2XX application/x-msgpack response",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: default application/json response",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: default application/x-msgpack response",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: default text/html response",
        "GET /channels/{channel_id}: 200 application/json response",
        "GET /channels/{channel_id}: default application/json response",
        "GET /channels/{channel_id}: default application/x-msgpack response",
        "GET /channels/{channel_id}: default text/html response",
        "GET /channels/{channel_id}/messages: 2XX application/json response",
        "GET /channels/{channel_id}/messages: 2XX application/x-msgpack response",
        "GET /channels/{channel_id}/messages: 2XX text/html response",
        "POST /channels/{channel_id}/messages: request body",
        "POST /channels/{channel_id}/messages: application/json request body",
        "POST /channels/{channel_id}/messages: application/x-msgpack request body",
        "POST /channels/{channel_id}/messages: application/x-www-form-urlencoded request body",
        "POST /channels/{channel_id}/messages: 2XX application/json response",
        "POST /channels/{channel_id}/messages: 2XX application/x-msgpack response",
        "POST /channels/{channel_id}/messages: 2XX text/html response",
        "POST /channels/{channel_id}/messages: default application/json response",
        "POST /channels/{channel_id}/messages: default application/x-msgpack response",
        "POST /channels/{channel_id}/messages: default text/html response",
        "DELETE /push/channelSubscriptions: default application/json response",
        "DELETE /push/channelSubscriptions: default application/x-msgpack response",
        "DELETE /push/channelSubscriptions: default text/html response",
        "GET /push/channelSubscriptions: default application/x-msgpack response",
        "GET /push/channelSubscriptions: default text/html response",
        "GET /push/channelSubscriptions: default application/json response",
        "GET /push/channelSubscriptions: 2XX application/json response",
        "POST /push/channelSubscriptions: request body",
        "POST /push/channelSubscriptions: default application/json response",
        "POST /push/channelSubscriptions: default application/x-msgpack response",
        "POST /push/channelSubscriptions: default text/html response",
        "GET /push/channels: 2XX application/json response",
        "GET /push/channels: 2XX application/x-msgpack response",
        "GET /push/channels: 2XX text/html response",
        "GET /push/channels: default application/json response",
        "GET /push/channels: default application/x-msgpack response",
        "GET /push/channels: default text/html response",
        "POST /push/publish: request body",
        "POST /push/publish: application/x-msgpack request body",
        "POST /push/publish: application/x-www-form-urlencoded request body",
        "POST /push/publish: application/json request body",
        "POST /push/publish: default application/x-msgpack response",
        "POST /push/publish: default text/html response",
        "POST /push/publish: default application/json response",
        "GET /stats: 2XX application/json response",
        "GET /stats: default application/x-msgpack response",
        "GET /stats: default text/html response",
        "GET /stats: default application/json response",
        "GET /time: 2XX application/json response",
        "GET /time: 2XX application/x-msgpack response",
        "GET /time: 2XX text/html response",
        "GET /time: default application/json response",
        "GET /time: default application/x-msgpack response",
        "GET /time: default text/html response"
      ],
      "All schemas have descriptions": [
        "Push",
        "TokenDetails",
        "ChannelDetails",
        "DeviceDetails",
        "PresenceMessage",
        "SignedTokenRequest",
        "TokenRequest",
        "Notification"
      ]
    }
  }
```

**Suggested Fix:**
Add missing documentation including descriptions, examples, and operation details

---

### Schema

#### P004: Request Schema Definitions (Failed) [critical]

Ensures all API requests have comprehensive schema definitions with proper data types, validation rules, and constraints

- **Status:** Failed
- **Message:** Request validation issues found: All string fields have length constraints: /push/deviceRegistrations/{device_id}: parameter X-Ably-Version, /push/deviceRegistrations/{device_id}: parameter format, DELETE /push/deviceRegistrations/{device_id}: parameter device_id, GET /push/deviceRegistrations/{device_id}: parameter device_id, PATCH /push/deviceRegistrations/{device_id}: parameter device_id, PATCH /push/deviceRegistrations/{device_id}.push.state: application/json schema, PATCH /push/deviceRegistrations/{device_id}.clientId: application/json schema, PATCH /push/deviceRegistrations/{device_id}.deviceSecret: application/json schema, PATCH /push/deviceRegistrations/{device_id}.formFactor: application/json schema, PATCH /push/deviceRegistrations/{device_id}.id: application/json schema, PATCH /push/deviceRegistrations/{device_id}.platform: application/json schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/json schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/json schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/json schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/json schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/json schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.push.state: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.clientId: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.deviceSecret: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.formFactor: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.id: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.platform: application/x-msgpack schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.push.state: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.clientId: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.deviceSecret: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.formFactor: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.id: application/x-www-form-urlencoded schema, PATCH /push/deviceRegistrations/{device_id}.platform: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}: parameter device_id, PUT /push/deviceRegistrations/{device_id}.platform: application/json schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/json schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/json schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/json schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/json schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/json schema, PUT /push/deviceRegistrations/{device_id}.push.state: application/json schema, PUT /push/deviceRegistrations/{device_id}.clientId: application/json schema, PUT /push/deviceRegistrations/{device_id}.deviceSecret: application/json schema, PUT /push/deviceRegistrations/{device_id}.formFactor: application/json schema, PUT /push/deviceRegistrations/{device_id}.id: application/json schema, PUT /push/deviceRegistrations/{device_id}.platform: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.state: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.clientId: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.deviceSecret: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.formFactor: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.id: application/x-msgpack schema, PUT /push/deviceRegistrations/{device_id}.push.state: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.clientId: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.deviceSecret: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.formFactor: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.id: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.platform: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-www-form-urlencoded schema, PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-www-form-urlencoded schema, /push/deviceRegistrations/{device_id}/resetUpdateToken: parameter X-Ably-Version, /push/deviceRegistrations/{device_id}/resetUpdateToken: parameter format, GET /push/deviceRegistrations/{device_id}/resetUpdateToken: parameter device_id, /stats: parameter X-Ably-Version, /stats: parameter format, GET /stats: parameter start, GET /stats: parameter end, GET /stats: parameter direction, GET /stats: parameter unit, /time: parameter X-Ably-Version, /time: parameter format, /channels/{channel_id}/presence: parameter X-Ably-Version, /channels/{channel_id}/presence: parameter format, GET /channels/{channel_id}/presence: parameter channel_id, GET /channels/{channel_id}/presence: parameter clientId, GET /channels/{channel_id}/presence: parameter connectionId, /keys/{keyName}/requestToken: parameter X-Ably-Version, /keys/{keyName}/requestToken: parameter format, POST /keys/{keyName}/requestToken: parameter keyName, /push/channelSubscriptions: parameter X-Ably-Version, /push/channelSubscriptions: parameter format, DELETE /push/channelSubscriptions: parameter channel, DELETE /push/channelSubscriptions: parameter deviceId, DELETE /push/channelSubscriptions: parameter clientId, GET /push/channelSubscriptions: parameter channel, GET /push/channelSubscriptions: parameter deviceId, GET /push/channelSubscriptions: parameter clientId, /push/channels: parameter X-Ably-Version, /push/channels: parameter format, /push/publish: parameter X-Ably-Version, /push/publish: parameter format, POST /push/publish.push.notification.sound: application/json schema, POST /push/publish.push.notification.title: application/json schema, POST /push/publish.push.notification.body: application/json schema, POST /push/publish.push.notification.collapseKey: application/json schema, POST /push/publish.push.notification.icon: application/json schema, POST /push/publish.push.web.notification.body: application/json schema, POST /push/publish.push.web.notification.collapseKey: application/json schema, POST /push/publish.push.web.notification.icon: application/json schema, POST /push/publish.push.web.notification.sound: application/json schema, POST /push/publish.push.web.notification.title: application/json schema, POST /push/publish.push.apns.notification.sound: application/json schema, POST /push/publish.push.apns.notification.title: application/json schema, POST /push/publish.push.apns.notification.body: application/json schema, POST /push/publish.push.apns.notification.collapseKey: application/json schema, POST /push/publish.push.apns.notification.icon: application/json schema, POST /push/publish.push.data: application/json schema, POST /push/publish.push.fcm.notification.collapseKey: application/json schema, POST /push/publish.push.fcm.notification.icon: application/json schema, POST /push/publish.push.fcm.notification.sound: application/json schema, POST /push/publish.push.fcm.notification.title: application/json schema, POST /push/publish.push.fcm.notification.body: application/json schema, POST /push/publish.recipient.deviceId: application/json schema, POST /push/publish.recipient.deviceToken: application/json schema, POST /push/publish.recipient.registrationToken: application/json schema, POST /push/publish.recipient.transportType: application/json schema, POST /push/publish.recipient.clientId: application/json schema, POST /push/publish.push.web.notification.icon: application/x-msgpack schema, POST /push/publish.push.web.notification.sound: application/x-msgpack schema, POST /push/publish.push.web.notification.title: application/x-msgpack schema, POST /push/publish.push.web.notification.body: application/x-msgpack schema, POST /push/publish.push.web.notification.collapseKey: application/x-msgpack schema, POST /push/publish.push.apns.notification.collapseKey: application/x-msgpack schema, POST /push/publish.push.apns.notification.icon: application/x-msgpack schema, POST /push/publish.push.apns.notification.sound: application/x-msgpack schema, POST /push/publish.push.apns.notification.title: application/x-msgpack schema, POST /push/publish.push.apns.notification.body: application/x-msgpack schema, POST /push/publish.push.data: application/x-msgpack schema, POST /push/publish.push.fcm.notification.body: application/x-msgpack schema, POST /push/publish.push.fcm.notification.collapseKey: application/x-msgpack schema, POST /push/publish.push.fcm.notification.icon: application/x-msgpack schema, POST /push/publish.push.fcm.notification.sound: application/x-msgpack schema, POST /push/publish.push.fcm.notification.title: application/x-msgpack schema, POST /push/publish.push.notification.body: application/x-msgpack schema, POST /push/publish.push.notification.collapseKey: application/x-msgpack schema, POST /push/publish.push.notification.icon: application/x-msgpack schema, POST /push/publish.push.notification.sound: application/x-msgpack schema, POST /push/publish.push.notification.title: application/x-msgpack schema, POST /push/publish.recipient.clientId: application/x-msgpack schema, POST /push/publish.recipient.deviceId: application/x-msgpack schema, POST /push/publish.recipient.deviceToken: application/x-msgpack schema, POST /push/publish.recipient.registrationToken: application/x-msgpack schema, POST /push/publish.recipient.transportType: application/x-msgpack schema, POST /push/publish.recipient.deviceId: application/x-www-form-urlencoded schema, POST /push/publish.recipient.deviceToken: application/x-www-form-urlencoded schema, POST /push/publish.recipient.registrationToken: application/x-www-form-urlencoded schema, POST /push/publish.recipient.transportType: application/x-www-form-urlencoded schema, POST /push/publish.recipient.clientId: application/x-www-form-urlencoded schema, POST /push/publish.push.apns.notification.sound: application/x-www-form-urlencoded schema, POST /push/publish.push.apns.notification.title: application/x-www-form-urlencoded schema, POST /push/publish.push.apns.notification.body: application/x-www-form-urlencoded schema, POST /push/publish.push.apns.notification.collapseKey: application/x-www-form-urlencoded schema, POST /push/publish.push.apns.notification.icon: application/x-www-form-urlencoded schema, POST /push/publish.push.data: application/x-www-form-urlencoded schema, POST /push/publish.push.fcm.notification.sound: application/x-www-form-urlencoded schema, POST /push/publish.push.fcm.notification.title: application/x-www-form-urlencoded schema, POST /push/publish.push.fcm.notification.body: application/x-www-form-urlencoded schema, POST /push/publish.push.fcm.notification.collapseKey: application/x-www-form-urlencoded schema, POST /push/publish.push.fcm.notification.icon: application/x-www-form-urlencoded schema, POST /push/publish.push.notification.icon: application/x-www-form-urlencoded schema, POST /push/publish.push.notification.sound: application/x-www-form-urlencoded schema, POST /push/publish.push.notification.title: application/x-www-form-urlencoded schema, POST /push/publish.push.notification.body: application/x-www-form-urlencoded schema, POST /push/publish.push.notification.collapseKey: application/x-www-form-urlencoded schema, POST /push/publish.push.web.notification.collapseKey: application/x-www-form-urlencoded schema, POST /push/publish.push.web.notification.icon: application/x-www-form-urlencoded schema, POST /push/publish.push.web.notification.sound: application/x-www-form-urlencoded schema, POST /push/publish.push.web.notification.title: application/x-www-form-urlencoded schema, POST /push/publish.push.web.notification.body: application/x-www-form-urlencoded schema, /channels/{channel_id}/presence/history: parameter X-Ably-Version, /channels/{channel_id}/presence/history: parameter format, GET /channels/{channel_id}/presence/history: parameter channel_id, GET /channels/{channel_id}/presence/history: parameter start, GET /channels/{channel_id}/presence/history: parameter end, GET /channels/{channel_id}/presence/history: parameter direction, /channels: parameter X-Ably-Version, /channels: parameter format, GET /channels: parameter prefix, GET /channels: parameter by, /channels/{channel_id}: parameter X-Ably-Version, /channels/{channel_id}: parameter format, GET /channels/{channel_id}: parameter channel_id, /channels/{channel_id}/messages: parameter X-Ably-Version, /channels/{channel_id}/messages: parameter format, POST /channels/{channel_id}/messages: parameter channel_id, POST /channels/{channel_id}/messages.encoding: application/json schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.body: application/json schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.collapseKey: application/json schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.icon: application/json schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.sound: application/json schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.title: application/json schema, POST /channels/{channel_id}/messages.extras.push.data: application/json schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.title: application/json schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.body: application/json schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.collapseKey: application/json schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.icon: application/json schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.sound: application/json schema, POST /channels/{channel_id}/messages.extras.push.notification.collapseKey: application/json schema, POST /channels/{channel_id}/messages.extras.push.notification.icon: application/json schema, POST /channels/{channel_id}/messages.extras.push.notification.sound: application/json schema, POST /channels/{channel_id}/messages.extras.push.notification.title: application/json schema, POST /channels/{channel_id}/messages.extras.push.notification.body: application/json schema, POST /channels/{channel_id}/messages.extras.push.web.notification.body: application/json schema, POST /channels/{channel_id}/messages.extras.push.web.notification.collapseKey: application/json schema, POST /channels/{channel_id}/messages.extras.push.web.notification.icon: application/json schema, POST /channels/{channel_id}/messages.extras.push.web.notification.sound: application/json schema, POST /channels/{channel_id}/messages.extras.push.web.notification.title: application/json schema, POST /channels/{channel_id}/messages.id: application/json schema, POST /channels/{channel_id}/messages.name: application/json schema, POST /channels/{channel_id}/messages.clientId: application/json schema, POST /channels/{channel_id}/messages.connectionId: application/json schema, POST /channels/{channel_id}/messages.data: application/json schema, POST /channels/{channel_id}/messages.id: application/x-msgpack schema, POST /channels/{channel_id}/messages.name: application/x-msgpack schema, POST /channels/{channel_id}/messages.clientId: application/x-msgpack schema, POST /channels/{channel_id}/messages.connectionId: application/x-msgpack schema, POST /channels/{channel_id}/messages.data: application/x-msgpack schema, POST /channels/{channel_id}/messages.encoding: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.web.notification.sound: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.web.notification.title: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.web.notification.body: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.web.notification.collapseKey: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.web.notification.icon: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.body: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.collapseKey: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.icon: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.sound: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.title: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.data: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.title: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.body: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.collapseKey: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.icon: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.sound: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.notification.sound: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.notification.title: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.notification.body: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.notification.collapseKey: application/x-msgpack schema, POST /channels/{channel_id}/messages.extras.push.notification.icon: application/x-msgpack schema, POST /channels/{channel_id}/messages.encoding: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.sound: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.title: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.body: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.collapseKey: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.apns.notification.icon: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.data: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.title: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.body: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.collapseKey: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.icon: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.fcm.notification.sound: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.notification.collapseKey: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.notification.icon: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.notification.sound: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.notification.title: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.notification.body: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.web.notification.body: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.web.notification.collapseKey: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.web.notification.icon: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.web.notification.sound: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.extras.push.web.notification.title: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.id: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.name: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.clientId: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.connectionId: application/x-www-form-urlencoded schema, POST /channels/{channel_id}/messages.data: application/x-www-form-urlencoded schema, GET /channels/{channel_id}/messages: parameter channel_id, GET /channels/{channel_id}/messages: parameter start, GET /channels/{channel_id}/messages: parameter end, GET /channels/{channel_id}/messages: parameter direction, /push/deviceRegistrations: parameter X-Ably-Version, /push/deviceRegistrations: parameter format, DELETE /push/deviceRegistrations: parameter deviceId, DELETE /push/deviceRegistrations: parameter clientId, GET /push/deviceRegistrations: parameter deviceId, GET /push/deviceRegistrations: parameter clientId, POST /push/deviceRegistrations.platform: application/json schema, POST /push/deviceRegistrations.push.recipient.clientId: application/json schema, POST /push/deviceRegistrations.push.recipient.deviceId: application/json schema, POST /push/deviceRegistrations.push.recipient.deviceToken: application/json schema, POST /push/deviceRegistrations.push.recipient.registrationToken: application/json schema, POST /push/deviceRegistrations.push.recipient.transportType: application/json schema, POST /push/deviceRegistrations.push.state: application/json schema, POST /push/deviceRegistrations.clientId: application/json schema, POST /push/deviceRegistrations.deviceSecret: application/json schema, POST /push/deviceRegistrations.formFactor: application/json schema, POST /push/deviceRegistrations.id: application/json schema, POST /push/deviceRegistrations.id: application/x-msgpack schema, POST /push/deviceRegistrations.platform: application/x-msgpack schema, POST /push/deviceRegistrations.push.recipient.deviceId: application/x-msgpack schema, POST /push/deviceRegistrations.push.recipient.deviceToken: application/x-msgpack schema, POST /push/deviceRegistrations.push.recipient.registrationToken: application/x-msgpack schema, POST /push/deviceRegistrations.push.recipient.transportType: application/x-msgpack schema, POST /push/deviceRegistrations.push.recipient.clientId: application/x-msgpack schema, POST /push/deviceRegistrations.push.state: application/x-msgpack schema, POST /push/deviceRegistrations.clientId: application/x-msgpack schema, POST /push/deviceRegistrations.deviceSecret: application/x-msgpack schema, POST /push/deviceRegistrations.formFactor: application/x-msgpack schema; All numeric fields have min/max values: GET /stats: parameter limit, GET /channels/{channel_id}/presence: parameter limit, GET /channels/{channel_id}/presence/history: parameter limit, GET /channels: parameter limit, POST /channels/{channel_id}/messages.timestamp: application/json schema, POST /channels/{channel_id}/messages.timestamp: application/x-msgpack schema, POST /channels/{channel_id}/messages.timestamp: application/x-www-form-urlencoded schema, GET /channels/{channel_id}/messages: parameter limit; All schemas specify data types: POST /keys/{keyName}/requestToken: application/json schema, POST /push/channelSubscriptions: application/json schema, POST /push/channelSubscriptions: application/x-msgpack schema, POST /push/channelSubscriptions: application/x-www-form-urlencoded schema
- **Tags:** schema, validation, request

**Checks Performed:**
- All path parameters have schemas
- All query parameters have schemas
- All header parameters have schemas
- All request bodies have content schemas
- All schemas specify data types
- All schemas have appropriate constraints
- All required fields are marked
- All enums have valid values
- All numeric fields have min/max values
- All string fields have length constraints

**Details:**
```json
{
    "checks": {
      "All enums have valid values": true,
      "All header parameters have schemas": true,
      "All numeric fields have min/max values": false,
      "All path parameters have schemas": true,
      "All query parameters have schemas": true,
      "All request bodies have content schemas": true,
      "All required fields are marked": true,
      "All schemas have appropriate constraints": true,
      "All schemas specify data types": false,
      "All string fields have length constraints": false
    },
    "messages": {},
    "missing_validation": {
      "All numeric fields have min/max values": [
        "GET /stats: parameter limit",
        "GET /channels/{channel_id}/presence: parameter limit",
        "GET /channels/{channel_id}/presence/history: parameter limit",
        "GET /channels: parameter limit",
        "POST /channels/{channel_id}/messages.timestamp: application/json schema",
        "POST /channels/{channel_id}/messages.timestamp: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.timestamp: application/x-www-form-urlencoded schema",
        "GET /channels/{channel_id}/messages: parameter limit"
      ],
      "All schemas specify data types": [
        "POST /keys/{keyName}/requestToken: application/json schema",
        "POST /push/channelSubscriptions: application/json schema",
        "POST /push/channelSubscriptions: application/x-msgpack schema",
        "POST /push/channelSubscriptions: application/x-www-form-urlencoded schema"
      ],
      "All string fields have length constraints": [
        "/push/deviceRegistrations/{device_id}: parameter X-Ably-Version",
        "/push/deviceRegistrations/{device_id}: parameter format",
        "DELETE /push/deviceRegistrations/{device_id}: parameter device_id",
        "GET /push/deviceRegistrations/{device_id}: parameter device_id",
        "PATCH /push/deviceRegistrations/{device_id}: parameter device_id",
        "PATCH /push/deviceRegistrations/{device_id}.push.state: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.clientId: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.deviceSecret: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.formFactor: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.id: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.platform: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/json schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.state: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.clientId: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.deviceSecret: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.formFactor: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.id: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.platform: application/x-msgpack schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.push.state: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.clientId: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.deviceSecret: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.formFactor: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.id: application/x-www-form-urlencoded schema",
        "PATCH /push/deviceRegistrations/{device_id}.platform: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}: parameter device_id",
        "PUT /push/deviceRegistrations/{device_id}.platform: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.push.state: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.clientId: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.deviceSecret: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.formFactor: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.id: application/json schema",
        "PUT /push/deviceRegistrations/{device_id}.platform: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.state: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.clientId: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.deviceSecret: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.formFactor: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.id: application/x-msgpack schema",
        "PUT /push/deviceRegistrations/{device_id}.push.state: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.clientId: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.deviceSecret: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.formFactor: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.id: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.platform: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceToken: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.registrationToken: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.transportType: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.clientId: application/x-www-form-urlencoded schema",
        "PUT /push/deviceRegistrations/{device_id}.push.recipient.deviceId: application/x-www-form-urlencoded schema",
        "/push/deviceRegistrations/{device_id}/resetUpdateToken: parameter X-Ably-Version",
        "/push/deviceRegistrations/{device_id}/resetUpdateToken: parameter format",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken: parameter device_id",
        "/stats: parameter X-Ably-Version",
        "/stats: parameter format",
        "GET /stats: parameter start",
        "GET /stats: parameter end",
        "GET /stats: parameter direction",
        "GET /stats: parameter unit",
        "/time: parameter X-Ably-Version",
        "/time: parameter format",
        "/channels/{channel_id}/presence: parameter X-Ably-Version",
        "/channels/{channel_id}/presence: parameter format",
        "GET /channels/{channel_id}/presence: parameter channel_id",
        "GET /channels/{channel_id}/presence: parameter clientId",
        "GET /channels/{channel_id}/presence: parameter connectionId",
        "/keys/{keyName}/requestToken: parameter X-Ably-Version",
        "/keys/{keyName}/requestToken: parameter format",
        "POST /keys/{keyName}/requestToken: parameter keyName",
        "/push/channelSubscriptions: parameter X-Ably-Version",
        "/push/channelSubscriptions: parameter format",
        "DELETE /push/channelSubscriptions: parameter channel",
        "DELETE /push/channelSubscriptions: parameter deviceId",
        "DELETE /push/channelSubscriptions: parameter clientId",
        "GET /push/channelSubscriptions: parameter channel",
        "GET /push/channelSubscriptions: parameter deviceId",
        "GET /push/channelSubscriptions: parameter clientId",
        "/push/channels: parameter X-Ably-Version",
        "/push/channels: parameter format",
        "/push/publish: parameter X-Ably-Version",
        "/push/publish: parameter format",
        "POST /push/publish.push.notification.sound: application/json schema",
        "POST /push/publish.push.notification.title: application/json schema",
        "POST /push/publish.push.notification.body: application/json schema",
        "POST /push/publish.push.notification.collapseKey: application/json schema",
        "POST /push/publish.push.notification.icon: application/json schema",
        "POST /push/publish.push.web.notification.body: application/json schema",
        "POST /push/publish.push.web.notification.collapseKey: application/json schema",
        "POST /push/publish.push.web.notification.icon: application/json schema",
        "POST /push/publish.push.web.notification.sound: application/json schema",
        "POST /push/publish.push.web.notification.title: application/json schema",
        "POST /push/publish.push.apns.notification.sound: application/json schema",
        "POST /push/publish.push.apns.notification.title: application/json schema",
        "POST /push/publish.push.apns.notification.body: application/json schema",
        "POST /push/publish.push.apns.notification.collapseKey: application/json schema",
        "POST /push/publish.push.apns.notification.icon: application/json schema",
        "POST /push/publish.push.data: application/json schema",
        "POST /push/publish.push.fcm.notification.collapseKey: application/json schema",
        "POST /push/publish.push.fcm.notification.icon: application/json schema",
        "POST /push/publish.push.fcm.notification.sound: application/json schema",
        "POST /push/publish.push.fcm.notification.title: application/json schema",
        "POST /push/publish.push.fcm.notification.body: application/json schema",
        "POST /push/publish.recipient.deviceId: application/json schema",
        "POST /push/publish.recipient.deviceToken: application/json schema",
        "POST /push/publish.recipient.registrationToken: application/json schema",
        "POST /push/publish.recipient.transportType: application/json schema",
        "POST /push/publish.recipient.clientId: application/json schema",
        "POST /push/publish.push.web.notification.icon: application/x-msgpack schema",
        "POST /push/publish.push.web.notification.sound: application/x-msgpack schema",
        "POST /push/publish.push.web.notification.title: application/x-msgpack schema",
        "POST /push/publish.push.web.notification.body: application/x-msgpack schema",
        "POST /push/publish.push.web.notification.collapseKey: application/x-msgpack schema",
        "POST /push/publish.push.apns.notification.collapseKey: application/x-msgpack schema",
        "POST /push/publish.push.apns.notification.icon: application/x-msgpack schema",
        "POST /push/publish.push.apns.notification.sound: application/x-msgpack schema",
        "POST /push/publish.push.apns.notification.title: application/x-msgpack schema",
        "POST /push/publish.push.apns.notification.body: application/x-msgpack schema",
        "POST /push/publish.push.data: application/x-msgpack schema",
        "POST /push/publish.push.fcm.notification.body: application/x-msgpack schema",
        "POST /push/publish.push.fcm.notification.collapseKey: application/x-msgpack schema",
        "POST /push/publish.push.fcm.notification.icon: application/x-msgpack schema",
        "POST /push/publish.push.fcm.notification.sound: application/x-msgpack schema",
        "POST /push/publish.push.fcm.notification.title: application/x-msgpack schema",
        "POST /push/publish.push.notification.body: application/x-msgpack schema",
        "POST /push/publish.push.notification.collapseKey: application/x-msgpack schema",
        "POST /push/publish.push.notification.icon: application/x-msgpack schema",
        "POST /push/publish.push.notification.sound: application/x-msgpack schema",
        "POST /push/publish.push.notification.title: application/x-msgpack schema",
        "POST /push/publish.recipient.clientId: application/x-msgpack schema",
        "POST /push/publish.recipient.deviceId: application/x-msgpack schema",
        "POST /push/publish.recipient.deviceToken: application/x-msgpack schema",
        "POST /push/publish.recipient.registrationToken: application/x-msgpack schema",
        "POST /push/publish.recipient.transportType: application/x-msgpack schema",
        "POST /push/publish.recipient.deviceId: application/x-www-form-urlencoded schema",
        "POST /push/publish.recipient.deviceToken: application/x-www-form-urlencoded schema",
        "POST /push/publish.recipient.registrationToken: application/x-www-form-urlencoded schema",
        "POST /push/publish.recipient.transportType: application/x-www-form-urlencoded schema",
        "POST /push/publish.recipient.clientId: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.apns.notification.sound: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.apns.notification.title: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.apns.notification.body: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.apns.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.apns.notification.icon: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.data: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.fcm.notification.sound: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.fcm.notification.title: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.fcm.notification.body: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.fcm.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.fcm.notification.icon: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.notification.icon: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.notification.sound: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.notification.title: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.notification.body: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.web.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.web.notification.icon: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.web.notification.sound: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.web.notification.title: application/x-www-form-urlencoded schema",
        "POST /push/publish.push.web.notification.body: application/x-www-form-urlencoded schema",
        "/channels/{channel_id}/presence/history: parameter X-Ably-Version",
        "/channels/{channel_id}/presence/history: parameter format",
        "GET /channels/{channel_id}/presence/history: parameter channel_id",
        "GET /channels/{channel_id}/presence/history: parameter start",
        "GET /channels/{channel_id}/presence/history: parameter end",
        "GET /channels/{channel_id}/presence/history: parameter direction",
        "/channels: parameter X-Ably-Version",
        "/channels: parameter format",
        "GET /channels: parameter prefix",
        "GET /channels: parameter by",
        "/channels/{channel_id}: parameter X-Ably-Version",
        "/channels/{channel_id}: parameter format",
        "GET /channels/{channel_id}: parameter channel_id",
        "/channels/{channel_id}/messages: parameter X-Ably-Version",
        "/channels/{channel_id}/messages: parameter format",
        "POST /channels/{channel_id}/messages: parameter channel_id",
        "POST /channels/{channel_id}/messages.encoding: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.body: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.collapseKey: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.icon: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.sound: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.title: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.data: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.title: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.body: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.collapseKey: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.icon: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.sound: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.collapseKey: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.icon: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.sound: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.title: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.body: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.body: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.collapseKey: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.icon: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.sound: application/json schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.title: application/json schema",
        "POST /channels/{channel_id}/messages.id: application/json schema",
        "POST /channels/{channel_id}/messages.name: application/json schema",
        "POST /channels/{channel_id}/messages.clientId: application/json schema",
        "POST /channels/{channel_id}/messages.connectionId: application/json schema",
        "POST /channels/{channel_id}/messages.data: application/json schema",
        "POST /channels/{channel_id}/messages.id: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.name: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.clientId: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.connectionId: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.data: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.encoding: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.sound: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.title: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.body: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.collapseKey: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.icon: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.body: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.collapseKey: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.icon: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.sound: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.title: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.data: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.title: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.body: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.collapseKey: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.icon: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.sound: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.sound: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.title: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.body: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.collapseKey: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.icon: application/x-msgpack schema",
        "POST /channels/{channel_id}/messages.encoding: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.sound: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.title: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.body: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.apns.notification.icon: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.data: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.title: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.body: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.icon: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.fcm.notification.sound: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.icon: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.sound: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.title: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.notification.body: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.body: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.collapseKey: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.icon: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.sound: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.extras.push.web.notification.title: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.id: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.name: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.clientId: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.connectionId: application/x-www-form-urlencoded schema",
        "POST /channels/{channel_id}/messages.data: application/x-www-form-urlencoded schema",
        "GET /channels/{channel_id}/messages: parameter channel_id",
        "GET /channels/{channel_id}/messages: parameter start",
        "GET /channels/{channel_id}/messages: parameter end",
        "GET /channels/{channel_id}/messages: parameter direction",
        "/push/deviceRegistrations: parameter X-Ably-Version",
        "/push/deviceRegistrations: parameter format",
        "DELETE /push/deviceRegistrations: parameter deviceId",
        "DELETE /push/deviceRegistrations: parameter clientId",
        "GET /push/deviceRegistrations: parameter deviceId",
        "GET /push/deviceRegistrations: parameter clientId",
        "POST /push/deviceRegistrations.platform: application/json schema",
        "POST /push/deviceRegistrations.push.recipient.clientId: application/json schema",
        "POST /push/deviceRegistrations.push.recipient.deviceId: application/json schema",
        "POST /push/deviceRegistrations.push.recipient.deviceToken: application/json schema",
        "POST /push/deviceRegistrations.push.recipient.registrationToken: application/json schema",
        "POST /push/deviceRegistrations.push.recipient.transportType: application/json schema",
        "POST /push/deviceRegistrations.push.state: application/json schema",
        "POST /push/deviceRegistrations.clientId: application/json schema",
        "POST /push/deviceRegistrations.deviceSecret: application/json schema",
        "POST /push/deviceRegistrations.formFactor: application/json schema",
        "POST /push/deviceRegistrations.id: application/json schema",
        "POST /push/deviceRegistrations.id: application/x-msgpack schema",
        "POST /push/deviceRegistrations.platform: application/x-msgpack schema",
        "POST /push/deviceRegistrations.push.recipient.deviceId: application/x-msgpack schema",
        "POST /push/deviceRegistrations.push.recipient.deviceToken: application/x-msgpack schema",
        "POST /push/deviceRegistrations.push.recipient.registrationToken: application/x-msgpack schema",
        "POST /push/deviceRegistrations.push.recipient.transportType: application/x-msgpack schema",
        "POST /push/deviceRegistrations.push.recipient.clientId: application/x-msgpack schema",
        "POST /push/deviceRegistrations.push.state: application/x-msgpack schema",
        "POST /push/deviceRegistrations.clientId: application/x-msgpack schema",
        "POST /push/deviceRegistrations.deviceSecret: application/x-msgpack schema",
        "POST /push/deviceRegistrations.formFactor: application/x-msgpack schema"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive schema validation including data types, constraints, and required fields

---

### Error Handling

#### P003: Error Handling Standards (Failed) [critical]

Validates comprehensive error response documentation and consistent error handling patterns

- **Status:** Failed
- **Message:** Error handling issues found: All operations document 4xx error responses: GET /channels/{channel_id}/messages, POST /channels/{channel_id}/messages, GET /channels/{channel_id}/presence, POST /keys/{keyName}/requestToken, GET /channels/{channel_id}/presence/history, DELETE /push/deviceRegistrations, GET /push/deviceRegistrations, POST /push/deviceRegistrations, GET /stats, GET /time, DELETE /push/channelSubscriptions, GET /push/channelSubscriptions, POST /push/channelSubscriptions, GET /push/channels, POST /push/publish, PATCH /push/deviceRegistrations/{device_id}, PUT /push/deviceRegistrations/{device_id}, DELETE /push/deviceRegistrations/{device_id}, GET /push/deviceRegistrations/{device_id}, GET /push/deviceRegistrations/{device_id}/resetUpdateToken, GET /channels, GET /channels/{channel_id}; All operations document 5xx error responses: GET /channels/{channel_id}/messages, POST /channels/{channel_id}/messages, GET /channels/{channel_id}/presence, POST /keys/{keyName}/requestToken, GET /channels/{channel_id}/presence/history, DELETE /push/deviceRegistrations, GET /push/deviceRegistrations, POST /push/deviceRegistrations, GET /stats, GET /time, DELETE /push/channelSubscriptions, GET /push/channelSubscriptions, POST /push/channelSubscriptions, GET /push/channels, POST /push/publish, PATCH /push/deviceRegistrations/{device_id}, PUT /push/deviceRegistrations/{device_id}, DELETE /push/deviceRegistrations/{device_id}, GET /push/deviceRegistrations/{device_id}, GET /push/deviceRegistrations/{device_id}/resetUpdateToken, GET /channels, GET /channels/{channel_id}
- **Tags:** errors, responses, standards

**Checks Performed:**
- All operations document 4xx error responses
- All operations document 5xx error responses
- Error responses include error codes
- Error responses include error messages
- Error responses include error details schema
- Common error responses are defined in components
- Error responses follow consistent format

**Details:**
```json
{
    "checks": {
      "All operations document 4xx error responses": false,
      "All operations document 5xx error responses": false,
      "Common error responses are defined in components": false,
      "Error responses follow consistent format": true
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {
      "All operations document 4xx error responses": [
        "GET /channels/{channel_id}/messages",
        "POST /channels/{channel_id}/messages",
        "GET /channels/{channel_id}/presence",
        "POST /keys/{keyName}/requestToken",
        "GET /channels/{channel_id}/presence/history",
        "DELETE /push/deviceRegistrations",
        "GET /push/deviceRegistrations",
        "POST /push/deviceRegistrations",
        "GET /stats",
        "GET /time",
        "DELETE /push/channelSubscriptions",
        "GET /push/channelSubscriptions",
        "POST /push/channelSubscriptions",
        "GET /push/channels",
        "POST /push/publish",
        "PATCH /push/deviceRegistrations/{device_id}",
        "PUT /push/deviceRegistrations/{device_id}",
        "DELETE /push/deviceRegistrations/{device_id}",
        "GET /push/deviceRegistrations/{device_id}",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken",
        "GET /channels",
        "GET /channels/{channel_id}"
      ],
      "All operations document 5xx error responses": [
        "GET /channels/{channel_id}/messages",
        "POST /channels/{channel_id}/messages",
        "GET /channels/{channel_id}/presence",
        "POST /keys/{keyName}/requestToken",
        "GET /channels/{channel_id}/presence/history",
        "DELETE /push/deviceRegistrations",
        "GET /push/deviceRegistrations",
        "POST /push/deviceRegistrations",
        "GET /stats",
        "GET /time",
        "DELETE /push/channelSubscriptions",
        "GET /push/channelSubscriptions",
        "POST /push/channelSubscriptions",
        "GET /push/channels",
        "POST /push/publish",
        "PATCH /push/deviceRegistrations/{device_id}",
        "PUT /push/deviceRegistrations/{device_id}",
        "DELETE /push/deviceRegistrations/{device_id}",
        "GET /push/deviceRegistrations/{device_id}",
        "GET /push/deviceRegistrations/{device_id}/resetUpdateToken",
        "GET /channels",
        "GET /channels/{channel_id}"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive error response documentation including codes, messages, and consistent error schemas

---

### Security

#### P005: Security Standards (Passed) [critical]

Validates comprehensive security requirements and authentication mechanisms

- **Status:** Passed
- **Message:** All security requirements are properly defined and consistent
- **Tags:** security, authentication, authorization

**Checks Performed:**
- Security schemes are defined
- Global security requirements are set
- Operation-level security is defined
- OAuth2 scopes are documented
- API keys are properly described
- Authentication headers are specified
- Security requirements are consistent

**Details:**
```json
{
    "checks": {
      "API keys are properly described": true,
      "Authentication headers are specified": true,
      "Global security requirements are set": true,
      "OAuth2 scopes are documented": true,
      "Operation-level security is defined": true,
      "Security requirements are consistent": true,
      "Security schemes are defined": true
    },
    "messages": {}
  }
```

---

### Versioning

#### P008: API Versioning Strategy (Failed) [warning]

Validates proper API versioning implementation and documentation

- **Status:** Failed
- **Message:** Versioning validation failed: Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Versioning strategy is documented: Info description does not mention versioning strategy; Breaking changes are documented: Info description does not reference breaking changes or a changelog
- **Tags:** versioning, compatibility, lifecycle

**Checks Performed:**
- API version is specified
- Version follows semantic versioning
- Versioning strategy is documented
- Deprecation notices are present
- Breaking changes are documented
- Version compatibility is specified
- Migration guides are referenced

**Details:**
```json
{
    "checks": {
      "API version is specified": true,
      "Breaking changes are documented": false,
      "Deprecation notices are present": true,
      "Migration guides are referenced": false,
      "Version compatibility is specified": false,
      "Version follows semantic versioning": true,
      "Versioning strategy is documented": false
    },
    "messages": {
      "Breaking changes are documented": "Info description does not reference breaking changes or a changelog",
      "Migration guides are referenced": "Info description does not reference migration or upgrade guides",
      "Version compatibility is specified": "Info description does not mention version compatibility",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

