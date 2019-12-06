#!make
SHELL := /bin/bash
.SHELLFLAGS := -ec

## General VARS
VERSION ?= dev
APPLICATION_NAME= mass-lobby
AWS_PROFILE=enekofb
AWS_ACCOUNT_ID=777171359344

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

builder: build-builder push-builder	

unit:
	@echo "Running unit tests"

docker-build:
	@echo "building docker for mass"

docker-push:
	@echo "pushing docker for mass"
	
build: unit docker-build docker-push
	@echo "Running build"

acceptance:
	@echo "Running acceptance tests"

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
