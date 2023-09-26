for file in ../migrations/up/*.sql;
do
  PGPASSWORD=qwerty psql -h localhost -U postgres -d short_url -a -f "$file";
done
