/*
Copyright 2018 Blindside Networks

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"net/http"
	"strings"
	"fmt"
	"sync/atomic"

	bbbAPI "github.com/xzl8028/xenia-plugin-bigbluebutton/server/bigbluebuttonapiwrapper/api"
	"github.com/xzl8028/xenia-plugin-bigbluebutton/server/bigbluebuttonapiwrapper/dataStructs"
	"github.com/xzl8028/xenia-server/model"
	"github.com/xzl8028/xenia-server/plugin"
	"github.com/robfig/cron"
)

type Plugin struct {

	plugin.XeniaPlugin

	c                            *cron.Cron
	configuration                atomic.Value
	Meetings                     []dataStructs.MeetingRoom
	MeetingsWaitingforRecordings []dataStructs.MeetingRoom
}

//OnActivate runs as soon as plugin activates
func (p *Plugin) OnActivate() error {
	// we save all the meetings infos that are stored on in our database upon deactivation
	// loads the details back so everything works
	p.LoadMeetingsFromStore()

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}
	config := p.config()
	if err := config.IsValid(); err != nil {
		return err
	}

	bbbAPI.SetAPI(config.BASE_URL+"/", config.SALT)

	//every 2 minutes, look through active meetings and check if recordings are done
	p.c = cron.New()
	p.c.AddFunc("@every 2m", p.Loopthroughrecordings)
	p.c.Start()

	// register slash command '/bbb' to create a meeting
	return p.API.RegisterCommand(&model.Command{
		Trigger:          "bbb",
		AutoComplete:     true,
		AutoCompleteDesc: "Create a BigBlueButton meeting",
	})
}


//following method is to create a meeting from '/bbb' slash command
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	meetingpointer := new(dataStructs.MeetingRoom)
	err := p.PopulateMeeting(meetingpointer, nil, "") 
	if err != nil {
		return nil, model.NewAppError("ExecuteCommand", "Please provide a 'Site URL' in Settings > General > Configuration", nil, err.Error(), http.StatusInternalServerError)
	}

	p.createStartMeetingPost(args.UserId, args.ChannelId, meetingpointer)
	p.Meetings = append(p.Meetings, *meetingpointer)
	return &model.CommandResponse{}, nil

}

//this is the router to handle our server calls
//methods are all in responsehandlers.go
func (p *Plugin) ServeHTTP(c *plugin.Context,w http.ResponseWriter, r *http.Request) {

	config := p.config()
	if err := config.IsValid(); err != nil {
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	path := r.URL.Path
	if path == "/joinmeeting" {
		p.handleJoinMeeting(w, r)
	} else if strings.HasPrefix(path, "/endmeeting") {
		p.handleEndMeeting(w, r)
	} else if path == "/create" {
		p.handleCreateMeeting(w, r)
	} else if strings.HasPrefix(path, "/recordingready") {
		p.handleRecordingReady(w, r)
	} else if path == "/getattendees" {
		p.handleGetAttendeesInfo(w, r)
	} else if path == "/publishrecordings" {
		p.handlePublishRecordings(w, r)
	} else if path == "/deleterecordings" {
		p.handleDeleteRecordings(w, r)
	} else if strings.HasPrefix(path, "/meetingendedcallback") {
		p.handleImmediateEndMeetingCallback(w, r, path)
	} else if path == "/ismeetingrunning" {
		p.handleIsMeetingRunning(w, r)
	} else if path == "/redirect"{
		// html file to automatically close a window
			fmt.Fprintf(w,`<!doctype html><html><head><script>
				 								window.onload = function load() {
													window.open('', '_self', '');
													window.close();
													};
											</script></head><body></body></html>`)
	}else {
		http.NotFound(w, r)
	}
	return
}

func (p *Plugin) OnDeactivate() error {
	//on deactivate, save meetings details, stop check recordings looper, destroy webhook
	p.SaveMeetingToStore()
	p.c.Stop()
	return nil
}

func main() {
	plugin.ClientMain(&Plugin{})
}
