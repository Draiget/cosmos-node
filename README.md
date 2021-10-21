### Initial setup steps

#### Users, keys and SSHd initial

```shell
useradd ansible
useradd ansible_tunnel 
mkdir -p /home/ansible/.ssh/
mkdir -p /home/ansible_tunnel/.ssh/
echo "${SSH_ANSIBLE_KEY}" | tee /home/ansible/.ssh/authorized_keys /home/ansible_tunnel/.ssh/authorized_keys > /dev/null

{
    echo "
AllowTcpForwarding no

AllowUsers ansible@127.0.0.1 ansible_tunnel

Match User ansible_tunnel
  AllowTcpForwarding yes
  PermitOpen 127.0.0.1:22
  ForceCommand echo 'Hello'
"
} >> /etc/ssh/sshd_config

groupadd disable2fa
usermod -aG disable2fa ansible
```

#### SeLinux

```shell
apt-get install selinux-basics selinux-policy-default auditd
selinux-activate
reboot
```

#### Two factor

`/etc/security/access-local.conf`:
```shell
+ : ALL : 192.168.2.0/24
- : ALL : ALL
```

###### Install

```shell
apt-get install -y libpam-google-authenticator

# echo 'auth required pam_google_authenticator.so' | cat - /etc/pam.d/sshd > temp && mv temp /etc/pam.d/sshd
sed -i.bak 's/ChallengeResponseAuthentication no/ChallengeResponseAuthentication yes/g' /etc/ssh/sshd_config

google-authenticator --time-based

service ssh restart
```

###### Pam ssh configuration
`/etc/pam.d/sshd`:
```shell
# Google Authenticator
#
auth [success=done default=ignore] pam_succeed_if.so user ingroup disable2fa
auth sufficient pam_google_authenticator.so nullok
```

###### Config main user
```
Do you want me to update your "~/.google_authenticator" file? (y/n) y

Do you want to disallow multiple uses of the same authentication
token? This restricts you to one login about every 30s, but it increases
your chances to notice or even prevent man-in-the-middle attacks (y/n) y

By default, a new token is generated every 30 seconds by the mobile app.
In order to compensate for possible time-skew between the client and the server,
we allow an extra token before and after the current time. This allows for a
time skew of up to 30 seconds between authentication server and client. If you
experience problems with poor time synchronization, you can increase the window
from its default size of 3 permitted codes (one previous code, the current
code, the next code) to 17 permitted codes (the 8 previous codes, the current
code, and the 8 next codes). This will permit for a time skew of up to 4 minutes
between client and server.
Do you want to do so? (y/n) y

If the computer that you are logging into isn't hardened against brute-force
login attempts, you can enable rate-limiting for the authentication module.
By default, this limits attackers to no more than 3 login attempts every 30s.
Do you want to enable rate-limiting? (y/n) y
```

#### Prometheus

```shell
docker run \
    --name prometheus \
    -p 9090:9090 \
    -v ~/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```