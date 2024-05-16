#!/bin/bash

# Define the pattern for the files to be cleaned up
file_pattern="*-connector.yaml"

# Find all files matching the pattern and delete Kubernetes resources
echo "Deleting Kubernetes resources..."
for filename in $file_pattern; do
    if [ -f "$filename" ]; then
        kubectl delete -f "$filename"
        echo "Deleted Kubernetes resource defined in $filename"
    else
        echo "No files found matching the pattern $file_pattern"
        exit 1
    fi
done

# Remove the files from the file system
echo "Cleaning up files..."
for filename in $file_pattern; do
    if [ -f "$filename" ]; then
        rm "$filename"
        echo "Removed file $filename"
    fi
done

echo "Cleanup completed."
