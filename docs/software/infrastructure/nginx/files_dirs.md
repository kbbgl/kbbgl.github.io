---
title: NGINX Files and Directories
slug: nginx-files-dirs
authors: [kbbgl]
tags: [nginx]
---

###  `/etc/nginx`

Default configuration root for the NGINX server.

* `/etc/nginx/nginx.conf` - Default configuration entry point used by the NGINX service. Sets up worker processes, tuning, logging, loading dynamic modules and references to other NIGNX configuration files.

* `/etc/nginx/conf.d/` - Directory that contains the default HTTP server configuration file. Files in this directory ending in `.conf` are included in the top-level `http` block from within the `/etc/nginx/nginx.conf` file. The best practice is to utulize `include` statements and organize your configuration in this way to keep your configuration files concise.

* `/var/log/nginx` - The `/var/log/nginx` directory is the default log location for NGINX and includes an `access.log` containing an entry for each request served by NGINX and an `error.log`.

