FROM redis:7-alpine

# set timezone
ENV TZ=Asia/Tehran
RUN ls /usr/share/zoneinfo && cp /usr/share/zoneinfo/Iran /etc/localtime && echo "Asia/Tehran" >  /etc/timezone

COPY ./redis.conf /usr/local/etc/redis/redis.conf

CMD ["redis-server","/usr/local/etc/redis/redis.conf"]