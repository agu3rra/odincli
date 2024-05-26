#!/bin/bash
# Run from the project root

# Create new branch
git checkout -b fabric.update

# Fetch and pull the latest changes
# git submodule update --remote --merge --depth=1
git submodule update --init --recursive --remote --merge --depth=1

# Remove unwanted files
cd fabric
find . -mindepth 1 -maxdepth 1 ! -name 'patterns' -exec rm -rf {} +
cd ..

# Add and commit changes
git add fabric
git commit -m "Updated fabric submodule"
