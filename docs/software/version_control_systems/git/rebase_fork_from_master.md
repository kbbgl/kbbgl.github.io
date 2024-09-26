# Rebase Fork from Upstream

```bash
# Check if there's an upstream set up
git remote -v

# If upsteam doesn't exist
git remote add upstream https://github.com/original-repo/goes-here.git

git fetch upstream

git rebase upstream/master

git push origin master --forc
```
