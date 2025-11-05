# Local TLS Certificates

Store development TLS key/cert files here.

- `dev.crt`: PEM-encoded certificate
- `dev.key`: PEM-encoded private key

To generate a self-signed pair for localhost:

```powershell
openssl req -x509 -nodes -days 365 -newkey rsa:2048 `
  -keyout certs/dev.key `
  -out certs/dev.crt `
  -subj "/CN=localhost"
```

Consider using [mkcert](https://github.com/FiloSottile/mkcert) for trust without warnings.

Once generated, access the app at `https://localhost:8443`.
