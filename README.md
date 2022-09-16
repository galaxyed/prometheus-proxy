# prometheus-proxy
The purposes are:

- Limit showing series by Labels
- Protect Metrics
- Prevent Metrics Comflicted while same label but different Instances

```
docker build . -t prometheus-proxy:dev
docker run -d -it -p 8000:8000 prometheus-proxy:dev
```
