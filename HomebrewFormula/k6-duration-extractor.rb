# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class K6DurationExtractor < Formula
  desc "Tool to extract the duration information per request from the K6 csv output."
  homepage "https://github.com/janritter/k6-duration-extractor"
  version "1.1.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.1/k6-duration-extractor_1.1.1_darwin_arm64.tar.gz"
      sha256 "5e9ba20aaaa692f5e887e1fc7df701f88062e71b84386d9269c691c38fd42222"

      def install
        bin.install "k6-duration-extractor"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.1/k6-duration-extractor_1.1.1_darwin_amd64.tar.gz"
      sha256 "4514103a27a90d700f02cce63e3a117742b06ffdef257cf5063b04e0db74a3f5"

      def install
        bin.install "k6-duration-extractor"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.1/k6-duration-extractor_1.1.1_linux_arm64.tar.gz"
      sha256 "891f6edbbdf40ea2548aa6d67c7d16edbaf4f715ad27c667d2f7a9a30fdf0631"

      def install
        bin.install "k6-duration-extractor"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.1/k6-duration-extractor_1.1.1_linux_amd64.tar.gz"
      sha256 "dc694e625217fcae9402698c7e115550d3a89635a57c438493a9e97c15939288"

      def install
        bin.install "k6-duration-extractor"
      end
    end
  end
end