## Generating ipv4 networks and ips

These commands will create directory **results-dir** and some files.
**results-dir/networks.txt** file will contain ipv4 networks.
All ip addresses from **results-dir/contains.txt** file will belong to these networks.

    $ go build -o bin/generate_ipv4 cmd/generate_ipv4/generate_ipv4.go
    $ bin/generate_ipv4 --networks-count 1000 --ips-count 10000 --results-dir ${NETWORKS_DIR}
