#!/bin/bash

# Define the directory name
TARGET_DIR="typescript"

# Check if the directory exists
if [ -d "$TARGET_DIR" ]; then
    echo "Changing to $TARGET_DIR directory..."
    
    # Navigate to the directory
    cd "$TARGET_DIR" || exit
    
    echo "Removing all files in dist/*..."
    
    # Remove all files inside the dist directory
    rm -rf dist/*
    
    echo "Running npm build..."
    
    # Run the npm build command
    npm run build
    
    # Check if the build was successful
    if [ $? -eq 0 ]; then
        echo "Build completed successfully."
    else
        echo "Build failed."
        exit 1
    fi

    # Install any necessary modules
    npm i
else
    echo "$TARGET_DIR directory does not exist."
    exit 1
fi
