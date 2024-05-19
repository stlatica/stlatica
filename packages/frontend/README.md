# Stlatica Frontend

## env

must install:

- make
- docker
- git lfs

## command reference

```
make

# or

make help
```

## work flow

### 1. develop...

You can run develop server with make & docker!

```
make dev // develop remix page

or

make ladle // for create components
```

Would you like to install any packages ?

```
make bash
> pnpm install xxxx
> exit
```

### 2. check & format code

```
make util/ci-checks
```

If there are no errors, you can create a pull request!
