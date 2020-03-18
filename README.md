# Simple Golang REST web service example

Simple Golang REST web service example easily deployable in OpenShift.

## Origins

Forked from https://github.com/WPFerg/services.

- converted to Golang modules project (added go.mod file)
- handlers changed a bit to support specific URL paths
- added Dockerfile & OpenShift template
  - copied from https://github.com/openshift/golang-ex and adjusted

## Deploy to OpenShift

- fork the project
- make some changes (if you want to)
- deploy using OpenShift template with proper parameters set
  - your repo as SOURCE_REPOSITORY_URL parameter
  - if you are not using nip.ip set proper APPLICATION_DOMAIN parameter

oc new-app openshift/templates/go-ws-example.yaml -p SOURCE_REPOSITORY_URL=https://github.com/your_repo_path -p APPLICATION_DOMAIN=go-ws-example.your.domain.net
