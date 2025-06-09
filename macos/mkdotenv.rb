class Mkdotenv < Formula
  version "0.3.2"
  desc "Simplify Your .env Files – One Variable at a Time!"

  homepage "https://github.com/pc-magas/mkdotenv"
  
  url "https://github.com/pc-magas/mkdotenv/releases/download/v#{version}/mkdotenv-macos.zip"
  
  sha256 "194e46e698642006eaecf050fbd757f8537bb98a4dbfecd3a71f70e56f23d861"
  
  license "GPL-v3"

  def install
    bin.install "mkdotenv"
  end

  test do
    system "#{bin}/mkdotenv", "--help"
  end
end
