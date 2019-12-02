#!make
SHELL := /bin/bash
.SHELLFLAGS := -ec

## General VARS
VERSION ?= dev
APPLICATION_NAME= mass-lobby
AWS_PROFILE=enekofb
AWS_ACCOUNT_ID=777171359344

## Test VARS
ACCEPTANCE_TESTS_DIRECTORY=acceptance

## Builder VARS
BUILD_BUILDER_DIRECTORY=build/builder
DOCKER_REGISTRY ?= docker.io
DOCKER_ORG = enekofb
DOCKER_REPO = ${APPLICATION}-builder

## Build VARS
BUILD_CLOUD_DIRECTORY=build/cloud
CLOUD_FORMATION_TEMPLATE_FILENAME=cloudformation-template.yaml
CLOUD_FORMATION_TEMPLATE_PATH=${BUILD_CLOUD_DIRECTORY}/${CLOUD_FORMATION_TEMPLATE_FILENAME}

default: test

## builder 
builder-build:
	cd ${BUILD_BUILDER_DIRECTORY} &&  docker build . -t $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_REPO):$(VERSION)

builder-push:
	docker push  $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_REPO):$(VERSION)

## test
acceptance-test:
	cd ${ACCEPTANCE_TESTS_DIRECTORY} &&  go test -v 

## build
cloud-deploy: 
	aws cloudformation deploy \
	--profile ${AWS_PROFILE} \
	--stack-name ${APPLICATION_NAME}-stack \
	--template-file ${CLOUD_FORMATION_TEMPLATE_PATH} \
	--capabilities CAPABILITY_NAMED_IAM \
	--parameter-overrides \
	BucketName=${AWS_ACCOUNT_ID}-${APPLICATION_NAME}-cloudformation-bucket

cloud-delete: 
	aws cloudformation delete-stack \
	--profile ${AWS_PROFILE} \
	--stack-name ${APPLICATION_NAME}-stack
	aws cloudformation wait stack-delete-complete \
	--profile ${AWS_PROFILE} \
	--stack-name ${APPLICATION_NAME}-stack

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	$(eval WEBSITE_PATH := $(GOPATH)/src/$(WEBSITE_REPO))
	@echo "==> Starting core website in Docker..."
	@docker run \
		--interactive \
		--rm \
		--tty \
		--publish "4567:4567" \
		--publish "35729:35729" \
		--volume "$(shell pwd)/website:/website" \
		--volume "$(shell pwd):/ext/terraform" \
		--volume "$(WEBSITE_PATH)/content:/terraform-website" \
		--volume "$(WEBSITE_PATH)/content/source/assets:/website/docs/assets" \
		--volume "$(WEBSITE_PATH)/content/source/layouts:/website/docs/layouts" \
		--workdir /terraform-website \
		hashicorp/middleman-hashicorp:${VERSION}

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	$(eval WEBSITE_PATH := $(GOPATH)/src/$(WEBSITE_REPO))
	@echo "==> Testing core website in Docker..."
	-@docker stop "tf-website-core-temp"
	@docker run \
		--detach \
		--rm \
		--name "tf-website-core-temp" \
		--publish "4567:4567" \
		--volume "$(shell pwd)/website:/website" \
		--volume "$(shell pwd):/ext/terraform" \
		--volume "$(WEBSITE_PATH)/content:/terraform-website" \
		--volume "$(WEBSITE_PATH)/content/source/assets:/website/docs/assets" \
		--volume "$(WEBSITE_PATH)/content/source/layouts:/website/docs/layouts" \
		--workdir /terraform-website \
		hashicorp/middleman-hashicorp:${VERSION}
	$(WEBSITE_PATH)/content/scripts/check-links.sh "http://127.0.0.1:4567" "/" "/docs/providers/*"
	@docker stop "tf-website-core-temp"

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: bin cover default dev e2etest fmt fmtcheck generate protobuf plugin-dev quickdev test-compile test testacc testrace vendor-status website website-test