ARG PMM_SERVER_IMAGE="perconalab/pmm-server:dev-latest"
FROM $PMM_SERVER_IMAGE

ARG PMM_SERVER_IMAGE
ARG GO_VERSION="1.18.5"

RUN curl -L https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz -o go.tar.gz && \
    tar -xzf go.tar.gz

RUN echo "Building with: GO: ${GO_VERSION}, PMM: ${PMM_SERVER_IMAGE}"

ENV GOPATH="/root/go"
ENV PATH="${GOPATH}/bin:${PATH}"

RUN mkdir -p $GOPATH/src/github.com/percona/pmm
WORKDIR $GOPATH/src/github.com/percona/pmm

COPY ./ ./
# setup.py, uses a task from Makefile.devcontainer but expect it to be in the fault file Makefile
COPY ./Makefile.devcontainer ./Makefile

RUN python ./.devcontainer/setup.py
