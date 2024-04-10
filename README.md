# Namecheap Update DNS

Updates your DNS record to reflect your current IP.

### Automatically Update DNS Records

ISPs, especially for residential, may change your IP address without notifying you. In use cases, such as a home server, this leaves you unable to reach your destination without updating you records in Namecheap's advanced DNS after finding out what your new IP address is. This program periodically checks your host records versus your current ip and updates them accordingly.

## Deployment

### Prerequisites

- Existing A or CNAME Record on Namecheap.com for your domain name.
- Enable and retrieve your API token. Add target server to whitelist (Logged in to Namecheap -> Profile -> Tools -> API Access).
- Either download binary for your system or compile from source (instructions below).

### Build & Run from source

```
git clone https://github.com/jqwez/namecheap-update-dns
cd namecheap-update-dns
make
cd bin
nm_updatedns config edit
# Fill out CLI form with your configuration details
nm_updatedns run
```

Just add to a service and you're all set!

## Status

The app currently fits my needs and is up and running. However, to deploy, it requires manually setting up a persistent systemd service.

### TODO

- App periodically runs itself
- Automatically create service
