#!/bin/bash

# Check if the output directory was provided as an argument
if [ -z "$1" ]; then
    echo "Usage: $0 <output_dir>"
    exit 1
fi

# Get the output directory from the first argument
OUTPUT_DIR=$1

# Check if the directory exists
if [ -d "$OUTPUT_DIR" ]; then
    echo "Removing all files inside $OUTPUT_DIR..."
    
    # Remove all files inside the directory
    rm -rf "$OUTPUT_DIR"/*
    
    echo "All files in $OUTPUT_DIR removed successfully."
else
    echo "$OUTPUT_DIR does not exist."
fi
