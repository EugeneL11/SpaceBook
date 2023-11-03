import { React, useState } from "react";

function Login(props) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const loginAction = () =>{
        //ask backend
    }

    return (
    <div class="flex flex-col items-center">
        <h1 class="text-6xl sm:text-7xl pt-16 pb-16">SpaceBook</h1>

        <div class="bg-white text-black text-center text-xl p-10 sm:p-14 rounded-md w-1/3 min-w-fit">
            <h4 class="text-3xl mb-10">Welcome Back!</h4>
            <div class="w-full">
                <label for="username" class="block mb-2 text-left">Username</label>
                <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter Username" type="text" value={username} onChange={e => setUsername(e.target.value)}></input>
            </div>
            <div class="w-full mt-8">
                <label for="password" class="block mb-2 text-left">Password</label>
                <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter Password" type="text" value={password} onChange={e => setPassword(e.target.value)}></input>
            </div>
            <button class="bg-purple-200 px-7 py-3 mt-10 rounded-lg" onClick={loginAction}>Log In</button>
        </div>

        <button class="text-xl mt-12">Sign Up</button>
    </div>
    );

}

export default Login;