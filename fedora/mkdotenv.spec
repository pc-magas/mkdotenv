%global debug_package %{nil}
Name:           mkdotenv
Version:        0.4.8
Release:        1%{?dist}
Summary:        Lightweight and efficient tool for managing your `.env` files.

License:        GPL-3
URL:            https://github.com/pc-magas/mkdotenv
Source0:        %{name}-%{version}-rpm.tar.gz

BuildRequires:  go, make

%description
MkDotenv is a lightweight and efficient tool for managing your `.env` files. 
Whether you're adding, updating, or replacing environment variables, MkDotenv makes it easy and hassle-free.

%prep
%setup -q

%build
    make VERSION="%{version}" INSTALL_BIN_DIR=/usr/bin INSTALL_MAN_DIR=/usr/share/man/man1

%install
    make install DESTDIR="%{buildroot}" INSTALL_BIN_DIR=/usr/bin INSTALL_MAN_DIR=/usr/share/man/man1

%files
/usr/bin/mkdotenv
/usr/share/man/man1/mkdotenv.1.gz

%changelog
* Sun Feb 16 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.1.7-1
- Initial RPM package

* Thu Mar 27 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.2.0-1
- 1. Split codebase into multiple files.
- 2. Use a seperate version file and define built version upon compile.
- 4. [BUGFIX] If input file is same as output file copy input file into a temporary one.
- 5. Improved Documentation
- 6. [BUGFIX] Out of bounds argument parsing
- 7. [BUGFIX] Values should not be an Argument

* Mon Apr 07 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.2.1-1
- 1. Improve Argument parsing

* Thu Apr 10 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.2.2-1
- Release for windows
- make runs `bin` target by default
- Fix lang upon rpm changelog

* Thu Apr 24 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.2.3-1
- Release for Alpine

* Mon May 12 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.1-1
- 1. Use common naming convention for golang module using repository's name
- 2. Upon rpm builds use Makefile
- 3. Ability to specify a version externally in Makefile.
- 4. Unit test value appending logic
	- 5. Validate variable name
- 6. Moving pcmagas/alpinebuild (used upon alpine image releases) docker image into a seperate repository.
- 7. Release for AUR and arch linux

* Mon May 26 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.1-1
- 1. [BUGFIX] Check Variable name
- 2. Use External Docker image for rpm builds.

* Sat May 31 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.2-1
- Release for MACOS

* Thu Jun 12 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.2-2
- Release for MACOS
- Use ghcr hosted images for fedora and alpine builds
- Native building of app usiong Make both on MacOs and Linux.

* Sun Jun 22 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.3-1
- Improved README (fixing typos and improve clarification).
- Arch package does installs oficial golang and does not requires it as dependency.


* Sun Jun 22 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.3-2
- Improved README (fixing typos and improve clarification).
- Arch package does installs oficial golang and does not requires it as dependency.
- Using ./fedora folder for rpm building.

* Thu Jun 26 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.3-3
- Improved README (fixing typos and improve clarification).
- Arch package does installs oficial golang and does not requires it as dependency.
- Using ./fedora folder for rpm building.
- Improvement upon debian building
- Fixing ppa release

* Thu Jun 26 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.3-3
- Improved README (fixing typos and improve clarification).
- Arch package does installs oficial golang and does not requires it as dependency.
- Using ./fedora folder for rpm building.
- Improvement upon debian building
- Fixing ppa release

* Tue Jul 01 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.3.4-1
- Add release number upon ppa build.
- Fix alpine build

* Fri Jul 18 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.0-1
- Fix alpine build.
- Use variable-value parameter for setting the value as variable value.
- Use variable-name parameter for setting the value as variable value.
- [NEW FEATURE] Flag to remove multiple occurences of the variable.
- [NEW FEATURE] Use - value upon in order to output modified .env contents upon stdout. Default behaviour is outputing upon .env

* Wed Aug 20 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.1-1
- Use Seperate script for debian builds

* Wed Aug 27 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.2-1
- Use Seperate script for debian builds
- Vendor dependencies upon build.
- Introducing tools for PPA and debian building
- Improve Documentation
- Use `pflag` for argument parsing

* Wed Sep 03 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.3-1
- Fix version upon manpage and remove non-existent arguments
- Fix pipeline upon releasing macos homebrew fromula.

* Sat Sep 06 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.4-1
- Alpine changes upon Makefile.

* Fri Sep 12 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.4-2
- Alpine changes upon Makefile.
- Autogenerate APKBUILD for alpine builds.

* Fri Sep 12 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.5-1
- Vendor dependencies upon alpine vendor source.

* Fri Sep 12 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.6-1
- Add extra supported values for building on alpine aports ci/cd.

* Fri Sep 12 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.7-1
- Autoenable gco support on some architectures. Windows only for x86_64 only.

* Sat Sep 13 2025 Dimitrios Desyllas <pcmagas@disroot.org> - 0.4.8-1
- Enable GCO on some architectures also allow default GCO architectures as well.
