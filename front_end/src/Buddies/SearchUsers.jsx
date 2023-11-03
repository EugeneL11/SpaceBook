import { React, useState } from "react";
import {userID} from "../Static.js"
function SearchUsers(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const toggleSearchUser = props.toggleSearchUser
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]

    return (<div className="flex flex-col">
        <h1>This is the Search Users component</h1>
        <button onClick={toggleHomepage}>Go to Homepage Screen</button>
        {exampleFriends.map((friend) =>(
            <button onClick={() => {toggleOtherProfile(friend,toggleSearchUser)}}> See Other Profile: {friend}</button>
        ))}
    </div>);

}

export default SearchUsers;