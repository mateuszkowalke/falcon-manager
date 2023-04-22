
# DEV WORKFLOW

Use `docker compose up` and `npm run dev`. This always has to run to generate appropriate graphQL and prisma schemas.
To make final checks before release run `docker compose up` - this simulates prod environment.

# CI

To release tag commit with v0.0.0 style tag, then push it using `git push origin v0.0.0`

# Running using docker image

docker run -d --name falcon-manager -e DATABASE_URL=postgresql://user:pass@pgdb:5432/falcon_manager --network=bridge --restart=always -p 3000:3000 mateuszkowalke/falcon-manager2:latest

# Production setting notes

Keep in mind any folders for static assets won't be created automatically when NODE_ENV=production.
This might also require setting appropriate permissions (`chmod -R 777`).
