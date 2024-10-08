#!/bin/bash
set -e
KERNEL_VERSION="6.8.0-45-generic"
IMAGE_NAME="my_rootfs.img"
MOUNT_POINT="/mnt/my_rootfs"

sudo apt-get update
sudo apt-get install -y qemu qemu-system-x86 linux-image-$KERNEL_VERSION

dd if=/dev/zero of=$IMAGE_NAME bs=1M count=10
mkfs.ext2 $IMAGE_NAME

mkdir -p $MOUNT_POINT
sudo mount -o loop $IMAGE_NAME $MOUNT_POINT

sudo sh -c "echo '#!/bin/sh' > $MOUNT_POINT/init"
sudo sh -c "echo 'echo \"hello world\"' >> $MOUNT_POINT/init"
sudo chmod +x $MOUNT_POINT/init

sudo umount $MOUNT_POINT

qemu-system-x86_64 -kernel /boot/vmlinuz-$KERNEL_VERSION -hda $IMAGE_NAME -append "console=ttyS0" -nographicc
