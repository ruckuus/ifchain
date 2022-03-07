## Interface chaining example
This to demonstrate interface chaining in Go.

![ifchain](https://user-images.githubusercontent.com/53867/156981370-9c78dadb-5627-45bf-8aad-b3d5df1062e2.png)


What has been done:
1. Define one interface as a [baseline](https://github.com/ruckuus/ifchain/blob/main/internal/ifchain/services.go#L10-L12).
2. The [structs](https://github.com/ruckuus/ifchain/blob/main/internal/ifchain/services.go#L14-L35) represent the layers needed, in this example: implementaion (Impl), validation (Validator), cache (Cache), and Orm (dealing with DB).
