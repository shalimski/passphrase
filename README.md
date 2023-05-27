# passphrase

[![Build and test](https://github.com/shalimski/passphrase/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/shalimski/passphrase/actions/workflows/go.yml)

Generator for passphrases.

usage:

```
    > passphrase --help
        -c int
            count of words. May be from 1 to 4. (default 4)
        -s string
            words separator. (default '')

    > passphrase
    TensionCorrectHonorableEntry

    > passphrase -c 2 -s '-'
    Bias-Respect

    > passphrase -s '777'
    Insanity777Drown777Antsy777Freedom
```
