FROM alpine

WORKDIR /app

COPY .env /app/
COPY booking-room /app/

ENTRYPOINT [ "/app/booking-room" ]