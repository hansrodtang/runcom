require 'formula'

class Runcom < Formula
  homepage "http://runcom.github.io/"
  head "https://github.com/hansrodtang/runcom.git"
  url "https://github.com/hansrodtang/runcom/archive/v1.0.tar.gz"
  sha1 "a821fcde92b03baf49a45970d0ae6b781a3b12c1"

  depends_on "go" => :build
  depends_on "bazaar" => :build
  depends_on :hg => :build

  def install
    ENV["GOBIN"] = bin
    ENV["GOPATH"] = buildpath
    ENV["GOHOME"] = buildpath
    system "go", "get"
    system "go", "build", "main.go"
    bin.install "main" => "runcom"
  end

  test do
    site = testpath/"hops-yeast-malt-water"
    system "#{bin}/runcom", "init", site
    assert File.exist?("#{site}/config.toml")
  end
end
