# Build go project
FROM golang:1.18 as go-build

# Add everything
ADD . /usr/src/multi-networkpolicy-tc

RUN cd /usr/src/multi-networkpolicy-tc && \
    go build ./cmd/multi-networkpolicy-tc/

# Build iproute2
FROM quay.io/centos/centos:stream8 as iproute-build

ARG IPROUTE2_TAG=v5.17.0

RUN dnf -q -y groupinstall "Development Tools"
RUN dnf -q -y install git libmnl-devel

RUN git clone --branch ${IPROUTE2_TAG} https://github.com/shemminger/iproute2.git && \
    cd /iproute2 && make && make install

# collect everything into target container
# TODO(adrianc): once we switch to native netlink to drive TC  we can switch back to distroless
#FROM gcr.io/distroless/base
FROM quay.io/centos/centos:stream8
LABEL io.k8s.display-name="Multus NetworkPolicy TC" \
      io.k8s.description="This is a component provides NetworkPolicy objects for secondary interfaces created with Multus CNI"
RUN dnf -q -y install git libmnl
COPY --from=go-build /usr/src/multi-networkpolicy-tc/multi-networkpolicy-tc /usr/bin
COPY --from=iproute-build /sbin/tc /sbin
WORKDIR /usr/bin

ENTRYPOINT ["multi-networkpolicy-tc"]
