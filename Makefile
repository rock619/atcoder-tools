build:
	poetry build

install: build
	pip3 install dist/atcoder-tools-2.12.0.tar.gz
