# days2xmaslet

It provides a blocklet that shows how many days until Christmas.

## Installation

```sh
make install BLOCKLETS_DIR=/path/to/blocklets CONFIG_DIR=/path/to/config
```

make sure that `BLOCKLETS_DIR` and `CONFIG_DIR` are set correctly and absolute paths are used.

## Uninstallation

```sh
make uninstall BLOCKLETS_DIR=/path/to/blocklets CONFIG_DIR=/path/to/config
```

## Configuration

editing `label` property into config file it is possible to change glyph shown in blocklet.

Actually the `label` rapresent the candy cane glyph into nerd font. https://www.nerdfonts.com/#home
