FROM golang:1.23.3-bookworm

ARG TARGETARCH

ARG NODE_VERSION="20.x"
ARG NODE_REPO_KEYRING="/usr/share/keyrings/nodesource.gpg"
RUN <<-EOR
  ARCH=$(dpkg --print-architecture)

  NODE_REPO_HOST="https://deb.nodesource.com"
  NODE_REPO_KEYRING="/usr/share/keyrings/nodesource.gpg"
  NODE_REPO_URL="${NODE_REPO_HOST}/node_${NODE_VERSION}"
  NODE_GPG_SRC="${NODE_REPO_HOST}/gpgkey/nodesource-repo.gpg.key"
  NODE_GPG_TGT="${NODE_REPO_KEYRING}"

  curl --fail --location --show-error --silent ${NODE_GPG_SRC} \
    | gpg --dearmor --output ${NODE_GPG_TGT}

  echo "deb [arch=${ARCH} signed-by=${NODE_REPO_KEYRING}] ${NODE_REPO_URL} nodistro main" \
    > /etc/apt/sources.list.d/nodesource.list
  chmod 644 ${NODE_REPO_KEYRING}

  echo "Package: nsolid" | tee /etc/apt/preferences.d/nsolid > /dev/null
  echo "Pin: origin deb.nodesource.com" | tee -a /etc/apt/preferences.d/nsolid > /dev/null
  echo "Pin-Priority: 600" | tee -a /etc/apt/preferences.d/nsolid > /dev/null

  echo "Package: nodejs" | tee /etc/apt/preferences.d/nodejs > /dev/null
  echo "Pin: origin deb.nodesource.com" | tee -a /etc/apt/preferences.d/nodejs > /dev/null
  echo "Pin-Priority: 600" | tee -a /etc/apt/preferences.d/nodejs > /dev/nul

  DEBIAN_FRONTEND=noninteractive apt-get --assume-yes --quiet --quiet update
  DEBIAN_FRONTEND=noninteractive apt-get --assume-yes --quiet --quiet install nodejs
  npm install -g cdk8s-cli
EOR

ENTRYPOINT [ "/bin/bash" ]