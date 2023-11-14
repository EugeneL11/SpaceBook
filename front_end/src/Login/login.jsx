import { React, useState } from "react";
import { userID } from "../Static.js";
function Login(props) {
    const toggleHomepage = props.toggleHomepage;
    const toggleRegister = props.toggleRegister;
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setError] = useState("")
    const loginAction = async () => {
        //ask backend
        if(username == "" || password == ""){
            setError("Actually put something in the text fields you jerk")
        }
        const result = await fetch(`http://localhost:8080/login/${encodeURIComponent(username)}/${encodeURIComponent(password)}`);

        const data = await result.json();
        console.log(data)
        if(data.error == "unable to find User"){
            setError("username or password incorrect")
        }
        else{
            userID = data.id
            console.log(userID)
            toggleHomepage();
        }
        
    };

    return (
        <div class="flex flex-col items-center">
            <h1 class="text-6xl py-11">SpaceBook</h1>

            <div class="bg-white text-black text-center text-xl p-10 sm:p-14 rounded-md w-2/3 sm:w-1/2 lg:w-2/5 xl:w-1/3 min-w-fit">
                <h4 class="text-3xl mb-10">Welcome Back!</h4>
                <div class="w-full">
                    <label for="username" class="block mb-2 text-left">
                        Username
                    </label>
                    <input
                        class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                        placeholder="Enter Username"
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}></input>
                </div>
                <div class="w-full mt-8">
                    <label for="password" class="block mb-2 text-left">
                        Password
                    </label>
                    <input
                        class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                        placeholder="Enter Password"
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}></input>
                </div>
                <button class="bg-purple-200 px-7 py-3 mt-10 rounded-lg" onClick={loginAction}>
                    Log In
                </button>
            </div>
            <h1>{errorMessage}</h1>
            <button class="text-xl mt-12" onClick={toggleRegister}>
                Sign Up
            </button>
        </div>
    );
}

export default Login;
