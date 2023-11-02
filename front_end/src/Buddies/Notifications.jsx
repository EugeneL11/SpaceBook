import { React, useState } from "react";
import {userID} from "../Static.js"
function Notifcations(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const toggleNotifications = props.toggleNotifications
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]

    return (<div className="flex flex-col">
        <h1>This is the Notifcations component</h1>
        <button onClick={toggleHomepage}>Go to Homepage Screen</button>
        {exampleFriends.map((friend) =>(
            <button onClick={() => {toggleOtherProfile(friend,toggleNotifications)}}> See Other Profile: {friend}</button>
        ))}
    </div>);

}

export default Notifcations;