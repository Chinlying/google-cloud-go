# Changes

## [1.11.2](https://github.com/googleapis/google-cloud-go/compare/iap/v1.11.1...iap/v1.11.2) (2025-05-29)


### Documentation

* **iap:** Add clarifications on IAP CorsSettings behavior ([#12325](https://github.com/googleapis/google-cloud-go/issues/12325)) ([8189e33](https://github.com/googleapis/google-cloud-go/commit/8189e3313ed62b99cc238c421ae9acfa32aaf9af))
* **iap:** Minor updates on GcipSettings description ([8189e33](https://github.com/googleapis/google-cloud-go/commit/8189e3313ed62b99cc238c421ae9acfa32aaf9af))

## [1.11.1](https://github.com/googleapis/google-cloud-go/compare/iap/v1.11.0...iap/v1.11.1) (2025-04-15)


### Bug Fixes

* **iap:** Update google.golang.org/api to 0.229.0 ([3319672](https://github.com/googleapis/google-cloud-go/commit/3319672f3dba84a7150772ccb5433e02dab7e201))

## [1.11.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.10.5...iap/v1.11.0) (2025-04-15)


### Features

* **iap:** Identity-aware Proxy (IAP) released a feature `Use IAP with Workforce Identity Federation`(https ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))


### Documentation

* **iap:** A comment for enum `PolicyType` is changed ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** A comment for field `cidrs` in message `.google.cloud.iap.v1.TunnelDestGroup` is changed ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** A comment for field `fqdns` in message `.google.cloud.iap.v1.TunnelDestGroup` is changed ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** A comment for field `name` in message `.google.cloud.iap.v1.TunnelDestGroup` is changed ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `access_denied_page_settings` in message `.google.cloud.iap.v1.ApplicationSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `access_settings` in message `.google.cloud.iap.v1.IapSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `allowed_domains_settings` in message `.google.cloud.iap.v1.AccessSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `application_settings` in message `.google.cloud.iap.v1.IapSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `attribute_propagation_settings` in message `.google.cloud.iap.v1.ApplicationSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `cors_settings` in message `.google.cloud.iap.v1.AccessSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `csm_settings` in message `.google.cloud.iap.v1.ApplicationSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `domains` in message `.google.cloud.iap.v1.AllowedDomainsSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `enable` in message `.google.cloud.iap.v1.AllowedDomainsSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `enable` in message `.google.cloud.iap.v1.AttributePropagationSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `expression` in message `.google.cloud.iap.v1.AttributePropagationSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `gcip_settings` in message `.google.cloud.iap.v1.AccessSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `max_age` in message `.google.cloud.iap.v1.ReauthSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `method` in message `.google.cloud.iap.v1.ReauthSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `oauth_settings` in message `.google.cloud.iap.v1.AccessSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `output_credentials` in message `.google.cloud.iap.v1.AttributePropagationSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `policy_type` in message `.google.cloud.iap.v1.ReauthSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `programmatic_clients` in message `.google.cloud.iap.v1.OAuthSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `reauth_settings` in message `.google.cloud.iap.v1.AccessSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))
* **iap:** Mark `tenant_ids` in message `.google.cloud.iap.v1.GcipSettings` as optional ([8d466c4](https://github.com/googleapis/google-cloud-go/commit/8d466c492fb1a15e1e857268397e795287fb844e))

## [1.10.5](https://github.com/googleapis/google-cloud-go/compare/iap/v1.10.4...iap/v1.10.5) (2025-03-13)


### Bug Fixes

* **iap:** Update golang.org/x/net to 0.37.0 ([1144978](https://github.com/googleapis/google-cloud-go/commit/11449782c7fb4896bf8b8b9cde8e7441c84fb2fd))

## [1.10.4](https://github.com/googleapis/google-cloud-go/compare/iap/v1.10.3...iap/v1.10.4) (2025-03-06)


### Bug Fixes

* **iap:** Fix out-of-sync version.go ([28f0030](https://github.com/googleapis/google-cloud-go/commit/28f00304ebb13abfd0da2f45b9b79de093cca1ec))

## [1.10.3](https://github.com/googleapis/google-cloud-go/compare/iap/v1.10.2...iap/v1.10.3) (2025-01-02)


### Bug Fixes

* **iap:** Update golang.org/x/net to v0.33.0 ([e9b0b69](https://github.com/googleapis/google-cloud-go/commit/e9b0b69644ea5b276cacff0a707e8a5e87efafc9))

## [1.10.2](https://github.com/googleapis/google-cloud-go/compare/iap/v1.10.1...iap/v1.10.2) (2024-10-23)


### Bug Fixes

* **iap:** Update google.golang.org/api to v0.203.0 ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))
* **iap:** WARNING: On approximately Dec 1, 2024, an update to Protobuf will change service registration function signatures to use an interface instead of a concrete type in generated .pb.go files. This change is expected to affect very few if any users of this client library. For more information, see https://togithub.com/googleapis/google-cloud-go/issues/11020. ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))

## [1.10.1](https://github.com/googleapis/google-cloud-go/compare/iap/v1.10.0...iap/v1.10.1) (2024-09-12)


### Bug Fixes

* **iap:** Bump dependencies ([2ddeb15](https://github.com/googleapis/google-cloud-go/commit/2ddeb1544a53188a7592046b98913982f1b0cf04))

## [1.10.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.11...iap/v1.10.0) (2024-08-20)


### Features

* **iap:** Add support for Go 1.23 iterators ([84461c0](https://github.com/googleapis/google-cloud-go/commit/84461c0ba464ec2f951987ba60030e37c8a8fc18))

## [1.9.11](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.10...iap/v1.9.11) (2024-08-08)


### Bug Fixes

* **iap:** Update google.golang.org/api to v0.191.0 ([5b32644](https://github.com/googleapis/google-cloud-go/commit/5b32644eb82eb6bd6021f80b4fad471c60fb9d73))

## [1.9.10](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.9...iap/v1.9.10) (2024-07-24)


### Bug Fixes

* **iap:** Update dependencies ([257c40b](https://github.com/googleapis/google-cloud-go/commit/257c40bd6d7e59730017cf32bda8823d7a232758))

## [1.9.9](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.8...iap/v1.9.9) (2024-07-10)


### Bug Fixes

* **iap:** Bump google.golang.org/grpc@v1.64.1 ([8ecc4e9](https://github.com/googleapis/google-cloud-go/commit/8ecc4e9622e5bbe9b90384d5848ab816027226c5))

## [1.9.8](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.7...iap/v1.9.8) (2024-07-01)


### Bug Fixes

* **iap:** Bump google.golang.org/api@v0.187.0 ([8fa9e39](https://github.com/googleapis/google-cloud-go/commit/8fa9e398e512fd8533fd49060371e61b5725a85b))

## [1.9.7](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.6...iap/v1.9.7) (2024-06-26)


### Bug Fixes

* **iap:** Enable new auth lib ([b95805f](https://github.com/googleapis/google-cloud-go/commit/b95805f4c87d3e8d10ea23bd7a2d68d7a4157568))

## [1.9.6](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.5...iap/v1.9.6) (2024-05-01)


### Bug Fixes

* **iap:** Bump x/net to v0.24.0 ([ba31ed5](https://github.com/googleapis/google-cloud-go/commit/ba31ed5fda2c9664f2e1cf972469295e63deb5b4))

## [1.9.5](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.4...iap/v1.9.5) (2024-03-14)


### Bug Fixes

* **iap:** Update protobuf dep to v1.33.0 ([30b038d](https://github.com/googleapis/google-cloud-go/commit/30b038d8cac0b8cd5dd4761c87f3f298760dd33a))

## [1.9.4](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.3...iap/v1.9.4) (2024-01-30)


### Bug Fixes

* **iap:** Enable universe domain resolution options ([fd1d569](https://github.com/googleapis/google-cloud-go/commit/fd1d56930fa8a747be35a224611f4797b8aeb698))

## [1.9.3](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.2...iap/v1.9.3) (2023-11-01)


### Bug Fixes

* **iap:** Bump google.golang.org/api to v0.149.0 ([8d2ab9f](https://github.com/googleapis/google-cloud-go/commit/8d2ab9f320a86c1c0fab90513fc05861561d0880))

## [1.9.2](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.1...iap/v1.9.2) (2023-10-26)


### Bug Fixes

* **iap:** Update grpc-go to v1.59.0 ([81a97b0](https://github.com/googleapis/google-cloud-go/commit/81a97b06cb28b25432e4ece595c55a9857e960b7))

## [1.9.1](https://github.com/googleapis/google-cloud-go/compare/iap/v1.9.0...iap/v1.9.1) (2023-10-12)


### Bug Fixes

* **iap:** Update golang.org/x/net to v0.17.0 ([174da47](https://github.com/googleapis/google-cloud-go/commit/174da47254fefb12921bbfc65b7829a453af6f5d))

## [1.9.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.8.1...iap/v1.9.0) (2023-09-20)


### Features

* **iap:** Adding programmatic_clients attribute to UpdateIapSettings API request ([2f3bb44](https://github.com/googleapis/google-cloud-go/commit/2f3bb443e9fa6968d20806f86b391dad85970afc))


### Documentation

* **iap:** Fixing Oauth typo ([2f3bb44](https://github.com/googleapis/google-cloud-go/commit/2f3bb443e9fa6968d20806f86b391dad85970afc))

## [1.8.1](https://github.com/googleapis/google-cloud-go/compare/iap/v1.8.0...iap/v1.8.1) (2023-06-20)


### Bug Fixes

* **iap:** REST query UpdateMask bug ([df52820](https://github.com/googleapis/google-cloud-go/commit/df52820b0e7721954809a8aa8700b93c5662dc9b))

## [1.8.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.7.3...iap/v1.8.0) (2023-05-30)


### Features

* **iap:** Update all direct dependencies ([b340d03](https://github.com/googleapis/google-cloud-go/commit/b340d030f2b52a4ce48846ce63984b28583abde6))

## [1.7.3](https://github.com/googleapis/google-cloud-go/compare/iap/v1.7.2...iap/v1.7.3) (2023-05-08)


### Bug Fixes

* **iap:** Update grpc to v1.55.0 ([1147ce0](https://github.com/googleapis/google-cloud-go/commit/1147ce02a990276ca4f8ab7a1ab65c14da4450ef))

## [1.7.2](https://github.com/googleapis/google-cloud-go/compare/iap/v1.7.1...iap/v1.7.2) (2023-04-11)


### Documentation

* **iap:** Few minor changes on doc description came out of cl/512701532 ([19f18c0](https://github.com/googleapis/google-cloud-go/commit/19f18c0a33d85e1949981d58bca2b765ce9787b5))

## [1.7.1](https://github.com/googleapis/google-cloud-go/compare/iap/v1.7.0...iap/v1.7.1) (2023-04-04)


### Documentation

* **iap:** Few minor changes on doc description came out of cl/512701532 ([c893c15](https://github.com/googleapis/google-cloud-go/commit/c893c158f1e6d03b0cde45dda2059c0e2aa9ead1))

## [1.7.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.6.0...iap/v1.7.0) (2023-03-15)


### Features

* **iap:** Add an enum ENROLLED_SECOND_FACTORS under IapSettings ([8775cae](https://github.com/googleapis/google-cloud-go/commit/8775cae47a9efb358ce34240853a1b09c7f6dc62))
* **iap:** Update iam and longrunning deps ([91a1f78](https://github.com/googleapis/google-cloud-go/commit/91a1f784a109da70f63b96414bba8a9b4254cddd))

## [1.6.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.5.0...iap/v1.6.0) (2023-01-04)


### Features

* **iap:** Add REST client ([06a54a1](https://github.com/googleapis/google-cloud-go/commit/06a54a16a5866cce966547c51e203b9e09a25bc0))

## [1.5.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.4.0...iap/v1.5.0) (2022-11-03)


### Features

* **iap:** rewrite signatures in terms of new location ([3c4b2b3](https://github.com/googleapis/google-cloud-go/commit/3c4b2b34565795537aac1661e6af2442437e34ad))

## [1.4.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.3.0...iap/v1.4.0) (2022-10-25)


### Features

* **iap:** start generating stubs dir ([de2d180](https://github.com/googleapis/google-cloud-go/commit/de2d18066dc613b72f6f8db93ca60146dabcfdcc))

## [1.3.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.2.0...iap/v1.3.0) (2022-05-24)


### Features

* **iap:** add the TunnelDestGroup-related methods and types feat: add ReauthSettings to the UpdateIapSettingsRequest ([6ef576e](https://github.com/googleapis/google-cloud-go/commit/6ef576e2d821d079e7b940cd5d49fe3ca64a7ba2))

## [1.2.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.1.0...iap/v1.2.0) (2022-02-23)


### Features

* **iap:** set versionClient to module version ([55f0d92](https://github.com/googleapis/google-cloud-go/commit/55f0d92bf112f14b024b4ab0076c9875a17423c9))

## [1.1.0](https://github.com/googleapis/google-cloud-go/compare/iap/v1.0.1...iap/v1.1.0) (2022-02-14)


### Features

* **iap:** add file for tracking version ([17b36ea](https://github.com/googleapis/google-cloud-go/commit/17b36ead42a96b1a01105122074e65164357519e))

### [1.0.1](https://www.github.com/googleapis/google-cloud-go/compare/iap/v1.0.0...iap/v1.0.1) (2022-02-03)


### Bug Fixes

* **iap:** remove unpublished service from service config ([6e56077](https://www.github.com/googleapis/google-cloud-go/commit/6e560776fd6e574320ce2dbad1f9eb9e22999185))

## 1.0.0

Stabilize GA surface.

## v0.1.0

This is the first tag to carve out iap as its own module. See
[Add a module to a multi-module repository](https://github.com/golang/go/wiki/Modules#is-it-possible-to-add-a-module-to-a-multi-module-repository).
