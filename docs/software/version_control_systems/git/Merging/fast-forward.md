# Fast-Forward Merge

This type of merge is possible when no commits are pushed in the main branch and all commits are pushed to the feature branch.

When performing a fast-forward merge, it takes the `master` branch pointer (`HEAD`) to the last commit in the feature branch:

After merging the feature branch, we can delete it if we have no intention of using it again:

```bash
git branch -d $BRANCH_NAME
```

To merge:

```bash
git branch feature_branch

git checkout feature_branch
echo "second file" >> second_file.txt

git add .
git commit -m "feature branch new commit"

git checkout master

git merge feature_branch

# d7bd8e8 is the master branch commit hash
# 7d0c2b0 is the feature branch commit hash
Updating d7bd8e8..7d0c2b0
Fast-forward
second_file.txt | 1 +
1 file changed, 1 insertion(+)
create mode 100644 second_file.txt


git log

commit 7d0c2b0529988a77f1418db87bf2810a3e11428d (HEAD -> master, temp)

Author: Kobbi Gal <kgal@m-c02fd2yumd6r.paloaltonetworks.local>
Date:   Tue Oct 12 19:01:55 2021 +0300

second commit

commit d7bd8e8519e2c5e23459bb981030a1ebf5f7d2a4
Author: Kobbi Gal <kgal@m-c02fd2yumd6r.paloaltonetworks.local>
Date:   Tue Oct 12 19:00:34 2021 +0300

first commit
```
