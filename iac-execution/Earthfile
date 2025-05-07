VERSION 0.8

ARG --global PROJECT_NAME="default"
ARG --global GROUP_NAME="default"
ARG --global BLUEPRINT_NAME

ARG --global RESOURCE_UUID
ARG --global RESOURCE_TYPE
ARG --global RESOURCE_VALUE
ARG --global RESOURCE_NAME

ARG --global TEMPLATE_PATH="."
ARG --global TEMPLATE_GIT_URI
ARG --global TEMPLATE_GIT_BRANCH

ARG --global GIT_USERNAME
ARG --global GIT_TOKEN

ARG --global ACTION="plan"
ARG --global DRY_RUN="false"

FROM alpine:latest
RUN apk add --no-cache \
  jq \
  yq \
  gettext \
  git

generate-vars:
  RUN cat ./$PROJECT_NAME/$GROUP_NAME/$BLUEPRINT_NAME/RESOURCE_UUID

get-terraform-template:
  RUN git clone --branch $TEMPLATE_GIT_BRANCH https://$GIT_USERNAME:$GIT_TOKEN@$TEMPLATE_GIT_URI

terraform-apply:
  # Use the official Terraform image as the base
  FROM hashicorp/terraform:latest

  # Define a build stage for Terraform execution
  # BUILD terraform-execution

  # Set the working directory inside the container
  WORKDIR /workspace

  # Copy the Terraform configuration files into the container
  COPY . .

  # Initialize Terraform
  RUN terraform init

  # Validate the Terraform configuration
  RUN terraform validate

  # Plan the Terraform deployment
  RUN terraform plan -out=tfplan

  # Apply the Terraform plan
  RUN terraform apply -auto-approve tfplan
