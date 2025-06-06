# Changelog




## [1.4.0](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.3.2...cloudcontrolspartner/v1.4.0) (2025-05-06)


### Features

* **cloudcontrolspartner:** A new field `organization_domain` is added to message `Customer` ([658d3d6](https://github.com/googleapis/google-cloud-go/commit/658d3d6cfb37bdf15ae40250fc1b9cb7b3b5ef80))
* **cloudcontrolspartner:** Enable partners to create, update and delete their customers ([658d3d6](https://github.com/googleapis/google-cloud-go/commit/658d3d6cfb37bdf15ae40250fc1b9cb7b3b5ef80))


### Documentation

* **cloudcontrolspartner:** Mark the enum value `EkmSolution.VIRTRU` as deprecated ([658d3d6](https://github.com/googleapis/google-cloud-go/commit/658d3d6cfb37bdf15ae40250fc1b9cb7b3b5ef80))

## [1.3.2](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.3.1...cloudcontrolspartner/v1.3.2) (2025-04-15)


### Bug Fixes

* **cloudcontrolspartner:** Update google.golang.org/api to 0.229.0 ([3319672](https://github.com/googleapis/google-cloud-go/commit/3319672f3dba84a7150772ccb5433e02dab7e201))

## [1.3.1](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.3.0...cloudcontrolspartner/v1.3.1) (2025-03-13)


### Bug Fixes

* **cloudcontrolspartner:** Update golang.org/x/net to 0.37.0 ([1144978](https://github.com/googleapis/google-cloud-go/commit/11449782c7fb4896bf8b8b9cde8e7441c84fb2fd))

## [1.3.0](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.2.2...cloudcontrolspartner/v1.3.0) (2025-01-30)


### Features

* **cloudcontrolspartner:** A new field `organization_domain` is added to message `.google.cloud.cloudcontrolspartner.v1beta.Customer` ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A new message `CreateCustomerRequest` is added ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A new message `DeleteCustomerRequest` is added ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A new message `UpdateCustomerRequest` is added ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A new method `CreateCustomer` is added to service `CloudControlsPartnerCore` ([#11504](https://github.com/googleapis/google-cloud-go/issues/11504)) ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A new method `DeleteCustomer` is added to service `CloudControlsPartnerCore` ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A new method `UpdateCustomer` is added to service `CloudControlsPartnerCore` ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))


### Documentation

* **cloudcontrolspartner:** A comment for enum value `VIRTRU` in enum `EkmSolution` is changed ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))
* **cloudcontrolspartner:** A comment for field `requested_cancellation` in message `.google.cloud.cloudcontrolspartner.v1beta.OperationMetadata` is changed ([232775b](https://github.com/googleapis/google-cloud-go/commit/232775bc8691c1a53b9a06dcacc87975e7c9e6d8))

## [1.2.2](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.2.1...cloudcontrolspartner/v1.2.2) (2025-01-02)


### Bug Fixes

* **cloudcontrolspartner:** Update golang.org/x/net to v0.33.0 ([e9b0b69](https://github.com/googleapis/google-cloud-go/commit/e9b0b69644ea5b276cacff0a707e8a5e87efafc9))

## [1.2.1](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.2.0...cloudcontrolspartner/v1.2.1) (2024-10-23)


### Bug Fixes

* **cloudcontrolspartner:** Update google.golang.org/api to v0.203.0 ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))
* **cloudcontrolspartner:** WARNING: On approximately Dec 1, 2024, an update to Protobuf will change service registration function signatures to use an interface instead of a concrete type in generated .pb.go files. This change is expected to affect very few if any users of this client library. For more information, see https://togithub.com/googleapis/google-cloud-go/issues/11020. ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))

## [1.2.0](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.1.1...cloudcontrolspartner/v1.2.0) (2024-09-19)


### Features

* **cloudcontrolspartner:** A new value `ACCESS_TRANSPARENCY_LOGS_SUPPORT_CASE_VIEWER` is added to enum `.google.cloud.cloudcontrolspartner.v1.PartnerPermissions.Permission` ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** A new value `ACCESS_TRANSPARENCY_LOGS_SUPPORT_CASE_VIEWER` is added to enum `.google.cloud.cloudcontrolspartner.v1beta.PartnerPermissions.Permission` ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** Field behavior for field `customer_onboarding_state` in message `.google.cloud.cloudcontrolspartner.v1.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** Field behavior for field `customer_onboarding_state` in message `.google.cloud.cloudcontrolspartner.v1beta.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** Field behavior for field `is_onboarded` in message `.google.cloud.cloudcontrolspartner.v1.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** Field behavior for field `is_onboarded` in message `.google.cloud.cloudcontrolspartner.v1beta.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))


### Bug Fixes

* **cloudcontrolspartner:** Field behavior for field `display_name` in message `.google.cloud.cloudcontrolspartner.v1.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** Field behavior for field `display_name` in message `.google.cloud.cloudcontrolspartner.v1beta.Customer` is changed ([#10865](https://github.com/googleapis/google-cloud-go/issues/10865)) ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))


### Documentation

* **cloudcontrolspartner:** A comment for field `display_name` in message `.google.cloud.cloudcontrolspartner.v1.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))
* **cloudcontrolspartner:** A comment for field `display_name` in message `.google.cloud.cloudcontrolspartner.v1beta.Customer` is changed ([9efa812](https://github.com/googleapis/google-cloud-go/commit/9efa8127fa72da9f084548822c567db2195f5a94))

## [1.1.1](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.1.0...cloudcontrolspartner/v1.1.1) (2024-09-12)


### Bug Fixes

* **cloudcontrolspartner:** Bump dependencies ([2ddeb15](https://github.com/googleapis/google-cloud-go/commit/2ddeb1544a53188a7592046b98913982f1b0cf04))

## [1.1.0](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.0.4...cloudcontrolspartner/v1.1.0) (2024-08-20)


### Features

* **cloudcontrolspartner:** Add support for Go 1.23 iterators ([84461c0](https://github.com/googleapis/google-cloud-go/commit/84461c0ba464ec2f951987ba60030e37c8a8fc18))

## [1.0.4](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.0.3...cloudcontrolspartner/v1.0.4) (2024-08-08)


### Bug Fixes

* **cloudcontrolspartner:** Update google.golang.org/api to v0.191.0 ([5b32644](https://github.com/googleapis/google-cloud-go/commit/5b32644eb82eb6bd6021f80b4fad471c60fb9d73))

## [1.0.3](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.0.2...cloudcontrolspartner/v1.0.3) (2024-07-24)


### Bug Fixes

* **cloudcontrolspartner:** Update dependencies ([257c40b](https://github.com/googleapis/google-cloud-go/commit/257c40bd6d7e59730017cf32bda8823d7a232758))

## [1.0.2](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.0.1...cloudcontrolspartner/v1.0.2) (2024-07-10)


### Bug Fixes

* **cloudcontrolspartner:** Bump google.golang.org/grpc@v1.64.1 ([8ecc4e9](https://github.com/googleapis/google-cloud-go/commit/8ecc4e9622e5bbe9b90384d5848ab816027226c5))

## [1.0.1](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v1.0.0...cloudcontrolspartner/v1.0.1) (2024-07-01)


### Bug Fixes

* **cloudcontrolspartner:** Bump google.golang.org/api@v0.187.0 ([8fa9e39](https://github.com/googleapis/google-cloud-go/commit/8fa9e398e512fd8533fd49060371e61b5725a85b))

## [1.0.0](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v0.2.1...cloudcontrolspartner/v1.0.0) (2024-06-26)


### Documentation

* **cloudcontrolspartner:** Mark the accessApprovalRequests.list method as deprecated ([d6c543c](https://github.com/googleapis/google-cloud-go/commit/d6c543c3969016c63e158a862fc173dff60fb8d9))
* **cloudcontrolspartner:** Mark the accessApprovalRequests.list method as deprecated ([d6c543c](https://github.com/googleapis/google-cloud-go/commit/d6c543c3969016c63e158a862fc173dff60fb8d9))


### Miscellaneous Chores

* **cloudcontrolspartner:** Release v1.0.0 ([#10443](https://github.com/googleapis/google-cloud-go/issues/10443)) ([2980bd1](https://github.com/googleapis/google-cloud-go/commit/2980bd1a8ef6c5a990a1311ab756a9669bedd291))

## [0.2.1](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v0.2.0...cloudcontrolspartner/v0.2.1) (2024-05-01)


### Bug Fixes

* **cloudcontrolspartner:** Bump x/net to v0.24.0 ([ba31ed5](https://github.com/googleapis/google-cloud-go/commit/ba31ed5fda2c9664f2e1cf972469295e63deb5b4))

## [0.2.0](https://github.com/googleapis/google-cloud-go/compare/cloudcontrolspartner/v0.1.0...cloudcontrolspartner/v0.2.0) (2024-03-25)


### Features

* **cloudcontrolspartner:** New clients ([#9643](https://github.com/googleapis/google-cloud-go/issues/9643)) ([75d0768](https://github.com/googleapis/google-cloud-go/commit/75d0768e1a779cdf829d12e4036e7a65671acf1b))

## 0.1.0 (2024-03-12)


### Features

* **cloudcontrolspartner:** New client(s) ([#9563](https://github.com/googleapis/google-cloud-go/issues/9563)) ([601f21a](https://github.com/googleapis/google-cloud-go/commit/601f21af3925fa43628739f314112ce4c754b4ce))

## Changes
