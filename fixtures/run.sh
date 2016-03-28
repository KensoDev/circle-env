import_date=$(date --date="2 day ago" +"%Y-%m-%d")
echo "Importing $import_date"
/usr/local/bin/gsutil -m cp -R gs://gogobot/results/<TEST_CODE>/$import_date /tmp/placer
/usr/bin/placer --database-name=test_database --username=test_db_username --hostname=<TEST_DB_HOST> --password=<TEST_DB_PASSWORD> --directory=/tmp/placer/$import_date
