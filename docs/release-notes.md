---
title: pingctl Release Notes
---

# Release Notes

## Release 1.0.5 (May 03, 2022)

* Fix some issues where pingctl did not function correctly on Ubuntu.
* Update pingctl to better comply with POSIX standards.

## Release 1.0.4 (Feb 03, 2022)

* Update pingctl to better comply with POSIX standards and better support adding source to shell
* Fixed some minor prompt syntax

## Release 1.0.3 (April 30, 2021)

* Fix version check and upgrade instructions
* Allow user to default ping-devops configuration if pingctl configuration not setup yet.  Easier
  for first time users.

## Release 1.0.2 (April 26, 2021)

Fixed minor typos with pingctl tool and docs.

## Release 1.0.1 (April 26, 2021)

Added support for Authorization Code (w/ pkce) and Implicit flows so that users of the tool can
login with a user account, with proper role access.  This is in addition to client_credentials.

## Release 1.0.0 (April 22, 2021)

Initial release of pingctl CLI tool, as well as documentation for all product features including
the following areas:

* CLI to PingOne Environments
    * List resources (users, groups, populations)
    * Add resources (incl groups to users)
    * Delete resources (incl groups from users)
    * Retrieve PingOne Token

* Ping Identity Evaluation Product Licenses
    * Retrieve licenses

* Generate Kubernetes Resources
    * devops-secret, tls-secrets, license-secrets

* Manage kubectl oidc token
    * get token into, clear, get token claims
