import { React, useState, useEffect} from "react";
import backPic from '../images/back.png';
import {userID} from "../Static.js"
function DMMessage(props) {
    const friendID = props.friendID
    const toggleDMMessage = () => {props.toggleDMMessage(friendID)}
    const toggleOtherProfile = () => props.toggleOtherProfile(friendID,toggleDMMessage)
    const toggleDMList = props.toggleDMList
    const exampleDMs = [
        {sender:"Eugene",
        text:"Hey Kevonos"},
        {sender:"Kev",
        text:"Go go go"},
        {sender:"Eugene",
        text:"Huh..??"},
        {sender:"Kev",
        text:"GOOOOOOOO"},
        {sender:"Eugene", 
        text:"Im confused...."}
    ]
    const [messages, setMessages] = useState(null)
        useEffect(()=>{
            // ask back end for dms
            setMessages(exampleDMs)
        },[])
    
    return (
    <div class="flex flex-col items-center min-h-screen">
        <div class="bg-pink-900 w-full py-5 mt-[-15px] flex justify-center items-center"> 
            <img  onClick = {toggleDMList}src={backPic} alt="a back arrow button" class="w-10 mr-auto pl-3"></img>
            <div  onClick = {toggleOtherProfile}class="text-white text-3xl absolute ">kevonosdiaz</div>
        </div>

        {messages ? (
            <div className="bg-white min-h-[70%] w-3/4 md:w-1/2 min-w-fit mt-10 py-12 px-16 rounded-xl">
                <div className="flex flex-col gap-7">
                    {messages.map((message, index) => (
                        message.sender === "Eugene" ? 
                        <div key={index} className="bg-purple-200 bg-opacity-50 w-fit ml-auto p-2 rounded-lg text-black text-right"> {message.text}</div> :
                        <div key={index} className="bg-purple-400 bg-opacity-50 w-fit mr-auto p-2 rounded-lg text-black text-left"> {message.text}</div>
                    ))}
                </div>
                
                <div className="flex items-center mt-20">
                    <input  
                        class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0 text-black"
                        placeholder="Enter a Wormhole Message"
                        type="text">
                    </input>
                    <button class="ml-2 px-4 py-2 bg-blue-500 text-white rounded">Send</button>
                </div>
                
            </div>
        ) : null}

    </div>
    );
}

export default DMMessage;