cask "crashtest" do
  version "0.0.10"
  sha256 "fd0e883e123d9fe02efa0dd4d5e560b92d4c2ba97a7c9b46308f5ab3d7170ddc"

  url "https://github.com/gomarcd/crashtest/releases/download/v#{version}/Crashtest_#{version}_universal.dmg",
      verified: "github.com/gomarcd/crashtest/"
  name "Crashtest"
  desc "Fast, minimal API tool"
  homepage "https://crashtest.app/"

  app "Crashtest.app"

  depends_on macos: ">= :high_sierra"
end