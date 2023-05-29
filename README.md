# xpwd

xpwd is a command line password / pass phrase generator in the spirit of XKCD.

xpwd is written in Go. Its initial implementation uses the dictionary built in to Mac/Unix platforms (/usr/share/dict/). It also currently uses cli commands internally for parsing the dictionary.

Here is the original [XKCD Password Strength Comic](http://xkcd.com/936/)
![XKCD Password Strength Comic](http://imgs.xkcd.com/comics/password_strength.png)

## Usage

- `xpwd`

### Options

- `-d`, `--dictionary`
  - dictionary to use [`oxford` (default) | `mac`]
- `-c`, `--count`
  - number of words in the passphrase (default 4, minimum 1)
- `-m`, `--min` 
  - minimum word length (default 6)
- `-x`, `--max`
  - maximum word length (default 11, maximum 16)
  - setting maximum word length to 0, defaults to 16

### Examples
- Three word passphrase from the Mac system dictionary 
  - `xpwd -d mac -c 3`
- Passphrase with four words between 8 and 9 characters long
  - `xpwd -m 8 -x 9`
