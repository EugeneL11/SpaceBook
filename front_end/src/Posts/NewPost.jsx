import { React, useState } from "react";

function NewPost(props) {
    const toggleHomepage = props.toggleHomepage
    return (<div className="flex flex-col">
        <h1>This is the New Post component</h1>
        <h1>It may need more than one screen which some bozo(me) has to implement the switching of screens for that</h1>
        <button onClick={toggleHomepage}> Go back to Homepage</button>
        </div>);
}

export default NewPost;