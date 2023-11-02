





import { React, useState, useEffect} from "react";
import backPic from './back.png';
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
        text:"Go go go"}
    ]
    const [messages, setMessages] = useState(null)
        useEffect(()=>{
            // ask back end for dms
            setMessages(exampleDMs)
        },[])
    
    return (
    <div class="flex flex-col min-h-screen">
        <div class="bg-pink-900 py-5 flex justify-center items-center"> 
            <img  onClick = {toggleDMList}src={backPic} alt="a back arrow button" class="w-10 mr-auto pl-3"></img>
            <div  onClick = {toggleOtherProfile}class="text-white text-3xl absolute ">kevonosdiaz</div>
        </div>
        {messages ? messages.map((message,index)=>(
            message.sender === "Eugene" ? 
            <div key={index} className="text-blue-700"> {message.text}</div> :
            <div key={index} className="text-green-700"> {message.text}</div>
        )) : null}
        {/* <div class="flex-1 bg-gray-700"></div> */}
    </div>
    );


}

export default DMMessage;