FROM golang

RUN apt-get update && apt-get install -y supervisor

COPY init/supervisord/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN mkdir -p /init/run /app/cmd /app/scripts
COPY cmd/simpleinferencer/simpleinferencer.exe /app/cmd/simpleinferencer.exe
COPY scripts/* /app/scripts/
RUN find /app/scripts/ -type f -exec chmod +x {} \;

EXPOSE 5001
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]