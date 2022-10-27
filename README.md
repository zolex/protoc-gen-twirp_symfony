# Protoc Plugin for Twirp Symfony

This project is based on [TwirPHP](https://github.com/twirphp/twirp) but generates code which seamlessly integrates with Symfony.

## Installation

Download prebuilt binaries for the protoc plugin from the [releases](https://github.com/zolex/protoc-gen-twirp_symfony) page.

Alternatively, you can use the following oneliner to install the plugin:

```bash
curl -Ls https://raw.githubusercontent.com/zolex/protoc-gen-twirp_symfony/master/install.sh | bash -s -- -b path/to/bin
```

## Generate Code

```bash
export PROTO_PATH=path/to/proto
export OUTPUT_PATH=path/to/output_dir
/usr/bin/protoc $(shell find ${PROTO_PATH} -type f -name "*.proto") \
    --php_out=${OUTPUT_PATH} \
    --twirp_symfony_out=${OUTPUT_PATH} \
    --plugin=protoc-gen-twirp_symfony=/usr/local/bin/protoc-gen-twirp_symfony \
    --proto_path ${PROTO_PATH} \
    --proto_path /usr/include
```

## Create a project

The generated code of this plugin needs the modix/twirp-bundle to run. View the [bundle docs](https://github.com/modix/TwirpBundle) for details on setting it up.

```bash
symfony new my-project
composer require modix/twirp-bundle
```
