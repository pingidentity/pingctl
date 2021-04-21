---
title: pingctl - a Ping Identity CLI
---

# pingctl pingone

## Description

Provides ability to manage PingOne environments.  Includes features:

* Listing, searching and retrieving PingOne resources (i.e. user, populations, groups)
* Add PingOne resources
* Deleting PingOne resources

## Usage

    pingctl pingone get                  # Get PingOne resource(s)
    pingctl pingone add                  # Add PingOne resource
    pingctl pingone delete               # Delete PingOne resource

    pingctl pingone add-user-group       # Add group to user
    pingctl pingone delete-user-group    # Delete group from user

    pingctl pingone token                # Obtain access token

## Options

### All subcommands

    -r
        Provide REST Calls

### get

    -o [ table | csv | json ]
        Output format (default: table)
        also set with env variable: PING_DEFAULT_OUTPUT

    -i {id}
        Search based on object guid

    -n {name}
        Search based on exact filter

    -f {filter}
        PingOne filter (SCIM based)
            ex: '.name.given eq "john"'
                '.email sw "john"'

### add

    -p {population name}
        Population to add user/group into.
        If not provided, 'Default' population used
