#!/bin/sh

pi()
{

  echo "
###########################################
# pi ${@}
###########################################"

./pi ${@}
}

pi pingone list users
pi pingone list groups
pi pingone list populations

pi pingone add population pi-test-pop
pi pingone add group pi-test-group -p pi-test-pop
pi pingone add user pi-test-user@_example.com pi-test-First pi-test-Last -p pi-test-pop

pi pingone add-user-group pi-test-user@_example.com pi-test-group

pi pingone get population pi-test-pop
pi pingone get user pi-test-user@_example.com
pi pingone get group pi-test-group

pi pingone delete-user-group pi-test-user@_example.com pi-test-group
pi pingone delete user pi-test-user@_example.com
pi pingone delete group pi-test-group
pi pingone delete population pi-test-pop