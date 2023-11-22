import { React, useState, useEffect} from "react";
import backPic from '../images/back.png';
import static1 from "../Static.js";
function DMMessage(props) {
    const friendID = props.friendID
    const toggleDMMessage = () => {props.toggleDMMessage(friendID)}
    const toggleOtherProfile = () => props.toggleOtherProfile(friendID,toggleDMMessage)
    const toggleDMList = props.toggleDMList
    const exampleDMs = [
        {sender:static1.userID,
        text:"Hey Kevonos"},
        {sender:friendID,
        text:"Go go go"},
        {sender:static1.userID,
        text:"Huh..??"},
        {sender:friendID,
        text:"GOOOOOOOO"},
        {sender:static1.userID, 
        text:"Im confused...."}
    ]
    const [messages, setMessages] = useState(null)
    const [messageValue, setmessageValue] = useState("")
        useEffect(()=>{
            // ask back end for dms
            setMessages(exampleDMs)
        },[])
    const sendMessage = () =>{
        // tell back end
        const newArr = [...messages, {sender: static1.userID, text: messageValue}]
        setMessages(newArr)
        setmessageValue("")
    }
    return (
    <div className="flex flex-col items-center min-h-screen">

        <button className="w-fit ml-10 mr-auto text-5xl hover:text-purple-300" onClick={toggleDMList}> {'←'} </button>

        {messages ? (
            <div className="bg-white min-h-[70%] w-3/4 md:w-1/2 min-w-fit mt-[-25px] pt-6 pb-12 px-16 rounded-xl">
                <div className="bg-purple-700 w-full rounded-md py-5 mb-10 flex justify-center items-center">
                    <div onClick={toggleOtherProfile} className="text-white text-3xl absolute hover:cursor-pointer">{friendID}</div>
                </div>
                <div className="flex flex-col gap-7">
                    {messages.map((message, index) => (
                        message.sender === static1.userID ? 
                        <div key={index} className="bg-purple-200 bg-opacity-50 w-fit ml-auto p-2 rounded-lg text-black text-right"> {message.text}</div> :
                        <div key={index} className="bg-purple-400 bg-opacity-50 w-fit mr-auto p-2 rounded-lg text-black text-left"> {message.text}</div>
                    ))}
                </div>
                
                <div className="flex items-center mt-20">
                    <input  
                        className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0 text-black"
                        placeholder="Enter a Wormhole Message"
                        type="text"
                        value = {messageValue}
                        onChange = {(e) => {setmessageValue(e.target.value)}}>
                    </input>
                    <button onClick ={sendMessage} className="ml-2 px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded">Send</button>
                </div>
                
            </div>
        ) : null}

    </div>
    );
}

export default DMMessage;