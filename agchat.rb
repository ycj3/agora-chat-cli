class Agchat < Formula
    desc "A command-line interface for Agora Chat"
    homepage "https://github.com/ycj3/agora-chat-cli"
    url "https://github.com/ycj3/agora-chat-cli/archive/refs/tags/v0.1.0.tar.gz"
    sha256 "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
    license "MIT"
  
    depends_on "go" => :build
  
    def install
        if Hardware::CPU.arm?
          system "go", "build", *std_go_args(ldflags: "-s -w -buildmode=pie")
        elsif Hardware::CPU.intel?
          system "go", "build", *std_go_args(ldflags: "-s -w")
        end
      end
  
    test do
      assert_match "agchat", shell_output("#{bin}/agchat --version")
    end
  end
  