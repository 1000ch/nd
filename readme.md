# nd

Node.js version manager.

## Install

Install via `go` command.

```sh
$ go get github.com/1000ch/nd
```

Configure `PATH` in your `.bashrc` or `.zshrc`, etc.

```sh
export PATH=$HOME/.nd/bin:$PATH
```

## Usage

### `nd install <version>`

Install Node.js <version>.

### `nd uninstall <version>`

Uninstall Node.js <version>.

### `nd use <version>`

Activate Node.js <version>.

### `nd list` / `nd ls`

Show installed Node.js versions.

### `nd list-remote` / `nd ls-remote`

Show installable Node.js versions.

## License

[MIT](https://1000ch.mit-license.org) © [Shogo Sensui](https://github.com/1000ch)
