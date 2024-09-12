FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -v -o sudoku_solver cmd/main.go

RUN apk update \
    && apk add --no-cache \
      bash \
      git \
      zsh \
      tzdata \
      ca-certificates \
      curl \
      openssl \
      gcompat \
    && rm -rf /tmp/* /var/cache/apk/*

ARG LINUX_USER_ID

RUN addgroup --gid $LINUX_USER_ID docker \
    && adduser --uid $LINUX_USER_ID --ingroup docker --home /home/docker --shell /bin/zsh --disabled-password --gecos "" docker

USER docker

RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh
RUN echo 'export ZSH=/home/docker/.oh-my-zsh' > ~/.zshrc \
    && echo 'ZSH_THEME="simple"' >> ~/.zshrc \
    && echo 'plugins=(npm)' >> ~/.zshrc \
    && echo 'source $ZSH/oh-my-zsh.sh' >> ~/.zshrc \
    && mkdir -p ~/.oh-my-zsh/themes \
    && echo 'PROMPT="%{$fg_bold[blue]%}%(!.%1~.%~)%{$fg_bold[yellow]%} >%{$reset_color%} "' > ~/.oh-my-zsh/themes/simple.zsh-theme

CMD ["tail", "-f", "/dev/null"]
