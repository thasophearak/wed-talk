
- **Step 0**: Generate certs

Make sure to have https://github.com/FiloSottile/mkcert installed
```
cd certs && mkcert example.com
```

- **Step 1**: Start Caddy

Make sure to have https://github.com/caddyserver/caddy installed
```
cd ../ && caddy start
```

- **Step 2**: POST Caddy config

```
curl localhost:2019/load \
-X POST \
-H "Content-Type: application/json" \
-d @caddy.json
```

- **Step 3**: Start services

```
make
```
