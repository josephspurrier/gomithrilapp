---
title: Smarter Testing
author: Joseph Spurrier
authorURL: http://github.com/josephspurrier
authorFBID: 1225770017
---

When I first wrote the code and tests in note.go and note_test.go, I wanted to
see if I could get to 100% code coverage. It's a good exercise if you've never
done it before to see what you have to do to get there. It requires a lot more
effort than it's worth for most applications, but still a good exercise.
[View commit](https://github.com/josephspurrier/govueapp/commit/5204659c7ab7a19c24b2d3c29f2fb03b7760f119).

In Go, table driven tests is very common. You can read able them
[here](https://github.com/golang/go/wiki/TableDrivenTests). It helps cut down
on all the boilerplate code you have to repeat everywhere.

It's also helpful if you have a setup and teardown process that is run with
each of your tests. Since that code was copy and pasted in many different places,
considated they to easy to use functions.
[View commit](https://github.com/josephspurrier/govueapp/commit/89020eaea7cd7922d5936c566513c30602e2e701).

Code samples below.

<!--truncate-->

### Before

Before the refactor, the code looked like this in many of the test functions. 
This test ensures that a user can register themselves as a new user.

```go
func TestRegisterSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := testutil.Services(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "Foo")
	form.Set("last_name", "Bar")
	form.Set("email", "fbar@example.com")
	form.Set("password", "guess123")
	w := tr.SendJSON(t, p, "POST", "/api/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(r.Body.RecordID))
}
```

### After

Notice the boilerplate code at the top of the function has been consolidated
to a few functions. It's also more clear that teardown will happen at
the end for everything that was setup whereas in the old code, it was a little
confusing on what was actually being set up and torn down.

Also, notice the asserts at the bottom of the function have been moved into a
separate function since there are many tests that will use the same test logic.

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