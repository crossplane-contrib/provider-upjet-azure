FROM alpine:3.22.1
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH
ARG CROSSPLANE_PROVIDER_VERSION

ENV USER_ID=65532
ENV ARM_PARTNER_ID a9cee75d-8f11-42e4-bc19-953757f4ea3c
ENV ARM_DISABLE_TERRAFORM_PARTNER_ID true
ENV AZURE_HTTP_USER_AGENT crossplane-provider-upjet-azure/${CROSSPLANE_PROVIDER_VERSION} (${TARGETARCH}-${TARGETOS})

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]