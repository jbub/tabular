package tabular

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RowTestSuite struct {
	suite.Suite
}

func (s *RowTestSuite) TestHasTag() {
	r := NewRow()
	r.AddTag("tag1")

	s.True(r.HasTag("tag1"))
	s.False(r.HasTag("tag2"))
}

func (s *RowTestSuite) TestHasAllTagsFalse() {
	r := NewRow()
	r.AddTag("tag1")

	has := r.HasAllTags("tag1", "tag2")
	s.False(has)
}

func (s *RowTestSuite) TestHasAllTagsTrue() {
	r := NewRow()
	r.AddTag("tag1")
	r.AddTag("tag2")

	has := r.HasAllTags("tag1", "tag2")
	s.True(has)
}

func (s *RowTestSuite) TestHasAllTagsEmpty() {
	r := NewRow()

	has := r.HasAllTags("tag1", "tag2")
	s.False(has)
}

func (s *RowTestSuite) TestHasAnyTagsTrue() {
	r := NewRow()
	r.AddTag("tag1")

	has := r.HasAnyTags("tag1", "tag2")
	s.True(has)
}

func (s *RowTestSuite) TestHasAnyTagsFalse() {
	r := NewRow()
	r.AddTag("tag3")

	has := r.HasAnyTags("tag1", "tag2")
	s.False(has)
}

func (s *RowTestSuite) TestHasAnyTagsEmpty() {
	r := NewRow()

	has := r.HasAnyTags("tag1", "tag2")
	s.False(has)
}

func TestRowTestSuite(t *testing.T) {
	suite.Run(t, new(RowTestSuite))
}
