package cmd

import (
	"context"
	"encoding/json"
	"time"

	"github.com/alwitt/httpmq-go/api"
	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"
)

/*
getMgntStreamsCLISubcmds fetch the list of subcommands for managing streams through API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the list of stream CLI subcommands
*/
func getMgntStreamsCLISubcmds(mgntBaseArgs *ManagementCLIArgs) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "list",
			Usage:       "List all streams",
			Description: "List all streams through httpmq management API",
			Action:      actionListStreams(mgntBaseArgs),
		},
		{
			Name:        "create",
			Usage:       "Define a new stream",
			Description: "Define a new stream through httpmq management API",
			Flags:       actionCreateStreamCLIFlags(&mgntBaseArgs.stream.createStream),
			Action:      actionCreateStream(mgntBaseArgs),
		},
		{
			Name:        "get",
			Usage:       "Fetch one stream",
			Description: "Read information regarding one stream through management API",
			Flags:       actionGetStreamCLIFlags(&mgntBaseArgs.stream.getStream),
			Action:      actionGetStream(mgntBaseArgs),
		},
		{
			Name:        "delete",
			Usage:       "Delete one stream",
			Description: "Delete one stream through management API",
			Flags:       actionDeleteStreamCLIFlags(&mgntBaseArgs.stream.deleteStream),
			Action:      actionDeleteStream(mgntBaseArgs),
		},
		{
			Name:        "change-subject",
			Aliases:     []string{"cs"},
			Usage:       "Change target subject of stream",
			Description: "Changed the target subjects of a stream through management API",
			Flags:       actionChangeSubjectsCLIFlags(&mgntBaseArgs.stream.changeSubject),
			Action:      actionChangeSubjects(mgntBaseArgs),
		},
		{
			Name:        "change-retention",
			Aliases:     []string{"ca"},
			Usage:       "Change stream message retention",
			Description: "Changed a stream's message retention policy. Only exposed message age here.",
			Flags:       actionChangeRetentionCLIFlags(&mgntBaseArgs.stream.changeRetention),
			Action:      actionChangeRetention(mgntBaseArgs),
		},
	}
}

// ==============================================================================

/*
actionListStreams query the management API for list of all streams

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionListStreams(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		reqID, streams, err := client.ListAllStreams(context.Background())
		if err != nil {
			log.WithError(err).Errorf("Failed to list all streams")
			return err
		}
		type response struct {
			RequestID string                                   `json:"request_id"`
			Streams   map[string]api.ApisAPIRestRespStreamInfo `json:"streams"`
		}
		resp := response{RequestID: reqID, Streams: streams}
		t, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.WithError(err).Errorf("Failed to JSON format stream list")
			return err
		}
		log.Infof("Set of known streams\n%s", string(t))
		return nil
	}
}

// ==============================================================================

// createStreamCLIArgs cli arguments needed for creating new stream
type createStreamCLIArgs struct {
	Name      string `validate:"required"`
	Subjects  cli.StringSlice
	MaxMsgAge time.Duration
}

/*
actionCreateStreamCLIFlags fetch the list of CLI arguments needed by create stream action

 @param args *createStreamCLIArgs - arguments needed for create stream action
*/
func actionCreateStreamCLIFlags(args *createStreamCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream stream name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
		&cli.StringSliceFlag{
			Name:        "subjects",
			Usage:       "Target subjects for the stream. If not set, the name will be the target.",
			Aliases:     []string{"s"},
			Destination: &args.Subjects,
			Required:    false,
		},
		&cli.DurationFlag{
			Name:        "max-message-age",
			Usage:       "For data retention, the max duration a message is kept in the stream.",
			Aliases:     []string{"o"},
			Value:       time.Hour,
			DefaultText: "1h",
			Destination: &args.MaxMsgAge,
			Required:    false,
		},
	}
}

/*
actionCreateStream create new JetStream stream through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionCreateStream(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.stream.createStream); err != nil {
			return err
		}
		params := api.ManagementJSStreamParam{
			Name:   mgntBaseArgs.stream.createStream.Name,
			MaxAge: api.PtrInt64(mgntBaseArgs.stream.createStream.MaxMsgAge.Nanoseconds()),
		}
		subjects := mgntBaseArgs.stream.createStream.Subjects.Value()
		if len(subjects) > 0 {
			params.Subjects = &subjects
		} else {
			params.Subjects = &[]string{mgntBaseArgs.stream.createStream.Name}
		}
		reqID, err := client.CreateStream(context.Background(), params)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to create new stream %s", mgntBaseArgs.stream.createStream.Name,
			)
			return err
		}
		log.Infof("Created new stream %s. Request ID %s", mgntBaseArgs.stream.createStream.Name, reqID)
		return nil
	}
}

// ==============================================================================

// getStreamCLIArgs cli arguments needed for query one stream
type getStreamCLIArgs struct {
	Name string `validate:"required"`
}

/*
actionGetStreamCLIFlags fetch the list of CLI arguments needed by fetch stream info

 @param args *fetchStreamCLIArgs - arguments needed for fetch stream action
*/
func actionGetStreamCLIFlags(args *getStreamCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream stream name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
	}
}

/*
actionGetStream fetch stream info through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionGetStream(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.stream.getStream); err != nil {
			return err
		}
		reqID, info, err := client.GetStream(context.Background(), mgntBaseArgs.stream.getStream.Name)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to read stream %s info", mgntBaseArgs.stream.getStream.Name,
			)
			return err
		}
		type response struct {
			RequestID string                         `json:"request_id"`
			Stream    *api.ApisAPIRestRespStreamInfo `json:"stream"`
		}
		resp := response{RequestID: reqID, Stream: info}
		t, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.WithError(err).Errorf("Failed to JSON format stream info")
			return err
		}
		log.Infof("Stream %s\n%s", mgntBaseArgs.stream.getStream.Name, string(t))
		return nil
	}
}

// ==============================================================================

// deleteStreamCLIArgs cli arguments needed for delete one stream
type deleteStreamCLIArgs struct {
	Name string `validate:"required"`
}

/*
actionDeleteStreamCLIFlags fetch the list of CLI arguments needed by delete stream info

 @param args *deleteStreamCLIArgs - arguments needed for delete stream action
*/
func actionDeleteStreamCLIFlags(args *deleteStreamCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream stream name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
	}
}

/*
actionDeleteStream delete one stream through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionDeleteStream(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.stream.deleteStream); err != nil {
			return err
		}
		reqID, err := client.DeleteStream(context.Background(), mgntBaseArgs.stream.deleteStream.Name)
		if err != nil {
			log.WithError(err).Errorf("Failed to delete stream %s", mgntBaseArgs.stream.deleteStream.Name)
			return err
		}
		log.Infof("Delete stream %s. Request ID %s", mgntBaseArgs.stream.deleteStream.Name, reqID)
		return nil
	}
}

// ==============================================================================

// createChangeSubjectsCLIArgs cli arguments needed for changing stream target subjects
type createChangeSubjectsCLIArgs struct {
	Name     string `validate:"required"`
	Subjects cli.StringSlice
}

/*
actionChangeSubjectsCLIFlags fetch the list of CLI arguments needed by change subject action

 @param args *createStreamCLIArgs - arguments needed for change subjects action
*/
func actionChangeSubjectsCLIFlags(args *createChangeSubjectsCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream stream name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
		&cli.StringSliceFlag{
			Name:        "subjects",
			Usage:       "Target subjects for the stream",
			Aliases:     []string{"s"},
			Destination: &args.Subjects,
			Required:    true,
		},
	}
}

/*
actionChangeSubjects change one stream's target subject through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionChangeSubjects(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.stream.changeSubject); err != nil {
			return err
		}
		reqID, err := client.ChangeStreamSubjects(
			context.Background(),
			mgntBaseArgs.stream.changeSubject.Name,
			mgntBaseArgs.stream.changeSubject.Subjects.Value(),
		)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to change stream %s subjects", mgntBaseArgs.stream.deleteStream.Name,
			)
			return err
		}
		log.Infof(
			"Change stream %s subjects. Request ID %s", mgntBaseArgs.stream.deleteStream.Name, reqID,
		)
		return nil
	}
}

// ==============================================================================

// createChangeRetentionCLIArgs cli arguments needed for stream data retention
type createChangeRetentionCLIArgs struct {
	Name      string `validate:"required"`
	MaxMsgAge time.Duration
}

/*
actionChangeRetentionCLIFlags fetch the list of CLI arguments needed by change retention action

 @param args *createChangeRetentionCLIArgs - arguments needed for change retention action
*/
func actionChangeRetentionCLIFlags(args *createChangeRetentionCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream stream name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
		&cli.DurationFlag{
			Name:        "max-message-age",
			Usage:       "For data retention, the max duration a message is kept in the stream.",
			Aliases:     []string{"o"},
			Value:       time.Hour,
			DefaultText: "1h",
			Destination: &args.MaxMsgAge,
			Required:    false,
		},
	}
}

/*
actionChangeRetention change one stream's data retention through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionChangeRetention(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.stream.changeRetention); err != nil {
			return err
		}
		params := api.ManagementJSStreamLimits{
			MaxAge: api.PtrInt64(mgntBaseArgs.stream.changeRetention.MaxMsgAge.Nanoseconds()),
		}
		reqID, err := client.UpdateStreamLimits(
			context.Background(), mgntBaseArgs.stream.changeRetention.Name, params,
		)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to change stream %s data retention", mgntBaseArgs.stream.changeRetention.Name,
			)
			return err
		}
		log.Infof(
			"Changed stream %s data retention. Request ID %s",
			mgntBaseArgs.stream.changeRetention.Name,
			reqID,
		)
		return nil
	}
}
