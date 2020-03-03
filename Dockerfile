FROM hashicorp/terraform:latest AS terraform

FROM quay.io/rebuy/rebuy-go-sdk:v2.0.0 as builder

COPY --from=terraform /bin/terraform /usr/local/bin
