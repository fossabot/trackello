// Copyright © 2016 Nick Klauer <klauer@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"time"

	"github.com/VojtechVitek/go-trello"
)

const (
	DateLayout      string = "2006-01-02T15:04:05Z"
	API_APPKEY      string = "TRELLO_APPKEY"
	API_TOKEN       string = "TRELLO_TOKEN"
	PREFERRED_BOARD string = "preferredBoard"
)

func CreateArgsForBoardActions() []*trello.Argument {
	var args []*trello.Argument
	twoWeeksAgo := time.Now().Add(-1 * time.Hour * 24 * 14).Format(DateLayout)
	args = append(args, trello.NewArgument("since", twoWeeksAgo))
	args = append(args, trello.NewArgument("limit", "500"))
	return args
}
