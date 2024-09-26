# Package Management Systems

## Types of packages

- **Binary**: compiled executables per architecture.  
- **Source**: packages with source code that can be compiled for particular architecture.

## Packaging Tool Levels

- High: automatic package dependency resolution.

  - `apt`
  - `yum`

- Low: manual dependency and installation of package dependencies.  
  - `rpm` (Red Hat Package Manager)

    - By default, manages local repository  packages.
    - Binary package format `<name>-<version>-<release>.<distro>.<architecture>.rpm`.
    - Source package format `<name>-<version>-<release>.<distro>.src.rpm`.
    - database is held in `/var/lib/rpm`
    - helper scripts are located in `/usr/lib/rpm/`.
    - can create `rpmrc` file to modify `rpm` defaults. `rpm` looks for the file in order `/usr/lib/rpm/rpmrc`, `/etc/rpmrc`, `~/.rpmrc` and reads all of them. Can use `--rcfile` as an argument to specify as well.
    - Useful query commands

        ```bash
        # Version of package?
        rpm -q bash
        
        # Which package did file come from?
        rpm -qf /bin/bash
        
        # What files were installed by the package?
        rpm -ql bash
        
        # Show package info
        rpm -qi bash
        
        # Show package info from file
        rpm -qip foo-1.0.0-1.noarch.rpm
        
        # List all packages installed
        rpm -qa
        
        # List package dependencies
        rpm -qp --requires some_package-0.0.1-noarch.rpm
        
        # What installed package provides a particular requisite package
        rpm -q --whatprovides libc.so.6
        
        # Verify package matches `rpm` database
        
        rpm -V bash
        
        # Install a package
        rpm -i bash-4.4.19-8.el8_0.x86_64
        
        # Dry run erase a package
        rpm -e nano --test
        
        # Update package(s)
        rpm -U bash-4.4.19-10.el8.x86_64.rpm another_pkg
        
        # Revert to older package
        rpm -U --oldpackage bash-4.3.19-10.el8.x86_64.rpm
        
        # Upgrade multiple packages in directory
        rpm -F *.rpm
        
        # Install kernel (don't upgrade since it removes older kernels!)
        rpm -ivh kernel-{version}.{arch}.rpm
        
        # Convert .rpm to cpio
        rpm2cpio foobar.rpm > foobar.cpio
        
        # List files in RPM package
        rpm2cpio -qlp bash-XXX.rpm
        
        ```

  - `dpkg` (Debian Package Manager)

    - `*.deb` prefix.
    - Database location `/var/lib/dpkg`.
    - Standard naming format `<name>_<version>-<revision_number>_<architecture>.deb`.
    - A source package consists of:
      - unmodified sourcecode in `.tar.gz` format.
      - A description file `.dsc`.
      - Patch/additional files `.debian.tag.gz` or `.diff.gz`.

    - Commands

        ```bash
        # List all installed packages
        dpkg -l 
        
        # List files installed in package
        dpkg -L curl
        
        # Show information about installed package
        dpkg -p udev
        
        # See what package owns a file
        dpkg -S /etc/init/networking.conf
        
        # Show file status
        dpkg -s wget
        
        # Verify installed package integrity
        dpkg -V wget
        
        # Install/upgrade package
        dpkg -i foobar.deb
        
        # Remove package (- config)
        dpkg -r package
        
        # Purge package (+ config)
        dpkg -P package
        ```

## Package Sources

Distros have package repositories that contain all information about most up-to-date software.

## Revision Control Systems

`git` has two data strcutures:

- object database (`
- /.git/objects`:
  - **Blobs**: chunks of binary data containing file contents. Can use `git cat-file` to translate binary data to human-friendly text.
  - **Trees**: sets of blobs that construct the directory structure.
  - **Commits**: changesets describing tree snapshots.
- directory cache : captures the current state of the directory tree.
