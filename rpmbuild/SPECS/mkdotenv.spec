%global debug_package %{nil}
Name:           mkdotenv
Version:        0.2.2
Release:        1%{?dist}
Summary:        Lightweight and efficient tool for managing your `.env` files.

License:        GPL-3
URL:            https://github.com/pc-magas/mkdotenv
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  go

%description
MkDotenv is a lightweight and efficient tool for managing your `.env` files. 
Whether you're adding, updating, or replacing environment variables, MkDotenv makes it easy and hassle-free.

%prep
%setup -q

%build
	ls -l 
	cd ./mkdotenv &&\
	rm -rf mkdotenv &&\
    go build -o ./mkdotenv -ldflags "-X mkdotenv/msg.version=%{version}" ./mkdotenv.go && \
	cd ..

%install
mkdir -p %{buildroot}/usr/bin
mkdir -p %{buildroot}/usr/share/man/man1
install -m 0755 mkdotenv/mkdotenv %{buildroot}/usr/bin/mkdotenv
install -m 0644 man/mkdotenv.1 %{buildroot}/usr/share/man/man1/mkdotenv.1

%files
/usr/bin/mkdotenv
/usr/share/man/man1/mkdotenv.1.gz

%changelog
* Sun Feb 16 2025 Dimitrios Desyllas <pcmagas@disroot.org> - %{version}-1
- Initial RPM package

* Thu Mar 27 2025 Dimitrios Desyllas <pcmagas@disroot.org> - %{version}-1
- 1. Split codebase into multiple files.
- 2. Use a seperate version file and define built version upon compile.
- 4. [BUGFIX] If input file is same as output file copy input file into a temporary one.
- 5. Improved Documentation
- 6. [BUGFIX] Out of bounds argument parsing
- 7. [BUGFIX] Values should not be an Argument

* Mon Apr 07 2025 Dimitrios Desyllas <pcmagas@disroot.org> - %{version}-1
- 1. Improve Argument parsing

* Thu Apr 10 2025 Dimitrios Desyllas <pcmagas@disroot.org> - %{version}-1
- Release for windows
- make runs `bin` target by default
- Fix lang upon rpm changelog
