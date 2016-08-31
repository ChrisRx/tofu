# tofu


## Building

```shell
make
```

## Example

Running the volume and block stores with the embedded etcd metadata store:

```shell
bin/tofu server
```

Put a file:

```shell
bin/tofu client put <filename>
```

Get that file:

```shell
bin/tofu client get <filename>
```
