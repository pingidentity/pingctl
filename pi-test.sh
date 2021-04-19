#!/bin/sh

pi()
{

  echo "
###########################################
# pi ${@}
###########################################"

./pi ${@}
}

pi pingone get users
pi pingone get groups
pi pingone get populations

pi pingone add population pi-test-pop
pi pingone add group pi-test-group -p pi-test-pop
pi pingone add user pi-test-user@_example.com pi-test-First pi-test-Last -p pi-test-pop

pi pingone add-user-group pi-test-user@_example.com pi-test-group

pi pingone get populations -n pi-test-pop
pi pingone get users -n pi-test-user@_example.com
pi pingone get groups -n pi-test-group

pi pingone delete-user-group pi-test-user@_example.com pi-test-group
pi pingone delete user pi-test-user@_example.com
pi pingone delete group pi-test-group
pi pingone delete population pi-test-pop