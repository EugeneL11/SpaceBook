import { React, useState } from "react";

function Settings(props) {
    const toggleLogin = props.toggleLogin
    const toggleMyProfile = props.toggleMyProfile
    return (<div className="flex flex-col">
        <h1>This is the Settings component</h1>
        <button onClick={toggleLogin}> Go to Login Screen</button>
        <button onClick={toggleMyProfile}>Go back to My Profile</button>
    </div>);

}

export default Settings;