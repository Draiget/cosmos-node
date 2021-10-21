#!/usr/bin/env sh

TARGET_NODE_HOST=${TARGET_NODE_HOST:-192.168.137.141}
TARGET_SSH_KEY=${TARGET_SSH_KEY:-"~/.ssh/p2p-test"}

# shellcheck disable=SC2086
ssh -N -L 8022:127.0.0.1:22 ansible_tunnel@${TARGET_NODE_HOST} -i "${TARGET_SSH_KEY}" -v
