#!/bin/sh

pingctl()
{

  echo "
###########################################
# pingctl ${@}
###########################################"

./pingctl ${@}
}

pingctl info
pingctl version

pingctl pingone get users
pingctl pingone get groups
pingctl pingone get populations

pingctl pingone get users -c "LastName:name.family,FirstName:name.given,Username:username" -s name.family

pingctl pingone add population pingctl-test-pop
pingctl pingone add group pingctl-test-group -p pingctl-test-pop
pingctl pingone add user pingctl-test-user@_example.com pingctl-test-First pingctl-test-Last -p pingctl-test-pop

pingctl pingone get users -p pingctl-test-pop

pingctl pingone add-user-group pingctl-test-user@_example.com pingctl-test-group

pingctl pingone get populations -n pingctl-test-pop
pingctl pingone get users -n pingctl-test-user@_example.com
pingctl pingone get groups -n pingctl-test-group

pingctl pingone delete-user-group pingctl-test-user@_example.com pingctl-test-group
pingctl pingone delete user pingctl-test-user@_example.com
pingctl pingone delete group pingctl-test-group
pingctl pingone delete population pingctl-test-pop

pingctl pingone token

pingctl license pingfederate 10.2
pingctl k8s generate devops-secret
pingctl k8s generate tls-secret example.com
pingctl k8s generate license-secret pingfederate 10.2
