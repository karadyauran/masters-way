# Stage 1: build
FROM golang:1.23-alpine AS build
WORKDIR /app
RUN apk add --no-cache make
COPY . .
RUN go mod download
RUN make build

# Stage 2: production
FROM golang:1.23-alpine
WORKDIR /app
COPY --from=build /app/main ./
COPY --from=build /app/entrypoint.sh ./
COPY --from=build /app/.env ./
# required for reset-db
COPY --from=build /app/internal/db/migration ./internal/db/migration
EXPOSE 8001
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]