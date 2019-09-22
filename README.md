# Firease

A cli to interact with firebase realtime database

## Installation

Checout the code and build

```bash
go build .
```

## Usage and Commands

### add document to collection
```bash
Usage:
  firease database add [json data] [flags]

Flags:
  -c, --col string   collection name
  -h, --help         help for add
  --nid   string   new document id [optional]

Global Flags:
      --safile string   Service account authentication json file

```
### read document from collection
```bash
Usage:
  firease database get [document url] [flags]

Flags:
  -h, --help                  help for get
  -u, --path  string   document path to fetch by id

Global Flags:
      --safile string   Service account authentication json file

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
