FROM alpine:latest

# Creates an app directory to hold your app’s source code
WORKDIR /

# Copies everything from your root directory into /app
COPY ./grpc-chat ./

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD [ "/grpc-chat" ]