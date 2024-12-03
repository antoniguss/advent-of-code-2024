#!/bin/bash

# Check if the template directory exists
TEMPLATE_DIR="template"
if [ ! -d "$TEMPLATE_DIR" ]; then
    echo "Template directory '$TEMPLATE_DIR' does not exist. Please create it first."
    exit 1
fi

# Get the current day of the month
DAY=$(date +%d)

# Format day to be two digits (01, 02, ..., 25)
DAY_FORMATTED=$(printf "%d" $DAY)

# Create the new directory for the day
DAY_DIR="solutions/day$DAY_FORMATTED"
if [ -d "$DAY_DIR" ]; then
    echo "Directory '$DAY_DIR' already exists. Skipping..."
    exit 0
fi

# Copy the template directory to the new day directory
cp -r "$TEMPLATE_DIR" "$DAY_DIR"

# Update the solution.go file to reflect the current day
sed -i '' "s/Day X/Day $DAY_FORMATTED/" "$DAY_DIR/solution.go"

echo "Setup for Day $DAY_FORMATTED complete."

