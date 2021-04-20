#!/bin/sh

pingctl()
{

  echo "
###########################################
# pingctl ${@}
###########################################"

./pingctl ${@}
}

pingctl pingone get users
pingctl pingone get groups
pingctl pingone get populations

pingctl pingone add population pingctl-test-pop
pingctl pingone add group pingctl-test-group -p pingctl-test-pop
pingctl pingone add user pingctl-test-user@_example.com pingctl-test-First pingctl-test-Last -p pingctl-test-pop

pingctl pingone add-user-group pingctl-test-user@_example.com pingctl-test-group

pingctl pingone get populations -n pingctl-test-pop
pingctl pingone get users -n pingctl-test-user@_example.com
pingctl pingone get groups -n pingctl-test-group

pingctl pingone delete-user-group pingctl-test-user@_example.com pingctl-test-group
pingctl pingone delete user pingctl-test-user@_example.com
pingctl pingone delete group pingctl-test-group
pingctl pingone delete population pingctl-test-pop