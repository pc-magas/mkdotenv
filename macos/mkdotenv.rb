class Mkdotenv < Formula
  version "0.3.2"
  desc "Simplify Your .env Files – One Variable at a Time!"

  homepage "https://github.com/pc-magas/mkdotenv"
  
  url "https://github.com/pc-magas/mkdotenv/releases/download/v#{version}/mkdotenv-macos.zip"
  
  sha256 "f4d9dadc8ee02cfcfde251d375dbf12b42d5bdb00792b29c07ca5af92b2649d3"
  
  license "GPL-v3"

  def install
    bin.install "mkdotenv"
  end

  test do
    system "#{bin}/mkdotenv", "--help"
  end
end
