import { React, useState } from "react";

function Login(props) {
    return (
    <div class="flex flex-col items-center">
        <h1 class="text-6xl sm:text-7xl pt-16 pb-16">SpaceBook</h1>

        <div class="bg-white text-black text-center text-xl p-10 sm:p-14 rounded-md">
            <h4 class="text-3xl mb-10">Welcome Back!</h4>
            <div class="text-left">Username</div>
            <div class="text-left pt-5">Password</div>
            <button class="bg-purple-200 px-7 py-3 mt-12 rounded-lg">Log In</button>
        </div>

        <a class="text-xl mt-12">Sign Up</a>
    </div>
    );

}

export default Login;