import { React, useState,useEffect } from "react";
import currentUser from "../Static.js";
import {serverpath} from "../Path.js";
import pPic from '../images/pp.png';
import Background from '../Background/background'
import axios from 'axios'
function DMList(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleDMMessage = props.toggleDMMessage
    const toggleNewDM = props.toggleNewDM

    const examplefriends = [
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"kingJames", 
        msg:"Appreciate it!!!"},
        {username:"Gene", 
        msg:"See ya"},
        {username:"kevonosdiaz", 
        msg:"Go Go Go Go Go"},
        {username:"duppy", 
        msg:"just grabbing a bakechef, be rigth there"},
        {username:"vicGPT", 
        msg:"The most optimal algorithm would be O(n log n)"} 
        ]
    
    const [msgs, setMsgs] = useState(null)
        useEffect(()=>{
            const path = `/userdms/${encodeURIComponent(currentUser.userID)}`
            axios.get(`${serverpath}${path}`).then((res) => {
                const data = res.data
                console.log(data)
                if (data.status === "no error") {
                    console.log("WHASUDS")
                    setMsgs(data.all_dms)
                } else {
                    console.log(data.status)
                }
            })

            // ask back end for dms
        },[])

    return (
    <div className="flex flex-col items-center min-h-screen">
        <div className="w-full flex items-center">
            <button className="mb-5 w-fit ml-6 mr-auto text-3xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
            <button className="mr-6 p-2 h-12 bg-blue-200 hover:bg-blue-400 rounded-md" onClick={toggleNewDM}>New Chat</button> 
        </div>
        <div className="bg-white text-black text-xl mt-2 mb-14 md:p-6 p-2 rounded-xl w-full sm:w-3/4 lg:w-2/3">
            <div className="flex-1 overflow-y-auto flex flex-col">
                <h1 className="text-center text-3xl pt-2 mb-6">Chats</h1>
                {msgs ? msgs.map((msgObject, index) => (
                    <div onClick = {() => {toggleDMMessage(msgObject.user_id, msgObject.user_name)}} key = {index} className="bg-purple-300 hover:bg-purple-400 hover:cursor-pointer flex py-2 md:py-3 px-6 md:px-12 border-2 m-2 border-black rounded-xl">
                        <img src={serverpath + msgObject.profile_picture_path} className="pr-12 w-28 object-contain h-auto" alt="a placeholder profile picture"></img>
                        <div className="overflow-x-auto whitespace-nowrap">
                            <div className="font-bold text-xl pb-2">{msgObject.user_name}</div>
                            <div className="text-lg ml-10">
                                {msgObject.most_recent_message}
                            </div>
                        </div>
                    </div> 
                )) : null}
            </div>
        </div>
    </div>
    );

}

export default DMList
