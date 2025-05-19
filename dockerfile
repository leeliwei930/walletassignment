FROM golang:1.23.6-alpine3.21 AS builder

ARG COMMIT_SHA
ARG BUILD_RELEASE_TAG

ENV COMMIT_SHA=${COMMIT_SHA}
ENV BUILD_RELEASE_TAG=${BUILD_RELEASE_TAG}

WORKDIR /app

COPY ../../ .


RUN CGO_ENABLED=0 go build -o /app/build/walletsvc

FROM alpine:3.21
WORKDIR /app


COPY --from=builder /app/locales /app/locales
COPY --from=builder /app/database/migrations /app/database/migrations
COPY --from=builder /app/build/walletsvc /app/walletsvc

RUN apk add --no-cache curl && \
	curl -sSf https://atlasgo.sh | sh && \
	rm -rf /var/cache/apk/*



EXPOSE 8009

# To run server command
# docker run -d -v "./.env:/app/.env" -p "8001:8001" pawswinq-app-backend:v1 /app/app-backend serve
CMD ["/app/walletsvc"]
