# `curl` Tricks

## Globbing

```bash
curl -s "https://jsonplaceholder.typicode.com/users/[1-3]" | jq -s .
curl -s "https://jsonplaceholder.typicode.com/users/[0-10:2]" | jq -s .

curl -s "https://jsonplaceholder.typicode.com/photos/{1,6,35}" | jq -s .

curl -s "https://jsonplaceholder.typicode.com/users/[1-3]" -o "file_#1.json"
```

## Configuration 

```bash
cat ~/.curlrc

# some headers
-H "Upgrade-Insecure-Requests: 1"
-H "Accept-Language: en-US,en;q=0.8"

# follow redirects
--location
```

Specify configuration file:
```bash
curl -K .curlrc https://google.com
```

For authentication:
```bash
cat ~/.netrc
machine https://authenticationtest.com/HTTPAuth/
login user
password pass
```

## Parallel Requests

```bash
curl -I --parallel --parallel-immediate --parallel-max 3 --config websites.txt

curl -I --parallel --parallel-immediate --parallel-max 3 stackoverflow.com google.com example.com
```