# Upstart

## Nginx 
- Nginx works purely as a reverse proxy
```bash
sudo nginx -s reload
```

- Be sure to create a symbolic link
```bash
sudo ln -s /etc/nginx/sites-available/spacezoo.tech /etc/nginx/sites-enabled/spacezoo.tech
```

## Systemctl Services
- Service for the golang program
```bash
sudo systemctl restart <service-name>
sudo systemctl status <service-name>
```