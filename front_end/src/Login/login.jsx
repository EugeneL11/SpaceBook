import { React, useState } from "react";
import {userID} from "../Static.js"
function Login(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleRegister = props.toggleRegister
    return (<div className="flex flex-col">
        <h1>This is the Login component</h1>
        <button onClick = {toggleRegister}>Go to Register</button>
        <button onClick={toggleHomepage}> Go to Homepage</button>
    </div>);

}

export default Login;