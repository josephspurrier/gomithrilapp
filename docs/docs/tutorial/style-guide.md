---
id: style-guide
title: Style Guide
---

As you work this these docs, we'll follow a few conventions.

## Hyperlink

Tools and concepts that have their own websites will typically be in a clickable [hyperlink](https://en.wikipedia.org/wiki/Hyperlink) the first time they are mentioned in the article. Files may also be hyperlinked so they can be referenced quickly.

## Bold

Names of items like database tables or filenames will be in **bold**.

## Inline Code

Inline code like `go build` or commands like `settings := config.LoadEnv(l, "")` will be in `backticks`.

## Codeblock

Most of the code you see will be in fenced coded blocks like this:

```go
func TestRegisterSuccess(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "Foo")
	form.Set("last_name", "Bar")
	form.Set("email", "fbar@example.com")
	form.Set("password", "password")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	r := testutil.EnsureCreated(t, w)
	assert.Equal(t, 36, len(r.Body.RecordID))
}
```

You may also see two different ways to execute the same command - one through probably Makefile and the other written out. This is to make it easy (`# Makefile`) to run commands without having to memorize or copy and paste them (`# Manual`).

```bash
# Makefile
make db-init

# Manual
docker run -d --name=govueapp_db_1 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql:5.7
```