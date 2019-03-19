# Tableflip Plugin

This plugin acts just like the `\shrug` slash command, but its better. It adds the `(╯°□°)╯︵ ┻━┻` to your message.

**Supported Mattermost Server Versions: 5.2+**

## Installation

1. Go to the [releases page of this Github repository](https://github.com/cmeadows/mattermost-plugin-tableflip/releases) and download the latest release for your Mattermost server.
2. Upload this file in the Mattermost System Console under **System Console > Plugins > Management** to install the plugin. To learn more about how to upload a plugin, [see the documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).
3. Activate the plugin at **System Console > Plugins > Management**.

## Usage

`/tableflip` prints as `(╯°□°)╯︵ ┻━┻` a message. You can also add an additional message for the flip to be appended to with `/tableflip some angry message` (rendering `some angry message (╯°□°)╯︵ ┻━┻` as your message).

cd server && env GOOS=linux GOARCH=amd64 $(GO) build -o dist/plugin-linux-amd64;
cd server && env GOOS=darwin GOARCH=amd64 $(GO) build -o dist/plugin-darwin-amd64;
cd server && env GOOS=windows GOARCH=amd64 $(GO) build -o dist/plugin-windows-amd64.exe;
