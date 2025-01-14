/*
*
context can be used to handle indeterministic stuff
(like waiting for endpoints to do something,
where anything could happen)
*/
package main

import (
	gocontext "panzerstadt/go-concepts/go_context"
	book "panzerstadt/go-concepts/go_mindset"
)

func main() {
	gocontext.TestContext()
	book.BookTest()
}
