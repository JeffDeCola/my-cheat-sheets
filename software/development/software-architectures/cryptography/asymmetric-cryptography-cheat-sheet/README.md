# ASYMMETRIC CRYPTOGRAPHY CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_Asymmetric cryptography uses a public key to encrypt a message
and a private key to decrypt._

Tables of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/asymmetric-cryptography-cheat-sheet#overview)

Documentation and Reference

* [asymmetric-cryptography](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/asymmetric-cryptography-cheat-sheet)
  **(You are here)**
* [hashing](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/hashing-cheat-sheet)
* [symmetric-cryptography](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/cryptography/symmetric-cryptography-cheat-sheet)
* [my-go-examples](https://github.com/JeffDeCola/my-go-examples#cryptography)
  on cryptography

## OVERVIEW

Cryptography is the art of changing a message from a readable format,
referred to as `plaintext`, into an unreadable one, or `ciphertext`.

Asymmetric cryptography uses two separate keys,

* A public key is used to encrypt the messages
* A private key is used to then decrypt the message

This is used all over the web for ssl, tls, pgp, ssh, etc...

This illustration from my go example
rsa-asymmetric-cryptography (GET LINK)
may help,

![IMAGE - rsa-asymmetric-cryptography.jpg - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/rsa-asymmetric-cryptography.jpg)
