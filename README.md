# espapi
a Go wrapper for the Getty Images ESP API

# Workflow
```
client = getClient(authParams)
token = client.getToken()
params := RequestParams{"GET", path, token, nil}
request = client.Request(params)
result = client.Response(request)
```
