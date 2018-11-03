package methods

import "github.com/thestrukture/IDE/types"

//
func FindmyBugs(args ...interface{}) (ajet []types.DebugObj) {
	packge := args[0]

	ajet = GetLogs(packge.(string))
	sapp := GetApp(GetApps(), packge.(string))

	if sapp.Pid != "" {
		activLog := types.DebugObj{Time: "Server", Bugs: []types.DebugNode{}}
		ajet = append([]types.DebugObj{activLog}, ajet...)
	}
	return

}