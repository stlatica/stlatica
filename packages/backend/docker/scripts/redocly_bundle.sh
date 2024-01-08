#!/bin/sh

redocly bundle /openapi/internalapi/openapi.yaml -o /openapi/tmp/internalapi/openapi.yaml
redocly bundle /openapi/externalapi/openapi.yaml -o /openapi/tmp/externalapi/openapi.yaml
