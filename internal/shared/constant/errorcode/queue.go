package errorcode

import (
	"errors"
	"fmt"
)

var (
	QueueMessageNeedRetry  = fmt.Errorf("QueueMessageNeedRetry")
	QueueAlreadySubscribed = errors.New("already subscribed to subject")
)
