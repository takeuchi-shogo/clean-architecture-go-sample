# Build stage
FROM golang:1
RUN go install github.com/air-verse/air@latest

WORKDIR /var/www/html

# COPY app/go.mod /var/www/html
# COPY app/go.sum /var/www/html
# RUN go mod download

COPY --from=golang /go /go
COPY --from=golang /usr/local/go /usr/local/go

COPY ./app /var/www/html/
# ./appよりも上位階層に配置する
# そうしなければソースコードを変更しないといけない
COPY ./.env /var/www/
# COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

COPY ./.bashrc /root/.bashrc

CMD [ "bin/dev" ]
