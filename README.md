# days2xmaslet

It provides a blocklet that shows how many days until Christmas.

## Installation

* Copy blocklet to blocklets directory (e.g. `$HOME/.config/i3blocks/blocklets`)
* Add blocklet to i3blocks config (e.g. `$HOME/.config/i3blocks/config`)

Here's a sample configuration, `${BLOCKLETS_DIR}` has to be set in i3 configuration or env:
```ini
[days2xmaslet]
command = ${BLOCKLETS_DIR}/days2xmaslet
interval = 1
fg_color = #ff0000
bg_color =
label = 🎄
```

