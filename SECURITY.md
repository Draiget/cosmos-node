### Core dumps
**/etc/security/limits.conf**:
```shell
* hard core 0
* soft core 0
```

**/etc/sysctl.conf**:
```shell

fs.suid_dumpable=0
kernel.core_pattern=|/bin/false
```

### Loging password crypt rounds
**/etc/login.defs**:
```shell
SHA_CRYPT_MIN_ROUNDS 10000
SHA_CRYPT_MAX_ROUNDS 10000
```

### UMask
**/etc/login.defs**:
```shell
UMASK 027
```

### sysctl
```shell
sysctl -w dev.tty.ldisc_autoload=0
sysctl -w kernel.core_uses_pid=1
sysctl -w kernel.kptr_restrict=2
sysctl -w kernel.modules_disabled=1
sysctl -w kernel.sysrq=0
sysctl -w kernel.unprivileged_bpf_disabled=1
sysctl -w net.ipv4.conf.all.accept_redirects=0
sysctl -w net.ipv4.conf.all.log_martians=1
sysctl -w net.ipv4.conf.all.rp_filter=1
sysctl -w net.ipv4.conf.all.send_redirects=0
sysctl -w net.ipv4.conf.default.accept_redirects=0
sysctl -w net.ipv4.conf.default.accept_source_route=0
sysctl -w net.ipv4.conf.default.log_martians=1
sysctl -w net.ipv6.conf.all.accept_redirects=0
sysctl -w net.ipv6.conf.default.accept_redirects=0
sysctl -w kernel.yama.ptrace_scope=1
sysctl -w net.core.bpf_jit_harden=2
```

### Disable dccp
```shell
echo "install dccp /bin/true" > /etc/modprobe.d/disable-dccp.conf
echo "install sctp /bin/true" > /etc/modprobe.d/disable-sctp.conf
echo "install rds /bin/true" > /etc/modprobe.d/disable-rds.conf
echo "install tipc /bin/true" > /etc/modprobe.d/disable-tipc.conf
```

### Services
```shell
[Service]
CapabilityBoundingSet=
LockPersonality=true
MemoryDenyWriteExecute=true
NoNewPrivileges=yes
PrivateTmp=yes
PrivateDevices=true
PrivateUsers=true
ProtectSystem=strict
ProtectHome=yes
ProtectClock=true
ProtectControlGroups=true
ProtectKernelLogs=true
ProtectKernelModules=true
ProtectKernelTunables=true
ProtectProc=invisible
ProtectHostname=true
RemoveIPC=true
RestrictNamespaces=true
RestrictAddressFamilies=AF_INET AF_INET6
RestrictSUIDSGID=true
RestrictRealtime=true
SystemCallArchitectures=native
SystemCallFilter=@system-service
```

### Permissions
```shell
chmod 750 /etc/sudoers.d
```

### Sysstat
```shell
apt-get install -y sysstat
sed -i.bak 's/ENABLED="false"/ENABLED="true"/g' /etc/default/sysstat
systemctl start sysstat
```

#### Disable drivers
```shell
echo "blacklist usb-storage" >> /etc/modprobe.d/blacklist.conf
echo "blacklist firewire-core" >> /etc/modprobe.d/blacklist.conf
```