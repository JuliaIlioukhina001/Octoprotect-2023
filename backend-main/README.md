# Octoprotect Backend

[API Documentation](./doc/API.md)

This backend enables basic communication between `Nexus` and `User`, secured by pairing procedure.

## Deployment

Pre-built Docker Images are available at [Gitlab Registry](https://git.uwaterloo.ca/octoprotect/backend/container_registry/217).

### Use Docker

```shell
docker run -d -p <host port>:8080 \
  --restart always \
  -e DB_TYPE=sqlite \
  -e SQLITE_PATH=/data/app.db \
  -e HTTP_LISTEN_ADDR=0.0.0.0:8080 \
  -e JWT_SECRET_KEY=<jwt secret key for nexus auth> \
  -e PUSHOVER_TOKEN=<pushover app token> \
  -e MAX_RETRY_NOTIFY=3 \
  -e ADMIN_TOKEN_BCRYPT_BASE64=<base64 of bcrypt hashed admin token> \
  -v octoprotect-data:/data \
  --name octoprotect-backend git.uwaterloo.ca:5050/octoprotect/backend:v0.2
```

### Use Kubernetes

Since we haven't implemented connection sync between multiple instances, we can only deploy the app as a pod.

The basic configurations are available in [kubernetes](./kubernetes) folder, and can refer the `deploy_review` section in [Gitlab CI Config](.gitlab-ci.yml).