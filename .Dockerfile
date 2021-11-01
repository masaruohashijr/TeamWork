FROM mongo
WORKDIR /app
COPY ./teams ./teams
COPY ./entrypoint.sh ./entrypoint.sh
RUN chmod 755 /app/entrypoint.sh
RUN chmod 755 /app/zarbat_tester
EXPOSE 5000
ENTRYPOINT ["sh","./entrypoint.sh"]