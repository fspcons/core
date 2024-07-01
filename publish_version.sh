#!/bin/bash

helpFunction() {
    echo ""
    echo "Usage: $0 -p versionPattern"
    echo -e "\t-p Defines which version will be incremented. Valid values are MAJOR, MINOR or PATCH"
    echo -e "\t-c Also adds and commits all pending changes before applying the tag. Valid values are true or false"
    exit 1 # Exit script after printing help
}

commit=false

while getopts "p:c:" opt; do
    case "$opt" in
    p) versionPattern="$OPTARG" ;;
    c) commit=true ;;
    ?) helpFunction ;; # Print helpFunction in case parameter is non-existent
    esac
done

# Print helpFunction in case parameters are empty
if [ -z "$versionPattern" ]; then
    echo "Required the parameter is empty"
    helpFunction
fi

git tag -l | xargs git tag -d && git fetch -t

# version pattern #.#.# -> MAJOR.MINOR.PATCH  (i.e. v1.2.3)
current_version=$(git describe --abbrev=0)

echo "Current version: $current_version"

major="$(cut -d'.' -f1 <<<"$current_version")"
minor="$(cut -d'.' -f2 <<<"$current_version")"
patch="$(cut -d'.' -f3 <<<"$current_version")"
major="${major:1}"

if [ -z "$major" ]; then
    major=0
fi
if [ -z "$minor" ]; then
    minor=0
fi
if [ -z "$patch" ]; then
    patch=0
fi

case $versionPattern in
MAJOR)
    ((major++))
    minor=0
    patch=0
    ;;
MINOR)
    ((minor++))
    patch=0
    ;;
PATCH) ((patch++)) ;;
?) echo "Error, invalid version pattern" ;;
esac

new_version="v$major.$minor.$patch"
echo "New version: $new_version"

if [ "$commit" = true ]; then
    echo 'Committing changes ...'

    git status
    git add .
    git commit -m "publishing new version $new_version"
fi

git tag -a "$new_version" -m "publishing new version $new_version"

git push origin "$new_version"

git push