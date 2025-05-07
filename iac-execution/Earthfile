VERSION 0.8

ARG --global TEMPLATE_GIT_PATH="."
ARG --global TEMPLATE_GIT_URI
ARG --global TEMPLATE_GIT_BRANCH

ARG --global GIT_USERNAME
ARG --global GIT_TOKEN

ARG --global ACTION="plan"
ARG --global DRY_RUN="false"

ARG --global ENCODED_VALUE

ARG --global AWS_ACCESS_KEY_ID
ARG --global AWS_SECRET_ACCESS_KEY
ARG --global AWS_SESSION_TOKEN
ARG --global AWS_DEFAULT_REGION=us-east-1

FROM alpine:latest
RUN apk add --no-cache \
  jq \
  yq \
  gettext \
  git

get-terraform-template:
  RUN git clone --branch $TEMPLATE_GIT_BRANCH https://$GIT_USERNAME:$GIT_TOKEN@$TEMPLATE_GIT_URI ./template
  SAVE ARTIFACT ./template

terraform-execute:
  # Use the official Terraform image as the base
  FROM hashicorp/terraform:latest

  # Define a build stage for Terraform execution
  # BUILD terraform-execution

  # Set the working directory inside the container
  WORKDIR /workspace

  # Copy the Terraform configuration files into the container
  COPY +get-terraform-template/template .
  RUN echo $ENCODED_VALUE | base64 -d > ./terraform.tfvars

  # Initialize Terraform
  RUN terraform init

  # Plan the Terraform deployment
  RUN terraform plan

  # Apply the Terraform plan
  RUN terraform $ACTION -auto-approve
