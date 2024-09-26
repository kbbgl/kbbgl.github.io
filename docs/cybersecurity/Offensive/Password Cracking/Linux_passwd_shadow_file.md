### Password
![passwd](https://www.cyberciti.biz/media/ssb.images/uploaded_images/passwd-file-791527.png)

1. Username (1-32 chars)
2. `x` indicates that it's an encrypted password in the `/etc/shadow` file.
3. `uid` - UID 0 (zero) is reserved for root and UIDs 1-99 are reserved for other predefined accounts. Further UID 100-999 are reserved by system for administrative and system accounts/groups.
4. `gid` - stored in `/etc/group`
5. Comment field.
6. User's `~` directory.
7. Absolute path to command/shell. Doesn't need to be a shell.

### Shadow
![shadow](https://www.cyberciti.biz/media/new/uploaded_images/shadow-file-795497.png)

1. Username.
2. Password : It is your encrypted password. The password should be minimum 8-12 characters long including special characters, digits, lower case alphabetic and more. Usually password format is set to `$id$salt$hashed`, The `$id` is the algorithm used On GNU/Linux as follows:

	`$1$` is `MD5`

	`$2a$` is `Blowfish`

	`$2y$` is `Blowfish`

	`$5$` is `SHA-256`

	`$6$` is `SHA-512`

3. Last password change (lastchanged) : Days since Jan 1, 1970 that password was last changed

4. Minimum : The minimum number of days required between password changes i.e. the number of days left before the user is allowed to change his/her password
5. Maximum : The maximum number of days the password is valid (after that user is forced to change his/her password)
6. Warn : The number of days before password is to expire that user is warned that his/her password must be changed
7. Inactive : The number of days after password expires that account is disabled
8. Expire : days since Jan 1, 1970 that account is disabled i.e. an absolute date specifying when the login may no longer be used.