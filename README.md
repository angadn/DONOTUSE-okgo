# OKGO
Error checking can become tedious work in Golang, with tons of nested if-blocks, or worse, jumpy goto-like code with `return` statements that terminate functions prematurely on errors. **OKGO** is a microlibrary that proposes a funkier solution.

### Example Usage
```
...
var (
    email      Email
    uniqueSpec UniqueEmailSpecification
    request    UserRequest
    err        error
)

okgo.NewOKGO().On(func() bool {
    email, err = NewEmail(str)
    return err == nil
}).On(func() bool {
    return uniqueSpec.Verify(email)
}).On(func() bool {
    err = json.NewDecoder(r).Decode(&request)
    return err == nil
}).Run()

if err == nil {
    request.Resolve()
} else {
    log.Print(err)
}
...
```

