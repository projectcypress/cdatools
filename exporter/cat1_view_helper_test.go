package exporter

import . "gopkg.in/check.v1"

func (s *MySuite) TestValueOrNullFlavor(c *C) {
	c.Assert(valueOrNullFlavor(nil), Equals, "nullFlavor='UNK'")
	c.Assert(valueOrNullFlavor(0), Equals, "value='196912311900'")
	c.Assert(valueOrNullFlavor(int64(0)), Equals, "value='196912311900'")
	c.Assert(valueOrNullFlavor("0"), Equals, "value='196912311900'")
}

func (s *MySuite) TestEscape(c *C) {
	c.Assert(escape("&"), Equals, "&amp;")
	c.Assert(escape(1), Equals, "1")
	c.Assert(escape(nil), Equals, "")
}

func (s *MySuite) TestValueOrDefault(c *C) {
	c.Assert(valueOrDefault(nil, "hey"), Equals, "hey")
	c.Assert(valueOrDefault("hey", "hey thar"), Equals, "hey")
}

func (s *MySuite) TestCodeDisplay(c *C) {

}

func (s *MySuite) TestOidForCode(c *C) {

}
