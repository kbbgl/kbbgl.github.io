# Apache HTTP Server

## Configuration

The location of the Apache configuration files move from distribution to distribution. The name of the primary configuration file changes. The systemd.service name changes. Configuration files may have include files and directories that get merged at server start. The package name usually contains http or apache.

On RedHat, CentOS, or Fedora:
Package: httpd
Service: httpd
Primary configuration file: `/etc/httpd/conf/httpd.conf`

On OpenSUSE:
Package: apache2
Service: apache2
Primary configuration file: `/etc/apache2/httpd.conf`

On Debian, Ubuntu, or Linux Mint:
Package: apache2
Service: apache2
Primary configuration file: `/etc/apache2/apache2.conf`

To allow for modification and flexibility in the apache configuration file, you can include other files and directories. This allows you to avoid one large configuration file and is useful for servers with multiple sites. Many distributions use this feature to enable or disable web server configurations by installing or removing packages.

The OpenSUSE distribution also allows for easy creation of additional include files and directories. To learn more, check out the /etc/sysconfig/apache2 file.

Some of the default include directories are:

CentOS:
`/etc/httpd/conf.d/*.conf`

OpenSUSE:
`/etc/apache2/conf.d/
/etc/apache2/*`

Ubuntu:
`/etc/apache2/*-enabled
/etc/apache2/*-available/`
Ubuntu has active (enabled) and inactive (available) directories. See `man -k a2e` for more details.

Other important files include the document root, log file locations, and module locations (enabled in the configuration file).

The default document root is:

CentOS:
`/var/www/html/`
OpenSUSE:
`/srv/www/htdocs/`
Ubuntu:
`/var/www/html/`
The default log file location is:

CentOS:
`/var/log/httpd/`
OpenSUSE:
`/var/log/apache2/`
Ubuntu:
`/var/log/apache2/`
To load a module, use the following syntax:
`LoadModule alias_module modules/mod_alias.so`

## Log Configuration

Apache has powerful logging features.

To create custom logs on an Apache server, you must first define a custom log format:
LogFormat "example-custom-fmt %h %l %u %t "%r" %>s %b" example-custom-fmt

Below you will see a list of Apache log variables.

Apache Log Variables

VARIABLE EXPLANATION
%h Remote host name
%l Remote login name
​%u Remote user
%t ​Time of request
%r First line of request
​%s Status
%b Size of response

Then you can create a log file, which uses your custom format:

```bash
CustomLog "logs/example-custom.log" example-custom-fmt
```

You can find a reference to all the available tokens in the [Apache Module mod_log_config module](http://httpd.apache.org/docs/2.2/mod/mod_log_config.html#formats).

## Other Configuration Options

The `mod_userdir` module is used to allow all or some users to share a part of their home directory via the web server, even without access to the main document root. The URIs will look something like <http://example.com/~user/index.html> and will commonly be placed in the `/home/user/public_html/` directory.

If a directory does not have an index file (index.html), the autoindex module will generate an index of the files in the directory, similar to what you see from the ls shell command.

## IP/Port Virtual Hosts

For multiple web sites using multiple addresses/ports, use `VirtualHost` stanzas, and a unique IP address and port pair.

To allow Apache to serve different sites on different IP addresses or ports, you should do the following:

Ensure all of the IP addresses and ports are defined in a `Listen` directive.
Add a stanza for each virtual host, as in the example below:

```bash
Listen 192.168.42.11:4374
<VirtualHost 192.168.42.11:4374>
   ServerAdmin webmaster@host1.example.com
    DocumentRoot /www/docs/host1.example.com
    ServerName host1.example.com
    ErrorLog logs/host1.example.com-error_log
    CustomLog logs/host1.example.com-access_log common
</VirtualHost>
```
