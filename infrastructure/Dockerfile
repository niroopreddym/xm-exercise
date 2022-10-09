FROM postgres
EXPOSE 5432
ENV POSTGRES_PASSWORD postgres 
ENV POSTGRES_DB companiesdb
COPY init.sql /docker-entrypoint-initdb.d/