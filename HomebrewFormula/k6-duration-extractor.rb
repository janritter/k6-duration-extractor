# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://www.rubydoc.info/github/Homebrew/brew/master/Formula

class K6DurationExtractor < Formula
    desc "Tool to extract the duration information per request from the K6 csv output"
    if OS.mac?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.0/darwin_amd64_k6-duration-extractor"
      sha256 "50d95c80e132a77a6c07d08548826bebc013e1f0b6201c2cdbf6318ac7b35b9e"
      version "1.1.0"
    elsif OS.linux?
      url "https://github.com/janritter/k6-duration-extractor/releases/download/1.1.0/linux_amd64_k6-duration-extractor"
      sha256 "1603beb7dd03c09299b952f5ce47b64e2c9ccb2ae4d70c6e8d5ffeedca308ef3"
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