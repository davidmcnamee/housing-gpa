# Student Housing


## Local Dev Setup

First, install these:
* bazelisk
* docker
* k3d
* skaffold

### Set up a GPG key to gain access to secrets:

```
gpg --gen-key
gpg --armor --export your.email@address.com > public-key.gpg
```

Then send `public-key.gpg` to someone with secret access, and they will run:
```

```
