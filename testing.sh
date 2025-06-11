#! /usr/bin/bash

# # Start the Go application in the background
go run . &
GO_PID=$!  # Capture the process ID (PID) of the Go program

# # Wait briefly for the server to start
sleep 2  

# Perform API requests
curl localhost:8080/stashitems/0 -v | jq
curl -X POST -H "Content-Type: application/json" -d '{"name": "baz"}' localhost:8080/stashitems -v
curl localhost:8080/stashitems/1 -v | jq
curl -X PUT "http://localhost:8080/stashitems/0" -H "Content-Type: application/json" -d '{"id":1, "name":"Updated Item"}'
curl localhost:8080/stashitems/0 -v | jq
curl localhost:8080/stashitems -v | jq
curl -X DELETE "http://localhost:8080/stashitems/0"
curl localhost:8080/stashitems -v | jq


# # Stop the Go application
kill $GO_PID