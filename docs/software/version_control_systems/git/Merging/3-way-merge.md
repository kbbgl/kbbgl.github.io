# 3-Way Merge

When both the master and the feature branch has commits, Git uses a 3-way merge.

In such a case, it creates a new commit and this new commit is based on 3 other commits:

- The commit the feature branch was created on (ancestor commit)
- The latest commit in the `master` branch
- The latest commit in the feature branch

The new commit will have 2 parents:

- The latest commit in the `master` branch
- The latest commit in the feature branch

We can then delete the feature branch.

```bash
git merge br-2
Merge made by the 'recursive' strategy.
new_files/fifth_file.txt | 1 +
new_files/sixth_file.txt | 1 +
seventh_file.txt         | 1 +
3 files changed, 3 insertions(+)

git --no-pager log

commit 37cd146dbc077047a4d92a75ef49ebb88a7fe604 (HEAD -> master)
Merge: d97f997 5d4a148

Merge branch 'br-2'

commit 5d4a1481501d84d5da6599ca32ef9a2f76162bf6 (br-2)

changed file five

git cat-file -p 37cd146dbc077047a4d92a75ef49ebb88a7fe604

# two parent commits, from master and from br-2
tree c5f4dca7f0425badd0717f688e0626fc87f50156
parent d97f997d73ee9ee4ff200c84b7aba1c03affcead
parent 5d4a1481501d84d5da6599ca32ef9a2f76162bf6

Merge branch 'br-2'

```
