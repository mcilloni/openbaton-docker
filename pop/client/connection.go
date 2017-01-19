package client

import (
	"github.com/mcilloni/openbaton-docker/pop"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// connection represents a session with the server.
// connection instances are cached and discarded in case they become invalid.
type session struct {
	conn *grpc.ClientConn

	tok string

	invalid bool
}

// newSession initialises a session, authenticating into the service
// and getting a token.
func newSession(creds Credentials) (*session, error) {
	sess := new(session)

	// WithInsecure allows for non-TLS connections.
	gconn, err := grpc.Dial(
		creds.Host,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(sess.interceptor),
	)

	if err != nil {
		return nil, err
	}

	sess.conn = gconn

	s := sess.stub()

	tk, err := s.Login(context.Background(), creds.toPop())
	if err != nil {
		return nil, err
	}
	sess.tok = tk.Value

	return sess, nil
}

// ctx returns a Context in which the token has been set as metadata.
func (sess *session) ctx(ctx context.Context) context.Context {
	return metadata.NewContext(ctx, metadata.Pairs(pop.TokenKey, sess.tok))
}

// interceptor intercepts each call, injects the token, executes the call and then checks if the token is valid.
// In case  it is invalid, it marks the current session as invalid.
func (sess *session) interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if sess.invalid {
		return ErrInvalidClient
	}

	// if we are not logging in, inject the token metadata in the context
	if method != "/pop.Pop/Login" {
		ctx = sess.ctx(ctx)
	}

	err := invoker(ctx, method, req, reply, cc, opts...)

	// If the error we got is permission denied, the token is not valid, so the connection structure
	// must be dropped.
	if grpc.Code(err) == codes.PermissionDenied {
		sess.invalid = true

		return ErrInvalidClient
	}

	return err
}

func (sess *session) stub() pop.PopClient {
	return pop.NewPopClient(sess.conn)
}
