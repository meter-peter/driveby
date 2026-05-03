# API Validation Report

Generated: 2026-05-03T21:05:21+03:00
Environment: production
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 0
- Failed Checks: 6
- Critical Issues: 5
- Warnings: 1
- Info: 0

### Categories
- Schema
- Security
- Versioning
- Specification
- Documentation
- Error Handling


### Failed Tags
- schema
- validation
- lifecycle
- openapi
- compliance
- documentation
- request
- authentication
- authorization
- usability
- specification
- quality
- errors
- responses
- standards
- security
- versioning
- compatibility


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "Error_400": invalid example: Error at "/errors/0/source": there must be at most 1 properties
Schema:
  {
    "description": "an object containing references to the source of the error",
    "maxProperties": 1,
    "properties": {
      "example": {
        "description": "a string indicating an example of the right value",
        "type": "string"
      },
      "parameter": {
        "description": "a string indicating which URI query parameter caused the issue",
        "type": "string"
      },
      "pointer": {
        "description": "a JSON Pointer [RFC6901] to the associated entity in the request document",
        "type": "string"
      }
    },
    "title": "Issue_Source",
    "type": "object"
  }

Value:
  {
    "example": "CDG",
    "parameter": "airport"
  }

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
      "Specification structure is valid": "invalid components: schema \"Error_400\": invalid example: Error at \"/errors/0/source\": there must be at most 1 properties\nSchema:\n  {\n    \"description\": \"an object containing references to the source of the error\",\n    \"maxProperties\": 1,\n    \"properties\": {\n      \"example\": {\n        \"description\": \"a string indicating an example of the right value\",\n        \"type\": \"string\"\n      },\n      \"parameter\": {\n        \"description\": \"a string indicating which URI query parameter caused the issue\",\n        \"type\": \"string\"\n      },\n      \"pointer\": {\n        \"description\": \"a JSON Pointer [RFC6901] to the associated entity in the request document\",\n        \"type\": \"string\"\n      }\n    },\n    \"title\": \"Issue_Source\",\n    \"type\": \"object\"\n  }\n\nValue:\n  {\n    \"example\": \"CDG\",\n    \"parameter\": \"airport\"\n  }\n"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All operations have detailed descriptions: GET /shopping/flight-offers, POST /shopping/flight-offers; All request/response bodies have examples: GET /shopping/flight-offers: 400 application/vnd.amadeus+json response, GET /shopping/flight-offers: default application/vnd.amadeus+json response, GET /shopping/flight-offers: 200 application/vnd.amadeus+json response, POST /shopping/flight-offers: application/vnd.amadeus+json request body, POST /shopping/flight-offers: default application/vnd.amadeus+json response, POST /shopping/flight-offers: 200 application/vnd.amadeus+json response, POST /shopping/flight-offers: 400 application/vnd.amadeus+json response; All schemas have descriptions: FlightOffer, OriginDestination, DateTimeRange, Error_500, CarrierEntry, Collection_Meta, Segment, Co2Emission, Collection_Meta_Link, Error_400, Price, AircraftEntry, GetFlightOffersQuery, LocationValue, SearchCriteria, Dictionaries, CurrencyEntry, LocationEntry, Traveler, Issue; All enums have descriptions: ServiceName: enum value PRIORITY_BOARDING, ServiceName: enum value AIRPORT_CHECKIN, Coverage: enum value MOST_SEGMENTS, Coverage: enum value AT_LEAST_ONE_SEGMENT, FlightOfferSource: enum value GDS, SliceDiceIndicator: enum value LOCAL_AVAILABILITY, SliceDiceIndicator: enum value SUB_OD_AVAILABILITY_1, SliceDiceIndicator: enum value SUB_OD_AVAILABILITY_2, FeeType: enum value TICKETING, FeeType: enum value FORM_OF_PAYMENT, FeeType: enum value SUPPLIER, TravelClass: enum value ECONOMY, TravelClass: enum value PREMIUM_ECONOMY, TravelClass: enum value BUSINESS, TravelClass: enum value FIRST, AdditionalServiceType: enum value CHECKED_BAGS, AdditionalServiceType: enum value MEALS, AdditionalServiceType: enum value SEATS, AdditionalServiceType: enum value OTHER_SERVICES, TravelerType: enum value ADULT, TravelerType: enum value YOUNG, TravelerType: enum value STUDENT, TravelerPricingFareOption: enum value STANDARD, TravelerPricingFareOption: enum value INCLUSIVE_TOUR, TravelerPricingFareOption: enum value SPANISH_MELILLA_RESIDENT, TravelerPricingFareOption: enum value SPANISH_CEUTA_RESIDENT, TravelerPricingFareOption: enum value SPANISH_CANARY_RESIDENT, TravelerPricingFareOption: enum value SPANISH_BALEARIC_RESIDENT, TravelerPricingFareOption: enum value AIR_FRANCE_METROPOLITAN_DISCOUNT_PASS, TravelerPricingFareOption: enum value AIR_FRANCE_DOM_DISCOUNT_PASS, TravelerPricingFareOption: enum value AIR_FRANCE_COMBINED_DISCOUNT_PASS, TravelerPricingFareOption: enum value AIR_FRANCE_FAMILY, TravelerPricingFareOption: enum value ADULT_WITH_COMPANION, TravelerPricingFareOption: enum value COMPANION
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
      "All enums have descriptions": false,
      "All operations have detailed descriptions": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": false,
      "License information is provided": false
    },
    "messages": {
      "Contact information is provided": "Contact information is missing",
      "License information is provided": "License information is missing"
    },
    "missing_docs": {
      "All enums have descriptions": [
        "ServiceName: enum value PRIORITY_BOARDING",
        "ServiceName: enum value AIRPORT_CHECKIN",
        "Coverage: enum value MOST_SEGMENTS",
        "Coverage: enum value AT_LEAST_ONE_SEGMENT",
        "FlightOfferSource: enum value GDS",
        "SliceDiceIndicator: enum value LOCAL_AVAILABILITY",
        "SliceDiceIndicator: enum value SUB_OD_AVAILABILITY_1",
        "SliceDiceIndicator: enum value SUB_OD_AVAILABILITY_2",
        "FeeType: enum value TICKETING",
        "FeeType: enum value FORM_OF_PAYMENT",
        "FeeType: enum value SUPPLIER",
        "TravelClass: enum value ECONOMY",
        "TravelClass: enum value PREMIUM_ECONOMY",
        "TravelClass: enum value BUSINESS",
        "TravelClass: enum value FIRST",
        "AdditionalServiceType: enum value CHECKED_BAGS",
        "AdditionalServiceType: enum value MEALS",
        "AdditionalServiceType: enum value SEATS",
        "AdditionalServiceType: enum value OTHER_SERVICES",
        "TravelerType: enum value ADULT",
        "TravelerType: enum value YOUNG",
        "TravelerType: enum value STUDENT",
        "TravelerPricingFareOption: enum value STANDARD",
        "TravelerPricingFareOption: enum value INCLUSIVE_TOUR",
        "TravelerPricingFareOption: enum value SPANISH_MELILLA_RESIDENT",
        "TravelerPricingFareOption: enum value SPANISH_CEUTA_RESIDENT",
        "TravelerPricingFareOption: enum value SPANISH_CANARY_RESIDENT",
        "TravelerPricingFareOption: enum value SPANISH_BALEARIC_RESIDENT",
        "TravelerPricingFareOption: enum value AIR_FRANCE_METROPOLITAN_DISCOUNT_PASS",
        "TravelerPricingFareOption: enum value AIR_FRANCE_DOM_DISCOUNT_PASS",
        "TravelerPricingFareOption: enum value AIR_FRANCE_COMBINED_DISCOUNT_PASS",
        "TravelerPricingFareOption: enum value AIR_FRANCE_FAMILY",
        "TravelerPricingFareOption: enum value ADULT_WITH_COMPANION",
        "TravelerPricingFareOption: enum value COMPANION"
      ],
      "All operations have detailed descriptions": [
        "GET /shopping/flight-offers",
        "POST /shopping/flight-offers"
      ],
      "All request/response bodies have examples": [
        "GET /shopping/flight-offers: 400 application/vnd.amadeus+json response",
        "GET /shopping/flight-offers: default application/vnd.amadeus+json response",
        "GET /shopping/flight-offers: 200 application/vnd.amadeus+json response",
        "POST /shopping/flight-offers: application/vnd.amadeus+json request body",
        "POST /shopping/flight-offers: default application/vnd.amadeus+json response",
        "POST /shopping/flight-offers: 200 application/vnd.amadeus+json response",
        "POST /shopping/flight-offers: 400 application/vnd.amadeus+json response"
      ],
      "All schemas have descriptions": [
        "FlightOffer",
        "OriginDestination",
        "DateTimeRange",
        "Error_500",
        "CarrierEntry",
        "Collection_Meta",
        "Segment",
        "Co2Emission",
        "Collection_Meta_Link",
        "Error_400",
        "Price",
        "AircraftEntry",
        "GetFlightOffersQuery",
        "LocationValue",
        "SearchCriteria",
        "Dictionaries",
        "CurrencyEntry",
        "LocationEntry",
        "Traveler",
        "Issue"
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
- **Message:** Request validation issues found: All string fields have length constraints: GET /shopping/flight-offers: parameter departureDate, GET /shopping/flight-offers: parameter returnDate, GET /shopping/flight-offers: parameter travelClass, POST /shopping/flight-offers: parameter X-HTTP-Method-Override, POST /shopping/flight-offers.searchCriteria.flightFilters.cabinRestrictions[].cabin: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.flightFilters.cabinRestrictions[].coverage: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.flightFilters.cabinRestrictions[].originDestinationIds[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.flightFilters.carrierRestrictions.excludedCarrierCodes[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.flightFilters.carrierRestrictions.includedCarrierCodes[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.sources[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.travelers[].associatedAdultId: application/vnd.amadeus+json schema, POST /shopping/flight-offers.travelers[].id: application/vnd.amadeus+json schema, POST /shopping/flight-offers.travelers[].travelerType: application/vnd.amadeus+json schema, POST /shopping/flight-offers.currencyCode: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].includedConnectionPoints[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].id: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].alternativeDestinationsCodes[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].arrivalDateTimeRange.time: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].arrivalDateTimeRange.date: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].destinationLocationCode: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].excludedConnectionPoints[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].originLocationCode: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].alternativeOriginsCodes[]: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].departureDateTimeRange.time: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].departureDateTimeRange.date: application/vnd.amadeus+json schema; All schemas specify data types: POST /shopping/flight-offers: application/vnd.amadeus+json schema; All numeric fields have min/max values: POST /shopping/flight-offers.searchCriteria.maxPrice: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.flightFilters.connectionRestriction.maxNumberOfConnections: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.flightFilters.maxFlightTime: application/vnd.amadeus+json schema, POST /shopping/flight-offers.searchCriteria.maxFlightOffers: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].originRadius: application/vnd.amadeus+json schema, POST /shopping/flight-offers.originDestinations[].destinationRadius: application/vnd.amadeus+json schema
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
        "POST /shopping/flight-offers.searchCriteria.maxPrice: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.connectionRestriction.maxNumberOfConnections: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.maxFlightTime: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.maxFlightOffers: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].originRadius: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].destinationRadius: application/vnd.amadeus+json schema"
      ],
      "All schemas specify data types": [
        "POST /shopping/flight-offers: application/vnd.amadeus+json schema"
      ],
      "All string fields have length constraints": [
        "GET /shopping/flight-offers: parameter departureDate",
        "GET /shopping/flight-offers: parameter returnDate",
        "GET /shopping/flight-offers: parameter travelClass",
        "POST /shopping/flight-offers: parameter X-HTTP-Method-Override",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.cabinRestrictions[].cabin: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.cabinRestrictions[].coverage: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.cabinRestrictions[].originDestinationIds[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.carrierRestrictions.excludedCarrierCodes[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.searchCriteria.flightFilters.carrierRestrictions.includedCarrierCodes[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.sources[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.travelers[].associatedAdultId: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.travelers[].id: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.travelers[].travelerType: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.currencyCode: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].includedConnectionPoints[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].id: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].alternativeDestinationsCodes[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].arrivalDateTimeRange.time: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].arrivalDateTimeRange.date: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].destinationLocationCode: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].excludedConnectionPoints[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].originLocationCode: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].alternativeOriginsCodes[]: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].departureDateTimeRange.time: application/vnd.amadeus+json schema",
        "POST /shopping/flight-offers.originDestinations[].departureDateTimeRange.date: application/vnd.amadeus+json schema"
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
- **Message:** Error handling issues found: Error responses include error details schema: GET /shopping/flight-offers: 400 response, POST /shopping/flight-offers: 400 response; All operations document 5xx error responses: GET /shopping/flight-offers, POST /shopping/flight-offers
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
      "All operations document 4xx error responses": true,
      "All operations document 5xx error responses": false,
      "Common error responses are defined in components": true,
      "Error responses follow consistent format": true,
      "Error responses include error details schema": false
    },
    "messages": {},
    "missing_errors": {
      "All operations document 5xx error responses": [
        "GET /shopping/flight-offers",
        "POST /shopping/flight-offers"
      ],
      "Error responses include error details schema": [
        "GET /shopping/flight-offers: 400 response",
        "POST /shopping/flight-offers: 400 response"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive error response documentation including codes, messages, and consistent error schemas

---

### Security

#### P005: Security Standards (Failed) [critical]

Validates comprehensive security requirements and authentication mechanisms

- **Status:** Failed
- **Message:** Security validation failed: Operation-level security is defined: Endpoints without security: POST /shopping/flight-offers, GET /shopping/flight-offers; Security schemes are defined: No security schemes defined in components.securitySchemes; Global security requirements are set: No top-level security requirements defined
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
      "Global security requirements are set": false,
      "OAuth2 scopes are documented": true,
      "Operation-level security is defined": false,
      "Security requirements are consistent": true,
      "Security schemes are defined": false
    },
    "messages": {
      "Global security requirements are set": "No top-level security requirements defined",
      "Operation-level security is defined": "Endpoints without security: POST /shopping/flight-offers, GET /shopping/flight-offers",
      "Security schemes are defined": "No security schemes defined in components.securitySchemes"
    }
  }
```

**Suggested Fix:**
Review security schemes, global security requirements, and per-operation security

---

### Versioning

#### P008: API Versioning Strategy (Failed) [warning]

Validates proper API versioning implementation and documentation

- **Status:** Failed
- **Message:** Versioning validation failed: Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Versioning strategy is documented: Info description does not mention versioning strategy
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

