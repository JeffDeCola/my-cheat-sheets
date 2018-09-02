# SSH AND KEYS CHEAT SHEET

`ssh and keys` _??????????????????????????????._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## WHAT ARE KEYS USED FOR

Its a way to receive encrypted data from a server.
Really just another way to 'login' to a server.

Sequence:

* You create both a public and private key.
* You give the server your public key.
* The server encrypts the data and sends it to you.
* You use your private key to decrypt it.

If the data went anywhere else, it would need your private
key to decrypt.

## GENERATE KEY (ssh-keygen)

A way to generate keys.

For example,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github"
```

* -t switch is the `rsa` type key (very popular).
* -b switch is the length in bits.
* -C is a comment for this key

This will create two keys and place them in
`~/.ssh` with default names as `id_rsa` and `id_rsa.pub`.

I'm not a fan of using a paraphrase (password), because I'll
forget.  Its up to you if you want an extra level of protection if
by chance someone was on your computer being mischievous or if
someone stole your keys.  Hey, maybe you pushed to github by accident.

## KEY FINGERPRINT (ssh-keygen)

Sometimes it gets confusing with lots of keys, so you
can check your public key fingerprint to make sure you're
server is using the right public key.

If you public key is located in `~/.ssh/id_rsa.pub`

```bash
ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub
```

## WHAT IS SSH

Another way to communicate to a remote server without
traditional `login`.

It stands for secure shell and uses port 22.

## SSH - ADD YOUR PRIVATE KEY TO YOUR SSH AGENT (ssh-add)

Now add your key,

```bash
ssh-add ~/.ssh/id_rsa
```