#!/bin/bash

set -e

# Function to display usage
usage() {
    echo "Usage: $0 --service-name <service-name>"
    echo "Example: $0 --service-name auth-api"
    exit 1
}

# Parse command line arguments
SERVICE_NAME=""
while [[ $# -gt 0 ]]; do
    case $1 in
        --service-name)
            SERVICE_NAME="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            usage
            ;;
    esac
done

# Validate service name
if [ -z "$SERVICE_NAME" ]; then
    echo "Error: --service-name is required"
    usage
fi

# Validate service name format (lowercase letters, numbers, and hyphens only)
if ! [[ $SERVICE_NAME =~ ^[a-z0-9-]+$ ]]; then
    echo "Error: Service name must contain only lowercase letters, numbers, and hyphens"
    exit 1
fi

echo "Setting up service: $SERVICE_NAME"

# Replace all instances of <REPLACE> with the service name, excluding README.md and hidden files
find . -type f -not -path "*/\.*" -not -path "*/bin/*" -not -name "README.md" -exec sed -i '' "s/<REPLACE>/$SERVICE_NAME/g" {} +

echo "✅ Setup complete! Service name has been replaced with: $SERVICE_NAME"
echo "Next steps:"
echo "1. Run 'go mod tidy' to update dependencies"
echo "2. Run 'make up' to start the database"
echo "3. Run 'make run' to start the service" 