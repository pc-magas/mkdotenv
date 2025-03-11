# #!/bin/bash

# SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

# cp -r ${SCRIPTPATH}/git-hooks/* ${SCRIPTPATH}/../.git/hooks/

# chmod +x ${SCRIPTPATH}/../.git/hooks/*

#!/bin/bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
HOOKS_DIR="${SCRIPTPATH}/git-hooks"
GIT_HOOKS_DIR="${SCRIPTPATH}/../.git/hooks"

# Ensure the hooks directory exists
mkdir -p "$GIT_HOOKS_DIR"

# Loop through all hook scripts and create symlinks
for hook in "$HOOKS_DIR"/*; do
    hook_name=$(basename "$hook")
    target_hook="$GIT_HOOKS_DIR/$hook_name"

    # Remove existing file or symlink
    if [ -e "$target_hook" ] || [ -L "$target_hook" ]; then
        rm -f "$target_hook"
    fi

    chmod +x
    # Create a symlink
    ln -s "$hook" "$target_hook"
    echo "Symlinked $hook -> $target_hook"
done

# Ensure hooks are executable
chmod +x "$HOOKS_DIR"/*

echo "Git hooks successfully symlinked!"