#!/bin/bash
# Run this script to update the fabric/patterns folder

# Navigate to the submodule directory
cd fabric

# Fetch and pull the latest changes
git fetch --depth 1
git pull origin main

# Remove unwanted files
find . -mindepth 1 -maxdepth 1 ! -name 'patterns' -exec rm -rf {} +

# Navigate back to the main project directory
cd ..

# Add and commit changes
git add fabric
git commit -m "Updated fabric submodule"
