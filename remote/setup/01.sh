#!/bin/bash
set -eu

# ==================================================================================== #
# VARIABLES
# ==================================================================================== #

TIMEZONE=Europe/Warsaw

USERNAME=movieapp

read -p "Enter password for movieapp DB user: " DB_PASSWORD

export LC_ALL=en_US.UTF-8 

# ==================================================================================== #
# SCRIPT LOGIC
# ==================================================================================== #

add-apt-repository --yes universe

apt update
apt --yes -o Dpkg::Options::="--force-confnew" upgrade

timedatectl set-timezone ${TIMEZONE}
apt --yes install locales-all

useradd --create-home --shell "/bin/bash" --groups sudo "${USERNAME}"

passwd --delete "${USERNAME}"
chage --lastday 0 "${USERNAME}"

rsync --archive --chown=${USERNAME}:${USERNAME} /root/.ssh /home/${USERNAME}

ufw allow 22
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable

apt --yes install fail2ban

curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 /usr/local/bin/migrate

apt --yes install postgresql

sudo -i -u postgres psql -c "CREATE DATABASE movieapp"
sudo -i -u postgres psql -d movieapp -c "CREATE EXTENSION IF NOT EXISTS citext"
sudo -i -u postgres psql -d movieapp -c "CREATE ROLE movieapp WITH LOGIN PASSWORD '${DB_PASSWORD}'"


echo "MOVIEAPP_DB_DSN='postgres://movieapp:${DB_PASSWORD}@localhost/movieapp'" >> /etc/environment

apt --yes install -y debian-keyring debian-archive-keyring apt-transport-https
curl -L https://dl.cloudsmith.io/public/caddy/stable/gpg.key | sudo apt-key add -
curl -L https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt | sudo tee -a /etc/apt/sources.list.d/caddy-stable.list
apt update
apt --yes install caddy

echo "Script complete! Rebooting..."
reboot