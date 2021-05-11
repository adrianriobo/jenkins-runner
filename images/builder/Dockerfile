# Build the manager binary
FROM registry.access.redhat.com/ubi8/go-toolset:1.14.12 as builder

USER root
WORKDIR /workspace
COPY . .
RUN make build

FROM registry.access.redhat.com/ubi8/ubi-minimal

LABEL MAINTAINER "Adrian Riobo" "<ariobolo@redhat.com>"

WORKDIR /
COPY --from=builder /workspace/out/jkrunner .
#USER 65532:65532

ENTRYPOINT ["/jkrunner"]