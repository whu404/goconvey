package convey

import (
	"github.com/smartystreets/goconvey/assertions"
)

// TODO: make examples file that invokes each of these so we know it's all working.

var (
	ShouldEqual       = assertions.ShouldEqual
	ShouldNotEqual    = assertions.ShouldNotEqual
	ShouldResemble    = assertions.ShouldResemble
	ShouldNotResemble = assertions.ShouldNotResemble
	ShouldPointTo     = assertions.ShouldPointTo
	ShouldNotPointTo  = assertions.ShouldNotPointTo
	ShouldBeNil       = assertions.ShouldBeNil
	ShouldNotBeNil    = assertions.ShouldNotBeNil
	ShouldBeTrue      = assertions.ShouldBeTrue
	ShouldBeFalse     = assertions.ShouldBeFalse

	ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
	ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
	ShouldBeLessThan             = assertions.ShouldBeLessThan
	ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
	ShouldBeBetween              = assertions.ShouldBeBetween
	ShouldNotBeBetween           = assertions.ShouldNotBeBetween

	ShouldContain    = assertions.ShouldContain
	ShouldNotContain = assertions.ShouldNotContain
	ShouldBeIn       = assertions.ShouldBeIn
	ShouldNotBeIn    = assertions.ShouldNotBeIn

	ShouldStartWith    = assertions.ShouldStartWith
	ShouldNotStartWith = assertions.ShouldNotStartWith
	ShouldEndWith      = assertions.ShouldEndWith
	ShouldNotEndWith   = assertions.ShouldNotEndWith
	ShouldBeBlank      = assertions.ShouldBeBlank
	ShouldNotBeBlank   = assertions.ShouldNotBeBlank

	ShouldPanic        = assertions.ShouldPanic
	ShouldNotPanic     = assertions.ShouldNotPanic
	ShouldPanicWith    = assertions.ShouldPanicWith
	ShouldNotPanicWith = assertions.ShouldNotPanicWith

	ShouldHaveSameTypeAs    = assertions.ShouldHaveSameTypeAs
	ShouldNotHaveSameTypeAs = assertions.ShouldNotHaveSameTypeAs

	ShouldHappenBefore         = assertions.ShouldHappenBefore
	ShouldHappenOnOrBefore     = assertions.ShouldHappenOnOrBefore
	ShouldHappenAfter          = assertions.ShouldHappenAfter
	ShouldHappenOnOrAfter      = assertions.ShouldHappenOnOrAfter
	ShouldHappenBetween        = assertions.ShouldHappenBetween
	ShouldHappenOnOrBetween    = assertions.ShouldHappenOnOrBetween
	ShouldNotHappenOnOrBetween = assertions.ShouldNotHappenOnOrBetween
	ShouldHappenWithin         = assertions.ShouldHappenWithin
	ShouldNotHappenWithin      = assertions.ShouldNotHappenWithin
)
