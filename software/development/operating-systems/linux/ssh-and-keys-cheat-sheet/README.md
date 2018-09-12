# SSH & KEYS CHEAT SHEET

`ssh and keys` _(secure socket shell) is a way to securely
communicate (via ssh keys) with a remote computer._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## KEYS

Its a way to receive encrypted data from a server.

### SEQUENCE

Its a way to receive encrypted data from a server.
Really just another way to 'login' to a server.

Sequence:

* You create both a public and private key.
* You give the remote computer your public key.
* You keep your private key and NEVER share.
* The server encrypts the data and sends it to you.
* You use your private key to decrypt it.

If the data went anywhere else, it would need your private
key to decrypt.

### GENERATE KEYS (ssh-keygen)

A way to generate public and private keys.

For example,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github"
```

* -t switch is the `rsa` type key (very popular).
* -b switch is the length in bits.
* -C is a comment for this key.

This will create two keys and place them in
`~/.ssh` with default names as `id_rsa` and `id_rsa.pub`.

I'm not a fan of using a paraphrase (password), because I'll
forget.  Its up to you if you want an extra level of protection if
by chance someone was on your computer being mischievous or if
someone stole your keys.  Hey, maybe you pushed your private key
to github by accident.

### KEY FINGERPRINT (ssh-keygen)

Sometimes it gets confusing with lots of keys, so you
can check your public key fingerprint to make sure your
server is using the right public key.

If you public key is located in `~/.ssh/id_rsa.pub`

```bash
ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub
```

### SSH - ADD YOUR PRIVATE KEY TO YOUR SSH AGENT (ssh-add)

Now add your key,

```bash
ssh-add ~/.ssh/id_rsa
```

## SSH

Another way to communicate to a remote computer without
traditional username/password `login`.

It stands for secure shell and uses port 22.

Remote will mean the ssh server and local will be the ssh client machine.

### SSH REMOTE SERVER SETUP

If you want to ssh into a machine, that machine must have a ssh server running.

To install on ubuntu, on the remover machine,

```bash
sudo apt-get install openssh-server
sudo systemctl start ssh
sudo systemctl enable ssh
```

Now from your local terminal (ssh client), push the public key
from your local machine to the remote ssh server,

```bash
ssh-copy-id username@<IP>
```

Or you can login to the remote ssh server machine and manually
add the public key to `.ssh/authorized_keys`.

### MAKE YOUR SSH SERVER SECURE - NO PASSWORDS

```bash
sudo nano /etc/ssh/sshd_config
```

And change the following to only allow ssh,

```
PermitRootLogin no
PasswordAuthentification no
```

Restart ssh service,

```
sudo systemctl reload sshd.service
```

If you ever want to add another ssh client, you need to
turn `PasswordAuthentification yes` and then upload
your ssh public key as above to `.ssh/authorized_keys`.

### SSH CLIENT - FROM YOUR TERMINAL - TO REMOTE LOGIN

From command line,

```go
ssh username@<IP>
```

Again, no password is needed.
