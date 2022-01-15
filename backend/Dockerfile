FROM rust:1.58

RUN cargo install cargo-watch

WORKDIR /usr/src/app

COPY . .

ENV ROCKET_ADDRESS=0.0.0.0
ENV ROCKET_PORT=8000

CMD ["cargo", "watch", "-x", "run"]
