# TrenchBoot user init

u-root's built-in init command will call uinit if it exits.

The `uinit.go` in this dir mounts storage and kexec to Xen using provided
kernel and rootfs for dom0.

To build the example:

```shell
    # make some changes to uinit.go in examples/uinit/uinit.go
    u-root -format=cpio -build=bb -o initramfs.cpio \
        ./cmds/* \
        github.com/3mdeb/uinit
```

This code was designed for PSEC 2019 demo. You can get more information about
it [here](https://www.platformsecuritysummit.com/#krol).

