import { React, useState } from "react";
import {userID} from "../Static.js"
function DMMessage(props) {
    const friendID = props.friendID
    /*const toggleDMMessage = props.toggleDMMessage(friendID)*/
    const toggleOtherProfile = () => props.toggleOtherProfile(friendID,null)
    const toggleDMList = props.toggleDMList
    return (<div className="flex flex-col">
        <h1>This is the DM Message component</h1>
        <h1>The current chat is with {friendID}</h1>
        <button onClick={toggleDMList}>Go to DM list</button>{
        <button onClick={toggleOtherProfile}> Go to Other Profile</button>}
        </div>);

}

export default DMMessage;