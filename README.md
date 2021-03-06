# Thixo: Template functions for Go templates

The Go language comes with a [built-in template
language](http://golang.org/pkg/text/template/), but not
very many template functions. Thixo is a library that provides more than 100 commonly
used template functions.

It is inspired by the template functions found in
[Twig](http://twig.sensiolabs.org/documentation) and in various
JavaScript libraries, such as [underscore.js](http://underscorejs.org/).

## Usage

**Go developers**: If you'd like to include Thixo as a library in your program,
our API documentation is available [at GoDoc.org](http://godoc.org/github.com/nikgalushko/thixo).

For standard usage, read on.

### Load the Thixo library

To load the Thixo `FuncMap`:

```go

import (
  "github.com/nikgalushko/thixo"
  "html/template"
)

// This example illustrates that the FuncMap *must* be set before the
// templates themselves are loaded.
tpl := template.Must(
  template.New("base").Funcs(thixo.FuncMap()).ParseGlob("*.html")
)


```

### Calling the functions inside of templates

By convention, all functions are lowercase. This seems to follow the Go
idiom for template functions (as opposed to template methods, which are
TitleCase). For example, this:

```
{{ "hello!" | upper | repeat 5 }}
```

produces this:

```
HELLO!HELLO!HELLO!HELLO!HELLO!
```

## Principles Driving Our Function Selection

We followed these principles to decide which functions to add and how to implement them:

- Use template functions to build layout. The following
  types of operations are within the domain of template functions:
  - Formatting
  - Layout
  - Simple type conversions
  - Utilities that assist in handling common formatting and layout needs (e.g. arithmetic)
- Template functions should not return errors unless there is no way to print
  a sensible value. For example, converting a string to an integer should not
  produce an error if conversion fails. Instead, it should display a default
  value.
- Simple math is necessary for grid layouts, pagers, and so on. Complex math
  (anything other than arithmetic) should be done outside of templates.
- Template functions only deal with the data passed into them. They never retrieve
  data from a source.
- Finally, do not override core Go template functions.

Based on https://github.com/Masterminds/sprig.
