# Twirp Symfony: Port of [TwirPHP](https://github.com/twirphp/twirp) for Symfony

## Installation

Download prebuilt binaries for the protoc plugin from the [releases](https://github.com/zolex/protoc-gen-twirp_symfony) page.


## Documentation

See the [official documentation](https://twirphp.github.io/).


## Development

Install dependencies:

```shell
go mod download
composer install
```

If you change something in the protoc plugin, regenerate the examples:

```shell
make generate
```

When all coding and testing is done, please run the test suite:

```shell
make check
```

For the best developer experience, install [Nix](https://builtwithnix.org/) and [direnv](https://direnv.net/).

Alternatively, install the following dependencies manually:

- Go >=1.17
- PHP 8.x
- Composer 2.x

Then run `make deps` to install the remaining dependencies.


## Security

If you discover any security related issues, please contact us at [twirphp@sagikazarmark.dev](mailto:twirphp@sagikazarmark.dev).


## License

The project is licensed under the [MIT License](LICENSE).

The original Twirp library is licensed under the Apache 2.0 License.
