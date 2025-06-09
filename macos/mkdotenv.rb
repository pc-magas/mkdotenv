class Mkdotenv < Formula
  version "0.3.2"
  desc "Simplify Your .env Files – One Variable at a Time!"

  homepage "https://github.com/pc-magas/mkdotenv"
  
  url "https://github.com/pc-magas/mkdotenv/releases/download/v#{version}/mkdotenv-macos.zip"
  

  sha256 "8f2c192e4a967d17f705ef432a4dea330781450c6294c8114d29a6a120835b61"
  
  license "GPL-v3"

  def install
    bin.install "mkdotenv"
  end

  test do
    system "#{bin}/mkdotenv", "--help"
  end
end
