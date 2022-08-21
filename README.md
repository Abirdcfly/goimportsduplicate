**Do it like https://github.com/kubernetes/kubernetes/pull/111101/files does.**

---

A linters to detect duplicate package imports, which are usually caused by git cherry-picking or other reasons, This greatly interferes with the source code reading.


for example:
```go
package main
import (
    "fmt"
    gofmt "fmt"  // duplicate import
)
....
```
