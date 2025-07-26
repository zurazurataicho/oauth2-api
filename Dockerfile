FROM golang:alpine as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

# # multi state build
# FROM mysql:8.0
#
# COPY create_db.sql /docker-entrypoint-initdb.d/create_db.sql
# COPY --from=build /usr/local/bin/app /usr/local/bin
#
# ENV MYSQL_ROOT_PASSWORD=root
#
# VOLUME ./db_data:/var/lib/mysql

CMD ["app"]
