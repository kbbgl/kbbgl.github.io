# Multiple GitHub Accounts using Aliases

```bash
cat ~/.config/gh/config.yml

git_protocol: ssh
aliases:
    personal: '!cp ~/.config/gh/hosts.yml.personal ~/.config/gh/hosts.yml && gh auth status'
    work: '!cp ~/.config/gh/hosts.yml.work ~/.config/gh/hosts.yml && gh auth status'


cat ~/.config/gh/hosts.yml.personal
github.com:
    oauth_token: ghp_[…]
    git_protocol: ssh
    user: me

cat ~/.config/gh/hosts.yml.work
github.com:
    oauth_token: ghp_[…]
    git_protocol: ssh
    user: me_2

```

To use personal account:

```bash
gh personal
github.com
  ✓ Logged in to github.com as me (/home/me/.config/gh/hosts.yml)
  ✓ Git operations for github.com configured to use ssh protocol.
  ✓ Token: *******************
```

To use work account:

```bash
gh work
```

Source: <https://gist.github.com/yermulnik/017837c01879ed3c7489cc7cf749ae47>
