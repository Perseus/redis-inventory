package adapter

import (
	"bytes"
	"testing"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProgressWriterTestSuite struct {
	suite.Suite
}

func (suite ProgressWriterTestSuite) TestStart() {
	pwMock := &ProgressMock{}
	prettyProgressWriter := &PrettyProgressWriter{
		pw:       pwMock,
		trackers: nil,
	}

	pwMock.On("AppendTracker", mock.Anything).Once().Run(func(args mock.Arguments) {
		tracker, ok := args.Get(0).(*progress.Tracker)

		suite.Assert().True(ok)
		suite.Assert().Equal(int64(42), tracker.Total)
	})

	pwMock.On("Render").Once()

	prettyProgressWriter.Start()

	time.Sleep(time.Millisecond)
	pwMock.AssertExpectations(suite.T())
}

func (suite ProgressWriterTestSuite) TestStop() {
	pwMock := &ProgressMock{}
	trackerMock := &TrackerMock{}
	prettyProgressWriter := &PrettyProgressWriter{
		pw:       pwMock,
		trackers: map[string]Tracker{"1": trackerMock},
	}

	pwMock.On("Stop").Once()
	trackerMock.On("MarkAsDone").Once()

	prettyProgressWriter.Stop()

	pwMock.AssertExpectations(suite.T())
	trackerMock.AssertExpectations(suite.T())
}

func (suite ProgressWriterTestSuite) TestIncrement() {
	pwMock := &ProgressMock{}
	trackerMock := &TrackerMock{}

	prettyProgressWriter := &PrettyProgressWriter{
		pw:       pwMock,
		trackers: map[string]Tracker{"1": trackerMock},
	}

	trackerMock.On("Increment", int64(1)).Once()

	prettyProgressWriter.Increment("1", 1)

	pwMock.AssertExpectations(suite.T())
	trackerMock.AssertExpectations(suite.T())
}

// Dumbest test I've ever written
func (suite ProgressWriterTestSuite) TestInit() {
	var buf bytes.Buffer
	pwMock := &ProgressMock{}
	style := progress.Style{}

	pwMock.On("SetAutoStop", mock.Anything)
	pwMock.On("SetTrackerLength", mock.Anything)
	pwMock.On("ShowETA", mock.Anything)
	pwMock.On("ShowOverallTracker", mock.Anything)
	pwMock.On("ShowTime", mock.Anything)
	pwMock.On("ShowTracker", mock.Anything)
	pwMock.On("ShowValue", mock.Anything)
	pwMock.On("SetMessageWidth", mock.Anything)
	pwMock.On("SetNumTrackersExpected", mock.Anything)
	pwMock.On("SetSortBy", mock.Anything)
	pwMock.On("SetStyle", mock.Anything)
	pwMock.On("SetTrackerPosition", mock.Anything)
	pwMock.On("SetUpdateFrequency", mock.Anything)
	pwMock.On("SetOutputWriter", &buf)
	pwMock.On("Style").Return(&style)

	prettyProgressWriter := &PrettyProgressWriter{
		pw:       pwMock,
		trackers: nil,
	}

	prettyProgressWriter.init(&buf)
}

func TestProgressWriterTestSuite(t *testing.T) {
	suite.Run(t, new(ProgressWriterTestSuite))
}
