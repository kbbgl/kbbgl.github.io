---
title: NGINX Cookbook
slug: nginx-cookbook
authors: [kbbgl]
tags: [nginx,proxy,load_balancing,routing,web,server,docker,devops,containers,images]
---

## Serving Static Content

```nginx
server {
 listen 80 default_server;
 server_name www.example.com;
 location / {
 root /usr/share/nginx/html;
 # alias /usr/share/nginx/html;
 index index.html index.htm;
 }
}
```

This configuration serves static files over HTTP on port 80 from the directory `/usr/share/nginx/html/`.

The `server_name` directive defines the hostname or the names of requests that should be directed to this server. If the configuration had not defined this context as the `default_server`, NGINX would direct requests to this server only if the HTTP host header matched the value provided to the `server_name` directive. With the `default_server` context set, you can omit the `server_name` directive if you do not yet have a domain name to use.

The `location` block defines a configuration based on the path in the URL. The path, or portion of the URL after the domain, is referred to as the uniform resource identifier (URI). NGINX will best match the URI requested to a `location` block. The example uses `/` to match all requests. The root directive shows NGINX where to look for static files when serving content for the given context. The URI of the request
is appended to the root directive’s value when looking for the requested file. If we had provided a URI prefix to the `location` directive, this would be included in the appended path, unless we used the alias directive rather than root. The `location` directive is able to match a wide range of expressions. 

## High-Performance Load Balancing

### HTTP Load Balancing

Distribute load between two or more HTTP servers.

```nginx
upstream backend {
	server 10.10.12.45:80		weight=1;
	server app.example.com:80	weight=2;
	server spare.example.com:80	backup;
}

server {
	location / {
		proxy_pass http://backend;
	}
}
```

This configuration balances load across two HTTP servers on port 80, and defines one as a `backup`, which is used when the primary two are unavailable. The `weight` parameter instructs NGINX to pass twice as many requests to the second server, and the `weight` parameter defaults to 1.


### TCP Load Balancing

```nginx title="/etc/nginx/nginx.conf"
user nginx;
worker_processes auto;
pid /run/nginx.pid;

stream {
	include /etc/nginx/stream.conf.d/*.conf;
}
```


```nginx title="/etc/nginx/stream.conf.d/mysql_reads.conf"
stream {
	upstream mysql_read {
		server 10.10.12.45:3306			weight=5;
		server app.example.com:3306				;
		server spare.example.com:3306	  backup;
	}

	server {
			listen 3306;
			proxy_pass mysql_read;
		}
}
```

The main difference between the `http` and `stream` context is that they operate on different layers of the OSI model. `http` context operates on the application layer (7) and is intended for working with the HTTP protocol whereas `stream` operates at the transport level (4).


### UDP Load Balancing

```nginx title="/etc/nginx/stream.conf.d/ntp.conf"
stream {
	upstream ntp {
		server ntp1.example.com:123		weight=5;
		server ntp2.example.com:123;
	}

	server {
			listen 123 udp;
			proxy_pass ntp;
		}
}
```

If the service over which you’re load balancing requires multiple packets to be sent back and forth between client and server, you can specify the `reuseport` parameter:

```nginx
stream {
	server {
		listen 1195 udp reuseport;
		proxy_pass 127.0.0.1:1194;
	}
}
```

The `reuseport` parameter instructs NGINX to create an individual listening socket for each worker process. This allows the kernel to distribute incoming connections between worker processes to handle multiple packets being sent between client and server.

When working with datagrams, there are some other directives that might apply where they would not in TCP, such as the `proxy_response` directive, which specifies to NGINX how many expected
responses can be sent from the upstream server. By default, this is unlimited until the
`proxy_timeout` limit is reached. The `proxy_timeout` directive sets the time between
two successive read-or-write operations on client or proxied server connections before the connection is closed.

### Load Balancing Methods

NGINX has different methods for load balancing such as Round-robin, least connections, least time, generic hash, random or IP hash.

* **Round Robin** is the default load balancing method which distributes the reqeusts in the order of the list of servers in the upstream pool. We can add the `weight` parameter to distribute requests within the server upstream pool. The higer the `weight`, the more requests will be routed to that server.

* **Least connections** (`last_conn`) balances by proxying the current request to the upstream server with the least number of open connections. We can also use the `weight` parameter.


```nginx
upstream backend {
	least_conn;
	server backend1.example.com;
	server backend2.example.com;
}
```

* **Least time** (`least_time`) only in NGINX Plus taking `header` or `last_byte` parameter.

* **Generic hash** (`hash`) which takes a hash defined by the admin with the given text, variables of the request or runtime, or both. NGINX distributes the load among the servers by producing a hash for the current request and placing it against the upstream servers. This gives more control over where requests are sent.

* **Random** (`random`) Randomly-distribute requests taking `weight` into consideration. It takes an optional `two [method]` parameter which directs NGINX to randomly select two servers and use the load balancing `method`  between the two.

### Health Checks

```nginx
upstream backend {
	server backend1.example.com:1234 max_fails=3 fail_timeout=3s;
	server backend2.example.com:1234 max_fails=3 fail_timeout=3s;
}
```

Passive monitoring watches for failed or timed-out connections as they
pass through NGINX as requested by a client. Passive health checks are enabled by default; the parameters mentioned here allow you to tweak their behavior. The default `max_fails` value is 1, and the default `fail_timeout` value is 10s.

https://docs.nginx.com/nginx/admin-guide/load-balancer/http-health-check
https://docs.nginx.com/nginx/admin-guide/load-balancer/tcp-health-check
https://docs.nginx.com/nginx/admin-guide/load-balancer/udp-health-check


## Traffic Management

NGINX and NGINX Plus are also classified as web-traffic controllers. You can use NGINX to intelligently route traffic and control flow based on many attributes. This chapter covers NGINX’s ability to split client requests based on percentages; utilize the geographical location of the clients; and control the flow of traffic in the form of rate, connection, and bandwidth limiting.

### Split Clients between different Versions (A/B Testing)

```nginx
split_clients "\$\{remote_addr\}AAA" $variant {
	20.0%	"backendv2";
	*		"backendv1";
}
```

The `split_clients` directive hashes the string "\$\{remote_addr\}AAA". The value of `variant` will be "backendv2" 20% of the client IP addresses and the rest "backendv1".

```nginx
http {
	split_clients \${\remote_addr\}" $static_site_root_folder {
		33.3%	"/var/www/sitev2/";
		*		"/var/www/sitev1/";
	}

	server {
		listen 80 _;
		root $static_site_root_folder;
		location / {
			index index.html
		}
	}
}
```


### Geolocation of Clients Physical Location

Install the module:
```bash
apt install nginx-module-geoip
```


Download the city and country databases:
```bash
#!/bin/bash

GEO_IP_DIR=/etc/nginx/geoip
COUNTRY_DB_URI="http://geolite.maxmind.com/download/geoip/database/GeoLiteCountry/GeoIP.dat.gz"
CITY_DB_URI="http://geolite.maxmind.com/download/geoip/database/GeoLiteCity/GeoLiteCity.dat.gz"

mkdir $GEO_IP_DIR&&cd $GEO_IP_DIR

wget $COUNTRY_DB_URI && gunzip GeoIP.dat.gz
wget $CITY_DB_URI && gunzip GeoLiteCity.dat.gz
```

Then load the module and specify the database location:

```nginx
load_module = "/usr/lib64/nginx/modules/ngx_http_geoip_module.so"

http {
	geoip_country /etc/nginx/geoip/GeoIP.dat
	geoip_city /etc/nginx/geoip/GeoLiteCity.dat
}
```

The `$geoip_country_code` (two letters, e.g. IL, US), `$geoip_country_code3` (three letters, e.g. USA), `$geoip_country_name` (full country name) embedded variables are exposed.

The `geoip_city` directive enables all the same variables as the `geoip_country` directive, just with different names, such as `$geoip_city_country_code`, `$geoip_city_country_code3`, and `$geoip_city_country_name`. Other variables include `$geoip_city`, `$geoip_latitude`, `$geoip_longitude`, `$geoip_city_continent_code`, and `$geoip_postal_code`, all of which are descriptive of the value they return. `$geoip_region` and `$geoip_region_name` describe the region, territory, state, province, federal land, and the like. Region is the two-letter code, where region name is the full name. `$geoip_area_code`, only valid in the US, returns the three-digit telephone area code.

http://nginx.org/en/docs/http/ngx_http_geoip_module.html
https://github.com/maxmind/geoipupdate


### Restrict Access Based on Country

```nginx
load_module "/usr/lib64/nginx/modules/ngx_http_geoip_module.so";

http {
	map $geoip_country_code $country_access {
		"US"	0;
		default	1;
	}

	server {
		if ($country_access = '1') {
			return 403;
		}
	}
}
```


### Limiting Connections Based on IP Address

We use the `limit_conn` directive:

```nginx
http {
	limit_conn_zone $binary_remote_addr zone=limitbyaddr:10m;
	limit_conn_status 429;

	server {
		limit_conn liitbyaddr 40;
	}
}
```

This configuration creates a shared memory zone named `limitbyaddr`. The key we use is the client IP address in binary form `$binary_remote_addr`. The size of the shared memory zone is 10MB.

The `limit_conn` directive takes the name of the zone and the nunber of connections allowed (40). 

Be cautious when using the client IP as the key to the zone because if the IP address represents an organization behind a VPN, the whole organization will be limited.

Testing limitations can be tricky. It’s often hard to simulate live traffic in an alternate environment for testing. In this case, you can set the `limit_req_dry_run` directive to on, then use the variable `$limit_req_status` in your access log. The `$limit_req_status` variable will evaluate to either `PASSED`, `DELAYED`, `REJECTED`, `DELAYED_DRY_RUN`, or `REJECTED_DRY_RUN`. With dry run enabled, you’ll be able to analyze the logs of live traffic and tweak your limits as needed before enabling, providing you with assurance that your limit configuration is correct.

### Limiting Rate of Requests by Client IP Address

```nginx
http {
	limit_req_zone $binary_remote_addr zone=limitbyaddr:10m rate=3r/s;
	limit_req_status 429;

	server {
		limit_req zone=limitbyaddr;
	}
}
```

We can also use `burst`, `delay`, `nodelay` to fine tune the rate-limiting.

This is a recommended way to prevent brute-forcing or DDoS.


### Limiting Bandwidth

```nginx
location /download/ {
	limit_rate_after 10m;
	limit_rate 1m;
}
```

URIs with the prefix `download` will be limited after 10MB to a rate of 1MB per second.


## Massively Scalable Content Caching

Scaling and distribution of caching servers in strategic locations should be close to the consumer for the best performance such as CDNs. Wherever the NGINX server is hosted, it will be cached in that location. 


### Caching Zones

Cache the content and define where it is stored.

```nginx
proxy_cache_path 
	/var/nginx/cache
	keys_zone=CACHE:60m
	levels=1:2
	inactive=3h
	max_size=20g;

proxy_cache CACHE;
```

This configuration creates a shared memory space `CACHE` with 60MB of memory and  a directory for cached responses on the filesystem. It defines the maximum size of the cache to 20GB.

It also sets the directory structure level and defines the release of cached responses after they have not been requested in 3 hours. 

The `levels` parameter defines how the file structure is created. The value is a colon- separated value that declares the length of subdirectory names, with a maximum of three levels.

We can also tell NGINX to not proxy requests that are currently being written to cache to an upstream server.

```nginx
proxy_cache_lock on;
proxy_cache_lock_age 10s;
proxy_cache_lock_timeout 3s;
```


### Bypass Cache

We use the `proxy_cache_bypass` directive with a nonempty or nonzero value.

```nginx
proxy_cache_bypass $http_cache_bypass;
```

NGINX will bypass the cache if the HTTP request header named `cache_bypass` is set to any value that is not 0.


### Cache Performance Boost: Client Side

```nginx
location ~* \.(css|js)$ {
	expires 1y;
	add_header Cache-Control "public";
}
```

The client can cache the content of CSS and JavaScript files. The `expires` directive instructs the client that their cached resource will no longer be valid after one year. The `add_header` directive adds the HTTP response header `Cache-Control` to the response and allows any caching server along the way to cache the resource. If we specify `private`, only the client is allowed to cache the value. 

## Programmability

### Expose JavaScript Functionality within NGINX

If we need to perform custom logic on requests and responses, we can use the NJS module.

```bash
apt install nginx-module-njs

mkdir -p /etc/nginx/njs
```

We create a file that will export two functions to NGINX:
```javascript title=/etc/nginx/njs/jwt.js
function jwt(data) {
        var parts = data.split('.').slice(0,2)
            .map(v=>Buffer.from(v, 'base64url').toString())
            .map(JSON.parse);
        return { headers:parts[0], payload: parts[1] };
    }
    function jwt_payload_subject(r) {
		// slice(7) to remove the "Bearer " from the Authorization header
        return jwt(r.headersIn.Authorization.slice(7)).payload.sub;
    }
    function jwt_payload_issuer(r) {
        return jwt(r.headersIn.Authorization.slice(7)).payload.iss;
    }
export default {jwt_payload_subject, jwt_payload_issuer}
```

Then we load the NJS module:
```nginx
load_module /etc/nginx/njs/modules/ngx_http_js_module.so

http {
	js_path "/etc/nginx/njs";
	js_import_main from jwt.js;
	js_set $jwt_payload_subject main.jwt_payload_subject;
	js_set $jwt_payload_issuer main.jwt_payload_issuer;

	server {
		listen 80 default_server;
		listen [::]:80 default_server;
		server_name _;
		location / {
			return 200 "$jwt_payload_subject $jwt_payload_issuer";
		}
	}
}
```

We can then test this out:

```bash
curl 'http://localhost/' -H \
    "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1\
    NiIsImV4cCI6MTU4NDcyMzA4NX0.eyJpc3MiOiJuZ2lueCIsInN1YiI6Im\
    FsaWNlIiwiZm9vIjoxMjMsImJhciI6InFxIiwie\
    nl4IjpmYWxzZX0.Kftl23Rvv9dIso1RuZ8uHaJ83BkKmMtTwch09rJtwgk"

alice nginx
```

https://nginx.org/en/docs/njs/


## Authentication

We can authenticate clients using NGINX.

### HTTP Basic Authentication

We can create an encrypted password using the C `crypt()` function with:

```bash
echo "name1:$(openssl passwd password1)" > conf.d/passwd
```

Then use the `auth_basic` and `auth_basic_user_file` directives to enable basic authentication:

```nginx
location / {
	auth_basic				"Private site";
	auth_basic_user_file	conf.d/passwd;
}
```

A popup with a "Private site" is going to be presented to an unauthenticated user.

We can then test it out:

```bash
curl --user name1:password1 http://localhost
```

Unauthenticated requests will be rejected with 401 and a `WWW-Authenticate` response header. The header will have a value of `Basic realm="your string"` which  causes the browser to prompt for a username and password. The username and password are the concatenated and delimited with a colon, `base64`-encoded and then sent in the request header named `Authorization`. The `Authorization` request header will specify a `Basic` and `user:pass3word` encoded string. The server decodes the header and verifies against the provided `auth_basic_user_file`. 

Use HTTPS with basic authentication.


### Third-Party Authentication System (Subrequest)

We use the `http_auth_request_module` to make a request to the authentication service before serving the request:

```nginx
location /private/ {
	auth_request		/auth;
	auth_request_set	$auth_status $upstream_status;
}

location = /auth {
	internal;
	proxy_pass				http://auth-server;
	proxy_pass_request_body	off;
	proxy_set_header		Content-Length "";
	proxt_set_header		X-Original-URI	$request_uri;
}
```

A request that will arrive in `/private/` will be passed to a subrequest sent to the `auth-server` and it will observer the subrequest response before routing the request to its destination.
If the HTTP request status code is 200, permission will be granted. If the status code is 401 or 403, the same will be returned to the original request. We're dropping the body in the request to the authentication service.

## Security Controls

### Access Based on IP Address

```nginx
location /admin/ {
	deny 10.0.0.1;
	allow 10.0.0.0/20;
	allow 2001:0db8::/32;
	deny all;
}
```

This configuration:
- Allows access from any IPv4 address in 10.0.0.0/20 except for 10.0.0.1.
- Allows access from IPv6 in the 2001:0db8::/32 subnet.
- Returns 403 for requests from any other address.


### Enable CORS

Resource such as JavaScript make CORS when the resource they're requesting is of a domain other than their own. The browser will not use a resource if it does not have appropriate headers that specifically allow it to use it.

We can alter the headers based on the request method:

```nginx
map $request_method $cors_method {
	OPTIONS 11;
	GET 	1;
	POST	1;
	default	0;
}

server {
	location / {
		if($cors_method ~ '1') {
			add_header 'Access-Control-Allow-Methods' 'GET,POST,OPTIONS';
			add_header 'Access-Control-Allow-Origin' '*.example.com';
			add_header 'Access-Control-Allow-Headers'
				'DNT
				,Keep-Alive,
				,User-Agent
				,X-Requested-With
				,If-Modified-Since
				,Cache-Control
				,ContentType'
				;
		}

		if($cors_method = '11') {
			add_header 'Access-Control-Max-Age', 1728000;
			add_header 'Content-Type', 'text/plain; charset=UTF-8';
			add_header 'Content-Length' 0;
			return 204;
		}
	}
}
```

The `OPTIONS` request method returns a preflight request to ask for the CORS rules enforced by the server. `OPTIONS, GET, POST` are allowed under CORS. Setting the `Access-Control-Allow-Origin` header allows for content being served from this server to also be used on pages of origins that match this header. The preflight request can be cached on the client for 1,728,000 seconds, or 20 days.


### Client-Side Encryption

Encrypt traffic between NGINX server and the client. We can utilize the `ngx_http_ssl_module` and `ngx_stream_ssl_module`.

```nginx
http {
	server {
		listen 8443 ssl;
		ssl_certificate /etc/nginx/ssl/example.crt;
		ssl_certificate_key /etc/nginx/ssl/example.key;

		# Set accepted protocol and cipher
		ssl_protocols TLSv1.2 TLSv1.3;
		ssl_ciphers HIGH:!aNULL:!MD5;

		# Client-Server negotiation caching
        ssl_session_cache shared:SSL:10m;
        ssl_session_timeout 10m;
	}
}
```

### Upstream Encryption

We can proxy over HTTPS by changing the protocol on the value passed to the `proxy_pass` directive:

```nginx
location / {
	proxy_pass https://upstream.example.com;
	proxy_ssl_verify on;
	proxy_ssl_verify_depth 2;
	proxy_ssl_protocols TLSv1.2;
}
```

### Securing a Location

Securing resources with a secret is a great way to ensure your files are protected. The argument to `secure_link_secret` is `md5` hashed and the hex digest of that `md5` is used in the URI. 

```nginx
location /resources {
	secure_link_secret mySecret;
	if ($secure_link = "") { return 403; }

	rewrite ^ /secured/$secure_link;
}

location /secured/ {
	internal;
	root /var/www;
}
```

The configuration creates an internal and public facing location block. The public-facing one will return 403 if the secure link is empty. The secure link is empty when the `md5` hash was not verified.


### Generating a Secure Link with a Secret

Let's say we have a file `/var/www/secured/index.html` and the same configuration from [Securing a Location](#securing-a-location). We can generate a secure link by concatenating the URI path and the secret.

```bash
echo -n 'index.htmlmySecret' | openssl md5 -hex
a53bee08a4bf0bbea978ddf736363a12
```

We can do the same in Python:
```python
import hashlib
hashlib.md5(b'index.htmlmySecret').hexdigest()
'a53bee08a4bf0bbea978ddf736363a12'
```

We can access the resource by constructing the request:

```bash
curl http://www.example.com/resources/a53bee08a4bf0bbea978ddf736363a12/index.html
```


### Securing a Location with Expire Date

We can secure a location with a link that expires at some future time and specific to a client.
This is a very secure option as it prevents sharing of links to access secured locations.

```nginx
location /resources {
	root /var/www;
	secure_link $arg_md5,$arg_expires;
	secure_link_md5 "$secure_link_expires$uri$remote_addrmySecret";
	if ($secure_link = "") { return 403;}
	if ($secure_link = "0") { return 410;}
}
```

`secure_link` takes the variable holding the `md5` hash (which in this example is an HTTP argument of `md5`) ad a variable that holds the time in which the link expires in Unix epoch time format. If the secure link equals 0, it means that the link expired and a 410 Gone is returned.

### Generating an Expiring Link

In the example below, we're able to generate a secure link in a special format that can be used in URLs. It's secure because the value of the variable is never sent to the client.

This script generates a secure link that expires after 1 hour:
```python title="securelink.py"
from datetime import datetime, timedelta from base64 import b64encode
import hashlib
    # Set environment vars
    resource = b'/resources/index.html'
    remote_addr = b'127.0.0.1'
    host = b'www.example.com'
    mysecret = b'mySecret'
    # Generate expire timestamp
    now = datetime.utcnow()
    expire_dt = now + timedelta(hours=1)
    expire_epoch = str.encode(expire_dt.strftime('%s'))
    # md5 hash the string
    uncoded = expire_epoch + resource + remote_addr + mysecret
    md5hashed = hashlib.md5(uncoded).digest()
    # Base64 encode and transform the string
    b64 = b64encode(md5hashed)
    unpadded_b64url = b64.replace(b'+', b'-')\
        .replace(b'/', b'_')\
        .replace(b'=', b'')
    # Format and generate the link
    linkformat = "{}{}?md5={}?expires={}"
    securelink = linkformat.format(
        host.decode(),
        resource.decode(),
        unpadded_b64url.decode(),
        expire_epoch.decode()
) print(securelink)
```

Another way is by generating a Unix epoch expiration date:

```bash
date -d "2030-12-31 00:00" +%s --utc
1924905600
```

If we use the value of `secure_link_md5` example provided in [Securing a Location with Expire Date](#securing-a-location-with-expire-date), we can construct the link as so:

```bash
echo -n '1924905600/resources/index.html127.0.0.1 mySecret' \
| openssl md5 -binary \
| openssl base64 \
| tr +/ -_ \
| tr -d =

sqysOw5kMvQBL3j9ODCyoQ
```

We can then access the resource:

```bash
curl "http://example.com/resources/index.html?md5=sqysOw5kMvQBL3j9ODCyoQ&expires=1924905600"
```


### Redirect HTTPS

If we need to redirect unencrypted requests to HTTPS:

```nginx
server {
	listen 80 default_server;
	listen [::]:80 default_server;
	server_name _;
	return 301 https://$host$request_uri;
}
```

### HTTP Strict Transport Security (HSTS)

Instruct browsers to never send requests over HTTP.

```nginx
add_header Strict-Transport-Security max-age=31536000
```

## HTTP/2

To use HTTP/2, we add the `http2` argument:

```nginx
server {
	listen 443 ssl http2 default_server;

	ssl_certificate 	server.crt;
	ssl_certificate_key server.key;
}
```


### gRPC

We can listen on HTTP/2 traffic and proxy that traffic to a machine 

```nginx
server {
	listen 80 http2;

	location / {
		grpc_pass	grpc://backend.local:50051;
	}
}
```

For TLS encryption that terminates at NGINX and passes the gPRC communication to the application over unencrypted HTTP/2:

```nginx
server {
	listen 443 ssl http2 default_server;

	ssl_certificate		server.crt;
	ssl_certificate_key	server.key;
	location / {
		grpc_pass grpc://backend.local:50051;
		# for end-to-end encryption
		# grpc_pass grpcs://backend.local:50051;
	}
}
```

We can also route the calls to different backend services:

```nginx
location /mypackage.service1 {
	grpc_pass	grpc://$grpc_service1;
}

location /mypackage.service2 {
	grpc_pass	grpc://$grpc_service2;
}

location / {
	root	/usr/share/nginx/html;
	index	index.html	index.htm;
}
```

### HTTP/2 Server Push

To preemtively push content to the client:

```nginx
server {
	listen 443 ssl http2 default_server;

	ssl_certificate		server.crt;
	ssl_certificate_key	server.key;
	root /usr/share/nginx/html;

	location /demo.html {
		http2_push	/style.css;
		http2_push	/image1.jpg;
	}
}
```
NGINX can automatically push resources to clients if proxied applications include an HTTP response header named `Link`. To enable this feature, we add `http2_push_preload on;` to the configuration.

## Containers/Microservices


### Using NGINX as an API Gateway

https://oreil.ly/75l-m

```nginx title=/etc/nginx/api_gateway.conf
server {
	listen 443 ssl;
	server_name api.company.com;
	default_type application/json

	# Error handling
	proxy_intercept_errors on;

	error_page 400 = @400;
	location @400 { return 400 '{"status": 400, "message": "Bad request"}\n;'}

	error_page 401 = @401;
	location @401 { return 401 '{"status": 401, "message": "Bad Unauthorized"}\n;'}

	error_page 403 = @403;
	location @403 { return 403 '{"status": 403, "message": "Forbidden"}\n;'}

	error_page 404 = @404;
	location @404 { return 404 '{"status": 404, "message": "Resource not found"}\n;'}
}
```

We can then import this into the main configuration:

```nginx title=/etc/nginx/nginx.conf
include /etc/nginx/api_gateway.conf
```

Now we can define the upstream service endpoints.

```nginx title=/etc/nginx/nginx.conf
upstream service_1 {
	server 10.0.0.12:80;
	server 10.0.0.13:80;
}

upstream service_2 {
	server 10.0.0.14:80;
	server 10.0.0.15:80;
}
```

In case services should also be defined as proxy location endpoints, we can define an endpoint as a variable:

```nginx
location = /_service_1 {
	internal;
	# Config common to service
	proxy_pass http://service_1/$request_uri;

}

location = /_service_2 {
	internal;
	# Config common to service
	proxy_pass http://service_2/$request_uri;

}
```

Then we need to build up `location` blocks that define specific URI paths for a given service.

```bash
mkdir /etc/nginx/api_conf.d
touch /etc/nginx/api_conf.d/service_1.conf
```

```nginx title=/etc/nginx/api_conf.d/service_1.conf
location /api/service_1/object {
	limit_except GET PUT { deny all; }
	rewrite ^ /_service_1 last;
}

location /api/service_1/object/[^/]*$ {
	limit_except GET POST { deny all; }
	rewrite ^ /_service_1 last;
}
```

The `rewrite` directive directs the request to the prior configured `location` block that proxies the request to a service. In the example above, `rewrite` instructs NGINX to reprocess the request with an altered URI. It specifies the expected HTTP methods and reroutes the request.

We do the above for all services and make sure to include them configuration files in `nginx.conf`:

```nginx title=/etc/nginx/nginx.conf
server {
	listen 443 ssl;
	server_name api.company.com

	default_type application/json;

	include api_conf.d/*.conf;
}
```

Then make sure to enable authentication such as simple preshared API keys:

```nginx
map $http_apikey $api_client_name {
	default "";

	"1234jdankfjna" "client_one";
	"5678jdankfjna" "client_two";
	"9012jdankfjna" "client_three";
}
```

Also, protect the backend services from attacks:

```nginx
http {
	limit_req_zone $http_apikey zone=limitbyapikey rate=100r/s;
	limit_req_status	429;

	location /api/service_2/object {
		limit_req zone=lomitbyapikey;

		if ($http_apikey = "") {
			return 401;
		}

		if ($api_client_name = "") {
			return 403;
		}

		limit_except GET PUT {deny_all;}
		rewrite ^ /_service_2 last;
	}
}
```

### Creating an NGINX Dockerfile

```
 .
    ├── Dockerfile
    └── nginx-conf
        ├── conf.d
        │   └── default.conf
        ├── fastcgi.conf
        ├── fastcgi_params
        ├── koi-utf
        ├── koi-win
        ├── mime.types
        ├── nginx.conf
        ├── scgi_params
        ├── uwsgi_params
        └── win-utf
```

```dockerfile
FROM centos:7

RUN yum -y install epel-release && yum -y install nginx

ADD /nginx-conf	/etc/nginx

EXPOSE 80 443

CMD ["nginx"]
```

Since we need to run NGINX in the foreground inside a container, we start it using
`-g "daemon off;"` or add `daemon off;` to NGINX configuration.
We also need to alter the NGINX configuration to log to `/dev/std[out|err]` so that the container logs will be routed to the Docker daemon.


### Using Environment Variables in NGINX

Use the `ngx_http_perl_module` to set variables in NGINX from the environment:

```nginx
daemon off;
env APP_DNS;
include /usr/share/nginx/modules/*.conf

http {
	perl_set $upstream_app 'sub { return $ENV{"APP_DNS"}; }';
	server {
		location / {
			proxy_pass https://$upstream_app;
		}
	}
}
```

To use `perl_set`, we must have the `ngx_http_perl_module` installed.

```bash
yum -y install nginx nginx-mod-http-perl
```

When installing modules from the package utility for CentOS, they’re placed in the `/usr/lib64/nginx/modules/` directory, and configuration files that dynamically load these modules are placed in the `/usr/share/nginx/modules/` directory.

By default, NGINX gets rid of env vars so we need to specifiy them explicitly using `env`. 
The `perl_set` directive takes the variable name and the string that renders the result as arguments.

### Kubernetes Ingress Controller



We can use the `nginx/nginx-ingress` image from Dockerhub. 

To set it up, we use the [kubernetes-ingress repository on GitHub](https://oreil.ly/KxF7i).


We first create a `Namespace` and a `ServiceAccount` for the ingress controller.
```bash
kubectl apply -f common/ns-and-sa.yaml
```

We then create a `Secret` with TLS certificate and key:
```bash
kubectl apply -f common/default-server-secret.yaml
```

We can also create a `ConfigMap` for customizing configuration.

```bash
kubectl apply -f common/nginx-config.yaml
```

If RBAC is enabled in the cluster, we need to create a `ClusterRole` and bind it to the `ServiceAccount`:

```bash
kubectl apply -f rbac/rbac.yaml
```

We can then choose to create a `Deployment` or `DaemonSet`. Use the `Deployment` if there are plans to dynamically change the number of ingress controller replicas or use `DaemonSet` to deploy an ingress controller on every `Node` or subset of `Node`s.

```bash
# FOR DEPLOYMENT
kubectl apply -f deployment/nginx-ingress.yaml

# FOR DAEMONSET
kubectl apply -f daemon-set/nginx-ingress.yaml
```

If a `DaemonSet` was created, ports 80 and 443are mapped to the same ports on the `Node` where the container is running.

For a `Deployment`, there are 2 options for accessing the ingress controller `Pod`s. Either we instruct Kubernetes to randomly assign a `Node` port that maps to the ingress controller `Pod` using a `Service` with type `NodePort`:

```bash
kubectl create -f service/nodeport.yaml
```

To statically configure the port that is opened for the `Pod`, alter the YAML and add the attribute:

```yaml
nodePort: {port}
```

Or create a `LoadBalancer`-type `Service`:

```bash
kubectl create -f service/loadbalancer.yaml

# for AWS/ELB
kubectl create -f service/loadbalancer-aws-elb.yaml
```

For AWS, Kubernetes creates a classic ELB in TCP mode with the PROXY Protocol enabled. We must configure NGINX to use the PROXY Protocol by adding the following to the `ConfigMap`:

```yaml title=common/nginx-config.yaml
proxy-protocol: "True"
real-ip-header: "proxy_protocol"
set-real-ip-from: "0.0.0.0/0"
```

And update the `ConfigMap`:
```bash
kubectl apply -f common/nginx-config.yaml
```


### Prometheus Exporter Module

We can configure Prometheus to collect NGINX statistics for monitoring purposes.

There's an NGINX Prometheus Exporter Module and can be found in [Docker Image on Docker Hub](https://oreil.ly/mC_i9).

The exported will be started for NGINX and only harvest the `stub_status` info. We must ensure the stub status is enabled.

To enable stub status:

```nginx
location /stub_status {
	stub_status;
	allow 127.0.0.1;
	deny all;
}
```

Then we can run:

```bash
docker run -p 9113:9113 nginx/nginx-prometheus-exporter:0.8.0 -nginx.scrape-uri http://${nginxEndpoint}:8080/stub_status
```

The `stub_status` module enables some basic monitoring such as number of active connections (`$connections_active`), accepted connections, connections handled and requests served. Also, the current number of connections being read (`$connections_reading`), written (`$connections_writing`) or in waiting (`$connections_waiting`).


## Debugging and Troubleshooting

NGINX allows us to divide access logs into different files and formats for different contexts and to change the log level of error logging to get a deep understanding of what's happening. Logs can be streamed to Syslog.

### Configuring Access Logs

```nginx
http {
	log_format	geoproxy
				'[$time_local] $remote_addr '
				'$realip_remote_addr $remote_user '
				'$proxy_protocol_server_addr $proxy_protocol_server_port '
				'$request_method $server_protocol '
				'$scheme $server_name $uri $status '
				'$request_time $body_bytes_sent '
				'$geoip_city_country_code3 $geoip_region '
				'"$geoip_city" $http_x_forwarded_for '
				'$upstream_status $upstream_response_time '
				'"$http_referer" "$http_user_agent"';
}
```

The log format configuration is named `geoproxy` 

We can set an optional `escape` parameter to `log_format` such as `default,json,none` to escape prohibited characters.

To use this log format:
```nginx
server {
	access_log /var/log/nginx/access.log geoproxy;
}
```

### Configuring Error Logs

```nginx
error_log /var/log/nginx/error.log warn;
```

The log level is optional (`warn` in this case) and is `error` by default. It can be set to `debug/info/notice/warn/error/crit/alert/emerg`.

This log will include information about configuration files not working correctly and errors produced by the application servers.


### Forwarding to Syslog

```nginx
error_log syslog:server=10.0.1.42 debug;
access_log syslog:server=10.0.1.42,tag=nginx,severity=info geoproxy;
```

### Request Tracing

We can use NGINX logs to have an end-to-end understanding of a request. `$request_id` provides a randomly generated UUIDv4 string to help in correlate a request across the upstream servers handling the request.

```nginx
log_format trace '$remote_addr - $remote_user [$time_local] '
				 '"$request" $status $body_bytes_sent '
				 '"$http_referer" "$http_user_agent" '
				 '"$http_x_forwarded_for" $request_id';

upstream backend {
	server 10.0.0.42;
}

server {
	listen 80;

	# Add the X-Request-ID header to the response to client
	add_header X-Request-ID $request_id;

	location / {
		proxy_pass http://backend;

		# Send the header to the application
		proxy_set_header X-Request-ID $request_id;
		access_log /var/log/nginx/access_trace.log trace;
	}
}
```

In the frontend client, the request will include the header. We'll need to capture this header value and add it to the application logs.
