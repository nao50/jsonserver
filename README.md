# jsonserver
Json server in go

## Authentication
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"email":"admin@test.com", "password":"password"}' localhost:8080/auth

# Response
{"X-Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs"}
```

---

## Json Endpoint
### POST
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"jsonName":"hogeJson"}' \
localhost:8080/jsons
```
### GET ALL
```bash
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/jsons
```
### GET
```bash
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/jsons/1
```
### PUT
```bash
$ curl -X PUT -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"jsonName":"piyoJson"}' \
localhost:8080/jsons/1
```
### DELETE
```bash
$ curl -X DELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/jsons/2
```

---

## NestedJson Endpoint
### POST
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"nestedJsonName": "nestedHogeJson","configuration":{"configurationName": "nestedHogeConfiguration","configuration":"High"}}' \
localhost:8080/nestedjsons
```
### GET ALL
```bash
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/nestedjsons
```
### GET
```bash
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/nestedjsons/1
```
### PUT
```bash
$ curl -X PUT -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"nestedJsonName": "nestedPiyoJson","configuration":{"configurationName": "nestedPiyoConfiguration","configuration":"Low"}}' \
localhost:8080/nestedjsons/1
```
### DELETE
```bash
$ curl -X DELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/nestedjsons/2
```

---

## ListInJson Endpoint
### POST
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"jsonName":"hogeJson"}' \
localhost:8080/listinjsons
```
### GET ALL
```bash
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/listinjsons
```
### PUT
```bash
$ curl -X PUT -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"jsonName":"piyoJson"}' \
localhost:8080/listinjsons/1
```
### DELETE
```bash
$ curl -X DELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/listinjsons/2
```

---

## NilJson Endpoint
### POST FULL
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"pureString":"hello", "pointerString":"hello", "omitemptyString":"hello"}' \
localhost:8080/niljsons
```
### POST 1
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"pureString":"hello"}' \
localhost:8080/niljsons
```
### POST 2
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"pointerString":"hello"}' \
localhost:8080/niljsons
```
### POST 3
```bash
$ curl -X POST \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
-H 'Content-Type: application/json' \
-d '{"omitemptyString":"hello"}' \
localhost:8080/niljsons
```
### GET ALL
```bash
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjkxOTY0MjEsIm5hbWUiOiJhZG1pbiJ9.5nkp-Xr2r5rCtp-5z1DL29OMUhXRFjkpdFRO0K8mWNs" \
localhost:8080/niljsons
```
