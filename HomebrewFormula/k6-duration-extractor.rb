# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class K6DurationExtractor < Formula
  desc "Tool to extract the duration information per request from the K6 csv output."
  homepage "https://github.com/janritter/k6-duration-extractor"
  version "1.1.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.2/k6-duration-extractor_1.1.2_darwin_arm64.tar.gz"
      sha256 "78bac63af02ddf7b2372e67c76ce28313cd8f67b837f4f68c2298497b604340c"

      def install
        bin.install "k6-duration-extractor"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.2/k6-duration-extractor_1.1.2_darwin_amd64.tar.gz"
      sha256 "0a76e9e03e0c37af2db5e622c8a61350a6f35cdfd36cdd7e31a8e529c765868d"

      def install
        bin.install "k6-duration-extractor"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.2/k6-duration-extractor_1.1.2_linux_arm64.tar.gz"
      sha256 "30cce11918e55acdb3cb5c5482a7481f8602eab19445cfe1a8573163055b79bd"

      def install
        bin.install "k6-duration-extractor"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.2/k6-duration-extractor_1.1.2_linux_amd64.tar.gz"
      sha256 "9ba24a1691f8a291a40c8030c481a06dda65dff8eba5ae33d64d98a2958168c6"

      def install
        bin.install "k6-duration-extractor"
      end
    end
  end
end
