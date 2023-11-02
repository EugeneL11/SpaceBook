import { React, useState } from "react";
import {userID} from "../Static.js"
function FriendsList(props) {
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]
    const toggleMyProfile = props.toggleMyProfile
    const toggleFriendsList = props.toggleFriendsList
    const toggleOtherProfile = props.toggleOtherProfile
    return (<div className="flex flex-col">
        <h1>This is the FriendsList component</h1>
        <button onClick ={toggleMyProfile}>Go back to My Profile</button>
        {exampleFriends.map((friend) =>(
            <button onClick={() => {toggleOtherProfile(friend,toggleFriendsList)}}> See Other Profile: {friend}</button>
        ))}
    </div>);

}

export default FriendsList;