#!/bin/bash

die () {
    echo >&2 "$@"
	    exit 1
		}

[ "$#" -eq 1 ] || die "1 argument required, $# provided"

src=$1

filename="$(basename -s .asm $src)"
tmp_dir=$(mktemp -d)

echo "Assembling $src..."
o_file="$tmp_dir/$filename.o"
nasm -f elf64 $src -o $o_file
echo "Finished assembling $o_file"

echo "Linking $o_file..."
output="$tmp_dir/$filename"
ld -o $output $o_file

echo "Compilation complete, see file $output"

