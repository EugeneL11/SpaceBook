import { React, useState } from "react";
import {userID} from "../Static.js"
function MyProfile(props) {
    const toggleFriendsList = props.toggleFriendsList
    const toggleSettings = props.toggleSettings
    const toggleHomepage = props.toggleHomepage
    return (<div className="flex flex-col">
        <h1>This is the My Profile component</h1>
        <button onClick={toggleSettings}>Go to Settings</button>
        <button onClick={toggleFriendsList}>Go to Friends List</button>
        <button onClick={toggleHomepage}>Go to Homepage</button>
    </div>);

}

export default MyProfile;