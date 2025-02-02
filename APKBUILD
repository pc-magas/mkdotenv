# Maintainer: Your Name <your.email@example.com>
pkgname=mkdotenv
pkgver=0.1.0
pkgrel=0
pkgdesc="MkDotenv is a lightweight and efficient tool for managing your .env files."
url="https://github.com/pc-magas/mkdotenv"
arch="all"
license="GPL-3"
depends=""
makedepends="make go"
source="https://github.com/pc-magas/mkdotenv/archive/refs/tags/v$pkgver.tar.gz"
builddir="$srcdir/mkdotenv-$pkgver"

build() {
    cd "$builddir"
    make
}

package() {
    cd "$builddir"
    install -Dm755 mkdotenv "$pkgdir/usr/bin/mkdotenv"
    cp man/mkdotenv.1 "$pkgdir/usr/local/share/man/man1/"
}

sha512sums="SKIP"
