import { React, useState } from "react";
import currentUser from "../Static.js";
import axios from "axios";
import {serverpath} from "../Path.js";

function Register(props) {
    const toggleLogin = props.toggleLogin;
    const toggleHomepage = props.toggleHomepage;

    const [fullName, setName] = useState("");
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setError] = useState(null);
    const registerAction = () => {
        //ask backend


        if (username == "" || password == "" || fullName == "" || email == "") {
            setError("Please Enter Something");
        } else {
            const path = `/register/${encodeURIComponent(email)}/${encodeURIComponent(password)}/${encodeURIComponent(
                fullName
            )}/${encodeURIComponent(username)}`;
            axios.post(`${serverpath}${path}`).then((res) => {
                const data = res.data;
                if (data.status === "email already in use") {
                    setError(data.status);
                } else if (data.status === "unable to create account at this time") {
                    setError(data.status);
                } else if (data.status === "user name not available") {
                    setError(data.status);
                } else {
                console.log(data)
                currentUser.userID = data.user.id;
                currentUser.userName = data.user.user_name;
                currentUser.planet = data.user.planet
                currentUser.pfp = data.user.profile_picture_path;
                currentUser.bio = data.user.bio
                    toggleHomepage();
                }
            });
        }
        /*
        const path = "/testInsert/val"
        fetch(`${localhost}${path}`).then(res => {res.json().then(data =>{
            console.log(data)
        })})*/
    };

    return (
        <div className="flex flex-col items-center">
            <h1 className="text-6xl py-11">SpaceBook</h1>

            <div className="bg-white text-black text-center text-lg p-6 sm:px-12 rounded-md w-3/4 sm:w-1/2 lg:w-2/5 xl:w-1/3 min-w-fit">
                <h4 className="text-3xl mb-2">Welcome!</h4>
                <div className="w-full">
                    <label for="username" className="block text-left">
                        Full Name
                    </label>
                    <input
                        className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                        placeholder="Enter Full Name"
                        type="text"
                        value={fullName}
                        onChange={(e) => setName(e.target.value)}></input>
                </div>
                <div className="w-full mt-4">
                    <label for="password" className="block text-left">
                        Email Address
                    </label>
                    <input
                        className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                        placeholder="Enter Email Address"
                        type="text"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}></input>
                </div>
                <div className="w-full mt-4">
                    <label for="username" className="block text-left">
                        Username
                    </label>
                    <input
                        className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                        placeholder="Enter Username"
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}></input>
                </div>
                <div className="w-full mt-4">
                    <label for="password" className="block text-left">
                        Password
                    </label>
                    <input
                        className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                        placeholder="Enter Password"
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}></input>
                </div>
                <h1 className="text-red-500">{errorMessage}</h1>

                <button className="bg-purple-200 px-7 py-3 mt-6 rounded-lg" onClick={registerAction}>
                    Sign Up
                </button>
            </div>

            <button className="text-xl mt-12" onClick={toggleLogin}>
                Log In
            </button>
            {/* <h1>{errorMessage}</h1> */}
        </div>
    );
}

export default Register;
