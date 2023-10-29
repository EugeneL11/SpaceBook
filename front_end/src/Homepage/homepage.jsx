import { React, useState } from "react";

function Homepage(props) {
    
    const toggleExpandPost = props.toggleExpandPost;
    const examplePosts = ["Go Go", "React-ions"]
    return (<div className="flex flex-col"> 
        <h1>This is the HomePage component</h1>
        {examplePosts.map((postID) =>(
            <div onClick={() => toggleExpandPost(postID)}> Expand this post: {postID}</div>
        ))}
    </div>);

}

export default Homepage;