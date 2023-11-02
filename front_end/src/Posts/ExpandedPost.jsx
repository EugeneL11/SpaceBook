import { React, useState } from "react";
import {admin} from "../Static.js"
function ExpandedPost(props) {
    const postID = props.postID
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const togglePost = () => props.toggleExpandPost(postID)
    return (<div className="flex flex-col">
        <h1>This is the Expanded Post component</h1>
        <h1>This is Post = {postID}</h1>
        <button onClick ={toggleHomepage}>Go back to HomePage</button>
        {exampleFriends.map((friend) =>(
            <button onClick={() => {toggleOtherProfile(friend,togglePost)}}> See Other Profile by clicking their comment: {friend}</button>
        ))}
    </div>);
}

export default ExpandedPost;