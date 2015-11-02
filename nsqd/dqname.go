// +build !windows

package nsqd

func getBackendName(topicName, channelName string, deferred bool) string {
	// backend names, for uniqueness, automatically include the topic... <topic>:<channel>
	backendName := topicName + ":" + channelName
	if deferred {
		backendName += ":deferred"
	} else {
		backendName += ":ready"
	}
	return backendName
}
