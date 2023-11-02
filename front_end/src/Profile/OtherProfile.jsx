import { React, useState } from "react";
import {admin} from '../Static.js'
function OtherProfile(props) {
    const userID = props.userID
    const goBackScreen = props.goBackScreen
    return (<div className="flex flex-col">
        <h1>This is the Other Profile component</h1>
        <h2> The user being viewed is {userID}</h2>
        <button onClick={goBackScreen}> Go Back to the Screen you came from</button>
    </div>);

}

export default OtherProfile;