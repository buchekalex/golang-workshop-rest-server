#!/bin/bash

# Use LC_ALL=C to handle non-ASCII characters
export LC_ALL=C

# Make exactly 1000 requests with curl
for i in $(seq 1 100); do
    # Generate random user data
    id=$RANDOM
    name=$(cat /dev/urandom | tr -dc 'a-zA-Z' | fold -w 10 | head -n 1)
    email="${name}@example.com"
    password=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)

    # Create JSON payload
    json=$(jq -n --arg id "$id" --arg name "$name" --arg email "$email" --arg password "$password" \
        '{id: $id, name: $name, email: $email, password: $password}')

    # Send POST request
    curl -s -X POST -H "Content-Type: application/json" -d "${json}" http://localhost:8088/app/users

    # Save email to file
    echo "http://localhost:8088/app/users/${email}" >> ./urls.txt
done

siege -f ./urls.txt -c 20
