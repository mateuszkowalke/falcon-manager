FROM rust:1.58

RUN cargo install cargo-watch

WORKDIR /usr/src/app

COPY Cargo.toml Cargo.lock ./

RUN cargo fetch

COPY src migrations ./

CMD ["cargo", "watch", "-x", "run"]
