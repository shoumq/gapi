#!/bin/bash

migrate -path ./migrations -database "postgres://postgres:Andrew1095@localhost:5432/gapi?sslmode=disable" "$@"