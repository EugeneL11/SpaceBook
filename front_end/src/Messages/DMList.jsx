import { React, useState,useEffect } from "react";
import currentUser from "../Static.js";
import pPic from '../images/pp.png';
import Background from '../Background/background'
import axios from 'axios'
function DMList(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleDMMessage = props.toggleDMMessage

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
            // ask back end for dms
            setMsgs(examplefriends)
        },[])

    return (
    <div className="flex flex-col items-center min-h-screen">
        <div className="bg-white text-black text-xl m-14 md:p-6 p-2 rounded-xl w-full sm:w-3/4 lg:w-1/2 ">
            <div className="flex-1 overflow-y-auto flex flex-col">
                <h1 className="text-center text-3xl pt-2 mb-6">Chats</h1>
                {msgs ? msgs.map((msgObject, index) => (
                    <div  onClick = {() => {toggleDMMessage(msgObject.username)}} key = {index} className="bg-purple-300 hover:bg-purple-400 flex py-2 md:py-3 px-6 md:px-12 border-2 m-2 border-black rounded-xl">
                        <img src={pPic} className="pr-12 w-28 object-contain h-auto" alt="a placeholder profile picture"></img>
                        <div className="overflow-x-auto whitespace-nowrap">
                            <div className="font-bold text-xl pb-2">{msgObject.username}</div>
                            <div className="text-lg ml-10">
                                {msgObject.msg}
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
