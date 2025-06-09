class Mkdotenv < Formula
  version "0.3.2"
  desc "Simplify Your .env Files – One Variable at a Time!"

  homepage "https://github.com/pc-magas/mkdotenv"
  
  # url "https://github.com/pc-magas/mkdotenv/releases/download/v#{version}/mkdotenv-macos.zip"
  
  url "file:///Volumes/SystemRoot/home/pcmagas/Kwdikas/go/mkdotenv/mkdotenv_app/macos/mkdotenv-macos.zip"

  sha256 "34ec00b2738c326f817a2560d07d32c61fc5e5dd927334f5c235786c4e5c60f4"
  
  license "GPL-v3"

  def install
    bin.install "mkdotenv"
  end

  test do
    system "#{bin}/mkdotenv", "--help"
  end
end
