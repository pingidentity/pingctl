---
title: pingctl Environment and Configuration variables
---

# Congiguration & Environment variables

Congiguration and Environment variables allow for users to cache secure and repetitive settings into
a `pingctl` config file.  The default location of the file is `~/.pingidentity/config`.

In cases where the a configuration item might be specified at any of the three levels of the
`pingctl` config file, the users current environment, or the command line arguments.  The rule is:

* Command-Line argument overrides
* `pingctl` config file
* Environment variable overrides

The available variables honored by `pingctl` are as follows:

| Variable                                   | Description                                                                                                     |
| ------------------------------------------ | --------------------------------------------------------------------------------------------------------------- |
| **PINGCTL_CONFIG**                         | Location of the `pingctl` configuration file. Defaults to: `~/.pingidentity/config`                             |
| **PINGCTL_DEFAULT_OUTPUT**                 | Specifies format of data returned. Defaults to: `table`                                                         |
| **PINGCTL_OUTPUT_COLUMNS_{resource_type}** | Specify custom format of table csv data to be returned.  See more detail [below](#ping_output-columns-and-sort) |
| **PINGCTL_OUTPUT_SORT_{resource_type}**    | Specify column to sort data.  See more detail [below](#ping_output-columns-and-sort)                            |

## PINGCTL_OUTPUT Columns and Sort

There are two classes of variables under the `PINGCTL_OUTPUT` name that provides:

* `PINGCTL_OUTPUT_COLUMNS_{resource}` - Specifies the columns to display whenever a `pingctl pingone get {resource}` command is used.

    Same as the `-c` option on the command-line (see [pingctl pingone get](commands/pingone.md) command).

    Format of value should be constructed with `HeadingName:jsonName,HeadingName:jsonName`.  The best way to understand is
    looking at the example of the default `USERS` resource:

    !!! note "Example PINGCTL_OUTPUT_COLUMNS_USERS setting and output"
        ```
        PINGCTL_OUTPUT_COLUMNS_USERS=LastName:name.family,FirstName:name.given
        ```

        will generate output, looking like:

        ```
        $ pingctl pingone get users
        LastName     FirstName
        --------     ---------
        Adham        Antonik
        Agnès        Enterle
        --
        2 'USERS' returned
        ```

        can also use the `-c` option as command line argument:

        ```
        $ pingctl pingone get users -c "LastName:name.family,FirstName:name.given,Username:username"
        LastName     FirstName    Username
        --------     ---------    --------
        Adham        Antonik      antonik_adham
        Agnès        Enterle      enterle_agnès
        --
        2 'USERS' returned
        ```


* `PINGCTL_OUTPUT_SORT_{resource}` - Specifies the column to sort on.

    Same as the `-s` option on the command-line (see [pingctl pingone get](commands/pingone.md) command).

    Format of value should be constructed with `jsonName`.  The name must be of the names in `PINGCTL_OUTPUT_COLUMNS_{resource}`.

    !!! note "Example PINGCTL_OUTPUT_SORT_USERS setting and output"
        ```
        PINGCTL_OUTPUT_SORT_USERS=name.family
        ```

        will generate output, looking like (note that the LastName, aka name.family, is what is sorted):

        ```
        $ pingctl pingone get users
        LastName     FirstName
        --------     ---------
        Adham        Antonik
        Agnès        Enterle
        --
        2 'USERS' returned
        ```

        can also use the `-s` option as command line argument:

        ```
        $ pingctl pingone get users -s "name.given"
        LastName     FirstName    Username
        --------     ---------    --------
        Adham        Antonik      antonik_adham
        Agnès        Enterle      enterle_agnès
        --
        2 'USERS' returned
        ```