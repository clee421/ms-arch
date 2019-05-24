#!/bin/bash

protoc auth/authpb/auth.proto --go_out=plugins=grpc:.