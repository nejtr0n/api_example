# README

### Staring
execute commands
```
make start
make all
```

### Testing
execute tests
```
make test
```

### Testing fetch csv
execute command
```
./src/bin/fetch_client --fetch_address=127.0.0.1:6000 --fetch_file=http://testData/test_1.csv
```

### Testing list products
execute command
```
./src/bin/list_client --fetch_address=127.0.0.1:6000 --offset=0 --limit=1 --field=1 --sort=1
```
(**field** is ***Enum*** number from *proto* file, **sort** types: *1* is asc, *-1* is desc)