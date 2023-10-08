---
title: pingctl - a Ping Identity CLI
---

# pingctl

## Description

Ping Identity Command Line Interface (CLI) tool used for PingOne and future command line tools and aliases.

## Usage

    pingctl <command> [options]

    Available Commands:
        info            Print pingctl config
        config          Manage pingctl config
        version         Version Details and Check
        clean           Remove ~/.pingidentity/pingctl

        kubernetes      Kubernetes Tools
        license         Ping Identity Licensing Tools
        pingone         PingOne Tools

Use `pingctl` for info on available commands.

Use `pingctl <command>` for info on a specific command.

## Options

   -h

Provide usage details.

## Available Commands

* [kubernetes](commands/kubernetes)
* [license](commands/license)
* [pingone](commands/pingone.md)
* info

    Provides a summary of variables defined with pingctl.

* config

    Provides an interactive interview allowing user to provide all the `pingctl` standard
    variables (i.e. PingOne and Ping DevOps) as well as custom variables

* version

    Gets the current version of the tool, and checks to see if an update is available.

* clean

    Cleans the cached pingctl work directory containing:

    * Latest PingOne Access Token
