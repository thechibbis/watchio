FROM rust:1.88.0-bookworm as builder_base
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler
RUN cargo install cargo-chef

FROM builder_base as planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

FROM builder_base as builder
ARG SERVICE_NAME

COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json

COPY . .
RUN cargo build --release --bin ${SERVICE_NAME}

FROM debian:bookworm-slim as runtime
ARG SERVICE_NAME

COPY --from=builder /app/target/release/${SERVICE_NAME} /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/${SERVICE_NAME}" ]
