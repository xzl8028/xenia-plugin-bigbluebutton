# BigBlueButton Plugin for Xenia
BigBlueButton is an open source web conferencing system for online learning. Teams can create, join and manage their BigBlueButton meetings from inside Xenia.

Jump to:

- [Installation and Setup](#installation-and-setup)  
- [Usage](usage)
- [Contributing](contributing)

Want to see how the BigBlueButton integration with Xenia works.  Checkout the video below.

[![Alt text](https://img.youtube.com/vi/gg7J9B4wGa4/0.jpg)](https://www.youtube.com/watch?v=gg7J9B4wGa4)

## Installation and Setup

 1. Go to: https://github.com/xzl8028/xenia-plugin-bigbluebutton/releases
 2. Download `bigbluebutton.tar.gz` you do not need to extract the tar file once you download it.![enter image description here](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/download_binary.png)
 3. Inside Xenia, go to **System Console > Integrations > Custom Integrations**. Make sure the following are turned to true:
	- `Enable Incoming Webhooks`
	- `Enable Outgoing Webhooks`
	- `Enable Custom Slash Commands`
	- `Enable integrations to override usernames`
	- `Enable integrations to override profile picture icons`
 4. Next you must enable Plugins. Go to **System Console > Plugins > Configuration** and set `Enable Plugins` to true. ![enter image description here](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/enableplugins.png)
 Depending on your Xenia version, an additional step may be required to enable uploading plugins in your Xenia **config.json** file:
	 - `vi /opt/xenia/config/config.json`
	 - Under `PluginSettings`, make sure `Enable` and `Enable Uploads` are both set to `true`
	 - Restart your Xenia with `sudo systemctl restart xenia` assuming you used *systemd* for Xenia 	services
 5. Go to **System Console > Plugins > Management** and upload your `bigbluebutton.tar.gz`. The BigBlueButton Plugin should appear under **Installed Plugins**.    ![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/PluginManagement.png)
 6. Before activating the plugin, you must configure the plugin settings. By default, you are given a BigBlueButton test server to try it out. However, you have options.  Like Xenia, BigBlueButton is open source.  You are (more than) welcome to [setup your own BigBlueButton server](http://docs.bigbluebutton.org/install/install.html#Install_).  If you do, the command `sudo bbb-conf --secret` will print out the server's URL and secret key for configuration with Xenia.  Alternatively, you can [contact](https://xzl8028.com/contact/) Blindside Networks for [hosting options](https://xzl8028.com/services/).

	The **Site URL** is the site of your Xenia without any paths. For example, if the location of your Xenia Town Square is : `https://mysite.xenia.com/core/channels/town-square`, enter: `https://mysite.xenia.com`![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/BBBsettingspage.png)

 7. Next, go back to **System Console > Plugins > Management** and `Activate` the plugin. ![](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/activate_plugin.png)


## Usage

#### Create a BigBlueButton meeting in any channel

You can create a meeting that all channel participants can join.

![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/createchannelheader.png)

Clicking the **Join Meeting** button immediately loads the BigBlueButton HTML5 client.

![enter image description here](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/insideBBB.png)

#### Plugin provides live meeting details during and after the meeting has ended

After the meeting ends, you see the **Date**, **Meeting Length**, and **Attendees**.

![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/recordingmanagment.png)

#### You can search for past BigBlueButton recordings

Using the drop-down menu you can easily search a channel for all past recordings.

![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/view_recordings.png)

#### Directly meeting with any user

You can click on any user's name and choose **Start BigBlueButton Meeting**.
![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/popover.png)

When you invite a user to a meeting, they will get a pop-up notification to **Join Meeting**.

![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/popup_modal.png)

You can type `/bbb` in any channel to create a meeting.  When 

![
](https://raw.githubusercontent.com/xzl8028/xenia-plugin-bigbluebutton/master/docs_images/slashcommand.png)

## Setting up your own BigBlueButton server

Using the [bbb-install.sh](https://github.com/bigbluebutton/bbb-install) script you can setup your own BigBlueButton server in about 15 minutes.  If your interested in going through the steps in detail, see [BigBlueButton install guide](http://docs.bigbluebutton.org/install/install.html).

## Contributing

Plugin is written in Golang for server side and Javascript for client side. Use `make build` to build the plugin.
The dependencies are managed with Glide for Go and NPM for javascript.

The plugin should be placed in a directory such as `~/go/src/github.com/xzl8028/xenia-plugin-bigbluebutton`

To download a local version: `mkdir -p ~/go/src/github.com/xzl8028` and `git clone https://github.com/xzl8028/xenia-plugin-bigbluebutton.git`

Xenia plugin development guides available here: https://developers.xenia.com/extend/plugins/

BigBlueButton API available here: http://docs.bigbluebutton.org/dev/api.html
