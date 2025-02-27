# grocksdb, RocksDB wrapper for Go

[![](https://github.com/linxGnu/grocksdb/workflows/CI/badge.svg)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/linxGnu/grocksdb)](https://goreportcard.com/report/github.com/linxGnu/grocksdb)
[![Coverage Status](https://coveralls.io/repos/github/linxGnu/grocksdb/badge.svg?branch=master)](https://coveralls.io/github/linxGnu/grocksdb?branch=master)
[![godoc](https://img.shields.io/badge/docs-GoDoc-green.svg)](https://godoc.org/github.com/linxGnu/grocksdb)

This is a `Fork` from [tecbot/gorocksdb](https://github.com/tecbot/gorocksdb). I respect the author work and community contribution.
The `LICENSE` still remains as upstream.

Why I made a patched clone instead of PR:
- Supports almost C API (unlike upstream). Catching up with latest version of Rocksdb as promise.
- This fork contains `no defer` in codebase (my side project requires as less overhead as possible). This introduces loose
convention of how/when to free c-mem, thus break the rule of [tecbot/gorocksdb](https://github.com/tecbot/gorocksdb).

## Install

### Prerequisite 

- librocksdb
- libsnappy
- libz
- liblz4
- libzstd
- libbz2 (optional)

Please follow this guide: https://github.com/facebook/rocksdb/blob/master/INSTALL.md to build above libs.

### Build 

After installing both `rocksdb` and `grocksdb`, you can build your app using the following commands:

```bash
CGO_CFLAGS="-I/path/to/rocksdb/include" \
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" \
  go build
```

Or just:
```bash
go build // if prerequisites are in linker paths
```

If your rocksdb was linked with bz2:
```bash
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd -lbz2" \
  go build
```

#### Customize the build flags
Currently, the default build flags without specifying `CGO_LDFLAGS` or the corresponding environment variables are `-lrocksdb -pthread -lstdc++ -ldl -lm -lzstd -llz4 -lz -lsnappy`

If you want to customize the build flags:

1. Use `-tags grocksdb_clean_link` to create a cleaner set of flags and build it based on the cleaner flag. The base build flags after using the tag are `-lrocksdb -pthread -lstdc++ -ldl`.
```bash
CGO_LDFLAGS="-L/path/to/rocksdb -lzstd" go build -tags grocksdb_clean_link
```
2. Use `-tags grocksdb_no_link` to ignore the build flags provided by the library and build it fully based on the custom flags.
```bash
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lzstd -llz4" go build -tags grocksdb_clean_link
```

## Usage

See also: [doc](https://godoc.org/github.com/linxGnu/grocksdb)

## API Support

Almost C API, excepts:
- [ ] putv/mergev/deletev/delete_rangev
- [ ] compaction_filter/compaction_filter_factory/compaction_filter_context
- [ ] transactiondb_property_value/transactiondb_property_int
- [ ] optimistictransactiondb_property_value/optimistictransactiondb_property_int
