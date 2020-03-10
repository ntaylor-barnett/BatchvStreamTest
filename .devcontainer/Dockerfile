#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.12

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Or your actual UID, GID on Linux if not the default 1000
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git procps lsb-release \
    #
    # Install gocode-gomod
    && go get -x -d github.com/stamblerre/gocode 2>&1 \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/ 
    #
    # Install Go tools
    # && go get -u -v \
    #     github.com/mdempsky/gocode \
    #     github.com/uudashr/gopkgs/cmd/gopkgs \
    #     github.com/ramya-rao-a/go-outline \
    #     github.com/acroca/go-symbols \
    #     github.com/godoctor/godoctor \
    #     golang.org/x/tools/cmd/guru \
    #     golang.org/x/tools/cmd/gorename \
    #     github.com/rogpeppe/godef \
    #     github.com/zmb3/gogetdoc \
    #     github.com/haya14busa/goplay/cmd/goplay \
    #     github.com/sqs/goreturns \
    #     github.com/josharian/impl \
    #     github.com/davidrjenni/reftools/cmd/fillstruct \
    #     github.com/fatih/gomodifytags \
    #     github.com/cweill/gotests/... \
    #     golang.org/x/tools/cmd/goimports \
    #     golang.org/x/lint/golint \
    #     golang.org/x/tools/cmd/gopls \
    #     github.com/alecthomas/gometalinter \
    #     honnef.co/go/tools/... \
    #     github.com/golangci/golangci-lint/cmd/golangci-lint \
    #     github.com/mgechev/revive \
    #     github.com/derekparker/delve/cmd/dlv 2>&1 
#RUN go get -u -v golang.org/x/tools/cmd/gopls 2>&1

RUN apt update \
    && apt-get update \
    # Install other needed tools
    && apt-get -y install net-tools openssh-client curl httpie jq make ruby dos2unix git python-pip binutils bison gcc build-essential \
    && pip install --user awscli \
    #
    # Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
    && groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    # [Optional] Add sudo support
    # && apt-get install -y sudo \
    # && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    # && chmod 0440 /etc/sudoers.d/$USERNAME \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

RUN git config --global url.https://github.com/apache/thrift.insteadof https://git.apache.org/thrift.git
RUN git config --global url.git@github.com:.insteadOf https://github.com/

RUN echo 'PATH=$PATH:$HOME/bin' >> /root/.bashrc 

RUN apt-get update \
    && apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common lsb-release \
    && curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | (OUT=$(apt-key add - 2>&1) || echo $OUT) \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" \
    && apt-get update \
    && apt-get install -y docker-ce-cli \
    #
    # Install Docker Compose
    && curl -sSL "https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
    && chmod +x /usr/local/bin/docker-compose

RUN apt-get -y install awscli

RUN curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
    && chmod +x /usr/local/bin/docker-compose \
    && docker-compose --version

RUN apt-get update && apt-get install -y apt-transport-https \
    && curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    && echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list \
    && apt-get update \
    && apt-get install -y kubectl

RUN curl -LO https://git.io/get_helm.sh \
    && chmod 700 get_helm.sh \
    && ./get_helm.sh

# RUN go get -u -v \
#         github.com/mdempsky/gocode \
#         github.com/uudashr/gopkgs/cmd/gopkgs \
#         github.com/ramya-rao-a/go-outline \
#         github.com/acroca/go-symbols \
#         github.com/godoctor/godoctor \
#         golang.org/x/tools/cmd/guru \
#         golang.org/x/tools/cmd/gorename \
#         github.com/rogpeppe/godef \
#         github.com/zmb3/gogetdoc \
#         github.com/haya14busa/goplay/cmd/goplay \
#         github.com/sqs/goreturns \
#         github.com/josharian/impl \
#         github.com/davidrjenni/reftools/cmd/fillstruct \
#         github.com/fatih/gomodifytags \
#         github.com/cweill/gotests/... \
#         golang.org/x/tools/cmd/goimports \
#         golang.org/x/lint/golint \
#         golang.org/x/tools/cmd/gopls \
#         github.com/alecthomas/gometalinter \
#         honnef.co/go/tools/... \
#         github.com/golangci/golangci-lint/cmd/golangci-lint \
#         github.com/mgechev/revive \
#         github.com/derekparker/delve/cmd/dlv
# ENV GOPATH /go
# ENV GO111MODULE=on
# RUN mkdir ~/.ssh && \
#     chmod 0700 ~/.ssh
#     # mkdir ~/.ssh/known_hosts && \
#     # chmod 0700 ~/.ssh/known_hosts
# RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
# RUN go get -u -v golang.org/x/tools/gopls 2>&1 
RUN cd /tmp \
    && PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip \
    && curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \
    && rm -f $PROTOC_ZIP