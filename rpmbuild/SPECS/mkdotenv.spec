Name:           mkdotenv
Version:        0.1.7
Release:        1%{?dist}
Summary:        Lightweight and efficient tool for managing your `.env` files.

License:        GPL-3
URL:            https://github.com/pc-magas/mkdotenv
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  make

%description
MkDotenv is a lightweight and efficient tool for managing your `.env` files. 
Whether you're adding, updating, or replacing environment variables, MkDotenv makes it easy and hassle-free.

%prep
%autosetup

# Download and install Go 1.23 if it's not available in Fedora repos
sudo rm -rf /usr/local/go
export GO_VERSION=1.23.6
curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | sudo tar -C /usr/local -xz 
export PATH=/usr/local/go/bin:$PATH

%build
make %{?_smp_mflags}

%install
mkdir -p %{buildroot}/usr/bin
mkdir -p %{buildroot}/usr/share/man/man1
install -m 0755 mkdotenv %{buildroot}/usr/bin/mkdotenv
install -m 0644 man/mkdotenv.1 %{buildroot}/usr/share/man/man1/mkdotenv.1

%files
%license LICENSE
/usr/bin/mkdotenv
/usr/share/man/man1/mkdotenv.1.gz

%changelog
* Sun Feb 16 2025 Dimitrios Desyllas <pcmagas@disroot.org> - %{version}-1
- Initial RPM package
