import { React, useState } from "react";

function Register(props) {
    const toggleLogin = props.toggleLogin
    const toggleHomepage = props.toggleHomepage
    return (<div className="flex flex-col">
        <h1>This is the Register component</h1>
        <button onClick = {toggleLogin}>Go to Login</button>
        <button onClick={toggleHomepage}> Go to Homepage</button>
        </div>);

}

export default Register;