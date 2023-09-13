# 🚀 modern-rm-go

🔒 Safely delete files with the option to recover them using a sleek CLI interface. This tool is fully compatible with `rm` and is inspired by [`rip`](https://github.com/nivekuil/rip). It's crafted with love in Go.

## 🌟 Features

- 🔄 Fully compatible with `rm` - supports all flags and arguments.
  - 🛡️ By default, `modern-rm` employs `rip` for file deletion.
    - 🔄 Use `rip -u` to undo a deletion.
  - 🛠️ If a flag exclusive to `rm` is provided, `modern-rm` will internally invoke `rm` for file deletion, ensuring full compatibility.

## 📦 Installation

🚧 The project is still in development and isn't packaged yet. You'll need to build it from source.

## 🛠️ Usage

Given its full compatibility with `rm`, it's recommended to set `modern-rm-go` as an alias for `rm`.

```sh

```sh
$ modern-rm-go -h

🗑️  modern-rm
🔒 Safely delete files with the option to recover them using a sleek CLI interface 💻
💯 Fully compatible with `rm` and built on `rip`.

Usage:
  modern-rm [files] [flags]

Flags:
  -d, --directory     📁 Remove directories (Invokes original rm).
  -f, --force         🚫 Ignore nonexistent files and arguments, never prompt
  -h, --help          📖 Show help.
  -i, --interactive   ❓ Prompt before every removal
  -I, --once          ❗ Prompt once before removing more than three files, or when removing recursively.
  -P, --overwrite     📝 Overwrite regular files before deleting them (Invokes original rm).
  -r, --recursive     Remove directories and their contents recursively
  -x, --same-fs       📌 Stay on the same filesystem (Invokes original rm).
  -W, --undelete      🔄 Attempt to undelete the named files (Invokes original rm).
  -u, --undo          🔙 Undo the last deletion.
  -v, --verbose       📊 Display detailed information about the deletion process.

Written by YeonGyu Kim (public.kim.yeon.gyu@gmail.com)
- https://github.com/code-yeongyu
```
