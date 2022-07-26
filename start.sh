#!/bin/bash

if [ -L "$0" ]; then
  DIR="$(cd "$($(pwd)/$(readlink "$0"))" && pwd)"
else
  DIR="$(cd "$(dirname "$0")" && pwd)"
fi

mkdir -p ~/Library/LaunchAgents

cat << EOL > ~/Library/LaunchAgents/taal-client.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
  <key>Label</key>
  <string>taal-client</string>
  <key>ProgramArguments</key>
  <array>
    <string>$DIR/taal-client</string>
  </array>
  <key>UserName</key>
  <string>$(whoami)</string>
  <key>RunAtLoad</key>
  <true/>
  <key>KeepAlive</key>
  <true/>

  <key>WorkingDirectory</key>
  <string>$DIR</string>

  <key>StandardOutPath</key>
  <string>$DIR/taal-client.log</string>
  <key>StandardErrorPath</key>
  <string>$DIR/taal-client.log</string>

</dict>
</plist>
EOL

launchctl load ~/Library/LaunchAgents/taal-client.plist
