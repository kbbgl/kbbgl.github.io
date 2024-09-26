# Merge Conflict

Occurs when **the same files are edited in 2 different branches**.

They can only occur in 3-way merge situation.

When a merge conflict occurs:

```bash
❯ git branch --show-current
master

# sixth_file.txt and fifth_file.txt were modified in both master and br3

❯ git merge br-3
Auto-merging new_files/sixth_file.txt
CONFLICT (content): Merge conflict in new_files/sixth_file.txt
Auto-merging new_files/fifth_file.txt
CONFLICT (content): Merge conflict in new_files/fifth_file.txt
Automatic merge failed; fix conflicts and then commit the result.

❯ git status
On branch master
You have unmerged paths.
(fix conflicts and run "git commit")
(use "git merge --abort" to abort the merge)

Unmerged paths:
(use "git add <file>..." to mark resolution)

both modified:   new_files/fifth_file.txt
both modified:   new_files/sixth_file.txt


❯ cat new_files/fifth_file.txt
Hello World 05
<<<<<<< HEAD
file modified in master.
=======
this line was changed in br-3.
>>>>>>> br-3
```

Listing the staging area:

```bash
❯ git ls-files -s
100644 211dc43543ca6c6b9724fff9986d86da73e4bd43 0       eighth_file.txt
100644 557db03de997c86a4a028e1ebd3a1ceb225be238 0       first_file.txt
100644 1c227b889361ac0306a749e380b32cab1ca353ef 0       fourth_file.txt
100644 f064900941db9f2663cc26503d18af50deda7eb7 1       new_files/fifth_file.txt
100644 59388dde1d79d727b4127343d0e67de097a83a00 2       new_files/fifth_file.txt
100644 058e19228f6a25cee8b6ade79325434d45e48089 3       new_files/fifth_file.txt
100644 e1d76690d44fbf2de1a38584a247e1512ec361d1 1       new_files/sixth_file.txt
100644 f72edbf6802ae9788296e1efa770a8057ac365fe 2       new_files/sixth_file.txt
100644 57f032fd71af3a11f7879b2084882e102be81228 3       new_files/sixth_file.txt
100644 3ee384936466a484e0089c82ce559a10dc9c46ea 0       second_file.txt
100644 2ce0b4c4bd1b27caf3bde75b3885e4e8157bdeea 0       seventh_file.txt
100644 22c0dee49b87ff7ed42c7bd37987163cbe5b0d60 0       third_file.txt
```

The number in the third column indicates the version of the blob/file.

`0` means unchanged

`1` means the initial/common version of the file, the file content before branching.

`2` means the version of the file in the receiving/master branch.

`3` means the version of the file in the feature branch.

```bash
❯ git ls-files -s | grep fifth_file.txt | cut -d" " -f2 | while read hash; do echo "===="; git cat-file -p $hash; echo "===="; done
====
Hello World 05
====
====
Hello World 05

file modified in master.
====
====
Hello World 05


this line was changed in br-3.
====


❯ cat new_files/fifth_file.txt
Hello World 05


<<<<<<< HEAD
file modified in master.
=======
this line was changed in br-3.
>>>>>>> br-3
```

To remove the conflicts:

1) Delete all lines in file that has conflict including `<<<` and `>>>` lines.
2) `git add files...`
3) `git commit`
