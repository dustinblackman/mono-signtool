# mono-signtool

This project is a drop in replacement to the issues with [Squirell.Windows](https://github.com/Squirrel/Squirrel.Windows/pull/505) and [windows-installer](https://github.com/electron/windows-installer/issues/27) that prevent signing windows executables on OSX/Linux machines with Wine. It works by replacing `signtool.exe`, and calls `osslsigncode` that lives in your OSX/Linux host from __within__ Wine. No forks, no messing with C#, just copy/paste.

`osslsigncode` must be located at `/usr/local/bin/osslsigncode` on your host system. The app supports dual signing `sha1` and `sha256` by default.

## Setup

An example setup on OSX using [electron windows-installer](https://github.com/electron/windows-installer/).

```sh
brew install osslsigncode
npm i --save-dev electron-winstaller
curl -Ls "https://github.com/dustinblackman/mono-signtool/releases/download/0.0.1/mono-signtool.tar.gz" | tar xz -C ./node_modules/electron-winstaller/vendor/
node setup.js
```

__setup.js__
```javascript
const winstaller = require('electron-winstaller');

winstaller.createWindowsInstaller({
  appDirectory: './app',
  outputDirectory: './dist',
  exe: `app.exe`,
  iconUrl: 'https://example.com/icon.ico',
  setupExe: 'setup.exe',
  certificateFile: './path/to/cert.p12',
  certificatePassword: 'super_secret_password'
});
```

## Build From Source

Built with Go 1.7.3.

```
git pull https://github.com/dustinblackman/mono-signtool.git
cd mono-signtool
make
```
