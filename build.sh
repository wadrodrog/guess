#!/usr/bin/bash

PACKAGE="guess"
PLATFORMS=(
    "windows/amd64" "linux/amd64" "darwin/amd64" "darwin/arm64"
    "android/arm64"
)
OUTPUT_DIR="target"

IFS='/'  # Delimeter for platforms array

for platform in "${PLATFORMS[@]}";
do
    # Split OS and Arch (by IFS value)
    read -ra plat <<< "$platform"

    GOOS=${plat[0]}
    GOARCH=${plat[1]}
    
    output_name=$PACKAGE'-'$GOOS'-'$GOARCH
    if [ "$GOOS" = "windows" ]; then
        output_name+='.exe'
    fi    

    echo "Building $output_name..."
    if ! env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUTPUT_DIR/$output_name" main.go
    then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

