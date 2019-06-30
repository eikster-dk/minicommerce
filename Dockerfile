FROM golang:1.12 as build-env

WORKDIR /app
ADD . /app

RUN go mod download && go build

FROM gcr.io/distroless/base
COPY --from=build-env /app/miniCommerce /miniCommerce
COPY --from=build-env /app/sku_DJx1hCHoxDAAtE.pdf /sku_DJx1hCHoxDAAtE.pdf
COPY --from=build-env /app/sku_DWJE6B88Ih3Wgg.pdf /sku_DWJE6B88Ih3Wgg.pdf

CMD ["/miniCommerce"]
