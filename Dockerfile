# Copyright (c) 2022 EPAM Systems, Inc.
# 
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

FROM golang:1.18 as builder

WORKDIR /go/src/dexctl
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/dexctl

FROM gcr.io/distroless/static-debian11

COPY --from=builder /go/bin/dexctl /
ENTRYPOINT ["/dexctl"]
