Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/nilkonto/terraform-provider-$PROVIDER_NAME`

```sh
$ mkdir -p $GOPATH/src/github.com/nilkonto; cd $GOPATH/src/github.com/nilkonto
$ git clone git@github.com:nilkonto/terraform-provider-$PROVIDER_NAME
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/nilkonto/terraform-provider-$PROVIDER_NAME
$ make build
```

Using the provider
----------------------

At present Terraform can automatically install only the providers distributed by HashiCorp. Third-party providers can be manually installed by placing their plugin executables in one of the following locations depending on the host operating system:

- On Windows, in the sub-path terraform.d/plugins beneath your user's "Application Data" directory.
- On all other systems, in the sub-path .terraform.d/plugins in your user's home directory.
terraform init will search this directory for additional plugins during plugin initialization.

The naming scheme for provider plugins is terraform-provider-NAME_vX.Y.Z, and Terraform uses the name to understand the name and version of a particular provider binary. Third-party plugins will often be distributed with an appropriate filename already set in the distribution archive so that it can be extracted directly into the plugin directory described above.


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-$PROVIDER_NAME
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
