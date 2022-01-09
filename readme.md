# Initial notes

## Remember to install diesel cli

```bash
cargo install diesel_cli
```
or
```bash
cargo install diesel_cli --no-default-features --features postgres
```
To use diesel cli tool it is needed to have the db running(see docker-compose info below) and DATABASE_URL env var configured.
Dev DATABASE_URL is the same as in dockerfile:
```bash
export DATABASE_URL=postgresql://testuser:testpass@pgdb:5432/falcon_manager?schema=public
```

## Docker-compose and dockerfiles are suitable only for development for now
