#!/bin/bash

# run container
docker compose up -d 

# get ip address of mysql container
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql   
# my ip in container => 172.26.0.2

# Test curl request
curl -X POST \                                                                              ─╯
     --location "Content-Type: application/json" \
     --data-raw '{
       "name": "Test Admin",
       "email": "testadmin@example.com",
       "password": "securepassword"
     }' \
     http://localhost:8080/api/admin/register | jq
