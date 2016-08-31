# tofu


## Building

```
make
```

## Example

Running the volume and block stores with the embedded etcd metadata store:

```Shell
bin/tofu server
```

Put a file:

```Shell
bin/tofu client put <filename>
```

Get that file:

```Shell
bin/tofu client get <filename>
```
