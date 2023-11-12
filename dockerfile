FROM golang:1.21-alpine
COPY .build/linux/* ./
EXPOSE 8080
CMD "./messenger_v2"