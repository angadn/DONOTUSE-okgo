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

okgo.NewOKGO().On(func() error {
    email, err = NewEmail(str)
    return err
}).On(func() error {
    return uniqueSpec.Verify(email)
}).On(func() error {
    return json.NewDecoder(r).Decode(&request)
}).Run()

if err == nil {
    request.Resolve()
} else {
    log.Print(err)
}
...
```

