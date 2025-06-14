#! /usr/bin/bash

# # Start the Go application in the background
go run . &
GO_PID=$!  # Capture the process ID (PID) of the Go program

# # Wait briefly for the server to start
sleep 2  

# Perform API requests
echo "Expect ID 1"
echo
curl localhost:8080/stashitems/1 -v | jq
echo

# curl -X POST -H "Content-Type: application/json" -d '{"name": "baz"}' localhost:8080/stashitems -v

echo "Expect ID 3"
echo
curl localhost:8080/stashitems/3 -v | jq
echo

# curl -X PUT "http://localhost:8080/stashitems/1" -H "Content-Type: application/json" -d '{"id":1, "name":"Updated Item"}'
# curl localhost:8080/stashitems/2 -v | jq

echo "Expect All IDs"
echo
curl localhost:8080/stashitems -v | jq
echo

# curl -X DELETE "http://localhost:8080/stashitems/1"
# curl localhost:8080/stashitems -v | jq


# # Stop the Go application
kill $GO_PID