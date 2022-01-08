FROM rust:1.57

WORKDIR /usr/src/app

COPY . .

RUN cargo install --path .

ENV ROCKET_ADDRESS=0.0.0.0
ENV ROCKET_PORT=8000

CMD ["app"]
