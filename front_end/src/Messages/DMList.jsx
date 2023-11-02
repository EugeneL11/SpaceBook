import { React, useState } from "react";
import {userID} from "../Static.js"
function DMList(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleDMMessage = props.toggleDMMessage
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]

    return (<div className="flex flex-col">
        <h1>This is the DM List component</h1>
        <button onClick={toggleHomepage}>Go to Homepage Screen</button>
        {exampleFriends.map((friend) =>(
            <button onClick={() => {toggleDMMessage(friend)}}> See DM: {friend}</button>
        ))}
    </div>);

}

export default DMList;
