# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://www.rubydoc.info/github/Homebrew/brew/master/Formula

class K6DurationExtractor < Formula
    desc "Tool to extract the duration information per request from the K6 csv output"
    if OS.mac?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.0/darwin_amd64_k6-duration-extractor"
      sha256 ""
      version "1.1.0"
    elsif OS.linux?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.0/linux_amd64_k6-duration-extractor"
      sha256 ""
      version "1.1.0"
    end

    def install
      if OS.mac?
        bin.install "darwin_amd64_k6-duration-extractor"
        mv bin/"darwin_amd64_k6-duration-extractor", bin/"k6-duration-extractor"
      elsif OS.linux?
        bin.install "linux_amd64_k6-duration-extractor"
        mv bin/"linux_amd64_k6-duration-extractor", bin/"k6-duration-extractor"
      end
    end

    test do
      system "false"
    end
  end