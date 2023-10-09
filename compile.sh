# Build using version, build date, commit hash and author name

# Default values
DEFAULT_VERSION="0.0.0"
DEFAULT_COMMIT_HASH="no-commits"
DEFAULT_AUTHOR="pilinux"

# Get the current version and commit hash from Git
if [ "$(git rev-list --count HEAD)" -gt 0 ]; then
    # Check if any Git tags exist
    if [ "$(git tag)" ]; then
        # Get the latest Git tag
        latest_tag=$(git describe --tags --abbrev=0)
        # Get the number of commits since the latest Git tag
        commits_since_tag=$(git rev-list --count "$latest_tag"..HEAD)
        # Get the commit hash of the latest commit
        commit=$(git rev-parse --short HEAD)
        # If there are no commits since the latest Git tag, use the tag as the version
        if [ "$commits_since_tag" -eq 0 ]; then
            version=$latest_tag
        else
            # If there are commits since the latest Git tag, use the tag and number of commits as the version
            version="$latest_tag-$commits_since_tag"
        fi
    else
        # If there are no Git tags, use the default tag and number of commits as the version
        version=$DEFAULT_VERSION-$(git rev-list --count HEAD)
        commit=$(git rev-parse --short HEAD)
    fi
else
    version=$DEFAULT_VERSION
    commit=$DEFAULT_COMMIT_HASH
fi

# Get the current date in UTC
date=$(date -u '+%Y-%m-%d_%H:%M:%S')

# List of target OS and architectures
targets=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/arm64"
)

# Print the version, commit hash and build date
echo "Building version $version"
echo "Commit hash $commit"
echo "Build date $date"

# Iterate over the target platforms and build binaries
for target in "${targets[@]}"; do
    # Split the target into OS and architecture
    IFS='/' read -ra target_split <<< "$target"
    os=${target_split[0]}
    arch=${target_split[1]}

    # Print the target OS and architecture
    echo "Building for $os/$arch"

    # Find all directories with a main function and build them
    for d in $(find . -name main.go -exec dirname {} \;); do
        echo "Building $d"

        # Extract the base name of the directory
        dirname=$(basename $d)

        # Build the binary
        echo "Building $dirname-$os-$arch"
        GOOS="$os" GOARCH="$arch" go build \
            -ldflags \
            " \
        -X 'github.com/pilinux/totext.AppVersion=$version' \
        -X 'github.com/pilinux/totext.BuildDate=$date' \
        -X 'github.com/pilinux/totext.CommitHash=$commit' \
        -X 'github.com/pilinux/totext.Author=$DEFAULT_AUTHOR' \
        " \
            -o "build/$dirname-$os-$arch" \
            -v \
            "$d"

        # Rename the output file for Windows
        if [ "$os" == "windows" ]; then
            mv "build/$dirname-$os-$arch" "build/$dirname-$os-$arch.exe"
        fi
    done
done
