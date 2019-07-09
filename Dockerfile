# We specify the base image we need for our
# go application
FROM golang:1.12.0-alpine3.9
# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD http.exe /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/http.exe"]