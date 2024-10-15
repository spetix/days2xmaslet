CONFIG_DIR ?= $(HOME)/.config/i3blocks
BLOCKLETS_DIR ?= $(HOME)/.config/i3blocks/blocklets

install-blocklets:
	@echo "Installing blocklets to $(BLOCKLETS_DIR)"
	@mkdir -p $(BLOCKLETS_DIR)
	@cp -r blocklets/days2xmaslet $(BLOCKLETS_DIR)
	@chmod +x $(BLOCKLETS_DIR)/days2xmaslet

install-config:
	@echo "Installing config to $(CONFIG_DIR)"
	@mkdir -p $(CONFIG_DIR)
	@cat ./config | envsubst >> $(CONFIG_DIR)/config
	
install: install-config install-blocklets


uninstall-blocklets:
	@echo "Uninstalling blocklets from $(BLOCKLETS_DIR)"
	@rm $(BLOCKLETS_DIR)/days2xmaslet

uninstall-config:
	@echo "Uninstalling config from $(CONFIG_DIR)"
	@sed -i "/^\[days2xmaslet\]/,/^\[/ { /^\[days2xmaslet\]/d; /^\[/!d }" $(CONFIG_DIR)/config

uninstall: uninstall-blocklets uninstall-config