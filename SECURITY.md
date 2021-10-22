# Tweaks
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
{ \
  echo " \
dev.tty.ldisc_autoload=0 \
kernel.core_uses_pid=1 \
fs.protected_fifos=2 \
kernel.kptr_restrict=2 \
kernel.modules_disabled=1 \
kernel.sysrq=0 \
kernel.unprivileged_bpf_disabled=1 \
net.ipv4.conf.all.accept_redirects=0 \
net.ipv4.conf.all.log_martians=1 \
net.ipv4.conf.all.rp_filter=1 \
net.ipv4.conf.all.send_redirects=0 \
net.ipv4.conf.default.accept_redirects=0 \
net.ipv4.conf.default.accept_source_route=0 \
net.ipv4.conf.default.log_martians=1 \
net.ipv6.conf.all.accept_redirects=0 \
net.ipv6.conf.default.accept_redirects=0 \
kernel.yama.ptrace_scope=1 \
net.core.bpf_jit_harden=2 \
" \
} > /etc/sysctl.d/99-custom.conf
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

### Disable drivers
```shell
echo "blacklist usb-storage" >> /etc/modprobe.d/blacklist.conf
echo "blacklist firewire-core" >> /etc/modprobe.d/blacklist.conf
```

### Services and packages
```shell
apt-get install -y debsums apt-show-versions unattended-upgrades apt-listchanges 
```

### passwdqc
```shell
apt-get install -y passwdqc
```

### Audit daemon
```shell
# passwd changes monitor
auditctl -a exit,always -F path=/etc/passwd -F perm=wa

# `open()` monitor
auditctl -a exit,always -F arch=x86_64 -S open -F auid=80

# Get rule list 
auditctl -l

# Save rules above (example)
echo "w /etc/passwd -p wa" > /etc/audit/rules.d/audit.rules
echo "-a always,exit -F arch=b64 -S open -F auid=80" > /etc/audit/rules.d/audit.rules
```

### DNS
Nameserver configuration:
```shell
apt-get install -y resolvconf
echo "nameserver 127.0.0.1" >> /etc/resolvconf/resolv.conf.d/base
echo "nameserver 1.1.1.1" >> /etc/resolvconf/resolv.conf.d/base
echo "nameserver 8.8.8.8" >> /etc/resolvconf/resolv.conf.d/base
echo "nameserver 8.8.4.4" >> /etc/resolvconf/resolv.conf.d/base
resolvconf -u
```

DNSSec forwarder:
```shell
apt-get install -y bind9
# Update /etc/bind/named.conf.options
named-checkconf
```

### Purge compilers
```shell
apt-get --purge remove gcc
```

### Process accounting
> Process accounting is the method of recording and summarizing commands executed on Linux. The modern
Linux kernel is capable of keeping process accounting records for the commands being run, the user who
executed the command, the CPU time, and much more.

Install steps, reboot required:
```shell
apt-get install -y acct
accton on
touch /var/log/pacct
chown root /var/log/pacct
chmod 0644 /var/log/pacct
```

### Issue banner
```shell
{ \
  echo " \
******************************************************************** \
*                                                                  * \
* This system is for the use of authorized users only.  Usage of   * \
* this system may be monitored and recorded by system personnel.   * \
*                                                                  * \
* Anyone using this system expressly consents to such monitoring   * \
* and is advised that if such monitoring reveals possible          * \
* evidence of criminal activity, system personnel may provide the  * \
* evidence from such monitoring to law enforcement officials.      * \
*                                                                  * \
******************************************************************** \
  " \
} | tee /etc/issue /etc/issue.net > /dev/null
```

# Service changes
### dbus
`systemctl edit dbus`
```shell
[Service]
LockPersonality=true
MemoryDenyWriteExecute=true
NoNewPrivileges=yes
PrivateTmp=yes
PrivateDevices=true
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
RestrictSUIDSGID=true
RestrictRealtime=true
SystemCallArchitectures=native
SystemCallFilter=@system-service
```

# Malware scanner

```shell
apt-get install -y rkhunter 

sed -i.bak 's/^WEB_CMD=/#WEB_CMD=/g' /etc/rkhunter.conf
sed -i.bak 's/^UPDATE_MIRRORS=0/UPDATE_MIRRORS=1/g' /etc/rkhunter.conf
sed -i.bak 's/^MIRRORS_MODE=1/MIRRORS_MODE=0/g' /etc/rkhunter.conf
rkhunter -C
rkhunter --update

# Check using below command
rkhunter --check
```

# File integrity tool
```shell
apt-get install -y aide
echo "!/home/cosmos/.gaia/data" >> /etc/aide/aide.conf
aideinit
cp /var/lib/aide/aide.db{.new,}
```

# Apparmor
```shell
mkdir -p /etc/default/grub.d
echo 'GRUB_CMDLINE_LINUX_DEFAULT="$GRUB_CMDLINE_LINUX_DEFAULT apparmor=1 security=apparmor"' | \
  tee /etc/default/grub.d/apparmor.cfg
update-grub
reboot
```