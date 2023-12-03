import { React, useState, useEffect} from "react";
import backPic from '../images/back.png';
import currentUser from "../Static.js";
import {serverpath} from "../Path.js";
import axios from 'axios'
function DMMessage(props) {
    const friendID = props.friendID
    const toggleDMMessage = () => {props.toggleDMMessage(friendID)}
    const toggleOtherProfile = () => props.toggleOtherProfile(friendID,toggleDMMessage)
    const toggleDMList = props.toggleDMList
    const exampleDMs = [
        {sender:currentUser.userID,
        text:"Hey Kevonos"},
        {sender:friendID,
        text:"Go go go"},
        {sender:currentUser.userID,
        text:"Huh..??"},
        {sender:friendID,
        text:"GOOOOOOOO"},
        {sender:currentUser.userID, 
        text:"Im confused...."}
    ]
    const [messages, setMessages] = useState(null)
    const [messageValue, setmessageValue] = useState("")
    const subsetSize = 10
    useEffect(()=>{
        const path = `/getmessages/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(friendID)}/${encodeURIComponent(subsetSize)}`
        console.log(path)
        axios.get(`${serverpath}${path}`).then((res) => {
            const data = res.data
            console.log(data)
            //setMessages() what do i set this as? it only returns a status and a bool for moreMessages
        })

        // ask back end for dms
        //setMessages(exampleDMs)
    },[])
    
    const handleKeyPress = (event) => {
        // Check if the Enter key was pressed (key code 13)
        if (event.key === 'Enter') {
            // Trigger the button click action
            sendMessage();
        }
    };
    const sendMessage = () =>{
        const sendPath = `/senddm/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(friendID)}/${encodeURIComponent(messageValue)}`
        axios.post(`${serverpath}${sendPath}`).then((res) => {
            const sendData = res.data
            console.log(sendData)
            //what does this axios post even do? we only get sent back a status
            if (sendData.status === "no error") {
                const newArr = [...messages, {sender: currentUser.userID, text: messageValue}]
                setMessages(newArr)
                setmessageValue("")
            } else {
                console.log(sendData.status)
            }
        })

        // tell back end
    }
    return (
    <div className="flex flex-col items-center min-h-screen">

        <button className="mb-5 w-fit ml-6 mr-auto text-3xl hover:text-purple-300" onClick={toggleDMList}> {'←'} </button>

        {messages ? (
            <div className="bg-white min-h-[70%] w-full sm:w-3/4 lg:w-1/2 min-w-fit mt-[-20px] pt-6 pb-12 px-10 lg:px-16 rounded-xl">
                <div className="bg-purple-700 w-full rounded-md py-5 mb-10 flex justify-center items-center">
                    <div onClick={toggleOtherProfile} className="text-white text-3xl absolute hover:cursor-pointer">{friendID}</div>
                </div>
                <div className="flex flex-col gap-1">
                    {messages.map((message, index) => (
                        message.sender === currentUser.userID ? 
                        <div key={index} className="bg-purple-200 bg-opacity-50 w-fit ml-auto p-2 rounded-lg text-black text-right"> {message.text}</div> :
                        <div key={index} className="bg-purple-400 bg-opacity-50 w-fit mr-auto p-2 rounded-lg text-black text-left"> {message.text}</div>
                    ))}
                </div>
                
                <div className="flex items-center mt-10">
                    <input  
                        className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0 text-black"
                        placeholder="Enter a Wormhole Message"
                        type="text"
                        value = {messageValue}
                        onKeyPress={handleKeyPress}
                        onChange = {(e) => {setmessageValue(e.target.value)}}>
                    </input>
                    <button onClick={sendMessage} className="ml-2 px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded">Send</button>
                </div>
                
            </div>
        ) : null}

    </div>
    );
}

export default DMMessage;