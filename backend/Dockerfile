FROM rust:1.63

WORKDIR /usr/src/backend
COPY . .

RUN cargo install --path .

EXPOSE 9090

CMD ["backend"]
