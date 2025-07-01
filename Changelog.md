# Version 0.3.4 2025-07-01

Add release number upon ppa build.
Fix alpine build

# Version 0.3.3 2025-06-26

Improved README (fixing typos and improve clarification).
Arch package does installs oficial golang and does not requires it as dependency.
Using ./fedora folder for rpm building.
Improvement upon debian building
Fixing ppa release

# Version 0.3.2 2025-06-12

Release for MACOS
Use ghcr hosted images for fedora and alpine builds
Native building of app usiong Make both on MacOs and Linux.

# Version 0.3.1 2025-05-26

1. [BUGFIX] Check Variable name
2. Use External Docker image for rpm builds.

# Version 0.3.0 2025-05-12

1. Use common naming convention for golang module using repoesitory's name
2. Upon rpm builds use Makefile
3. Ability to specify a version externally in Makefile.
4. Unit test value appending logic
5. Validate variable name
6. Moving pcmagas/alpinebuild (used upon alpine image releases) docker image into a seperate repository.
7. Release for AUR and arch linux

# Version 0.2.3 2025-04-24

Release for Alpine

# Version 0.2.2 2025-04-10

Release for windows
make runs `bin` target by default
Fix lang upon rpm changelog

# Version 0.2.1 2025-04-07

1. Improve Argument parsing

# Version 0.2.0 2025-03-27

1. Split codebase into multiple files.
2. Use a seperate version file and define built version upon compile.
4. [BUGFIX] If input file is same as output file copy input file into a temporary one.
5. Improved Documentation
6. [BUGFIX] Out of bounds argument parsing
7. [BUGFIX] Values should not be an Argument

# Version 0.1.7 2025-02-23

1. Release rpm image

# Version 0.1.5 2025-02-12

1. Build for PPA

# Version 0.1.0 2025-02-02

1. Build Debian package
2. Add build from sources upon documentation
3. Build Docker image
4. If no arguments provided print usage.

# Version 0.0.1 2025-01-31

1. Basic .env manipulation
2. Write .env files
3. Print .env to stdout
4. Receive .env contents from unix pipes.
