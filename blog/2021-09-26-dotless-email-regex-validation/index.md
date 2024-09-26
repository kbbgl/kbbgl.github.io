---
slug: dotless-email-regex-validation-swagger
title: How To Set Up Dotless Email Validation in Swagger
description: A detailed explanation what dotless emails are and how to set them up in Swagger.
authors: [kbbgl]
tags: [email, regex, web-dev, swagger]
image: ./swagger-logo.png
---

## Introduction

One of our customers recently attempted (for reason unknown to us) to log into our platform using [Okta Single Sign On (SSO)](https://www.okta.com/products/single-sign-on/) and [OpenID-Connect](https://openid.net/connect/) with emails missing a top-level domain (TLD) while they could log in just fine with a standard email address.

While most of us in the day-to-day use the standard email address format with a top-level domain added, e.g. `user@domain.org`, this customer required to be able to log in with user@domain. Our platform rejected the user and they were unable to log in.

My first question was whether this email address (`user@domain`) was even allowed. If it was, I needed to find why it's being rejected.

## Email Address Syntax Compliance

Reaching out to the web to figure out whether the customer's request was valid, I was surprised to find that the only necessary parts of an email address are according to [section 2.3.10 of the RFC 2821](https://datatracker.ietf.org/doc/html/rfc2821#section-2.3.10):

- A 'local' part, anything on the left of the @ sign.
- A 'domain host' part, anything on the right of the @ sign.

I found that the [Internet Corporation for Assigned Names (ICANN) highly discourages the use of unspecified TLD](https://www.icann.org/en/system/files/files/ua-factsheet-a4-17dec15-en.pdf) and that it [requires a specified record in the apex of the TLD zone in the DNS](https://www.icann.org/en/announcements/details/new-gtld-dotless-domain-names-prohibited-30-8-2013-en). Aside from that, it is compliant with [RFC 5322](https://datatracker.ietf.org/doc/html/rfc5322#section-3).

So the customer's request is a valid one! This meant that I needed to perform 2 changes in two different areas:

- Modify the name server configuration on the customer's Windows instance.

- Find where within our code there's validation for the email address and modify it.

## Adding TLD Zone Record to DNS

After a few days of awaiting access to the DNS, I now needed to understand what record I needed to add to the DNS. The [ICANN link above](https://www.icann.org/en/announcements/details/new-gtld-dotless-domain-names-prohibited-30-8-2013-en) suggested that:

> Dotless names would require the inclusion of, for example, an A, AAAA, or MX, record in the apex of a TLD zone in the DNS (i.e., the record relates to the TLD-string itself).

I found out that the [A record indicates the IPv4 address](https://www.cloudflare.com/learning/dns/dns-records/dns-a-record/) of a domain and that the the [AAAA record was the same but for IPv6](https://il.godaddy.com/en/help/add-an-aaaa-record-19214). The relevant record seemed to be [MX as it directs emails to SMTP servers](https://www.cloudflare.com/learning/dns/dns-records/dns-mx-record/).

Now that I knew that I needed to add an MX record, I found a [PowerShell cmdlet called `Add-DnsServerResourceRecordMX`](https://docs.microsoft.com/en-us/powershell/module/dnsserver/add-dnsserverresourcerecordmx?view=windowsserver2019-ps) which would allow me to do just that.

After reading through the documentation and doing some trial and error on my test machine, I formulated the following command to successfully add the MX record:

```powershell
Add-DnsServerResourceRecordMX -Preference 1 -Name “.” -MailExchange “GS-MEX001.domain.suffix” -ZoneName “domain.suffix”
```

Here's an explanation of the cmdlet argments:

- I set `-Preference` to `1` ensure that all SMTP requests sent from this Windows instance to mailservers will go through this `MX` record.

- `-Name` specifies the name of the host for which this record will apply. In this case, we used `.` because we're going to be using the parent domain as the suffix instead of the TLD.

- `-MailExchange` specifes the Fully Qualified Domain Name (FQDN) for the mail exchanger, e.g. the SMTP server address. Obviously the one used in the command above is made up.

- `-ZoneName` specifies the DNS zone. DNS is split into different zones which allow administrators to have granular control of different DNS components such as certificates, authoritative nameservers and providers.

After running this command on the production server, I saw that they were successfully added by running the following cmdlet:

```powershell
Get-DnsServerSetting -All
```

So now, if someone sends an email to the address `@domain.suffix`, the sender email server will review and choose the `MX` record we created above and will direct the packet to it.

Our work with nameservers is done! Next step is to look at some platform code!

## Changing Swagger Validation

Reviewing our backend logic for the particular NodeJS microservice which handles all user authentication and management, I found out that when a user is able to successfully  log into the system using an SSO mechanism, the microservice generates an object for this user and saves in it in the database. As part of this user object generation, there is validation of the request body of the email property. This is why, when attempting to generate a user without a TLD on the platform, the API returns a `422 - Unprocessable Entity`:

```bash
curl -X POST "http://test/api/users" -H "Accept: application/json" -H "Content-Type: application/json" -d
{
   "first_name": "John",
   "last_name": "Doe",
   "email": "jdoe@domain"
}
```

```json
{
   "error": {
        "message": "Request validation failed: Parameter (user) failed schema validation",
        "status": 422,
        "httpMessage": "Unprocessable Entity"
   }
}
```

After further debugging, I was able to pinpoint where this user creation validation is enforced: in the [Swagger](https://swagger.io/) configuration:

```yaml
user:
  type: object
  properties:
    email:
      type: string
      pattern: >-
        ^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$

```

Removing this pattern-based validation and restarting the microservice, I was able to successfully generate a user!

```bash
curl -X POST "http://test/api/users" -H "Accept: application/json" -H "Content-Type: application/json" -d
{
   "firstName": "John",
   "lastName": "Doe",
   "email": "jdoe@domain"
}
```

```json
{
   "email" : "jdoe@domain",
   "id" : "a1b2c3"
   "active": true,
   "created" : "2021-09-26T15:02:01.328Z",
   "lastUpdated": "2021-09-26T15:02:01.328Z",
   "firstName": "John",
   "lastName": "Doe",
   "username": "jdoe@domain"
}
```

Applying the same changes in the customer's environment resolved their issue!

**After note:** Please be mindful before making this change in a production environment. Removing the whole email validation in this case wasn't dangerous since we did it in an internal/dev environment to prove a point. A change such as the one described here can make your application API vulnerable to different attack vectors.
