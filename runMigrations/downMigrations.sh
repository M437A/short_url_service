for file in ../migrations/down/*.sql;
do
  psql -h localhost -U postgres -d short_url -a -f "$file";
done
