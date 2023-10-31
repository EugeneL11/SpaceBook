import { React, useState } from "react";
import backPic from './back.png';

function DMMessage(props) {
    return (
    <div class="flex flex-col min-h-screen">
        <div class="bg-pink-900 py-5 flex justify-center items-center"> 
            <img src={backPic} alt="a back arrow button" class="w-10 mr-auto pl-3"></img>
            <div class="text-white text-3xl absolute ">kevonosdiaz</div>
        </div>

        <div class="flex-1 bg-gray-700"></div>
    </div>
    );

}

export default DMMessage;