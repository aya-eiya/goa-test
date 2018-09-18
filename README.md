# Error patterns for goa ^v1.4.0

Test project for checking some errors for goa(^1.4.0)

1. ArrayOf for Payload causes defined but not be used error
2. non-required Integer Type for Payload causes strConv.Itoa(int) type mismatch error

## Run test

just run test.sh

```
$> sh test.sh
```

and the shell will show the error

```
# github.com/aya-eiya/goa-test/apps/app
apps/app/controllers.go:72:2: rawArrs declared and not used
# github.com/aya-eiya/goa-test/apps/client
apps/client/my_resorce.go:63:28: cannot use payload.IntVal (type *int) as type int in argument to strconv.Itoa
```

