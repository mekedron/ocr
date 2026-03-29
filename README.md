# ocr

Capture a screenshot, OCR it with Apple Vision framework, and print the recognized text to stdout. Single binary, zero dependencies.

## Install

```bash
brew install mekedron/tap/ocr
```

Or build from source (requires macOS):

```bash
git clone https://github.com/mekedron/ocr.git
cd ocr
make build
# binary is at ./bin/ocr
```

## Usage

```bash
# Capture screenshot and OCR (English)
ocr

# Specify language(s)
ocr -l ru-RU
ocr -l en-US+ru-RU

# Copy to clipboard
ocr | pbcopy
ocr -l en-US+ru-RU | pbcopy

# List supported languages
ocr languages

# Show version
ocr -v
```

When you run `ocr`, the macOS screenshot selection UI appears. Select a region, and the recognized text is printed to stdout. Press Escape to cancel.

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-l, --lang` | `en-US` | OCR language(s), e.g. `en-US+ru-RU` |
| `-x, --silent` | | Do not play sounds |
| `-v, --version` | | Show version and exit |

## Languages

Apple Vision does not auto-detect languages. You must specify them with `-l`. To recognize multiple languages at once, join them with `+`:

```bash
ocr -l en-US+ru-RU
ocr -l en-US+de-DE+fr-FR
```

To see which languages are supported:

```bash
ocr languages
```

## Keyboard Shortcut

You can bind `ocr` to a global hotkey using [Hammerspoon](https://www.hammerspoon.org/) (`brew install hammerspoon`).

Add this to your `~/.hammerspoon/init.lua`:

```lua
-- Cmd+Shift+2: screenshot OCR → clipboard
hs.hotkey.bind({ "cmd", "shift" }, "2", function()
  local output = hs.execute("PATH=/opt/homebrew/bin:$PATH /opt/homebrew/bin/ocr -x -l en-US+ru-RU")
  if output and #output > 0 then
    hs.pasteboard.setContents(output)
  end
end)
```

Adjust the language (`-l en-US+ru-RU`) and keybinding to your preference. Use `-x` to suppress the screenshot sound.

## License

MIT
