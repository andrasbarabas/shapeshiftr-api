# Stage: Base
FROM golang:1.19-buster as base

LABEL maintainer="András Barabás <barabasandras1@gmail.com>"

ARG PROJECT_NAME=shapeshiftr-api
ENV PROJECT_NAME=${PROJECT_NAME}

ARG APPLICATION_PORT=8080
ENV APPLICATION_PORT=${APPLICATION_PORT}

WORKDIR /go/src/${PROJECT_NAME}

COPY . .

# Stage: Development
FROM base as development

EXPOSE ${APPLICATION_PORT}

CMD ["go", "run", "./cmd/main.go"]

# Stage: Test
FROM base as test

CMD ["go", "test", "-cover", "./..."]

# Stage: Lint
FROM base as lint

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1

CMD ["./bin/golangci-lint", "run", "./..."]

# Stage: Build
FROM base as build

RUN ["sh", "-c", "go build -o \"${PROJECT_NAME}\" -v ./cmd/main.go"]


# Stage: Production
FROM golang:1.19.4-alpine as production

ARG PROJECT_NAME
ARG APPLICATION_PORT

WORKDIR /go/src/${PROJECT_NAME}

RUN apk add --no-cache libc6-compat=1.2.3-r4

COPY --from=build "/go/src/${PROJECT_NAME}/${PROJECT_NAME}" "/go/src/${PROJECT_NAME}/${PROJECT_NAME}"

EXPOSE ${APPLICATION_PORT}

CMD ["sh", "-c", "./${PROJECT_NAME}"]
