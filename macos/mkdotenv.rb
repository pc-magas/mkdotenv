class Mkdotenv < Formula
  version "0.4.2"
  desc "Simplify Your .env Files – One Variable at a Time!"

  homepage "https://github.com/pc-magas/mkdotenv"
  
  url "https://github.com/pc-magas/mkdotenv/releases/download/v#{version}/mkdotenv-macos.zip"
  
  sha256 "b3e323469c8bf35a9db2dd30fabc32c48f93ad7cb9e654443ee26fdfab8d95a5"
  
  license "GPL-3.0-or-later"

  def install
    bin.install "mkdotenv"
  end

  test do
    system "#{bin}/mkdotenv", "--help"
  end
end
