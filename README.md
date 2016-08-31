# tofu


## Building

```
make
```

## Example

Running the volume and block stores with the embedded etcd metadata store:

```
bin/tofu server
```

Put a file:

```
bin/tofu client put <filename>
```

Get that file:

```
bin/tofu client get <filename>
```
