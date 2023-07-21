

run:



help:
	@echo ""
	@echo "make i   | make init      : 初始化项目"
	@echo ""

i:init
init:
	sh shell/install-goctl.sh
	sh shell/install-gopls.sh
	sh shell/install-vscode-extension.sh
