FROM golang:1.25.5 as build

ENV BIN_FILE /opt/subscription/subscription-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/application/main.go

FROM alpine:3.9
LABEL SERVICE="subscription"

ENV BIN_FILE="/opt/subscription/subscription-app"
ENV CONFIG_PATH="/opt/subscription/config.yaml"

COPY --from=build ${BIN_FILE} ${BIN_FILE}
COPY --from=build /go/src/config/config.yaml ${CONFIG_PATH}

CMD ${BIN_FILE}