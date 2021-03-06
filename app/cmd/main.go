package cmd

import (
	"github.com/spf13/cobra"
)

var memoCmd = &cobra.Command{
	Use:   "memo",
	Short: "Run Memo app",
}

func Execute() {
	memoCmd.AddCommand(webCmd)
	memoCmd.AddCommand(actionNodeCmd)
	memoCmd.AddCommand(decodeCmd)
	memoCmd.AddCommand(userNodeCmd)
	memoCmd.AddCommand(scannerCmd)
	memoCmd.AddCommand(scanRecentCmd)
	memoCmd.AddCommand(fixPostEmojisCmd)
	memoCmd.AddCommand(fixNameEmojisCmd)
	memoCmd.AddCommand(viewPostCmd)
	memoCmd.AddCommand(backfillRootTxCmd)
	memoCmd.AddCommand(fixLeadingCharsCmd)
	memoCmd.AddCommand(addLikeNotificationsCmd)
	memoCmd.AddCommand(addReplyNotificationsCmd)
	memoCmd.AddCommand(parseTransactionCmd)
	memoCmd.AddCommand(addFollowerNotificationsCmd)
	memoCmd.AddCommand(profilePicDevCmd)
	memoCmd.AddCommand(populateFeedCmd)
	memoCmd.AddCommand(populateTopicInfoCmd)
	memoCmd.AddCommand(populateUserStatsCmd)
	memoCmd.AddCommand(minifyCmd)
	memoCmd.AddCommand(getUserInfoCmd)
	memoCmd.Execute()
}
