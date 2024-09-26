# Email Overview

Email programs and daemons have multiple roles and utilize various protocols.

## MUA (Mail User Agent)

It is the main role of your email client. A MUA is where you read your email and compose emails to others. The MUA uses IMAP (Internet Message Access Protocol) or POP3 (Post Office Protocol) to download your email.

## MSP (Mail Submission Program)

It is the role your email client has when you click 'Send'. The MSP submits your message to the MTA (Mail Transfer Agent) using SMTP (Simple Mail Transfer Protocol).

## MTA (Mail Transfer Agent)

It is the main role of your email server. The MTA is responsible for starting the process of sending the message to the recipient. It does this by looking up the recipient and sending the message to their MTA using SMTP.

## MDA (Mail Delivery Agent)

It is the role your email server takes when it receives an email for you. The MDA is responsible for storing the message for future retrieval. The MTA uses SMTP, LMTP (Local Mail Transfer Protocol), or another protocol to transfer the message to the MDA.

## SMTP, POP3 and IMAP Protocols

### Simple Mail Transfer Protocol (SMTP)

The Simple Mail Transfer Protocol (SMTP) is a TCP/IP protocol used as an Internet standard for electronic mail transmission.

It uses a plain "English" syntax such as HELO, MAIL, RCPT, DATA, or QUIT. SMTP is easily tested using telnet.

Sample mail exchange using telnet:

```bash
telnet localhost 25

Trying 127.0.0.1...
Connected to localhost.
Escape character is 'ˆ]'.
220 SERVER1 ESMTP Postfix (Ubuntu)

helo localhost
250 SERVER1

mail from:root@localhost
250 2.1.0 Ok

rcpt to:root@localhost
250 2.1.5 Ok

data
354 End data with <CR><LF>.<CR><LF>
This is neato
stuff ...

.
250 2.0.0 Ok: queued as 9C4DCEE3382

quit
221 2.0.0 Bye
```

### Post Office Protocol (POP3)

The Post Office Protocol (POP) is one of the main protocols used by MUAs to fetch mail. By default, the protocol downloads the messages and deletes them from the server. It is simpler yet less flexible protocol.

### Internet Message Access Protocol (IMAP)

The Internet Message Access Protocol (IMAP) is the other main protocol used by MUA to fetch mail. When using IMAP, the messages are managed on the server and left there. Copies are downloaded to the MUA. This protocol is more complex and more flexible than POP3.

## Email Life Cycle

The email life cycle is a fairly simple one.

1. You compose an email using your MUA.
2. Your MUA connects to your outbound MTA via SMTP, and sends the message to be delivered.
3. Your outbound MTA connects to the inbound MTA of the recipient via SMTP, and sends the message along. (Note: This step can happen more than once).
4. Once the message gets to the final destination MTA, it is delivered to the MDA. This can happen over SMTP, LMTP or other protocols.
5. The MDA stores the message (on disk as a file, or in a database, etc).
6. The recipient connects (via IMAP, POP3 or a similar protocol) to their email server, and fetches the message. The IMAP or POP daemon fetches the message out of the storage and sends it to the MUA.
7. The message is then read by the recipient.

## MTA, MDA, MUA and IMAP/POP Implementations

Mail Transfer Agent (MTA) Implementations
The Mail Transfer Agent (MTA) implementations use different software suites.

Sendmail was one of the first, if not the first SMTP implementation. Sendmail is difficult to configure with m4 macros. Older versions had security problems.

Exim is an alternative to Sendmail.

Postfix has an architecture characterized by separate binaries and division of privileges. It is a drop-in replacement for Sendmail.

## Mail Delivery Agents (MDA) Implementations

Many MTA software suites have components which act as Mail Delivery Agents (MDAs) as well. Some examples of these software suites include:

- Sendmail
- Postfix
- procmail
- Sieve
- Cyrus
- Spam Assassin.

When the MDA is activated to deliver a mail item to the local machine, there are some delivery options to be considered:

In postfix /etc/postfix/main.cf contains location information for the alias database. This has two common locations and names:

- `/etc/aliases` for system-wide redirection of mail
- `~/.forward` for user-configurable redirection of mail.

`/etc/aliases` has several options for its storage type; generally, data is stored in a .bdm or hash format as described by `/etc/postfix/main.cf`.

The database is built by running:

```bash
# postalias /etc/aliases
```

The format of the /etc/aliases and ~/.forward files is:
`name: value1, value2, value3...`

where name is the name the email was intended for initially (root, sysadmin). The value may be:

a different email address
a local file name, messages are appended
a command, via the pipe
`:include:/filename` to send to all the mail addresses in filename.
See man 5 aliases for more details.

## Mail User Agent (MUA) Implementations

Mail User Agent (MUA) is used to manage a user's email. Some of the software suites used include:

- Thunderbird
- Mutt
- Evolution
- Outlook/Outlook Express.

## IMAP/POP Implementations

Mail programs use IMAP and POP protocols to access mail stored on remote computers. Some of the servers used for implementing IMAP and POP are the following:

- UW: Reference implementation of the IMAP protocol.
- Cyrus IMAP: Enterprise-ready IMAP. It's more complex and more difficult to configure.
- Dovecot: Easy to use due to its powerful configuration options.
- Courier": Designed to be "plug-and-play" installable and does not require any site-specific configuration.
