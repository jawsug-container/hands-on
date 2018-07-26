FROM docker:18.06.0-ce-dind

ENV DOCKER_COMPOSE_VERSION=1.22.0 \
    AWSCLI_VERSION=1.15.64 \
    JUPYTER_VERSION=1.0.0

RUN apk --no-cache add python3 git

# Install GNU libc
RUN apk --no-cache add curl \
    && version="2.27-r0" \
    && repo="https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${version}" \
    && curl --location --silent --show-error -O "${repo}/glibc-${version}.apk" \
    && curl --location --silent --show-error -O "${repo}/glibc-bin-${version}.apk" \
    && curl --location --silent --show-error -O "${repo}/glibc-i18n-${version}.apk" \
    && apk add --allow-untrusted "glibc-${version}.apk" \
       "glibc-bin-${version}.apk" "glibc-i18n-${version}.apk" \
    && /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib \
    && /usr/glibc-compat/bin/localedef -i en_US -f UTF-8 en_US.UTF-8 \
    && rm -rf "glibc-${version}.apk" "glibc-bin-${version}.apk" "glibc-i18n-${version}.apk"

# Install docker-compose
RUN repo="https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}" \
    && curl --location --silent --show-error -o /usr/local/bin/docker-compose \
       "${repo}/docker-compose-$(uname -s)-$(uname -m)" \
    && chmod +x /usr/local/bin/docker-compose

# Install AWS CLI
RUN apk --no-cache add groff less jq \
    && apk --no-cache add --virtual build-deps py3-pip \
    && pip3 install "awscli == ${AWSCLI_VERSION}" \
    && pip3 install yq \
    && find / -type d -name \__pycache__ -depth -exec rm -rf {} \; \
    && rm -rf /root/.cache \
    && apk del --purge -r build-deps

# Install Jupyter notebook
RUN apk --no-cache add bash tini \
    && apk --no-cache add --virtual build-deps build-base python3-dev \
    && pip3 install "jupyter == ${JUPYTER_VERSION}" \
    && pip3 install backcall bash_kernel \
    && python3 -m bash_kernel.install \
    && find / -type d -name tests -depth -exec rm -rf {} \; \
    && find / -type d -name \__pycache__ -depth -exec rm -rf {} \; \
    && rm -rf /root/.cache

# Configure git to use CodeCommit as git repositories
RUN git config --global credential.helper '!aws codecommit credential-helper $@' \
    && git config --global credential.UseHttpPath true

WORKDIR /root/notebook
ADD notebooks_config/jupyter_notebook.py /root/.jupyter/jupyter_notebook_config.py
ADD notebooks /root/notebook
ADD application /root/notebook/application
ADD infrastructure /root/notebook/infrastructure

ADD notebooks_config/entrypoint.sh /
RUN chmod +x /entrypoint.sh \
    && mkdir /root/config

ENTRYPOINT ["/entrypoint.sh"]
CMD ["tini", "--", "jupyter", "notebook"]
