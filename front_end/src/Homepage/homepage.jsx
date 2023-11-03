import { React, useState } from "react";
import {userID} from "../Static.js"
function Homepage(props) {
    
    const toggleExpandPost = props.toggleExpandPost;
    const toggleOtherProfile = props.toggleOtherProfile;
    const toggleHomePage = props.toggleHomePage;
    const examplePosts = [{
        userID : "Kevon",
        postID: "Go Go"}, {
            userID: "Duppy",
            postID: "React-ions"}]
    return (<div className="flex flex-col"> 
        <h1>This is the HomePage component</h1>
        {examplePosts.map((post,index) =>(
            <div key = {index}>
                <div onClick={() => toggleOtherProfile(post.userID,toggleHomePage)  }> Expand the user: {post.userID}</div>
                <div onClick={() => toggleExpandPost(post.postID)}> Expand this post: {post.postID}</div>
            </div>
        ))}
    </div>);

}

export default Homepage;