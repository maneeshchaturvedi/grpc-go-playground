#!/bin/bash
protoc unary/greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc unary/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.