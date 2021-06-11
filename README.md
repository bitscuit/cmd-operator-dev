# CertManagerDeployment Operator

This repository is for hacking. When this project nears finalization, it will be moved to a different, more-permanent home with a fresh git history.

# Running demo

### Pre-requisites
[Pipe Viewer](https://catonmat.net/unix-utilities-pipe-viewer)
  - `sudo apt install pv`

The demo script will use [demo-magic](https://github.com/paxtonhare/demo-magic) to nicely present execution of the commands. The demo-magic script is already included in this repo, but you will need the [Pipe Viewer](https://catonmat.net/unix-utilities-pipe-viewer) tool to use it.

Alternatively, you may choose to inspect the `hack/bring-up.sh` and `hack/tear-down.sh` scripts to see what commands are being run, and run those manually.

### Running

You should not need to rebuild the image or push the image to a registry. The operator deployment is currently set to deploy the image from `quay.io/bitscuit/cmd-operator-dev-controller:latest`, which should work fine.

```console
make install
make demo
```

This will deploy the operator using `oc` CLI, and call the `hack/bring-up.sh` script to create the Issuer, alpha1 Certificate, and operand. Since the script is using `demo-magic`, the commands will be printed out first and then executed. You need to press enter after each command has been executed to continue.

During the demo, the operator will automatically create a v1 Certificate resource matching the contents of the alpha1 Certificate
  - Note that while the entire alpha1 Certificate spec is available, the operator currently only converts some of the fields

When the operand is deployed, a secret for the v1 Certificate will be created, thus the core flow of cert-manager has been accomplished.

### Clean up

```console
make undemo
```

This will call the `hack/tear-down.sh` script, which deletes everything that was created during the demo. Similarly, it uses demo-magic, so you need to press enter after each command has been executed.
